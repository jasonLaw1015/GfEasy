// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Api

import (
	"goEasy/app/model/BaseSysRoleMenuModel"
	"goEasy/app/service/BaseSysRoleMenuService"
	"goEasy/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var BaseSysRoleMenu = new(baseSysRoleMenu)

//控制器
type baseSysRoleMenu struct{}

// @Summary 添加信息
// @Description 添加信息
// @Tags Api-BaseSysRoleMenu-Admin角色菜单关联表管理
// @Param data body BaseSysRoleMenuModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysRoleMenu/add [post]
// @Security
func (c *baseSysRoleMenu) Add(r *ghttp.Request) {
	var req *BaseSysRoleMenuModel.AddReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseSysRoleMenuService.S.Add(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "新增成功", g.Map{"id": id})
}

// @Summary 修改信息
// @Description 修改信息
// @Tags Api-BaseSysRoleMenu-Admin角色菜单关联表管理
// @Param data body BaseSysRoleMenuModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysRoleMenu/update [post]
// @Security
func (c *baseSysRoleMenu) Update(r *ghttp.Request) {
	var req *BaseSysRoleMenuModel.UpdateReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseSysRoleMenuService.S.Update(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"id": id})
}

// @Summary 删除信息
// @Description 删除信息
// @Tags Api-BaseSysRoleMenu-Admin角色菜单关联表管理
// @Param ids body BaseSysRoleMenuModel.DeleteReqParams  true "ids[1,2,3..]"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysRoleMenu/delete [post]
// @Security
func (c *baseSysRoleMenu) Delete(r *ghttp.Request) {
	var req *BaseSysRoleMenuModel.DeleteReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := BaseSysRoleMenuService.S.Delete(req.Ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

// @Summary 单个信息
// @Description 单个信息
// @Tags Api-BaseSysRoleMenu-Admin角色菜单关联表管理
// @Param data body BaseSysRoleMenuModel.InfoReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysRoleMenu/info [get]
// @Security
func (c *baseSysRoleMenu) Info(r *ghttp.Request) {
	var req *BaseSysRoleMenuModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	result, err := BaseSysRoleMenuService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 信息列表
// @Description 信息列表
// @Tags Api-BaseSysRoleMenu-Admin角色菜单关联表管理
// @Param data body BaseSysRoleMenuModel.ListReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysRoleMenu/list [get]
// @Security
func (c *baseSysRoleMenu) List(r *ghttp.Request) {
	var req *BaseSysRoleMenuModel.ListReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	condition := g.Map{}

	result, err := BaseSysRoleMenuService.S.List(condition)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取列表信息成功", result)
}

// @Summary 信息分页列表
// @Description 信息分页列表
// @Tags Api-BaseSysRoleMenu-Admin角色菜单关联表管理
// @Param data body BaseSysRoleMenuModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysRoleMenu/page [get]
// @Security
func (c *baseSysRoleMenu) Page(r *ghttp.Request) {
	var req *BaseSysRoleMenuModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := BaseSysRoleMenuService.S.Page(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	pagination := g.Map{
		"page":  page,
		"size":  size,
		"total": total,
	}
	response.SusJson(true, r, "获取列表信息成功", g.Map{
		"total":      total,
		"list":       list,
		"pagination": pagination,
	})
}
