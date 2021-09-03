// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package BaseSysUserModel

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table baseSysUser.
type Entity struct {
	Id           int         `orm:"id,primary" json:"id"`             // ID
	CreateTime   *gtime.Time `orm:"createTime" json:"createTime"`     // 创建时间
	UpdateTime   *gtime.Time `orm:"updateTime" json:"updateTime"`     // 更新时间
	DepartmentId int64       `orm:"departmentId" json:"departmentId"` // 部门ID
	Name         string      `orm:"name" json:"name"`                 // 姓名
	Username     string      `orm:"username" json:"username"`         // 用户名
	Password     string      `orm:"password" json:"password"`         // 密码
	PasswordV    int         `orm:"passwordV" json:"passwordV"`       // 密码版本, 作用是改完密码，让原来的token失效
	Salt         string      `orm:"salt" json:"salt"`                 // 密码盐
	NickName     string      `orm:"nickName" json:"nickName"`         // 昵称
	HeadImg      string      `orm:"headImg" json:"headImg"`           // 头像
	Phone        string      `orm:"phone" json:"phone"`               // 手机
	Email        string      `orm:"email" json:"email"`               // 邮箱
	Status       int         `orm:"status" json:"status"`             // 状态 0:禁用 1：启用
	Remark       string      `orm:"remark" json:"remark"`             // 备注
	SocketId     string      `orm:"socketId" json:"socketId"`         // socketId
}

var (
	// Table is the table name of baseSysUser.
	Table = "base_sys_user"
	M     = g.DB("default").Model(Table).Safe()
)

/**
{
    "keyWord": "商品标题", // 模糊搜索，搜索的字段对应keyWordLikeFields
    "types": 1, // 全等于筛选，对应fieldEq
    "page": 2, // 第几页
    "size": 1, // 每页返回个数
    "sort": "desc", // 排序方向
    "order": "id" // 排序字段
	"startTime": "2021-08-02 00:00:00",//用创建时间判断
    "endTime": "2021-09-30 23:59:59",//用创建时间判断
	"isExport": true,//是否是都出excel
}
*/
//分页请求参数
type PageReqParams struct {
	KeyWord       string `p:"keyWord"`
	Page          int    `p:"page" default:"1"`
	Size          int    `p:"size" default:"20"`
	Status        int    `p:"status" default:"-1"`
	Order         string `p:"order"`
	Sort          string `p:"sort"`
	IsExport      bool   `p:"isExport"`
	StartTime     string `p:"startTime"`
	EndTime       string `p:"endTime"`
	DepartmentIds []int  `p:"departmentIds"`
}
type AddReqParams struct {
	DepartmentId int64    `p:"departmentId" `
	Name         string   `p:"name" `
	Username     string   `p:"username" v:"required#用户名不能为空"`
	Password     string   `p:"password" v:"required#密码不能为空"`
	PasswordV    int      `p:"passwordV"`
	Salt         string   `p:"salt"`
	NickName     string   `p:"nickName" `
	HeadImg      string   `p:"headImg" `
	Phone        string   `p:"phone" `
	Email        string   `p:"email" `
	Status       int      `p:"status" v:"required#状态 0:禁用 1：启用不能为空"`
	Remark       string   `p:"remark" `
	SocketId     string   `p:"socketId" `
	RoleIdList   []string `p:"menuIdList" `
}
type UpdateReqParams struct {
	Id           int      `p:"id" v:"required#ID不能为空"`
	DepartmentId int64    `p:"departmentId" `
	Name         string   `p:"name" `
	Username     string   `p:"username" `
	Password     string   `p:"password" `
	PasswordV    int      `p:"passwordV" `
	Salt         string   `p:"salt" `
	NickName     string   `p:"nickName" `
	HeadImg      string   `p:"headImg" `
	Phone        string   `p:"phone" `
	Email        string   `p:"email" `
	Status       int      `p:"status" `
	Remark       string   `p:"remark" `
	SocketId     string   `p:"socketId" `
	RoleIdList   []string `p:"menuIdList" `
}
type MoveReqParams struct {
	DepartmentId int   `p:"departmentId" v:"required#部门ID不能为空"`
	AdminUserIds []int `p:"adminUserIds" v:"required#管理员ID不能为空"`
}
type InfoReqParams struct {
	Id int `p:"id" v:"required#id不能为空"`
}
type ListReqParams struct {
	Name   string `p:"name"`
	Status int    `p:"status"`
}
type DeleteReqParams struct {
	Ids []int `p:"ids" v:"required#ids不能为空"`
}

type InfoResponse struct {
	Id           int         `json:"id"`           // ID
	CreateTime   *gtime.Time `json:"createTime"`   // 创建时间
	UpdateTime   *gtime.Time `json:"updateTime"`   // 更新时间
	DepartmentId int64       `json:"departmentId"` // 部门ID
	Name         string      `json:"name"`         // 姓名
	Username     string      `json:"username"`     // 用户名
	Password     string      `json:"password"`     // 密码
	PasswordV    int         `json:"passwordV"`    // 密码版本, 作用是改完密码，让原来的token失效
	Salt         string      `json:"salt"`         // 密码盐
	NickName     string      `json:"nickName"`     // 昵称
	HeadImg      string      `json:"headImg"`      // 头像
	Phone        string      `json:"phone"`        // 手机
	Email        string      `json:"email"`        // 邮箱
	Status       int         `json:"status"`       // 状态 0:禁用 1：启用
	Remark       string      `json:"remark"`       // 备注
	SocketId     string      `json:"socketId"`     // socketId
	RoleIdList   []int64     `json:"roleIdList"`   //角色id数组
}
