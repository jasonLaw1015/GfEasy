// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Admin

import (
	"gfEasy/app/model/BaseSysRoleDepartmentModel"
	"gfEasy/app/service/BaseSysRoleDepartmentService"
	"gfEasy/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var BaseSysRoleDepartment = new(baseSysRoleDepartment)

//控制器
type baseSysRoleDepartment struct{}

// @Summary 添加信息
// @Description 添加信息
// @Tags Admin-BaseSysRoleDepartment-Admin角色部门关联表管理
// @Param data body BaseSysRoleDepartmentModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysRoleDepartment/add [post]
// @Security
func (c *baseSysRoleDepartment) Add(r *ghttp.Request) {
	var req *BaseSysRoleDepartmentModel.AddReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseSysRoleDepartmentService.S.Add(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "新增成功", g.Map{"id": id})
}

// @Summary 修改信息
// @Description 修改信息
// @Tags Admin-BaseSysRoleDepartment-Admin角色部门关联表管理
// @Param data body BaseSysRoleDepartmentModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysRoleDepartment/update [post]
// @Security
func (c *baseSysRoleDepartment) Update(r *ghttp.Request) {
	var req *BaseSysRoleDepartmentModel.UpdateReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseSysRoleDepartmentService.S.Update(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"id": id})
}

// @Summary 删除信息
// @Description 删除信息
// @Tags Admin-BaseSysRoleDepartment-Admin角色部门关联表管理
// @Param ids body BaseSysRoleDepartmentModel.DeleteReqParams  true "ids[1,2,3..]"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysRoleDepartment/delete [post]
// @Security
func (c *baseSysRoleDepartment) Delete(r *ghttp.Request) {
	var req *BaseSysRoleDepartmentModel.DeleteReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := BaseSysRoleDepartmentService.S.Delete(req.Ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

// @Summary 单个信息
// @Description 单个信息
// @Tags Admin-BaseSysRoleDepartment-Admin角色部门关联表管理
// @Param data body BaseSysRoleDepartmentModel.InfoReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysRoleDepartment/info [get]
// @Security
func (c *baseSysRoleDepartment) Info(r *ghttp.Request) {
	var req *BaseSysRoleDepartmentModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	result, err := BaseSysRoleDepartmentService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 信息列表
// @Description 信息列表
// @Tags Admin-BaseSysRoleDepartment-Admin角色部门关联表管理
// @Param data body BaseSysRoleDepartmentModel.ListReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysRoleDepartment/list [get]
// @Security
func (c *baseSysRoleDepartment) List(r *ghttp.Request) {
	var req *BaseSysRoleDepartmentModel.ListReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	condition := g.Map{}

	result, err := BaseSysRoleDepartmentService.S.List(condition)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取列表信息成功", result)
}

// @Summary 分页信息
// @Description 分页信息
// @Tags Admin-BaseSysRoleDepartment-Admin角色部门关联表管理
// @Param data body BaseSysRoleDepartmentModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysRoleDepartment/page [get]
// @Security
func (c *baseSysRoleDepartment) Page(r *ghttp.Request) {
	var req *BaseSysRoleDepartmentModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := BaseSysRoleDepartmentService.S.Page(req)
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
