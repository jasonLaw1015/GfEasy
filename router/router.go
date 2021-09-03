package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"goEasy/app/controller/Test"
)

func init() {
	s := g.Server()
	s.SetNameToUriType(ghttp.URI_TYPE_CAMEL)

	s.Group("/", func(group *ghttp.RouterGroup) {
		//group.Middleware(middleWare.CtxAuth) //后台上下文绑定&权限验证

		group.GET("/test1", Test.TestCaml.Test1)
		//group.POST("/handleMenus2Camel", Test.Tools.HandleMenus2Camel)
		group.ALL("/test", Test.Tools)
	})
}
