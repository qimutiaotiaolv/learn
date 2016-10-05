package utils

import (
	"Game/beans"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

func HttpGetArea(ip string) (beans.AreaBean, error) {
	serverUrl := beego.AppConfig.String("urls::search_area") + "?ip=" + ip
	req := httplib.Get(serverUrl)
	var bean beans.AreaBean
	err := req.ToJSON(&bean)
	return bean, err
}
