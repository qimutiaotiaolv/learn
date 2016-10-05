package beans

import (
	"encoding/json"
	"github.com/astaxie/beego"
	// "errors"
)

type AreaData struct {
	Country    string `json:"country"`
	Country_id string `json:"country_id"`
	Area       string `json:"area"`
	Area_id    string `json:"area_id"`
	Region     string `json:"region"`
	Region_id  string `json:"region_id"`
	City       string `json:"city"`
	City_id    string `json:"city_id"`
	County     string `json:"county"`
	County_id  string `json:"county_id"`
	Isp        string `json:"isp"`
	Isp_id     string `json:"isp_id"`
	Ip         string `json:"ip"`
}

type AreaBean struct {
	Code int32     `json:"code"`
	Data *AreaData `json:"data"`
}

func Json2AreaBean(buffer []byte) (AreaBean, bool) {
	// buffer := []byte(jsonString)
	var bean AreaBean
	if err := json.Unmarshal(buffer, bean); err != nil {
		beego.BeeLogger.Error("Json2AreaBean() => Json(AreaBean)解析失败,错误信息:%s", err.Error())
		return bean, false
	}
	return bean, true
}

func AreaBean2Json(bean *AreaBean) (string, bool) {
	if buffer, err := json.Marshal(bean); err != nil {
		beego.BeeLogger.Error("AreaBean2Json(bean *AreaBean) => Json(AreaBean)解析失败,错误信息:%s", err.Error())
		return "", false
	} else {
		return string(buffer), true
	}
}
