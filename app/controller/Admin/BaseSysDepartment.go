// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Admin

import (
	"goEasy/app/model/BaseSysDepartmentModel"
	"goEasy/app/service/BaseSysDepartmentService"
	"goEasy/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var BaseSysDepartment = new(baseSysDepartment)

//控制器
type baseSysDepartment struct{}

// @Summary 添加信息
// @Description 添加信息
// @Tags Admin-BaseSysDepartment-admin部门管理
// @Param data body BaseSysDepartmentModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysDepartment/add [post]
// @Security
func (c *baseSysDepartment) Add(r *ghttp.Request) {
	var req *BaseSysDepartmentModel.AddReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseSysDepartmentService.S.Add(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "新增成功", g.Map{"id": id})
}

// @Summary 修改信息
// @Description 修改信息
// @Tags Admin-BaseSysDepartment-admin部门管理
// @Param data body BaseSysDepartmentModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysDepartment/update [post]
// @Security
func (c *baseSysDepartment) Update(r *ghttp.Request) {
	var req *BaseSysDepartmentModel.UpdateReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseSysDepartmentService.S.Update(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"id": id})
}

// @Summary 删除信息
// @Description 删除信息
// @Tags Admin-BaseSysDepartment-admin部门管理
// @Param ids body BaseSysDepartmentModel.DeleteReqParams  true "ids[1,2,3..]"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysDepartment/delete [post]
// @Security
func (c *baseSysDepartment) Delete(r *ghttp.Request) {
	var req *BaseSysDepartmentModel.DeleteReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := BaseSysDepartmentService.S.Delete(req.Ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

// @Summary 单个信息
// @Description 单个信息
// @Tags Admin-BaseSysDepartment-admin部门管理
// @Param data body BaseSysDepartmentModel.InfoReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysDepartment/info [get]
// @Security
func (c *baseSysDepartment) Info(r *ghttp.Request) {
	var req *BaseSysDepartmentModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	result, err := BaseSysDepartmentService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 信息列表
// @Description 信息列表
// @Tags Admin-BaseSysDepartment-admin部门管理
// @Param data body BaseSysDepartmentModel.ListReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysDepartment/list [get]
// @Security
func (c *baseSysDepartment) List(r *ghttp.Request) {
	var req *BaseSysDepartmentModel.ListReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	condition := g.Map{}

	if req.Name != "" {
		condition["name = ?"] = req.Name
	}

	result, err := BaseSysDepartmentService.S.List(condition)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取列表信息成功", result)
}

// @Summary 分页信息
// @Description 分页信息
// @Tags Admin-BaseSysDepartment-admin部门管理
// @Param data body BaseSysDepartmentModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysDepartment/page [get]
// @Security
func (c *baseSysDepartment) Page(r *ghttp.Request) {
	var req *BaseSysDepartmentModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := BaseSysDepartmentService.S.Page(req)
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
