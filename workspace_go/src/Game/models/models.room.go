package models

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo/bson"
)

type RoomType int32

const (
	ROOMTYPE_SAMPLE = iota //普通房间
	ROOMTYPE_LAN           //局域网房间
	ROOMTYPE_OTHER         //其他房间
)

type RoomModel struct {
	RoomId            string `json:"room_id"`
	Ip                string
	Type              RoomType `json:"room_type"`
	GameId            string   `json:"game_id"`
	CountryId         string   `json:"country_id"`
	ArearId           string   `json:"arear_id"`
	RegionId          string   `json:"region_id"`
	CityId            string   `json:"city_id"`
	IspId             string   `json:"isp_id"`
	IpRegion          string   `json:"ipregion"`
	MaxPlayerCount    int32    `json:"max_playercount"`
	CurPlayerCount    int32    `json:"cur_playercount"`
	Longitude         float32  //经度
	Latitude          float32  //纬度
	DeviceInfo        string
	PlayerTockenArray []string
	State             int32 //0: 房间已满开始游戏 1: 房间未满等待玩家
}

func NewRoomModel() *RoomModel {
	maxCount, err := beego.AppConfig.Int("game::horse_room_max")
	if err != nil {
		beego.BeeLogger.Error("NewRoomModel() *Room => 配置文件解析失败:%s", err.Error())
		return nil
	}
	model := &RoomModel{
		RoomId:            bson.NewObjectId().Hex(),
		CurPlayerCount:    0,
		MaxPlayerCount:    int32(maxCount),
		PlayerTockenArray: make([]string, 0, maxCount),
	}
	return model
}
