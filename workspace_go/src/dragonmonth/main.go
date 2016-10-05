package main

import (
	_ "dragonmonth/models"
	"dragonmonth/mq"
	_ "dragonmonth/routers"
	"dragonmonth/shared"
	_ "dragonmonth/util"
	"encoding/json"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	"os"
	"os/exec"
	"time"
)

func main() {
	beego.Run()
}

func nsq_init() {
	nsqdrun_sig := make(chan byte)
	nsqlookuprun_sig := make(chan byte)
	//启动nsqlookup
	go func() {
		nsqlookup_path := beego.AppConfig.String("nsq::nsqlookup_path")
		nsqlookup_params_json := beego.AppConfig.String("nsq::nsqlookup_params")
		beego.BeeLogger.Info("nsqlookup_path:%s,nsqlookup_params_json:%s", nsqlookup_path, nsqlookup_params_json)
		var nsqlookup_params []string
		err := json.Unmarshal([]byte(nsqlookup_params_json), &nsqlookup_params)
		if err != nil {
			beego.BeeLogger.Error("nsqlookup::JSON解析错误:%#v", err)
			return
		}
		cmd := exec.Command(nsqlookup_path, nsqlookup_params...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Start()
		if err != nil {
			beego.BeeLogger.Error("nsqlookup::cmd.Start()错误:%#v", err)
			return
		}
		time.Sleep(3 * time.Second)
		nsqlookuprun_sig <- '0'
		err = cmd.Wait()
		if err != nil {
			beego.BeeLogger.Error("nsqlookup::cmd.Wait()错误:%#v", err)
			return
		}
	}()
	//启动nsqd
	go func() {
		<-nsqlookuprun_sig
		beego.BeeLogger.Info("开始启动nsqd")
		nsq_path := beego.AppConfig.String("nsq::nsqd_path")
		nsq_params_json := beego.AppConfig.String("nsq::nsqd_params")
		beego.BeeLogger.Info("nsq_path:%s,nsq_params_json:%s", nsq_path, nsq_params_json)
		var nsq_params []string
		err := json.Unmarshal([]byte(nsq_params_json), &nsq_params)
		if err != nil {
			beego.BeeLogger.Error("nsqd::JSON解析错误:%#v", err)
			return
		}
		cmd := exec.Command(nsq_path, nsq_params...)
		err = cmd.Start()
		if err != nil {
			beego.BeeLogger.Error("nsqd::cmd.Start()错误:%#v", err)
			return
		}
		time.Sleep(3 * time.Second)
		nsqdrun_sig <- '0'
		err = cmd.Wait()
		if err != nil {
			beego.BeeLogger.Error("nsqd::cmd.Wait()错误:%#v", err)
			return
		}
	}()
	//启动nsqadmin
	go func() {
		<-nsqdrun_sig
		beego.BeeLogger.Info("开始启动nsqadmin")
		nsqadmin_path := beego.AppConfig.String("nsq::nsqadmin_path")
		nsqadmin_params_json := beego.AppConfig.String("nsq::nsqadmin_params")
		beego.BeeLogger.Info("nsqadmin_path:%s,nsqadmin_params_json:%s", nsqadmin_path, nsqadmin_params_json)
		var nsqadmin_params []string
		err := json.Unmarshal([]byte(nsqadmin_params_json), &nsqadmin_params)
		if err != nil {
			beego.BeeLogger.Error("nsqadmin::JSON解析错误:%#v", err)
			return
		}
		cmd := exec.Command(nsqadmin_path, nsqadmin_params...)
		err = cmd.Start()
		if err != nil {
			beego.BeeLogger.Error("nsqadmin::cmd.Start()错误:%#v", err)
			return
		}
		err = cmd.Wait()
		if err != nil {
			beego.BeeLogger.Error("nsqadmin::cmd.Wait()错误:%#v", err)
			return
		}
	}()
}

func shared_init() {
	shared.SharedData.MessageQueueConsumer = mq.SharedConsumer
	shared.SharedData.MessageQueueProducer = mq.SharedProducer
}

func appConfig_init() {
	//是否开启进程内监控模块，默认 false 关闭。
	// beego.BConfig.Listen.AdminEnable = true
	//监控程序监听的地址，默认值是 8088
	beego.BConfig.Listen.AdminPort = 8088
	// 是否在日志里面显示文件名和输出日志行号，默认 true。此参数不支持配置文件配置。
	beego.BConfig.Log.FileLineNum = true
}

func redis_init() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600
}

func init() {
	// go func() {
	// 	nsq_init()
	// }()
	appConfig_init()
	shared_init()
	redis_init()
}
