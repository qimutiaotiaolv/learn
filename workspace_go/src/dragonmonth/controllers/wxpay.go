package controllers

import (
	"dragonmonth/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/lunny/weixinpay"
)

var SharedWXPay *weixinpay.Merchant

type WXPay struct {
	beego.Controller
}

func (this *WXPay) Get() {
	this.Playorder()
}

/**
 * 192.168.199.120:9533/wxpayment/playorder?orderid=12323&clientip=192.168.1.1
 * @param  {[type]} this *WXPay)       WXPayment( [description]
 * @return {[type]}      [description]
 */
func (this *WXPay) Playorder() {
	orderid := this.GetString("orderid")
	clientip := this.GetString("clientip")
	if orderid == "" || clientip == "" {
		responseModel := models.NewResponseModel(1, "传参错误", nil)
		responseJson, err := responseModel.Json()
		if err != nil {
			this.Ctx.WriteString("")
			beego.BeeLogger.Error("WXPay::WXPayment:%s", err.Error())
			return
		}
		this.Ctx.WriteString(responseJson)
		return
	}
	// this.Ctx.WriteString("orderid:" + orderid + ",clientip:" + clientip)
	result, err := SharedWXPay.PlaceOrder(orderid, "测试商品", "这是测试商品的描述", clientip, "192.168.199.120:9533/wxpayment/wxpayNotify", 1)
	if err != nil {
		responseModel := models.NewResponseModel(2, "微信生成预支付交易单失败:"+err.Error(), nil)
		responseJson, err := responseModel.Json()
		if err != nil {
			this.Ctx.WriteString("")
			beego.BeeLogger.Error("WXPay::WXPayment:%s", err.Error())
			return
		}
		this.Ctx.WriteString(responseJson)
		return
	}
	this.Ctx.WriteString(fmt.Sprintf("生成预支付交易单成功:%#v", result))
}

/**
 * 微信支付(生成预支付交易单)异步回到
 * @param  {[type]} this *WXPay)       WXPayNotify( [description]
 * @return {[type]}      [description]
 */
func (this *WXPay) WXPayNotify() {

}

func init() {
	appid := beego.AppConfig.String("wxpay::appId")
	appkey := beego.AppConfig.String("wxpay::appkey")
	mchid := beego.AppConfig.String("wxpay::mchid")
	appsecret := beego.AppConfig.String("wxpay::appsecret")
	beego.BeeLogger.Info("appkey:%s\nappid:%s\nmchid:%s\nappsecret:%s\n", appkey, appid, mchid, appsecret)
	SharedWXPay = weixinpay.NewMerchant(appid, appkey, mchid, appsecret)
}
