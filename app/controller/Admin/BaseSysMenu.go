// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Admin

import (
	"goEasy/app/model/BaseSysMenuModel"
	"goEasy/app/service/BaseSysMenuService"
	"goEasy/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var BaseSysMenu = new(baseSysMenu)

//控制器
type baseSysMenu struct{}

// @Summary 添加信息
// @Description 添加信息
// @Tags Admin-BaseSysMenu-admin菜单管理
// @Param data body BaseSysMenuModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysMenu/add [post]
// @Security
func (c *baseSysMenu) Add(r *ghttp.Request) {
	var req *BaseSysMenuModel.AddReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseSysMenuService.S.Add(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "新增成功", g.Map{"id": id})
}

// @Summary 修改信息
// @Description 修改信息
// @Tags Admin-BaseSysMenu-admin菜单管理
// @Param data body BaseSysMenuModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysMenu/update [post]
// @Security
func (c *baseSysMenu) Update(r *ghttp.Request) {
	var req *BaseSysMenuModel.UpdateReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseSysMenuService.S.Update(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"id": id})
}

// @Summary 删除信息
// @Description 删除信息
// @Tags Admin-BaseSysMenu-admin菜单管理
// @Param ids body BaseSysMenuModel.DeleteReqParams  true "ids[1,2,3..]"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysMenu/delete [post]
// @Security
func (c *baseSysMenu) Delete(r *ghttp.Request) {
	var req *BaseSysMenuModel.DeleteReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := BaseSysMenuService.S.Delete(req.Ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

// @Summary 单个信息
// @Description 单个信息
// @Tags Admin-BaseSysMenu-admin菜单管理
// @Param data body BaseSysMenuModel.InfoReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysMenu/info [get]
// @Security
func (c *baseSysMenu) Info(r *ghttp.Request) {
	var req *BaseSysMenuModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	result, err := BaseSysMenuService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 信息列表
// @Description 信息列表
// @Tags Admin-BaseSysMenu-admin菜单管理
// @Param data body BaseSysMenuModel.ListReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysMenu/list [get]
// @Security
func (c *baseSysMenu) List(r *ghttp.Request) {
	var req *BaseSysMenuModel.ListReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	condition := g.Map{}

	if req.Name != "" {
		condition["name = ?"] = req.Name
	}

	result, err := BaseSysMenuService.S.List(condition)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取列表信息成功", result)
}

// @Summary 分页信息
// @Description 分页信息
// @Tags Admin-BaseSysMenu-admin菜单管理
// @Param data body BaseSysMenuModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysMenu/page [get]
// @Security
func (c *baseSysMenu) Page(r *ghttp.Request) {
	var req *BaseSysMenuModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := BaseSysMenuService.S.Page(req)
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
