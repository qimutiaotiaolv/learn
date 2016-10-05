package main

import (
	_ "Game/mq"
	_ "Game/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
