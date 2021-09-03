package router

import (
	"goEasy/app/controller/Admin"
	"goEasy/middleWare"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.SetNameToUriType(ghttp.URI_TYPE_CAMEL)

	s.SetServerRoot(g.Cfg().GetString("server.uploadServerRoot", "public"))
	// 分组路由注册方式
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Middleware(middleWare.CtxAuth) //后台上下文绑定&权限验证
		group.Middleware(middleWare.Log)     //后台日志记录
		/**insert_router_start**/
		group.ALL("/baseComm", Admin.BaseComm)

		group.ALL("/appGoodsInfo", Admin.AppGoodsInfo)

		group.ALL("/baseAppSpaceInfo", Admin.BaseAppSpaceInfo)

		group.ALL("/baseAppSpaceType", Admin.BaseAppSpaceType)

		group.ALL("/baseSysConf", Admin.BaseSysConf)

		group.ALL("/baseSysDepartment", Admin.BaseSysDepartment)

		group.ALL("/baseSysLog", Admin.BaseSysLog)

		group.ALL("/baseSysMenu", Admin.BaseSysMenu)

		group.ALL("/baseSysParam", Admin.BaseSysParam)

		group.ALL("/baseSysRole", Admin.BaseSysRole)

		group.ALL("/baseSysRoleDepartment", Admin.BaseSysRoleDepartment)

		group.ALL("/baseSysRoleMenu", Admin.BaseSysRoleMenu)

		group.ALL("/baseSysUser", Admin.BaseSysUser)

		group.ALL("/baseSysUserRole", Admin.BaseSysUserRole)

		group.ALL("/taskInfo", Admin.TaskInfo)

		group.ALL("/taskLog", Admin.TaskLog)

		group.ALL("/demoGo", Admin.DemoGo)

		group.ALL("/genCodeLog", Admin.GenCodeLog)

		group.ALL("/genCodeConfig", Admin.GenCodeConfig)
		/**insert_router_end**/

	})

}
