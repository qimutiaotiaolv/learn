package routers

import (
	"dragonmonth/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/fileupload", &controllers.SaveUploadFileController{})
	beego.Router("/payment", &controllers.Payment{})
	beego.Router("/payment/return", &controllers.Payment{}, "get,post:Return")
	beego.Router("/payment/notify", &controllers.Payment{}, "get,post:Notify")
	beego.Router("/register/registe", &controllers.Register{}, "get:Registe")
	beego.Router("/register/sendcode", &controllers.Register{}, "get:SendCode")
	beego.Router("/register/getsession", &controllers.Register{}, "get:TestGetSession")
	beego.Router("/login/userlogin", &controllers.LoginController{}, "get,post:UserLogin")
	beego.Router("/login/userloginout", &controllers.LoginController{}, "post:UserLogout")
	beego.Router("/login/update", &controllers.LoginController{}, "post:UpdateUserData")
	beego.Router("/wxpayment", &controllers.WXPay{})
	//客服
	// beego.Router("/customer", &controllers.CustomerServices{}, "get,post:Customer")

}
