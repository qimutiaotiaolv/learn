package models

import (
	// "github.com/astaxie/beego"
	"time"
)

/*
*保存在mongoDB中
*
 */
type GoodsBase struct {
	GoodsId        string `xorm:"pk notnull index unique"`
	SpecId         string //规格ID，不同规格对应不同goodsID
	Name           string
	State2ImageFid map[string]string //图片：['0':'fid'] 图片使用分布式小文件系统，只保存fid,当用户请求时需要手动拼接
	State          int               //商品状态，0：正常，1：满送，2：满减，3：秒杀
	RuleId         string            //如果State!=0,则对应活动规则ID
	MarketPrice    float32           //市场价
	OriginalPrice  float32           //电商商城原价
	CurrentPrice   float32           //电商商城现价（没有优惠活动时,CurrentPrice=OriginalPrice,有优惠活动时，次价格=OriginalPrice-优惠）
	Stock          int               //库存
	Created        time.Time         `xorm:"created"`
	Updated        time.Time         `xorm:"updated"`
}
