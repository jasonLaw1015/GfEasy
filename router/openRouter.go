package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"goEasy/app/controller/Admin"
)

func init() {
	s := g.Server()
	s.SetNameToUriType(ghttp.URI_TYPE_CAMEL)
	// 分组路由注册方式
	//此分组不需要token访问
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.ALL("/baseOpen", Admin.BaseOpen)
		//group.ALL("/baseComm", Admin.BaseComm)
		/**insert_router_end**/

	})

}
