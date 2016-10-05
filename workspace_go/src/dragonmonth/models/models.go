package models

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var SharedOrmEngine *xorm.Engine

func init() {
	driverName := beego.AppConfig.String("rdb::driver")
	dataSourceName := beego.AppConfig.String("rdb::source")
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		beego.BeeLogger.Error("xorm引擎初始化失败:%#v", err)
		return
	}
	SharedOrmEngine = engine
	runmode := beego.AppConfig.String("runmode")
	if runmode == "dev" {
		SharedOrmEngine.ShowSQL = true
		SharedOrmEngine.ShowDebug = true
		SharedOrmEngine.ShowErr = true
		SharedOrmEngine.ShowWarn = true
	}
	// engine内部支持连接池接口和对应的函数。
	// 如果需要设置连接池的空闲数大小，可以使用engine.SetMaxIdleConns()来实现。
	// 如果需要设置最大打开连接数，则可以使用engine.SetMaxOpenConns()来实现。
	maxConNumber, err := beego.AppConfig.Int("rdb::maxConNumber")
	if err != nil {
		beego.BeeLogger.Error("最大连接数解析错误:%s", err.Error())
	} else if maxConNumber <= 0 { //当小于等于0,则设置成最5000
		SharedOrmEngine.SetMaxOpenConns(5000)
	} else {
		SharedOrmEngine.SetMaxOpenConns(maxConNumber)
	}

	maxConPoolNumber, err := beego.AppConfig.Int("rdb::maxConPoolNumber")
	if err != nil {
		beego.BeeLogger.Error("最大空闲连接数解析错误:%s", err.Error())
	} else if maxConPoolNumber <= 0 { //当小于等于0,则设置成最5000
		SharedOrmEngine.SetMaxOpenConns(5000)
	} else {
		SharedOrmEngine.SetMaxOpenConns(maxConPoolNumber)
	}

	//初始化名称映射与前缀
	SharedOrmEngine.SetTableMapper(core.SnakeMapper{})  //table name 使用驼峰命名
	SharedOrmEngine.SetColumnMapper(core.SnakeMapper{}) //字段使用普通的
	//同步数据库
	SharedOrmEngine.Sync2(new(Order), new(OrderGoods), new(Payment), new(User))
}
