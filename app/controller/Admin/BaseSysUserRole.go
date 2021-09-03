// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Admin

import (
	"goEasy/app/model/BaseSysUserRoleModel"
	"goEasy/app/service/BaseSysUserRoleService"
	"goEasy/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var BaseSysUserRole = new(baseSysUserRole)

//控制器
type baseSysUserRole struct{}

// @Summary 添加信息
// @Description 添加信息
// @Tags Admin-BaseSysUserRole-Admin用户角色关联表管理
// @Param data body BaseSysUserRoleModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysUserRole/add [post]
// @Security
func (c *baseSysUserRole) Add(r *ghttp.Request) {
	var req *BaseSysUserRoleModel.AddReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseSysUserRoleService.S.Add(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "新增成功", g.Map{"id": id})
}

// @Summary 修改信息
// @Description 修改信息
// @Tags Admin-BaseSysUserRole-Admin用户角色关联表管理
// @Param data body BaseSysUserRoleModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysUserRole/update [post]
// @Security
func (c *baseSysUserRole) Update(r *ghttp.Request) {
	var req *BaseSysUserRoleModel.UpdateReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseSysUserRoleService.S.Update(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"id": id})
}

// @Summary 删除信息
// @Description 删除信息
// @Tags Admin-BaseSysUserRole-Admin用户角色关联表管理
// @Param ids body BaseSysUserRoleModel.DeleteReqParams  true "ids[1,2,3..]"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysUserRole/delete [post]
// @Security
func (c *baseSysUserRole) Delete(r *ghttp.Request) {
	var req *BaseSysUserRoleModel.DeleteReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := BaseSysUserRoleService.S.Delete(req.Ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

// @Summary 单个信息
// @Description 单个信息
// @Tags Admin-BaseSysUserRole-Admin用户角色关联表管理
// @Param data body BaseSysUserRoleModel.InfoReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysUserRole/info [get]
// @Security
func (c *baseSysUserRole) Info(r *ghttp.Request) {
	var req *BaseSysUserRoleModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	result, err := BaseSysUserRoleService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 信息列表
// @Description 信息列表
// @Tags Admin-BaseSysUserRole-Admin用户角色关联表管理
// @Param data body BaseSysUserRoleModel.ListReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysUserRole/list [get]
// @Security
func (c *baseSysUserRole) List(r *ghttp.Request) {
	var req *BaseSysUserRoleModel.ListReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	condition := g.Map{}

	result, err := BaseSysUserRoleService.S.List(condition)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取列表信息成功", result)
}

// @Summary 分页信息
// @Description 分页信息
// @Tags Admin-BaseSysUserRole-Admin用户角色关联表管理
// @Param data body BaseSysUserRoleModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysUserRole/page [get]
// @Security
func (c *baseSysUserRole) Page(r *ghttp.Request) {
	var req *BaseSysUserRoleModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := BaseSysUserRoleService.S.Page(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	pagination := g.Map{
		"page":  page,
		"size":  size,
		"total": total,
	}
	response.SusJson(true, r, "获取分页信息成功", g.Map{
		"total":      total,
		"list":       list,
		"pagination": pagination,
	})
}
