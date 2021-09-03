package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"goEasy/app/controller/Ws"
)

func init() {
	s := g.Server()
	s.SetNameToUriType(ghttp.URI_TYPE_CAMEL)
	// 分组路由注册方式
	//此分组不需要token访问
	s.Group("/", func(group *ghttp.RouterGroup) {

		group.ALL("/ws", Ws.WebsocketInfo)

	})

	//s.BindHandler("/ws", Ws.WebsocketInfo.Index)
}
