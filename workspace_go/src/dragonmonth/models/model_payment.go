package models

import (
	// "github.com/astaxie/beego"
	// "github.com/go-xorm/xorm"
	"time"
)

/*
*orm到mysql
 */
type Payment struct {
	PayID    string
	OrderID  string
	UserID   string
	PayState int
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}
