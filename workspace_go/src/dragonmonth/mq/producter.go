package mq

import (
	"dragonmonth/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/nsqio/go-nsq"
)

var SharedProducer *Producer

type Producer struct {
	_producer *nsq.Producer
}

func NewProducer() *Producer {
	p := &Producer{}
	p.init()
	return p
}

func (this *Producer) init() {
	addr := beego.AppConfig.String("nsq::nsqd_address_tcp")
	p, err := nsq.NewProducer(addr, nsq.NewConfig())
	if err != nil {
		beego.BeeLogger.Error("Producer::init() 连接nsqd失败:%s", err.Error())
		return
	}
	this._producer = p
	beego.BeeLogger.Info("Producer::init() 连接nsqd成功!")
}

func (this *Producer) SendMessage(topic string, message []byte) error {
	return this._producer.Publish(topic, message)
}

func (this *Producer) Colse() {
	this._producer.Stop()
}

func init() {
	if SharedProducer == nil {
		SharedProducer = NewProducer()
	}
}

func Producer_MessageSender_ImageWaitConver2Webp(messageModel *MessageImageWaitConver2Webp) {
	topic := beego.AppConfig.String("nsq-topic::waitconver2webp")
	messageBytes, err := json.Marshal(messageModel)
	if err != nil {
		beego.BeeLogger.Error("Producer_MessageSender_ImageWaitConver2Webp::json解析错误:%s", err.Error())
		return
	}
	err = SharedProducer.SendMessage(topic, messageBytes)
	if err != nil {
		beego.BeeLogger.Error("Producer_MessageSender_ImageWaitConver2Webp::写入消息队列失败:%s", err.Error())
		return
	}
}

func Producer_MessageSender_ImageWait2SendWeedfs(messageModel *MessageImageWaitConver2Webp) {
	topic := beego.AppConfig.String("nsq-topic::sendimage2nfs")
	messageBytes, err := json.Marshal(messageModel)
	if err != nil {
		beego.BeeLogger.Error("Producer_MessageSender_ImageWait2SendWeedfs::json解析错误:%s", err.Error())
		return
	}
	err = SharedProducer.SendMessage(topic, messageBytes)
	if err != nil {
		beego.BeeLogger.Error("Producer_MessageSender_ImageWait2SendWeedfs::写入消息队列失败:%s", err.Error())
		return
	}

}

func Producer_MessageSender_Register_Insert2Database(messageModel *models.MessageInsertUser) {
	topic := beego.AppConfig.String("nsq-topic::insertuserinfo")
	messageBytes, err := json.Marshal(messageModel)
	beego.BeeLogger.Error("Producer_MessageSender_Register_Insert2Database::转化的JSON:%s", string(messageBytes))
	if err != nil {
		beego.BeeLogger.Error("Producer_MessageSender_Register_Insert2Database::json解析错误:%s", err.Error())
		return
	}
	err = SharedProducer.SendMessage(topic, messageBytes)
	if err != nil {
		beego.BeeLogger.Error("Producer_MessageSender_Register_Insert2Database::写入消息队列失败:%s", err.Error())
		return
	}
}
