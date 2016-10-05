package routers

import (
	"Game/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.WebSocketController{})
	beego.Router("/join", &controllers.WebSocketController{}, "get,post:Join")
}
