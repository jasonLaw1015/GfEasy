// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Admin

import (
	"goEasy/app/model/BaseSysLogModel"
	"goEasy/app/service/BaseSysLogService"
	"goEasy/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var BaseSysLog = new(baseSysLog)

//控制器
type baseSysLog struct{}

// @Summary 添加信息
// @Description 添加信息
// @Tags Admin-BaseSysLog-admin日志管理
// @Param data body BaseSysLogModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysLog/add [post]
// @Security
func (c *baseSysLog) Add(r *ghttp.Request) {
	var req *BaseSysLogModel.AddReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseSysLogService.S.Add(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "新增成功", g.Map{"id": id})
}

// @Summary 修改信息
// @Description 修改信息
// @Tags Admin-BaseSysLog-admin日志管理
// @Param data body BaseSysLogModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysLog/update [post]
// @Security
func (c *baseSysLog) Update(r *ghttp.Request) {
	var req *BaseSysLogModel.UpdateReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseSysLogService.S.Update(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"id": id})
}

// @Summary 删除信息
// @Description 删除信息
// @Tags Admin-BaseSysLog-admin日志管理
// @Param ids body BaseSysLogModel.DeleteReqParams  true "ids[1,2,3..]"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysLog/delete [post]
// @Security
func (c *baseSysLog) Delete(r *ghttp.Request) {
	var req *BaseSysLogModel.DeleteReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := BaseSysLogService.S.Delete(req.Ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

// @Summary 单个信息
// @Description 单个信息
// @Tags Admin-BaseSysLog-admin日志管理
// @Param data body BaseSysLogModel.InfoReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysLog/info [get]
// @Security
func (c *baseSysLog) Info(r *ghttp.Request) {
	var req *BaseSysLogModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	result, err := BaseSysLogService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 信息列表
// @Description 信息列表
// @Tags Admin-BaseSysLog-admin日志管理
// @Param data body BaseSysLogModel.ListReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysLog/list [get]
// @Security
func (c *baseSysLog) List(r *ghttp.Request) {
	var req *BaseSysLogModel.ListReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	condition := g.Map{}

	result, err := BaseSysLogService.S.List(condition)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取列表信息成功", result)
}

// @Summary 分页信息
// @Description 分页信息
// @Tags Admin-BaseSysLog-admin日志管理
// @Param data body BaseSysLogModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseSysLog/page [get]
// @Security
func (c *baseSysLog) Page(r *ghttp.Request) {
	var req *BaseSysLogModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := BaseSysLogService.S.Page(req)
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

func (c *baseSysLog) SetKeep(r *ghttp.Request) {

	err := BaseSysLogService.S.SetKeep("logKeep",
		r.GetRequestString("value"))
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "设置成功", nil)
}

func (c *baseSysLog) GetKeep(r *ghttp.Request) {

	data, err := BaseSysLogService.S.GetKeep("logKeep")
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取成功", data.CValue)
}

func (c *baseSysLog) Clear(r *ghttp.Request) {
	isAll := r.GetRequestBool("isAll")
	err := BaseSysLogService.S.Clear(isAll)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "处理成功", nil)
}
