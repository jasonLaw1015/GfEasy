package router

import (
	"gfEasy/app/controller/Api"
	"gfEasy/middleWare"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.SetNameToUriType(ghttp.URI_TYPE_CAMEL)
	// 分组路由注册方式
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(middleWare.CtxAuth) //后台上下文绑定&权限验证，需自己改造下这个中间件适配自己需求

		/**api_insert_router_start**/

		group.ALL("/appGoodsInfo", Api.AppGoodsInfo)

		group.ALL("/baseAppSpaceInfo", Api.BaseAppSpaceInfo)

		group.ALL("/baseAppSpaceType", Api.BaseAppSpaceType)

		group.ALL("/baseSysConf", Api.BaseSysConf)

		group.ALL("/baseSysDepartment", Api.BaseSysDepartment)

		group.ALL("/baseSysLog", Api.BaseSysLog)

		group.ALL("/baseSysMenu", Api.BaseSysMenu)

		group.ALL("/baseSysParam", Api.BaseSysParam)

		group.ALL("/baseSysRole", Api.BaseSysRole)

		group.ALL("/baseSysRoleDepartment", Api.BaseSysRoleDepartment)

		group.ALL("/baseSysRoleMenu", Api.BaseSysRoleMenu)

		group.ALL("/baseSysUser", Api.BaseSysUser)

		group.ALL("/baseSysUserRole", Api.BaseSysUserRole)

		group.ALL("/taskInfo", Api.TaskInfo)

		group.ALL("/taskLog", Api.TaskLog)

		group.ALL("/baseComm", Api.BaseComm)

		group.ALL("/baseOpen", Api.BaseOpen)

		group.ALL("/demoGo", Api.DemoGo)

		group.ALL("/genCodeLog", Api.GenCodeLog)

		group.ALL("/genCodeConfig", Api.GenCodeConfig)
		/**api_insert_router_end**/

	})
}
