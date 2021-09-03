// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Ws

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

var WebsocketInfo = new(websocketInfo)

//控制器
type websocketInfo struct{}

// @Summary websocket demo
// @Description websocket demo
// @Tags Websocket-基本使用
// @Router /ws [post]
// @Security
func (c *websocketInfo) Index(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		glog.Error(err)
		r.Exit()
	}
	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			return
		}
		if err = ws.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}
