package middleWare

import (
	"gfEasy/app/model/ContextModel"
	"gfEasy/app/service/BaseSysLogService"
	"gfEasy/app/service/ContextService"
	"gfEasy/library/utils"
	"gfEasy/library/utils/cache"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"net/http"

	"gfEasy/library/utils/jwt"
)

//跨域处理中间件
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// ctx解析token绑定上下文
func CtxAuth(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &ContextModel.Context{
		Data: make(g.Map),
	}
	ContextService.S.Init(r, customCtx)
	user := new(ContextModel.CtxUser)

	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		r.Response.WriteStatusExit(http.StatusUnauthorized)
	}
	data, err := jwt.ParseToken(tokenStr, []byte(g.Cfg().GetString("jwt.sign", "gfEasy")))
	if err != nil {
		g.Log().Error(err.Error())
		r.Response.WriteStatusExit(http.StatusUnauthorized, "token不正确")
	}
	err = gconv.Struct(data, &user)
	if err != nil {
		g.Log().Error(err.Error())
		r.Response.WriteStatusExit(http.StatusUnauthorized)
	}
	//保存到上下文
	if user != nil {
		customCtx.User = &ContextModel.CtxUser{
			AdminUserId:     user.AdminUserId,
			UserName:        user.UserName,
			RoleIds:         user.RoleIds,
			PasswordVersion: user.PasswordVersion,
			Exp:             user.Exp,
			IsAdmin:         1,
			IsRefresh:       user.IsRefresh,
			Avatar:          user.Avatar,
		}
	}

	//判断token跟redis的缓存的token是否一样
	adminAdminUserIdStr := gconv.String(user.AdminUserId)
	cacheToken, err := cache.GCacheRedis.GetVar("admin:" + adminAdminUserIdStr + ":token")

	if cacheToken.IsEmpty() {
		r.Response.WriteStatusExit(http.StatusUnauthorized, "token已过期")
	}
	if tokenStr != cacheToken.String() {
		r.Response.WriteStatusExit(http.StatusUnauthorized, "跟最新颁发token不对")
	}

	// token密码版本
	cachePV, err := cache.GCacheRedis.GetVar("admin:" + adminAdminUserIdStr + ":passwordVersion")
	if user.PasswordVersion != cachePV.String() {
		r.Response.WriteStatusExit(http.StatusUnauthorized, "token版本不正确")
	}
	//admin用户不需要鉴权，拥有所有权限
	if user.AdminUserId == 1 {
		r.Middleware.Next()
		return
	}
	//判断权限
	cachePerms, err := cache.GCacheRedis.GetVar("admin:" + adminAdminUserIdStr + ":perms")
	if err != nil || cachePerms.IsEmpty() {
		r.Response.WriteStatusExit(http.StatusUnauthorized)
	}
	uri := r.Request.URL.Path
	check := utils.CheckPerms(cachePerms.String(), uri)
	if !check {
		r.Response.WriteStatusExit(http.StatusForbidden)
	}

	// 将自定义的上下文对象传递到模板变量中使用
	//r.Assigns(g.Map{
	//	"ContextService": customCtx,
	//})
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

//要在ctx后使用,记录操作者日志
func Log(r *ghttp.Request) {
	BaseSysLogService.S.Record(r)

	r.Middleware.Next()
}

// 权限判断中间件
//func Auth(r *ghttp.Request) {
//	userInfo := ContextService.S.Get(r.GetCtx()).User
//	g.Dump(userInfo)
//	r.Middleware.Next()
//}
