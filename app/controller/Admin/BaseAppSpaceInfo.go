// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Admin

import (
	"goEasy/app/model/BaseAppSpaceInfoModel"
	"goEasy/app/service/BaseAppSpaceInfoService"
	"goEasy/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var BaseAppSpaceInfo = new(baseAppSpaceInfo)

//控制器
type baseAppSpaceInfo struct{}

// @Summary 添加信息
// @Description 添加信息
// @Tags Admin-BaseAppSpaceInfo-图片空间信息管理
// @Param data body BaseAppSpaceInfoModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseAppSpaceInfo/add [post]
// @Security
func (c *baseAppSpaceInfo) Add(r *ghttp.Request) {
	var req *BaseAppSpaceInfoModel.AddReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseAppSpaceInfoService.S.Add(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "新增成功", g.Map{"id": id})
}

// @Summary 修改信息
// @Description 修改信息
// @Tags Admin-BaseAppSpaceInfo-图片空间信息管理
// @Param data body BaseAppSpaceInfoModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseAppSpaceInfo/update [post]
// @Security
func (c *baseAppSpaceInfo) Update(r *ghttp.Request) {
	var req *BaseAppSpaceInfoModel.UpdateReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseAppSpaceInfoService.S.Update(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"id": id})
}

// @Summary 删除信息
// @Description 删除信息
// @Tags Admin-BaseAppSpaceInfo-图片空间信息管理
// @Param ids body BaseAppSpaceInfoModel.DeleteReqParams  true "ids[1,2,3..]"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseAppSpaceInfo/delete [post]
// @Security
func (c *baseAppSpaceInfo) Delete(r *ghttp.Request) {
	var req *BaseAppSpaceInfoModel.DeleteReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := BaseAppSpaceInfoService.S.Delete(req.Ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

// @Summary 单个信息
// @Description 单个信息
// @Tags Admin-BaseAppSpaceInfo-图片空间信息管理
// @Param data body BaseAppSpaceInfoModel.InfoReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseAppSpaceInfo/info [get]
// @Security
func (c *baseAppSpaceInfo) Info(r *ghttp.Request) {
	var req *BaseAppSpaceInfoModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	result, err := BaseAppSpaceInfoService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 信息列表
// @Description 信息列表
// @Tags Admin-BaseAppSpaceInfo-图片空间信息管理
// @Param data body BaseAppSpaceInfoModel.ListReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseAppSpaceInfo/list [get]
// @Security
func (c *baseAppSpaceInfo) List(r *ghttp.Request) {
	var req *BaseAppSpaceInfoModel.ListReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	condition := g.Map{}

	result, err := BaseAppSpaceInfoService.S.List(condition)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取列表信息成功", result)
}

// @Summary 分页信息
// @Description 分页信息
// @Tags Admin-BaseAppSpaceInfo-图片空间信息管理
// @Param data body BaseAppSpaceInfoModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseAppSpaceInfo/page [get]
// @Security
func (c *baseAppSpaceInfo) Page(r *ghttp.Request) {
	var req *BaseAppSpaceInfoModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := BaseAppSpaceInfoService.S.Page(req)
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
