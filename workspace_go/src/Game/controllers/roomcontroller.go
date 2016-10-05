package controllers

import (
	"Game/beans"
	"Game/models"
	"Game/mq"
	"Game/utils"
	"container/list"
	"github.com/astaxie/beego"
	"strings"
)

var (
	useRoomSubscriberChan    = make(chan *RoomSubscriber, 10)
	createRoomSubscriberChan = make(chan *RoomSubscriber, 10)
	waitRoomList             = list.New()
	playingRoomList          = list.New()
)

type RoomSubscriber struct {
	Ws        *websocket.Conn
	RoomModel *models.RoomModel
}

func init() {
	go loop_room()
}

func loop_room() {
	for {
		select {
		case roomSuber <- createRoomSubscriberChan:
			playingRoomList.Remove(roomSuber) //正在游戏的队列删除
			waitRoomList.PushBack(roomSuber)
			newRoomCreateChan <- true
		case roomSuber <- useRoomSubscriberChan:
			waitRoomList.Remove(roomSuber)
			playingRoomList.PushBack(roomSuber)
			/**
			 * 此时，玩家已经记录到房间里，这里开始游戏逻辑
			 */
		}
	}
}

/**
 * 创建房间
 * @param {[type]} ws        *websocket.Conn       [创建房间的客户端的socket]
 * @param {[type]} bean      *beans.CreateRoomBean [protibuf中的Bean]
 * @param {[type]} requestId string                [protobuf中的请求Id]
 */
func CreateRoomHandler(ws *websocket.Conn, bean *beans.CreateRoomBean, requestId string) {
	areaBean, err := utils.HttpGetArea(bean.GetIp())
	if err != nil {
		beego.BeeLogger.Error("CreateRoomHandler(ws *websocket.Conn, bean *beans.CreateRoomBean, requestId string),获取设备地理位置失败:%s", err.Error())
		return
	}
	roomModel := models.NewRoomModel()
	roomModel.Ip = bean.GetIp()
	roomModel.Type = models.ROOMTYPE_SAMPLE
	roomModel.GameId = bean.GetGameId()
	roomModel.CountryId = areaBean.Data.County_id
	roomModel.ArearId = areaBean.Data.Area_id
	roomModel.RegionId = areaBean.Data.Region_id
	roomModel.CityId = areaBean.Data.City_id
	roomModel.IspId = areaBean.Data.Isp_id
	index := strings.LastIndex(bean.GetIp(), ".")
	roomModel.IpRegion = (roomModel.Ip)[0:index]
	roomModel.Latitude = bean.GetLatitude()
	roomModel.Longitude = bean.GetLongitude()
	roomModel.DeviceInfo = bean.GetDeviceInfo()
	beego.BeeLogger.Error("创建房间:%#v", roomModel)
	roomSuber := &RoomSubscriber{
		Ws:        ws,
		RoomModel: roomModel,
	}
	createRoomSubscriberChan <- roomSuber
	mq.Producer_PushMessage_CreateRoom(roomModel)
}
