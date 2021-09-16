// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Api

import (
	"gfEasy/app/model/BaseSysUserModel"
	"gfEasy/app/service/BaseSysUserService"
	"gfEasy/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var BaseSysUser = new(baseSysUser)

//控制器
type baseSysUser struct{}

// @Summary 添加信息
// @Description 添加信息
// @Tags Api-BaseSysUser-admin用户表管理
// @Param data body BaseSysUserModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysUser/add [post]
// @Security
func (c *baseSysUser) Add(r *ghttp.Request) {
	var req *BaseSysUserModel.AddReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseSysUserService.S.Add(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "新增成功", g.Map{"id": id})
}

// @Summary 修改信息
// @Description 修改信息
// @Tags Api-BaseSysUser-admin用户表管理
// @Param data body BaseSysUserModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysUser/update [post]
// @Security
func (c *baseSysUser) Update(r *ghttp.Request) {
	var req *BaseSysUserModel.UpdateReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseSysUserService.S.Update(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"id": id})
}

// @Summary 删除信息
// @Description 删除信息
// @Tags Api-BaseSysUser-admin用户表管理
// @Param ids body BaseSysUserModel.DeleteReqParams  true "ids[1,2,3..]"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysUser/delete [post]
// @Security
func (c *baseSysUser) Delete(r *ghttp.Request) {
	var req *BaseSysUserModel.DeleteReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := BaseSysUserService.S.Delete(req.Ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

// @Summary 单个信息
// @Description 单个信息
// @Tags Api-BaseSysUser-admin用户表管理
// @Param data body BaseSysUserModel.InfoReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysUser/info [get]
// @Security
func (c *baseSysUser) Info(r *ghttp.Request) {
	var req *BaseSysUserModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	result, err := BaseSysUserService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 信息列表
// @Description 信息列表
// @Tags Api-BaseSysUser-admin用户表管理
// @Param data body BaseSysUserModel.ListReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysUser/list [get]
// @Security
func (c *baseSysUser) List(r *ghttp.Request) {
	var req *BaseSysUserModel.ListReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	condition := g.Map{}

	if req.Name != "" {
		condition["name = ?"] = req.Name
	}

	if req.Status != -1 {
		condition["status = ?"] = req.Status
	}

	result, err := BaseSysUserService.S.List(condition)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取列表信息成功", result)
}

// @Summary 信息分页列表
// @Description 信息分页列表
// @Tags Api-BaseSysUser-admin用户表管理
// @Param data body BaseSysUserModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysUser/page [get]
// @Security
func (c *baseSysUser) Page(r *ghttp.Request) {
	var req *BaseSysUserModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := BaseSysUserService.S.Page(req)
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
