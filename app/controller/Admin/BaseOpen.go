// ==========================================================================
// 生成日期：2021-08-09 16:12:06
// 生成人：JasonLaw
// ==========================================================================
package Admin

import (
	"gfEasy/app/model/BaseOpenModel"
	"gfEasy/app/model/BaseSysParamModel"
	"gfEasy/app/service/BaseOpenService"
	"gfEasy/library/response"
	"gfEasy/library/utils/captcha"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//不需要登录的后台接口
var BaseOpen = new(baseOpen)

//控制器
type baseOpen struct{}

// @Summary 获取验证码
// @Description 获取验证码
// @Tags Admin-BaseOpen-无需token的接口管理
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseOpen/Captcha [get]
// @Security
func (c *baseOpen) Captcha(r *ghttp.Request) {
	captchaId, data, err := captcha.GetVerifyImgString()
	if err != nil {
		response.FailJson(true, r, "获取验证码失败", err.Error())
	}
	response.SusJson(true, r, "获取信息成功", g.Map{"captchaId": captchaId, "data": data})
}

// @Summary 后台登录接口
// @Description 后台登录接口
// @Tags Admin-BaseOpen-无需token的接口管理
// @Param data body BaseOpenModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseOpen/login [post]
// @Security
func (c *baseOpen) Login(r *ghttp.Request) {
	var req *BaseOpenModel.LoginReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	result, err := BaseOpenService.S.Login(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "登录成功", result)
}

// @Summary 刷新token
// @Description 刷新token
// @Tags Admin-BaseOpen-无需token的接口管理
// @Param data body BaseOpenModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseOpen/refreshToken [post]
// @Security
func (c *baseOpen) RefreshToken(r *ghttp.Request) {
	var req *BaseOpenModel.LoginReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := BaseOpenService.S.RefreshToken(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", nil)
}

// @Summary 根据配置参数key获得网页内容(富文本)
// @Description 根据配置参数key获得网页内容(富文本)
// @Tags Admin-BaseOpen-无需token的接口管理
// @Param data body BaseOpenModel.SysParamsReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/baseOpen/html [get]
// @Security
func (c *baseOpen) Html(r *ghttp.Request) {
	var paramEntity *BaseSysParamModel.Entity
	var req *BaseOpenModel.SysParamsReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	err := BaseSysParamModel.M.Where("name", req.Name).Scan(&paramEntity)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", paramEntity)
}
