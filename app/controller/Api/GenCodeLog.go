// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Api

import (
	"gfEasy/app/model/GenCodeLogModel"
	"gfEasy/app/service/GenCodeLogService"
	"gfEasy/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var GenCodeLog = new(genCodeLog)

//控制器
type genCodeLog struct{}

// @Summary 添加信息
// @Description 添加信息
// @Tags Api-GenCodeLog-代码生成日志管理
// @Param data body GenCodeLogModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/genCodeLog/add [post]
// @Security
func (c *genCodeLog) Add(r *ghttp.Request) {
	var req *GenCodeLogModel.AddReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := GenCodeLogService.S.Add(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "新增成功", g.Map{"id": id})
}

// @Summary 修改信息
// @Description 修改信息
// @Tags Api-GenCodeLog-代码生成日志管理
// @Param data body GenCodeLogModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/genCodeLog/update [post]
// @Security
func (c *genCodeLog) Update(r *ghttp.Request) {
	var req *GenCodeLogModel.UpdateReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := GenCodeLogService.S.Update(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"id": id})
}

// @Summary 删除信息
// @Description 删除信息
// @Tags Api-GenCodeLog-代码生成日志管理
// @Param ids body GenCodeLogModel.DeleteReqParams  true "ids[1,2,3..]"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/genCodeLog/delete [post]
// @Security
func (c *genCodeLog) Delete(r *ghttp.Request) {
	var req *GenCodeLogModel.DeleteReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := GenCodeLogService.S.Delete(req.Ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

// @Summary 单个信息
// @Description 单个信息
// @Tags Api-GenCodeLog-代码生成日志管理
// @Param data body GenCodeLogModel.InfoReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/genCodeLog/info [get]
// @Security
func (c *genCodeLog) Info(r *ghttp.Request) {
	var req *GenCodeLogModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	result, err := GenCodeLogService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 信息列表
// @Description 信息列表
// @Tags Api-GenCodeLog-代码生成日志管理
// @Param data body GenCodeLogModel.ListReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/genCodeLog/list [get]
// @Security
func (c *genCodeLog) List(r *ghttp.Request) {
	var req *GenCodeLogModel.ListReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	condition := g.Map{}

	if req.Title != "" {
		condition["title = ?"] = req.Title
	}

	if req.Description != "" {
		condition["description = ?"] = req.Description
	}

	result, err := GenCodeLogService.S.List(condition)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取列表信息成功", result)
}

// @Summary 信息分页列表
// @Description 信息分页列表
// @Tags Api-GenCodeLog-代码生成日志管理
// @Param data body GenCodeLogModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/genCodeLog/page [get]
// @Security
func (c *genCodeLog) Page(r *ghttp.Request) {
	var req *GenCodeLogModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := GenCodeLogService.S.Page(req)
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
