package controllers

import (
	"dragonmonth/models"
	// "dragonmonth/shared"
	// "dragonmonth/util"
	// "encoding/json"
	"crypto/md5"
	"encoding/hex"
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/httplib"
)

type LoginController struct {
	beego.Controller
}

/**
 *用户登录
 *192.168.199.120:9533/login/userlogin?tel=13301103716&pwd=admin121yyx
 * @param  {[type]} this *Login)       UserLogin( [description]
 * @return {[type]}      [description]
 */
func (this *LoginController) UserLogin() {
	tel := this.GetString("tel")
	pwd := this.GetString("pwd")
	if tel == "" || pwd == "" {
		response := models.NewResponseModel(1, "参数错误", nil)
		responseString, err := response.Json()
		if err != nil {
			this.Ctx.WriteString("LoginController::UserLogin->JSON解析错误:" + err.Error())
			beego.BeeLogger.Error("LoginController::UserLogin->JSON解析错误:%s", err.Error())
			return
		}
		this.Ctx.WriteString(responseString)
		return
	}
	//md5加密
	m := md5.New()
	m.Write([]byte(pwd))
	encodePwd := hex.EncodeToString(m.Sum(nil))
	user := &models.User{}
	user.Telephone = tel
	user.Password = encodePwd
	isHas, err := models.SharedOrmEngine.Get(user)
	if err != nil {
		response := models.NewResponseModel(2, "数据库错误:"+err.Error(), nil)
		responseString, err := response.Json()
		if err != nil {
			this.Ctx.WriteString("LoginController::UserLogin->JSON解析错误:" + err.Error())
			beego.BeeLogger.Error("LoginController::UserLogin->JSON解析错误:%s", err.Error())
			return
		}
		this.Ctx.WriteString(responseString)
		return
	}
	if !isHas {
		response := models.NewResponseModel(3, "用户不存在或密码错误", nil)
		responseString, err := response.Json()
		if err != nil {
			this.Ctx.WriteString("LoginController::UserLogin->JSON解析错误:" + err.Error())
			beego.BeeLogger.Error("LoginController::UserLogin->JSON解析错误:%s", err.Error())
			return
		}
		this.Ctx.WriteString(responseString)
		return
	}
	response := models.NewResponseModel(0, "登录成功", nil)
	responseDatas := models.NewApiLoginResponse(user.UserId)
	response.Datas = responseDatas
	responseString, err := response.Json()
	if err != nil {
		this.Ctx.WriteString("LoginController::UserLogin->JSON解析错误:" + err.Error())
		beego.BeeLogger.Error("LoginController::UserLogin->JSON解析错误:%s", err.Error())
		return
	}
	// shared.SharedData.SharedSessionLong.Set(responseDatas.Tocken, responseDatas)
	this.SetSession(responseDatas.Tocken, responseDatas)
	this.Ctx.WriteString(responseString)
}

/**
 * 退出登录
 * 192.168.199.120:9533/login/userloginout?tocken=kjalksd-ajk3894758
 * @param  {[type]} this *LoginController) UserLogout( [description]
 * @return {[type]}      [description]
 */
func (this *LoginController) UserLogout() {

}

/**
 * 修改信息
 * 192.168.199.120:9533/login/update?tocken=kl3249891$%^&type=0&value=
 * @param  {[type]} this *LoginController) UpdateUserData( [description]
 * @return {[type]}      [description]
 */
func (this *LoginController) UpdateUserData() {

}
