package mq

import (
	"dragonmonth/models"
	"dragonmonth/util"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/ginuerzh/weedo"
	"github.com/nsqio/go-nsq"
	"os"
)

var SharedConsumer *Consumer

type Consumer struct {
	WeedClient *weedo.Client
}

func NewConsumer() *Consumer {
	weed_master := beego.AppConfig.String("weedfs::master")
	weed_filer := beego.AppConfig.String("weedfs::filer")
	consumer := &Consumer{
		WeedClient: weedo.NewClient(weed_master, weed_filer),
	}
	consumer.init()
	return consumer
}

func (this *Consumer) init() {
	RegisteConsumer_MessageHadler_ImageWait2SendWeedfs()
	RegisteConsumer_MessageHadler_ImageWaitConver2Webp()
	// RegisteConsumer_MessageHadler_ImageWait2SendWeedfs()
	// RegisteConsumer_MessageHadler_DatabaseWriter()
	RegisteConsumer_MessageHadler_Database_Insert_User()
}

func (this *Consumer) RegisteConsumer(topic, channel string, handlerFunc func(message *nsq.Message) error) (*nsq.Consumer, error) {
	consumer, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	// consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
	// 	return handlerFunc(message)
	// }))
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(func(message *nsq.Message) error {
		return handlerFunc(message)
	}), 4)
	addr := beego.AppConfig.String("nsq::nsqlookup_address_http")
	err = consumer.ConnectToNSQLookupd(addr)
	return consumer, err
}

func init() {
	if SharedConsumer == nil {
		SharedConsumer = NewConsumer()
	}
}

/*
*
*注册消费者，负责从消息队列里取出消息，并根据消息中的配置将图片转换成webp并保存至临时目录
*
*
 */
func RegisteConsumer_MessageHadler_ImageWaitConver2Webp() {
	topic := beego.AppConfig.String("nsq-topic::waitconver2webp")
	channel := beego.AppConfig.String("nsq-channel::waitconver2webp")
	_, err := SharedConsumer.RegisteConsumer(topic, channel, func(message *nsq.Message) error {
		var messageModel MessageImageWaitConver2Webp
		err := json.Unmarshal(message.Body, &messageModel)
		if err != nil {
			beego.BeeLogger.Error("RegisteConsumer_MessageHadler_ImageWaitConver2Webp::json解析错误:%s", err.Error())
			return err
		}
		config := util.NewImageConverConfig(messageModel.Config.WidthPix, messageModel.Config.HeightPix, messageModel.Config.Quality, messageModel.Config.Lossless, messageModel.OriginalImagePath, messageModel.SaveWebpTempPath)
		util.Encode2Webp(config, func(message *util.ImageConverConfig) {
			beego.BeeLogger.Info("转换完毕,线程ID:%d", os.Getpid())
			Producer_MessageSender_ImageWait2SendWeedfs(&messageModel)
		})
		return nil
	})
	if err != nil {
		beego.BeeLogger.Error("RegisteConsumer_MessageHadler_ImageWaitConver2Webp:注册consumer错误:%s", err.Error())
	}
}

/*
*
*注册消费者，负责从消息队列里取出消息，并根据消息中的配置将webp图片从零时目录中读取出并发送至weedfs
*
*
 */
func RegisteConsumer_MessageHadler_ImageWait2SendWeedfs() {
	topic := beego.AppConfig.String("nsq-topic::sendimage2nfs")
	channel := beego.AppConfig.String("nsq-channel::sendimage2nfs")
	_, err := SharedConsumer.RegisteConsumer(topic, channel, func(message *nsq.Message) error {
		var messageModel MessageImageWaitConver2Webp
		err := json.Unmarshal(message.Body, &messageModel)
		if err != nil {
			beego.BeeLogger.Error("RegisteConsumer_MessageHadler_ImageWaitConver2Webp::json解析错误:%s", err.Error())
			return err
		}
		filePath := messageModel.SaveWebpTempPath
		file, err := os.Open(filePath)
		if err != nil {
			beego.BeeLogger.Error("RegisteConsumer_MessageHadler_ImageWait2SendWeedfs::文件打开失败:%s", err.Error())
			return err
		}
		fid, size, err := SharedConsumer.WeedClient.AssignUpload("", "", file)
		if err != nil {
			beego.BeeLogger.Error("RegisteConsumer_MessageHadler_ImageWait2SendWeedfs::上传weedfs失败:%s", err.Error())
			return err
		}
		master := "http://" + beego.AppConfig.String("weedfs::master")
		beego.BeeLogger.Info("文件上传成功  fid:%s,size:%d,fileUrl:%s/%s", fid, size, master, fid)
		return nil
	})
	if err != nil {
		beego.BeeLogger.Error("RegisteConsumer_MessageHadler_ImageWaitConver2Webp:注册consumer错误:%s", err.Error())
	}
}

/**
 * mysql中插入用户
 *
 */
func RegisteConsumer_MessageHadler_Database_Insert_User() {
	topic := beego.AppConfig.String("nsq-topic::insertuserinfo")
	channel := beego.AppConfig.String("nsq-channel::insertuserinfo")
	_, err := SharedConsumer.RegisteConsumer(topic, channel, func(message *nsq.Message) error {
		var messageModel models.MessageInsertUser
		messageModel.Bean = new(models.User)
		err := json.Unmarshal(message.Body, &messageModel)
		if err != nil {
			beego.BeeLogger.Error("RegisteConsumer_MessageHadler_Database_Insert_User::json解析错误:%s", err.Error())
			return err
		}
		// beego.BeeLogger.Error("收到数据:%#s", messageModel)
		userModel, ok := messageModel.Bean.(*models.User)
		if !ok {
			return nil
		}
		beego.BeeLogger.Error("收到数据:%#v", userModel)
		_, err = models.SharedOrmEngine.InsertOne(messageModel.Bean)
		if err != nil {
			beego.BeeLogger.Error("RegisteConsumer_MessageHadler_Database_Insert_User::数据库插入错误:%s", err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		beego.BeeLogger.Error("RegisteConsumer_MessageHadler_Database_Insert_User:注册consumer错误:%s", err.Error())
	}
}
