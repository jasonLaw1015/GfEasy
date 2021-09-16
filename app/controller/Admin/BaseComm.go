// ==========================================================================
// 生成日期：2021-08-09 16:12:06
// 生成人：JasonLaw
// ==========================================================================
package Admin

import (
	"gfEasy/app/model/BaseCommModel"
	"gfEasy/app/model/BaseSysUserModel"
	"gfEasy/app/service/BaseSysRoleService"
	"gfEasy/app/service/BaseSysUserService"
	"gfEasy/app/service/ContextService"
	"gfEasy/library/response"
	"gfEasy/library/utils/cache"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//通用接口 一般写不需要权限过滤的接口
var BaseComm = new(baseComm)

//控制器
type baseComm struct{}

// @Summary 修改个人信息
// @Description 修改个人信息
// @Tags Admin-BaseComm-公共接口-需要token
// @Param data body BaseCommModel.PersonUpdateParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseComm/personUpdate [post]
// @Security
func (c *baseComm) PersonUpdate(r *ghttp.Request) {
	var req *BaseCommModel.PersonUpdateParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	adminUser := ContextService.S.Get(r.GetCtx()).User
	var userEntity *BaseSysUserModel.Entity
	err := BaseSysUserModel.M.Where("id=?", adminUser.AdminUserId).Scan(&userEntity)

	if err != nil {
		response.FailJson(true, r, err.Error())
	}

	if userEntity.Password != gmd5.MustEncryptString(req.NowPassword+userEntity.Salt) {
		response.FailJson(true, r, "当前密码不正确")
	}
	data := g.Map{
		"NickName":  req.NickName,
		"HeadImg":   req.HeadImg,
		"passwordV": userEntity.PasswordV + 1,
		"Password":  gmd5.MustEncryptString(req.NewPassword + userEntity.Salt),
	}
	_, err = BaseSysUserModel.M.OmitEmpty().Data(data).Where("id=?", adminUser.AdminUserId).Update()

	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", nil)
}

// @Summary 管理员登出
// @Description 管理员登出
// @Tags Admin-BaseComm-公共接口-需要token
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseComm/logout [post]
// @Security
func (c *baseComm) Logout(r *ghttp.Request) {
	adminUser := ContextService.S.Get(r.GetCtx()).User
	//缓存cache数据
	adminUserStr := gconv.String(adminUser.AdminUserId)
	cache.GCacheRedis.Remove("admin:" + adminUserStr + ":passwordVersion")
	cache.GCacheRedis.Remove("admin:" + adminUserStr + ":token")
	cache.GCacheRedis.Remove("admin:" + adminUserStr + ":perms")
	cache.GCacheRedis.Remove("admin:" + adminUserStr + ":token:refresh")
	response.SusJson(true, r, "更新成功", adminUser.AdminUserId)
}

// @Summary 管理员个人信息
// @Description 管理员个人信息
// @Tags Admin-BaseComm-公共接口-需要token
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseComm/person [get]
// @SecurityService.BaseSysRole
func (c *baseComm) Person(r *ghttp.Request) {
	//var req *BaseCommModel.InfoReqParams
	////获取参数&参数校验
	//if err := r.Parse(&req); err != nil {
	//	response.FailJson(true, r, err.Error())
	//}
	adminUser := ContextService.S.Get(r.GetCtx()).User

	result, err := BaseSysUserService.S.Info(adminUser.AdminUserId)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 权限菜单
// @Description 权限菜单
// @Tags Admin-BaseComm-公共接口-需要token
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseComm/permmenu [get]
// @Security
func (c *baseComm) Permmenu(r *ghttp.Request) {
	//var req *BaseCommModel.InfoReqParams
	////获取参数&参数校验
	//if err := r.Parse(&req); err != nil {
	//	response.FailJson(true, r, err.Error())
	//}
	adminUser := ContextService.S.Get(r.GetCtx()).User

	perms, err := BaseSysRoleService.S.GetPerms(adminUser.RoleIds)
	menus, err := BaseSysRoleService.S.GetMenus(adminUser.RoleIds)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", g.Map{
		"perms": perms,
		"menus": menus,
	})
}

// @Summary 文件上传模式，本地或者云存储
// @Description 文件上传模式，本地或者云存储
// @Tags Admin-BaseComm-公共接口-需要token
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseComm/UploadMode [get]
// @Security
func (c *baseComm) UploadMode(r *ghttp.Request) {
	mode := g.Map{
		"LOCAL": "local", //本地存储
		"CLOUD": "cloud", //oss存储
		"OTHER": "other", //其他
	}
	response.SusJson(true, r, "成功", g.Map{
		"mode": mode["LOCAL"],
	})
}

// @Summary 文件上传，本地存储
// @Description 文件上传，本地存储
// @Tags Admin-BaseComm-公共接口-需要token
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseComm/Upload [post]
// @Security
func (c *baseComm) Upload(r *ghttp.Request) {
	files := r.GetUploadFiles("file")
	names, err := files.Save(g.Cfg().GetString("server.uploadServerRoot"), true)
	if err != nil {
		r.Response.WriteExit(err)
	}

	response.SusJson(true, r, "成功", g.Cfg().GetString("server.uploadPath")+names[0])
}
