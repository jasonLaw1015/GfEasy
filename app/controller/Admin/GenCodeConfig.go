// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Admin

import (
	"goEasy/app/model/GenCodeConfigModel"
	"goEasy/app/service/GenCodeConfigService"
	"goEasy/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var GenCodeConfig = new(genCodeConfig)

//控制器
type genCodeConfig struct{}

// @Summary 添加信息
// @Description 添加信息
// @Tags Admin-GenCodeConfig-代码生成配置管理
// @Param data body GenCodeConfigModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/genCodeConfig/add [post]
// @Security
func (c *genCodeConfig) Add(r *ghttp.Request) {
	var req *GenCodeConfigModel.AddReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := GenCodeConfigService.S.Add(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "新增成功", g.Map{"id": id})
}

// @Summary 修改信息
// @Description 修改信息
// @Tags Admin-GenCodeConfig-代码生成配置管理
// @Param data body GenCodeConfigModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/genCodeConfig/update [post]
// @Security
func (c *genCodeConfig) Update(r *ghttp.Request) {
	var req *GenCodeConfigModel.UpdateReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := GenCodeConfigService.S.Update(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"id": id})
}

// @Summary 删除信息
// @Description 删除信息
// @Tags Admin-GenCodeConfig-代码生成配置管理
// @Param ids body GenCodeConfigModel.DeleteReqParams  true "ids[1,2,3..]"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/genCodeConfig/delete [post]
// @Security
func (c *genCodeConfig) Delete(r *ghttp.Request) {
	var req *GenCodeConfigModel.DeleteReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := GenCodeConfigService.S.Delete(req.Ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

// @Summary 单个信息
// @Description 单个信息
// @Tags Admin-GenCodeConfig-代码生成配置管理
// @Param data body GenCodeConfigModel.InfoReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/genCodeConfig/info [get]
// @Security
func (c *genCodeConfig) Info(r *ghttp.Request) {
	var req *GenCodeConfigModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	result, err := GenCodeConfigService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 信息列表
// @Description 信息列表
// @Tags Admin-GenCodeConfig-代码生成配置管理
// @Param data body GenCodeConfigModel.ListReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/genCodeConfig/list [get]
// @Security
func (c *genCodeConfig) List(r *ghttp.Request) {
	var req *GenCodeConfigModel.ListReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	condition := g.Map{}

	result, err := GenCodeConfigService.S.List(condition)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取列表信息成功", result)
}

// @Summary 分页信息
// @Description 分页信息
// @Tags Admin-GenCodeConfig-代码生成配置管理
// @Param data body GenCodeConfigModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/genCodeConfig/page [get]
// @Security
func (c *genCodeConfig) Page(r *ghttp.Request) {
	var req *GenCodeConfigModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := GenCodeConfigService.S.Page(req)
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

func (c *genCodeConfig) FirstInfo(r *ghttp.Request) {

	result, err := GenCodeConfigService.S.FirstInfo()
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 创建或更新信息
// @Description 创建或更新信息
// @Tags Admin-GenCodeConfig-代码生成配置管理
// @Param data body GenCodeConfigModel.SaveReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/genCodeConfig/save [post]
// @Security
func (c *genCodeConfig) Save(r *ghttp.Request) {
	var req *GenCodeConfigModel.SaveReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	res, err := GenCodeConfigService.S.Save(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"res": res})
}
