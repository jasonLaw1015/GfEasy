// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Admin

import (
	"goEasy/app/model/BaseSysRoleModel"
	"goEasy/app/service/BaseSysRoleService"
	"goEasy/app/service/ContextService"
	"goEasy/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var BaseSysRole = new(baseSysRole)

//控制器
type baseSysRole struct{}

// @Summary 添加信息
// @Description 添加信息
// @Tags Admin-BaseSysRole-admin角色管理
// @Param data body BaseSysRoleModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysRole/add [post]
// @Security
func (c *baseSysRole) Add(r *ghttp.Request) {
	var req *BaseSysRoleModel.AddReqParams
	adminUser := ContextService.S.Get(r.GetCtx()).User
	req.AdminUserId = adminUser.AdminUserId
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	id, err := BaseSysRoleService.S.Add(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "新增成功", g.Map{"id": id})
}

// @Summary 修改信息
// @Description 修改信息
// @Tags Admin-BaseSysRole-admin角色管理
// @Param data body BaseSysRoleModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysRole/update [post]
// @Security
func (c *baseSysRole) Update(r *ghttp.Request) {
	var req *BaseSysRoleModel.UpdateReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	adminUser := ContextService.S.Get(r.GetCtx()).User
	req.AdminUserId = adminUser.AdminUserId
	id, err := BaseSysRoleService.S.Update(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"id": id})
}

// @Summary 删除信息
// @Description 删除信息
// @Tags Admin-BaseSysRole-admin角色管理
// @Param ids body BaseSysRoleModel.DeleteReqParams  true "ids[1,2,3..]"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysRole/delete [post]
// @Security
func (c *baseSysRole) Delete(r *ghttp.Request) {
	var req *BaseSysRoleModel.DeleteReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := BaseSysRoleService.S.Delete(req.Ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

// @Summary 单个信息
// @Description 单个信息
// @Tags Admin-BaseSysRole-admin角色管理
// @Param data body BaseSysRoleModel.InfoReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysRole/info [get]
// @Security
func (c *baseSysRole) Info(r *ghttp.Request) {
	var req *BaseSysRoleModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	result, err := BaseSysRoleService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 信息列表
// @Description 信息列表
// @Tags Admin-BaseSysRole-admin角色管理
// @Param data body BaseSysRoleModel.ListReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysRole/list [get]
// @Security
func (c *baseSysRole) List(r *ghttp.Request) {
	var req *BaseSysRoleModel.ListReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	condition := g.Map{}

	if req.Name != "" {
		condition["name = ?"] = req.Name
	}

	result, err := BaseSysRoleService.S.List(condition)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取列表信息成功", result)
}

// @Summary 分页信息
// @Description 分页信息
// @Tags Admin-BaseSysRole-admin角色管理
// @Param data body BaseSysRoleModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysRole/page [get]
// @Security
func (c *baseSysRole) Page(r *ghttp.Request) {
	var req *BaseSysRoleModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := BaseSysRoleService.S.Page(req)
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
