package ContextModel

import (
	"github.com/gogf/gf/frame/g"
)

const (
	// CtxKey 上下文变量存储键名，前后端系统共享
	CtxKey = "GfEasyContext"
)

// Context 请求上下文结构
type Context struct {
	User *CtxUser // 上下文用户信息
	Data g.Map    // 自定KV变量，业务模块根据需要设置，不固定
}

// CtxUser 请求上下文中的用户信息
type CtxUser struct {
	AdminUserId     int    `json:"adminUserId"`     // 用户id
	UserName        string `json:"userName"`        // 用户名
	RoleIds         string `json:"roleIds"`         // 部门id
	PasswordVersion string `json:"passwordVersion"` // 密码版本
	Exp             string `json:"exp"`             // token过期时间
	IsAdmin         int    `json:"isAdmin"`         // 是否后台管理员 1 是  0   否
	IsRefresh       bool   `json:"isRefresh "`      // 是否刷新token  true 是  false   否
	Avatar          string `json:"avatar"`          //头像
}

// GetAdminUserId 获取登录用户id
func (ctxUser *CtxUser) GetAdminUserId() (id int) {
	return ctxUser.AdminUserId
}

// GetDept 获取登录用户所属部门
//func (ctxUser *CtxUser) GetDept() (err error, dept *model.SysDept) {
//	err = g.DB().Model(SysDept.Table).Fields(SysDept.C.DeptId, SysDept.C.DeptName).WherePri(ctxUser.DeptId).Scan(&dept)
//	if dept == nil {
//		dept = &model.SysDept{}
//	}
//	return
//}
