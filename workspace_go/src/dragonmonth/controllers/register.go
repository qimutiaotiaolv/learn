package controllers

import (
	"crypto/md5"
	"dragonmonth/models"
	"dragonmonth/mq"
	"dragonmonth/util"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type Register struct {
	beego.Controller
}

/**
 * /register/registe?tel=13301103716&code=11111&pwd=admin121yyx&repwd=admin121yyx
 * @param  {[type]} this *Register)    Registe( [description]
 * @return {[type]}      [description]
 */
func (this *Register) Registe() {
	model := models.RequestRegisteInfo{}
	err := this.ParseForm(&model)
	if err != nil {
		response := models.NewResponseModel(2, "参数错误", nil)
		this.Data["json"] = response
		this.ServeJSON()
		return
	}
	var u models.User
	if count, err := models.SharedOrmEngine.Alias("o").Where("o.Telephone = ?", model.Tel).Count(&u); err == nil {
		if count != 0 {
			response := models.NewResponseModel(3, "手机号已被占用", nil)
			this.Data["json"] = response
			this.ServeJSON()
			this.DelSession(model.Tel + "-security") //删除无用验证码
			return
		}
	}
	sessionCode := this.GetSession(model.Tel + "-security")
	if sessionCode == model.Code { //注册成功
		response := models.NewResponseModel(0, "注册成功", nil)
		//md5加密
		m := md5.New()
		m.Write([]byte(model.Pwd))
		encodePwd := hex.EncodeToString(m.Sum(nil))
		user := models.NewUser("", model.Tel, encodePwd, "", model.Gender)
		messageModel := models.NewMessageInsertUser(models.MQ_INSERT, "", user)
		mq.Producer_MessageSender_Register_Insert2Database(messageModel)
		// if _, err := models.SharedOrmEngine.Insert(user); err != nil {
		// 	response := models.NewResponseModel(5, "数据库写入错误:"+err.Error(), nil)
		// 	this.Data["json"] = response
		// 	this.ServeJSON()
		// 	return
		// }

		this.Data["json"] = response
		this.ServeJSON()
		this.DelSession(model.Tel + "-security") //注册成功删除验证码

	} else {
		response := models.NewResponseModel(1, "验证码错误", nil)
		this.Data["json"] = response
		this.ServeJSON()
	}

}

/**
 * 发送短信验证码 /register/sendcode?tel=13301103716
 * @param  {[type]} this *Register)    SendCode( [description]
 * @return {[type]}      [description]
 */
func (this *Register) SendCode() {
	tel := this.GetString("tel")
	//验证电话号码的正确性，客户端会首先验证一次
	user := &models.User{
		Telephone: tel,
	}
	isHas, err := models.SharedOrmEngine.Get(user)
	if err != nil {
		response := models.NewResponseModel(7, "数据库错误,请稍后重试:"+err.Error(), nil)
		responseString, err := response.Json()
		if err != nil {
			this.Ctx.WriteString("LoginController::UserLogin->JSON解析错误:" + err.Error())
			beego.BeeLogger.Error("LoginController::UserLogin->JSON解析错误:%s", err.Error())
			return
		}
		this.Ctx.WriteString(responseString)
		return
	}
	if isHas { //用户已经存在
		response := models.NewResponseModel(8, "手机号已被占用,请用其他手机号重试", nil)
		responseString, err := response.Json()
		if err != nil {
			this.Ctx.WriteString("Register::SendCode->JSON解析错误:" + err.Error())
			beego.BeeLogger.Error("Register::SendCode->JSON解析错误:%s", err.Error())
			return
		}
		this.Ctx.WriteString(responseString)
		return
	}
	//验证是否不到30秒重复发送
	security := util.NewSecurity(6)
	message := beego.AppConfig.String("captcha::template") + security
	apikey := beego.AppConfig.String("captcha::apikey")
	mobile := tel

	//向云片网发送请求
	req := httplib.Post(beego.AppConfig.String("captcha::requestUrl"))
	req.Param("text", message)
	req.Param("apikey", apikey)
	req.Param("mobile", mobile)
	var result models.PYSecurityCodeMsgSendedResult
	err = req.ToJSON(&result)
	if err != nil {
		s, _ := req.String()
		response := models.NewResponseModel(1, s, nil)
		buf, err := json.Marshal(response)
		if err != nil {
			this.Ctx.WriteString("失败")
		}
		this.Ctx.WriteString(string(buf))
		//{"code":0,"msg":"OK","result":{"count":1,"fee":0.055,"sid":7684069618}}
		return
	}

	msg := "验证码发送成功"
	if result.Code != 0 {
		msg = "验证码发送失败:"
		rs, err := req.String()
		if err == nil {
			msg = msg + rs
		}
		response := models.NewResponseModel(2, msg, nil)
		buf, err := json.Marshal(response)
		if err != nil {
			this.Ctx.WriteString("失败")
		}
		this.Ctx.WriteString(string(buf))
		return
	}

	this.SetSession(tel+"-security", security)
	response := models.NewResponseModel(0, msg, result)
	buf, err := json.Marshal(response)
	if err != nil {
		this.Ctx.WriteString("失败")
	}
	this.Ctx.WriteString(string(buf))
}

func (this *Register) TestGetSession() {
	tel := this.GetString("tel")
	sessionObj := this.GetSession(tel + "-security")
	session, ok := sessionObj.(string)
	if !ok {
		res := fmt.Sprintf("获取session失败:%#v", sessionObj)
		this.Ctx.WriteString(res)
		return
	}
	this.Ctx.WriteString(session)
}
