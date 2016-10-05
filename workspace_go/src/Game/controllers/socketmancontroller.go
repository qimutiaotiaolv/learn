package controllers

import (
	"Game/beans"
	// "Game/models"
	"github.com/astaxie/beego"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"net/http"
)

type SocketmanController struct {
	beego.Controller
}

func (this *SocketmanController) WebSocketJoin() {
	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		beego.BeeLogger.Error("SocketmanController::WebSocketJoin(),Not a websocket handshake:%s", err.Error())
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.BeeLogger.Error("SocketmanController::WebSocketJoin(),Cannot setup WebSocket connection:%s", err.Error())
		return
	}

	for {
		_, buffer, err := ws.ReadMessage()
		if err != nil {
			beego.BeeLogger.Error("SocketmanController::WebSocketJoin(),读取消息失败:%s", err.Error())
			continue
		}
		bean := &beans.ClientRequestBean{}
		err = proto.Unmarshal(buffer, bean)
		if err != nil {
			beego.BeeLogger.Error("SocketmanController::WebSocketJoin(),protobuf解析失败:%s", err.Error())
			continue
		}
		requestId := bean.GetRequestId()
		switch bean.GetOptionCode() {
		case beans.RequestOperationCode_REQUEST_OPERATIONCODE_CREATEROOM: //创建房间
			createRoomBean := bean.GetCreateroomBean()
			CreateRoomHandler(ws, createRoomBean, requestId)
			break
		case beans.RequestOperationCode_REQUEST_OPERATIONCODE_JOINROOM: //加入房间
			break
		case beans.RequestOperationCode_REQUEST_OPERATIONCODE_PLAYERDEVICCEBEAN: //玩家陀螺仪信息
			break
		}
	}
}

func BroadcastWebSocket() {

}
