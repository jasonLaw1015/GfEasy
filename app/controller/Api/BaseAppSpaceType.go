// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Api

import (
	"goEasy/app/model/BaseAppSpaceTypeModel"
	"goEasy/app/service/BaseAppSpaceTypeService"
	"goEasy/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var BaseAppSpaceType = new(baseAppSpaceType)

//控制器
type baseAppSpaceType struct{}

// @Summary 添加信息
// @Description 添加信息
// @Tags Api-BaseAppSpaceType-图片空间类型管理
// @Param data body BaseAppSpaceTypeModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseAppSpaceType/add [post]
// @Security
func (c *baseAppSpaceType) Add(r *ghttp.Request) {
	var req *BaseAppSpaceTypeModel.AddReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseAppSpaceTypeService.S.Add(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "新增成功", g.Map{"id": id})
}

// @Summary 修改信息
// @Description 修改信息
// @Tags Api-BaseAppSpaceType-图片空间类型管理
// @Param data body BaseAppSpaceTypeModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseAppSpaceType/update [post]
// @Security
func (c *baseAppSpaceType) Update(r *ghttp.Request) {
	var req *BaseAppSpaceTypeModel.UpdateReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := BaseAppSpaceTypeService.S.Update(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"id": id})
}

// @Summary 删除信息
// @Description 删除信息
// @Tags Api-BaseAppSpaceType-图片空间类型管理
// @Param ids body BaseAppSpaceTypeModel.DeleteReqParams  true "ids[1,2,3..]"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseAppSpaceType/delete [post]
// @Security
func (c *baseAppSpaceType) Delete(r *ghttp.Request) {
	var req *BaseAppSpaceTypeModel.DeleteReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := BaseAppSpaceTypeService.S.Delete(req.Ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

// @Summary 单个信息
// @Description 单个信息
// @Tags Api-BaseAppSpaceType-图片空间类型管理
// @Param data body BaseAppSpaceTypeModel.InfoReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseAppSpaceType/info [get]
// @Security
func (c *baseAppSpaceType) Info(r *ghttp.Request) {
	var req *BaseAppSpaceTypeModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	result, err := BaseAppSpaceTypeService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 信息列表
// @Description 信息列表
// @Tags Api-BaseAppSpaceType-图片空间类型管理
// @Param data body BaseAppSpaceTypeModel.ListReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseAppSpaceType/list [get]
// @Security
func (c *baseAppSpaceType) List(r *ghttp.Request) {
	var req *BaseAppSpaceTypeModel.ListReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	condition := g.Map{}

	if req.Name != "" {
		condition["name = ?"] = req.Name
	}

	result, err := BaseAppSpaceTypeService.S.List(condition)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取列表信息成功", result)
}

// @Summary 信息分页列表
// @Description 信息分页列表
// @Tags Api-BaseAppSpaceType-图片空间类型管理
// @Param data body BaseAppSpaceTypeModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /api/baseAppSpaceType/page [get]
// @Security
func (c *baseAppSpaceType) Page(r *ghttp.Request) {
	var req *BaseAppSpaceTypeModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := BaseAppSpaceTypeService.S.Page(req)
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
