package controllers

import (
	"Game/beans"
	"Game/models"
	"Game/mq"
	"Game/utils"
	"container/list"
	"github.com/astaxie/beego"
)

var (
	newRoomCreateChan           = make(bool, 10)
	addPlayerSubscriberChan     = make(chan *Player, 10)
	playingPlayerSubscriberChan = make(chan *Player, 10) //已经加入房间
	waitingPlayerSubscriberChan = make(chan *Player, 10) //等待加入房间
	LeavePlayerSubscriberChan   = make(chan *Player, 10)
	waitingPlayerList           = list.New()
	playingPlayerList           = list.New()
	tocken2playerMap            = make(map[string]*Player)
)

/**
 *陀螺仪信息
 */
// message PlayerDeviceBean{
//     optional string game_id = 10;
//     optional float angle_alpha = 1 [default = 0.0];
//     optional float angle_beta = 2 [default = 0.0];
//     optional float angle_gamma = 3 [default = 0.0];
//     optional float acce_x = 4 [default = 0.0];
//     optional float acce_y = 5 [default = 0.0];
//     optional float acce_z = 6 [default = 0.0];
//     optional float acce_alpha = 7 [default = 0.0];
//     optional float acce_beta = 8 [default = 0.0];
//     optional float acce_gamma = 9 [default = 0.0];
// }

/**
 *
 */
type Player struct {
	Tocken string
	Ws     *websocket.Conn
	RoomId string
	State  int32 //玩家状态 0:已经加入房间 1:等待加入房间 2:退出房间
	GameId string
}

/**
 * 用户新加入后调用
 * @param {[type]} player *Player [description]
 */
func addPlayer(player *Player) {
	tocken2playerMap[player.Tocken] = player
	switch player.State {
	case 0:
		playingPlayerList.PushBack(player)
		break
	case 1:
		waitingPlayerList.PushBack(player)
		break
	case 2:
		break
	}
}

func loop_player() {
	for {
		select {
		case _ <- newRoomCreateChan:
			update_player_list()
		case player <- addPlayerSubscriberChan:
			addPlayer(player)
		case player <- playingPlayerSubscriberChan:
		case player <- waitingPlayerSubscriberChan:
		case player <- LeavePlayerSubscriberChan:

		}
	}
}

func init() {
	loop_player()
}

/**
 * 选出一个合适的房间
 */
func roomFilter(bean *beans.JoinRoomBean) (*RoomSubscriber, bool) {
	if waitRoomList.Len() == 0 {
		return nil, false
	}
	var subRoom *RoomSubscriber
	isHas := false
	switch {
	case bean.GetGameId() == beego.AppConfig.String("game::horse_game_id"):
		for sub := waitRoomList.Back(); sub != nil; sub = sub.Prev() {
			roomModel := sub.Value(*models.RoomModel)
			if roomModel.CurPlayerCount < roomModel.MaxPlayerCount {
				subRoom = sub
				isHas = true
				break
			}
		}
		break
	}
	return subRoom, isHas
}

func update_player_list() {
	for player := waitingPlayerList.Back(); player != nil; player = player.Prev() {
		var subRoom *RoomSubscriber
		for sub := waitRoomList.Back(); sub != nil; sub = sub.Prev() {
			roomModel := sub.Value(*models.RoomModel)
			if roomModel.CurPlayerCount < roomModel.MaxPlayerCount {
				subRoom = sub
				break
			}
		}

	}
}

/**
 * 加入房间
 * @param {[type]} ws        *websocket.Conn     [description]
 * @param {[type]} bean      *beans.JoinRoomBean [description]
 * @param {[type]} requestId string              [description]
 */
func JoinRoomHandler(ws *websocket.Conn, bean *beans.JoinRoomBean, requestId string) {
	player := &Player{
		Ws:     ws,
		Tocken: bean.GetTocken(),
		RoomId: "",
		State:  1,
		GameId: bean.GetGameId(),
	}
	addPlayer(player)
	// roomSub, isHas := roomFilter(bean)
	// //没有可用房间
	// if !isHas {
	// 	/**
	// 	 *没有可用的房间，放入等待队列
	// 	 */
	// 	player.State = 1
	// 	addPlayer(player)
	// 	return
	// }
	// roomModel := roomSub.Value(*models.RoomModel)
	// roomModel.PlayerTockenArray = append(roomModel.PlayerTockenArray, bean.GetTocken())
	// roomModel.CurPlayerCount = len(roomModel.PlayerTockenArray)
	// if roomModel.CurPlayerCount == roomModel.MaxPlayerCount { //玩家加入房间成功,且房间已满,开始游戏
	// 	roomModel.State = 0
	// 	useRoomSubscriberChan <- roomSub
	// } else {
	// 	roomModel.State = 1
	// 	/**
	// 	 * 玩家加入房间成功，等待其他玩家中
	// 	 */
	// }
	// player.State = 0
	// addPlayer(player)
}

/**
 * 当获取玩家陀螺仪信息后条用此方法
 * @param {[type]} ws        *websocket.Conn         [description]
 * @param {[type]} bean      *beans.PlayerDeviceBean [description]
 * @param {[type]} requeseId string                  [description]
 */
func GetPlayerDeviceGyroscopeHandler(ws *websocket.Conn, bean *beans.PlayerDeviceBean, requeseId string) {

}
