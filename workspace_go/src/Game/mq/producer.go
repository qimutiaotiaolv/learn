package mq

import (
	"Game/messages"
	"Game/models"
	// "encoding/json"
	"github.com/astaxie/beego"
	"github.com/golang/protobuf/proto"
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

func init() {
	if SharedProducer == nil {
		SharedProducer = NewProducer()
	}
}

func (this *Producer) SendMessage(topic string, message []byte) error {
	return this._producer.Publish(topic, message)
}

// type RoomMessage struct {
// 	RoomId            *string  `protobuf:"bytes,1,opt,name=room_id" json:"room_id,omitempty"`
// 	RoomType          *int32   `protobuf:"varint,2,opt,name=room_type" json:"room_type,omitempty"`
// 	GameId            *string  `protobuf:"bytes,3,opt,name=game_id" json:"game_id,omitempty"`
// 	CountryId         *string  `protobuf:"bytes,4,opt,name=country_id" json:"country_id,omitempty"`
// 	ArearId           *string  `protobuf:"bytes,5,opt,name=arear_id" json:"arear_id,omitempty"`
// 	RegionId          *string  `protobuf:"bytes,6,opt,name=region_id" json:"region_id,omitempty"`
// 	CityId            *string  `protobuf:"bytes,7,opt,name=city_id" json:"city_id,omitempty"`
// 	IspId             *string  `protobuf:"bytes,8,opt,name=isp_id" json:"isp_id,omitempty"`
// 	IpRegion          *string  `protobuf:"bytes,9,opt,name=ip_region" json:"ip_region,omitempty"`
// 	MaxPlayercount    *int32   `protobuf:"varint,10,opt,name=max_playercount" json:"max_playercount,omitempty"`
// 	CurPlayercount    *int32   `protobuf:"varint,11,opt,name=cur_playercount" json:"cur_playercount,omitempty"`
// 	PlayerTockenArray []string `protobuf:"bytes,12,rep,name=player_tocken_array" json:"player_tocken_array,omitempty"`
// 	XXX_unrecognized  []byte   `json:"-"`
// }

func Producer_PushMessage_CreateRoom(roomModel *models.RoomModel) {
	roomType := int32(roomModel.Type)
	roomMsg := &messages.RoomMessage{
		RoomId:            proto.String(roomModel.RoomId),
		RoomType:          proto.Int32(roomType),
		GameId:            proto.String(roomModel.GameId),
		CountryId:         proto.String(roomModel.CountryId),
		ArearId:           proto.String(roomModel.ArearId),
		RegionId:          proto.String(roomModel.RegionId),
		CityId:            proto.String(roomModel.CityId),
		IspId:             proto.String(roomModel.IspId),
		IpRegion:          proto.String(roomModel.IpRegion),
		MaxPlayercount:    proto.Int32(roomModel.MaxPlayerCount),
		CurPlayercount:    proto.Int32(roomModel.CurPlayerCount),
		Longitude:         proto.Float32(roomModel.Longitude),
		Latitude:          proto.Float32(roomModel.Latitude),
		DeviceInfo:        proto.String(roomModel.DeviceInfo),
		PlayerTockenArray: roomModel.PlayerTockenArray,
	}
	data, err := proto.Marshal(roomMsg)
	if err != nil {
		beego.Error("Producer::Producer_PushMessage_CreateRoom() Protobuf解析失败:%s", err.Error())
		return
	}
	topic := beego.AppConfig.String("nsq-topic::message_createroom")
	err = SharedProducer.SendMessage(topic, data)
	if err != nil {
		beego.BeeLogger.Error("Producer::Producer_PushMessage_CreateRoom() 写入消息队列失败:%s", err.Error())
		return
	}
}

func Producer_PushMessage_DestoryRoom(roomId string) {

}

func Producer_PushMessage_UpdateRoom(roomMsg *messages.RoomMessage) {

}

func Producer_PushMessage_JoinRoom(playerTocken string) {

}
