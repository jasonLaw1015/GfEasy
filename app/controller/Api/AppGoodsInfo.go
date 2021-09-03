// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Api

import (
	"goEasy/app/model/AppGoodsInfoModel"
	"goEasy/app/service/AppGoodsInfoService"
	"goEasy/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var AppGoodsInfo = new(appGoodsInfo)

//控制器
type appGoodsInfo struct{}

// @Summary 添加信息
// @Description 添加信息
// @Tags Api-AppGoodsInfo-管理
// @Param data body AppGoodsInfoModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/appGoodsInfo/add [post]
// @Security
func (c *appGoodsInfo) Add(r *ghttp.Request) {
	var req *AppGoodsInfoModel.AddReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := AppGoodsInfoService.S.Add(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "新增成功", g.Map{"id": id})
}

// @Summary 修改信息
// @Description 修改信息
// @Tags Api-AppGoodsInfo-管理
// @Param data body AppGoodsInfoModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/appGoodsInfo/update [post]
// @Security
func (c *appGoodsInfo) Update(r *ghttp.Request) {
	var req *AppGoodsInfoModel.UpdateReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := AppGoodsInfoService.S.Update(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"id": id})
}

// @Summary 删除信息
// @Description 删除信息
// @Tags Api-AppGoodsInfo-管理
// @Param ids body AppGoodsInfoModel.DeleteReqParams  true "ids[1,2,3..]"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/appGoodsInfo/delete [post]
// @Security
func (c *appGoodsInfo) Delete(r *ghttp.Request) {
	var req *AppGoodsInfoModel.DeleteReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := AppGoodsInfoService.S.Delete(req.Ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

// @Summary 单个信息
// @Description 单个信息
// @Tags Api-AppGoodsInfo-管理
// @Param data body AppGoodsInfoModel.InfoReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/appGoodsInfo/info [get]
// @Security
func (c *appGoodsInfo) Info(r *ghttp.Request) {
	var req *AppGoodsInfoModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	result, err := AppGoodsInfoService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 信息列表
// @Description 信息列表
// @Tags Api-AppGoodsInfo-管理
// @Param data body AppGoodsInfoModel.ListReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/appGoodsInfo/list [get]
// @Security
func (c *appGoodsInfo) List(r *ghttp.Request) {
	var req *AppGoodsInfoModel.ListReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	condition := g.Map{}

	if req.Title != "" {
		condition["title = ?"] = req.Title
	}

	if req.Types != -1 {
		condition["types = ?"] = req.Types
	}

	if req.Status != -1 {
		condition["status = ?"] = req.Status
	}

	result, err := AppGoodsInfoService.S.List(condition)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取列表信息成功", result)
}

// @Summary 信息分页列表
// @Description 信息分页列表
// @Tags Api-AppGoodsInfo-管理
// @Param data body AppGoodsInfoModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/appGoodsInfo/page [get]
// @Security
func (c *appGoodsInfo) Page(r *ghttp.Request) {
	var req *AppGoodsInfoModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := AppGoodsInfoService.S.Page(req)
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
