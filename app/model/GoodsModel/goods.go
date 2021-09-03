// ==========================================================================
// 生成日期：2021-08-08 23:47:39
// 生成人：JasonLaw
// ==========================================================================
package GoodsModel

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table goods.
type Entity struct {
	Id         uint        `orm:"id,primary" json:"id"`         // ID
	Name       string      `orm:"name" json:"name"`             // 标题
	Inventory  int         `orm:"inventory" json:"inventory"`   // 库存
	Pic        string      `orm:"pic" json:"pic"`               // 图片
	CreateTime *gtime.Time `orm:"createTime" json:"createTime"` // 创建时间
	UpdateTime *gtime.Time `orm:"updateTime" json:"updateTime"` // 更新时间
	Status     int         `orm:"status" json:"status"`         // 状态#1:启用,2:禁用
	Types      int         `orm:"types" json:"types"`           // 类型#1:上架,2:下架
	StartTime  *gtime.Time `orm:"startTime" json:"startTime"`   // 活动开始时间
	EndTime    *gtime.Time `orm:"endTime" json:"endTime"`       // 活动结束时间
}

var (
	// Table is the table name of goods.
	Table = "goods"
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
	Name      string      `p:"name" `
	Inventory int         `p:"inventory" v:"required#库存不能为空"`
	Pic       string      `p:"pic" v:"required#图片不能为空"`
	Status    int         `p:"status" v:"required#状态#1:启用,2:禁用不能为空"`
	Types     int         `p:"types" v:"required#类型#1:上架,2:下架不能为空"`
	StartTime *gtime.Time `p:"startTime" v:"required#活动开始时间不能为空"`
	EndTime   *gtime.Time `p:"endTime" v:"required#活动结束时间不能为空"`
}
type UpdateReqParams struct {
	Id        uint        `p:"id" v:"required#ID不能为空"`
	Name      string      `p:"name" `
	Inventory int         `p:"inventory" `
	Pic       string      `p:"pic" `
	Status    int         `p:"status" `
	Types     int         `p:"types" `
	StartTime *gtime.Time `p:"startTime" `
	EndTime   *gtime.Time `p:"endTime" `
}
type InfoReqParams struct {
	Id int `p:"id" v:"required#id不能为空"`
}
type ListReqParams struct {
	Name   string `p:"name"`
	Status int    `p:"status"`
	Types  int    `p:"types"`
}
type DeleteReqParams struct {
	Ids []int `p:"ids" v:"required#ids不能为空"`
}
