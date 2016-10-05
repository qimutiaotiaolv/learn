package controllers

import (
	"github.com/ascoders/alipay"
	"github.com/astaxie/beego"
)

var SharedPayment *alipay.Client

type Payment struct {
	beego.Controller
}

func (this *Payment) Get() {
	act := this.GetString("action")
	switch act {
	case "jump2PayPage":
		orderID := this.GetString("orderID")
		this.JumpPaymentPageWithOrderID(orderID)
		break
	}
}

/**
 * 同步回调
 * @param  {[type]} this *Payment)     Return( [description]
 * @return {[type]}      [description]
 */
func (this *Payment) Return() {
	result := SharedPayment.Return(&this.Controller)
	this.Ctx.WriteString("(同步回调)支付成功")
	beego.BeeLogger.Info("收到支付宝同步回调:%#v", result)
	if result.Status == 1 { //付款成功，处理订单
		//处理订单
	}
}

/**
 * 异步回调
 * @param  {[type]} this *Payment)     Notify( [description]
 * @return {[type]}      [description]
 */
func (this *Payment) Notify() {
	result := SharedPayment.Notify(&this.Controller)
	beego.BeeLogger.Info("收到支付宝异步回调:%#v", result)
	if result.Status == 1 { //付款成功，处理订单
		//处理订单
	}
}

func (this *Payment) JumpPaymentPageWithOrderID(orderID string) {
	//数据库查询
	form := SharedPayment.Form(alipay.Options{
		OrderId:  orderID,
		Fee:      0.02,
		NickName: "测试人员",
		Subject:  "罚款了",
	})
	if form == "" {
		beego.BeeLogger.Error("生成form错误")
		return
	}
	beego.BeeLogger.Info("%s", form)
	this.Data["form"] = form
	this.TplName = "payment.html"
}

func init() {
	partner := beego.AppConfig.String("alipay::partner")
	key := beego.AppConfig.String("alipay::key")
	safeCode := beego.AppConfig.String("alipay::safeCode")
	// beego.BeeLogger.Info("合作者ID:%s\n合作者私钥:%s\n", partner, key)
	SharedPayment = alipay.NewClient(partner, key, safeCode, "http://192.168.199.120:8088/payment/return", "http://192.168.199.120:8088/payment/notify", "xunbaivip@163.com")
}
