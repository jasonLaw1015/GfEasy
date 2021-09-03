// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Api

import (
	"goEasy/app/model/BaseSysParamModel"
	"goEasy/app/service/BaseSysParamService"
	"goEasy/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var BaseSysParam = new(baseSysParam)

//控制器
type baseSysParam struct{}

// @Summary 添加信息
// @Description 添加信息
// @Tags Api-BaseSysParam-admin参数配置管理
// @Param data body BaseSysParamModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysParam/add [post]
// @Security
func (c *baseSysParam) Add(r *ghttp.Request) {
	var req *BaseSysParamModel.AddReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseSysParamService.S.Add(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "新增成功", g.Map{"id": id})
}

// @Summary 修改信息
// @Description 修改信息
// @Tags Api-BaseSysParam-admin参数配置管理
// @Param data body BaseSysParamModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysParam/update [post]
// @Security
func (c *baseSysParam) Update(r *ghttp.Request) {
	var req *BaseSysParamModel.UpdateReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseSysParamService.S.Update(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"id": id})
}

// @Summary 删除信息
// @Description 删除信息
// @Tags Api-BaseSysParam-admin参数配置管理
// @Param ids body BaseSysParamModel.DeleteReqParams  true "ids[1,2,3..]"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysParam/delete [post]
// @Security
func (c *baseSysParam) Delete(r *ghttp.Request) {
	var req *BaseSysParamModel.DeleteReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := BaseSysParamService.S.Delete(req.Ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

// @Summary 单个信息
// @Description 单个信息
// @Tags Api-BaseSysParam-admin参数配置管理
// @Param data body BaseSysParamModel.InfoReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysParam/info [get]
// @Security
func (c *baseSysParam) Info(r *ghttp.Request) {
	var req *BaseSysParamModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	result, err := BaseSysParamService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 信息列表
// @Description 信息列表
// @Tags Api-BaseSysParam-admin参数配置管理
// @Param data body BaseSysParamModel.ListReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysParam/list [get]
// @Security
func (c *baseSysParam) List(r *ghttp.Request) {
	var req *BaseSysParamModel.ListReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	condition := g.Map{}

	if req.Name != "" {
		condition["name = ?"] = req.Name
	}

	result, err := BaseSysParamService.S.List(condition)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取列表信息成功", result)
}

// @Summary 信息分页列表
// @Description 信息分页列表
// @Tags Api-BaseSysParam-admin参数配置管理
// @Param data body BaseSysParamModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseSysParam/page [get]
// @Security
func (c *baseSysParam) Page(r *ghttp.Request) {
	var req *BaseSysParamModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := BaseSysParamService.S.Page(req)
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
