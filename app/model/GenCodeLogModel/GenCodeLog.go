// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package GenCodeLogModel

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table genCodeLog.
type Entity struct {
	Id          int         `orm:"id,primary" json:"id"`           // ID
	CreateTime  *gtime.Time `orm:"createTime" json:"createTime"`   // 创建时间
	UpdateTime  *gtime.Time `orm:"updateTime" json:"updateTime"`   // 更新时间
	Title       string      `orm:"title" json:"title"`             // 标题
	Description string      `orm:"description" json:"description"` // 描述##IsSearchParams
	Tables      string      `orm:"tables" json:"tables"`           // 生成的表
	Params      string      `orm:"params" json:"params"`           // 当前请求参数
	Res         string      `orm:"res" json:"res"`                 // 请求结果
}

var (
	// Table is the table name of genCodeLog.
	Table = "gen_code_log"
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
	KeyWord   string `p:"keyWord"`
	Page      int    `p:"page" default:"1"`
	Size      int    `p:"size" default:"20"`
	Order     string `p:"order"`
	Sort      string `p:"sort"`
	IsExport  bool   `p:"isExport"`
	StartTime string `p:"startTime"`
	EndTime   string `p:"endTime"`
}
type AddReqParams struct {
	Title       string `p:"title" v:"required#标题不能为空"`
	Description string `p:"description" `
	Tables      string `p:"tables" v:"required#生成的表不能为空"`
	Params      string `p:"params" v:"required#当前请求参数不能为空"`
	Res         string `p:"res" v:"required#请求结果不能为空"`
}
type UpdateReqParams struct {
	Id          int    `p:"id" v:"required#ID不能为空"`
	Title       string `p:"title" `
	Description string `p:"description" `
	Tables      string `p:"tables" `
	Params      string `p:"params" `
	Res         string `p:"res" `
}
type InfoReqParams struct {
	Id int `p:"id" v:"required#id不能为空"`
}
type ListReqParams struct {
	Title       string `p:"title"`
	Description string `p:"description"`
}
type DeleteReqParams struct {
	Ids []int `p:"ids" v:"required#ids不能为空"`
}
