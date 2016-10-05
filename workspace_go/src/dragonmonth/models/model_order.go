package models

import (
	// "github.com/astaxie/beego"
	// "github.com/go-xorm/xorm"
	"time"
)

/*
*orm到mysql
*
 */
type Order struct {
	OrderId          string
	UserId           string //对应的用户ID
	OrderGoodsIdList []string
	OrderState       int       //0：待付款 1：交易处理中 2：付款成功代发货 3：已发货 4：已收货待评价 5：已评价交易完成 6：已取消
	AllOriginaPrice  float32   //没有活动的原始价格 all OrderGoods.OriginalPrice sum
	AllCurrentPrice  float32   //实付总价  all OrderGoods.PayPrice sum
	CreatedAt        time.Time `xorm:"created"`
	UpdatedAt        time.Time `xorm:"updated"`
}

/*
*orm到mysql
*
 */
type OrderGoods struct {
	OrderGoodsId  string
	GoodsBaseId   string
	BuyCount      int       //购买数量
	OriginalPrice float32   //原始价格 BuyCount*GoodsBase.OriginalPrice
	PayPrice      float32   //需付价格 BuyCount*GoodsBase.CurrentPrice
	CreatedAt     time.Time `xorm:"created"`
	UpdatedAt     time.Time `xorm:"updated"`
}
