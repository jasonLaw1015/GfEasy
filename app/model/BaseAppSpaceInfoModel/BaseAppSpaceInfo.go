// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package BaseAppSpaceInfoModel

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table baseAppSpaceInfo.
type Entity struct {
	Id         int         `orm:"id,primary" json:"id"`         // ID
	CreateTime *gtime.Time `orm:"createTime" json:"createTime"` // 创建时间
	UpdateTime *gtime.Time `orm:"updateTime" json:"updateTime"` // 更新时间
	Url        string      `orm:"url" json:"url"`               // 地址
	Type       string      `orm:"type" json:"type"`             // 类型
	ClassifyId int64       `orm:"classifyId" json:"classifyId"` // 分类ID
}

var (
	// Table is the table name of baseAppSpaceInfo.
	Table = "base_app_space_info"
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
	KeyWord    string `p:"keyWord"`
	ClassifyId int64  `p:"classifyId"`
	Page       int    `p:"page" default:"1"`
	Size       int    `p:"size" default:"20"`
	Order      string `p:"order"`
	Sort       string `p:"sort"`
	IsExport   bool   `p:"isExport"`
	StartTime  string `p:"startTime"`
	EndTime    string `p:"endTime"`
}
type AddReqParams struct {
	Url        string `p:"url" v:"required#地址不能为空"`
	Type       string `p:"type" v:"required#类型不能为空"`
	ClassifyId int64  `p:"classifyId" `
}
type UpdateReqParams struct {
	Id         int    `p:"id" v:"required#ID不能为空"`
	Url        string `p:"url" `
	Type       string `p:"type" `
	ClassifyId int64  `p:"classifyId" `
}
type InfoReqParams struct {
	Id int `p:"id" v:"required#id不能为空"`
}
type ListReqParams struct {
}
type DeleteReqParams struct {
	Ids []int `p:"ids" v:"required#ids不能为空"`
}
