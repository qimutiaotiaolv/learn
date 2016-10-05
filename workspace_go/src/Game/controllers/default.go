// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package controllers

import (
	// "encoding/json"
	"Game/beans"
	"github.com/astaxie/beego"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"net/http"
	// "samples/WebIM/models"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	beego.Controller
}

// Get method handles GET requests for WebSocketController.
func (this *WebSocketController) Get() {
	// // Safe check.
	// uname := this.GetString("uname")
	// if len(uname) == 0 {
	// 	this.Redirect("/", 302)
	// 	return
	// }

	// this.TplName = "websocket.html"
	// this.Data["IsWebSocket"] = true
	// this.Data["UserName"] = uname
	this.Ctx.WriteString("Game通讯框架测试")
}

// Join method handles WebSocket requests for WebSocketController.
func (this *WebSocketController) Join() {
	tocken := this.GetString("tocken")
	if len(tocken) == 0 {
		beego.Error("获取Tocken:", tocken)
		this.Redirect("/", 302)
		return
	}

	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		beego.Error("ot a websocket handshak:", err)
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	// Join chat room.
	// Join(uname, ws)
	// defer Leave(uname)

	// Message receive loop.
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		bean := &beans.SendMessage{}
		err = proto.Unmarshal(p, bean)
		if err != nil {
			beego.BeeLogger.Error("proto解析失败%s", err.Error())
			beego.BeeLogger.Error("原始字符串", string(p))
			return
		}
		// model := models.Event{
		// 	Type:      models.EventType(bean.GetType()),
		// 	User:      bean.GetUser(),
		// 	Timestamp: int(bean.GetTimestamp()),
		// 	Content:   bean.GetContent(),
		// }
		devBean := bean.GetDeviceMessage()
		beego.BeeLogger.Error("angle_alpha:%.2f\nangle_beta:%.2f\nangle_gamma:%.2f\nacce_x:%.2f\nacce_y:%.2f\nacce_z:%.2f\nacce_alpha:%.2f\nacce_beta:%.2f\nacce_gamma:%.2f\n",
			devBean.GetAngleAlpha(),
			devBean.GetAngleBeta(),
			devBean.GetAngleGamma(),
			devBean.GetAcceX(),
			devBean.GetAcceY(),
			devBean.GetAcceZ(),
			devBean.GetAcceAlpha(),
			devBean.GetAcceBeta(),
			devBean.GetAcceGamma())
		// publish <- model
		// publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}
}

// broadcastWebSocket broadcasts messages to WebSocket users.
func broadcastWebSocket() {
	// bean := &models.EventBean{
	// 	Type:      proto.Int32(int32(event.Type)),
	// 	User:      proto.String(event.User),
	// 	Timestamp: proto.Int64(int64(event.Timestamp)),
	// 	Content:   proto.String(event.Content),
	// }
	// data, err := proto.Marshal(bean)
	// // data, err := json.Marshal(event)
	// if err != nil {
	// 	beego.Error("Fail to marshal event:", err)
	// 	return
	// }

	// for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
	// 	// Immediately send event to WebSocket users.
	// 	ws := sub.Value.(Subscriber).Conn
	// 	if ws != nil {
	// 		if ws.WriteMessage(websocket.BinaryMessage, data) != nil {
	// 			// User disconnected.
	// 			unsubscribe <- sub.Value.(Subscriber).Name
	// 		}
	// 	}
	// }
}
