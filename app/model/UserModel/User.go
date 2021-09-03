// ==========================================================================
// 生成日期：2021-08-03 16:06:40
// 生成人：JasonLaw
// ==========================================================================
package UserModel

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table user.
type Entity struct {
	Id         uint        `orm:"id,primary" json:"id"`         //
	Passport   string      `orm:"passport" json:"passport"`     // 账户
	Password   string      `orm:"password" json:"password"`     // 密码
	Nickname   string      `orm:"nickname" json:"nickname"`     // 昵称
	CreateAt   *gtime.Time `orm:"createAt" json:"createAt"`     // 112
	UserName   string      `orm:"user_name" json:"user_name"`   // 下横线字段
	Content    string      `orm:"content" json:"content"`       // 内容
	UpdateAt   *gtime.Time `orm:"updateAt" json:"updateAt"`     // 开1
	CreateTime *gtime.Time `orm:"createTime" json:"createTime"` //
	UpdateTime *gtime.Time `orm:"updateTime" json:"updateTime"` //
	Name       string      `orm:"name" json:"name"`             // 名称·
	Types      int         `orm:"types" json:"types"`           // 类型
	Status     int         `orm:"status" json:"status"`         // 状态
}

var (
	// Table is the table name of user.
	Table = "user"
	M     = g.DB("default").Model(Table).Safe()
)

/**
{
    "keyWord": "商品标题", // 模糊搜索，搜索的字段对应keyWordLikeFields
    "type": 1, // 全等于筛选，对应fieldEq
    "page": 2, // 第几页
    "size": 1, // 每页返回个数
    "sort": "desc", // 排序方向
    "order": "id" // 排序字段
}
*/
//分页请求参数
type PageReqParams struct {
	KeyWord string `p:"keyWord"`
	Page    int    `p:"page" default:"1"`
	Size    int    `p:"size" default:"20"`
	Types   string `p:"types"`
	Order   string `p:"order"`
	Sort    string `p:"sort"`
}
type AddReqParams struct {
	Passport string      `p:"passport" v:"required#账户不能为空"`
	Password string      `p:"password" `
	Nickname string      `p:"nickname" v:"required#昵称不能为空"`
	CreateAt *gtime.Time `p:"createAt" `
	UserName string      `p:"user_name" `
	Content  string      `p:"content" `
	UpdateAt *gtime.Time `p:"updateAt" `
	Name     string      `p:"name" v:"required#名称·不能为空"`
	Types    int         `p:"types" v:"required#类型不能为空"`
	Status   int         `p:"status" v:"required#状态不能为空"`
}
type UpdateReqParams struct {
	Id       uint        `p:"id" v:"required#不能为空"`
	Passport string      `p:"passport" v:"required#账户不能为空"`
	Password string      `p:"password" `
	Nickname string      `p:"nickname" v:"required#昵称不能为空"`
	CreateAt *gtime.Time `p:"createAt" `
	UserName string      `p:"user_name" `
	Content  string      `p:"content" `
	UpdateAt *gtime.Time `p:"updateAt" `
	Name     string      `p:"name" v:"required#名称·不能为空"`
	Types    int         `p:"types" v:"required#类型不能为空"`
	Status   int         `p:"status" v:"required#状态不能为空"`
}
type InfoReqParams struct {
	Id int `p:"id" v:"required#id不能为空"`
}
type ListReqParams struct {
	Name   string `p:"name"`
	Types  int    `p:"types"`
	Status int    `p:"status"`
}
type DeleteReqParams struct {
	Ids []int `p:"ids" v:"required#ids不能为空"`
}
