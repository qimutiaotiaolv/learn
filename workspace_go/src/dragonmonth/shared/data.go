package shared

import (
	"dragonmonth/mq"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

var SharedData *SharedDataStruct

type SharedDataStruct struct {
	MessageQueueProducer *mq.Producer
	MessageQueueConsumer *mq.Consumer
	SharedSessionLong    *session.Manager
}

func init() {
	if SharedData == nil {
		SharedData = new(SharedDataStruct)
		var err error
		SharedData.SharedSessionLong, err = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 2, "providerConfig": ""}`)
		if err != nil {
			beego.BeeLogger.Error("Shared::init()->session连接失败:%s", err.Error())
		} else {
			go SharedData.SharedSessionLong.GC()
		}

	}
}
