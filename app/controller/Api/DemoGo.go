// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Api

import (
	"goEasy/app/model/DemoGoModel"
	"goEasy/app/service/DemoGoService"
	"goEasy/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var DemoGo = new(demoGo)

//控制器
type demoGo struct{}

// @Summary 添加信息
// @Description 添加信息
// @Tags Api-DemoGo-示例go管理
// @Param data body DemoGoModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/demoGo/add [post]
// @Security
func (c *demoGo) Add(r *ghttp.Request) {
	var req *DemoGoModel.AddReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := DemoGoService.S.Add(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "新增成功", g.Map{"id": id})
}

// @Summary 修改信息
// @Description 修改信息
// @Tags Api-DemoGo-示例go管理
// @Param data body DemoGoModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/demoGo/update [post]
// @Security
func (c *demoGo) Update(r *ghttp.Request) {
	var req *DemoGoModel.UpdateReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := DemoGoService.S.Update(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"id": id})
}

// @Summary 删除信息
// @Description 删除信息
// @Tags Api-DemoGo-示例go管理
// @Param ids body DemoGoModel.DeleteReqParams  true "ids[1,2,3..]"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/demoGo/delete [post]
// @Security
func (c *demoGo) Delete(r *ghttp.Request) {
	var req *DemoGoModel.DeleteReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := DemoGoService.S.Delete(req.Ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

// @Summary 单个信息
// @Description 单个信息
// @Tags Api-DemoGo-示例go管理
// @Param data body DemoGoModel.InfoReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/demoGo/info [get]
// @Security
func (c *demoGo) Info(r *ghttp.Request) {
	var req *DemoGoModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	result, err := DemoGoService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 信息列表
// @Description 信息列表
// @Tags Api-DemoGo-示例go管理
// @Param data body DemoGoModel.ListReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/demoGo/list [get]
// @Security
func (c *demoGo) List(r *ghttp.Request) {
	var req *DemoGoModel.ListReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	condition := g.Map{}

	if req.Title != "" {
		condition["title = ?"] = req.Title
	}

	if req.SubTitle != "" {
		condition["subTitle = ?"] = req.SubTitle
	}

	if req.Types != -1 {
		condition["types = ?"] = req.Types
	}

	if req.Status != -1 {
		condition["status = ?"] = req.Status
	}

	if req.Other != -1 {
		condition["other = ?"] = req.Other
	}

	if req.OtherStr != "" {
		condition["otherStr = ?"] = req.OtherStr
	}

	result, err := DemoGoService.S.List(condition)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取列表信息成功", result)
}

// @Summary 信息分页列表
// @Description 信息分页列表
// @Tags Api-DemoGo-示例go管理
// @Param data body DemoGoModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/demoGo/page [get]
// @Security
func (c *demoGo) Page(r *ghttp.Request) {
	var req *DemoGoModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := DemoGoService.S.Page(req)
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
