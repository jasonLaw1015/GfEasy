// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package DemoGoModel

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table demoGo.
type Entity struct {
	Id             int         `orm:"id,primary" json:"id"`                 // ID
	CreateTime     *gtime.Time `orm:"createTime" json:"createTime"`         // 创建时间
	UpdateTime     *gtime.Time `orm:"updateTime" json:"updateTime"`         // 更新时间
	Title          string      `orm:"title" json:"title"`                   // 标题
	SubTitle       string      `orm:"subTitle" json:"subTitle"`             // 副标题##IsSearchParams
	Pic            string      `orm:"pic" json:"pic"`                       // 商品主图
	Types          int         `orm:"types" json:"types"`                   // 类型#1:上架,2:下架
	Status         int         `orm:"status" json:"status"`                 // 状态#1:启用,2:禁用
	Tupian         string      `orm:"tupian" json:"tupian"`                 // 图片##IsPicColumn
	Other          int         `orm:"other" json:"other"`                   // 其他状态#1:已激活,2:未激活#IsDictColumn,IsSearchParams
	Sort           int         `orm:"sort" json:"sort"`                     // 排序
	OtherStr       string      `orm:"otherStr" json:"otherStr"`             // 其他状态2#s1:已激活,s2a2:未激活#IsDictColumn,IsSearchParams
	AppGoodsInfoId int         `orm:"appGoodsInfoId" json:"appGoodsInfoId"` // appGoodsInfoID
	TestName       string      `orm:"testName" json:"testName"`             // 组名
}

var (
	// Table is the table name of demoGo.
	Table = "demo_go"
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
	Types     int    `p:"types" default:"-1"`
	Status    int    `p:"status" default:"-1"`
	Other     int    `p:"other" default:"-1"`
	OtherStr  string `p:"otherStr" default:"-1"`
	Order     string `p:"order"`
	Sort      string `p:"sort"`
	IsExport  bool   `p:"isExport"`
	StartTime string `p:"startTime"`
	EndTime   string `p:"endTime"`
}
type AddReqParams struct {
	Title          string `p:"title" v:"required#标题不能为空"`
	SubTitle       string `p:"subTitle" `
	Pic            string `p:"pic" v:"required#商品主图不能为空"`
	Types          int    `p:"types" v:"required#类型#1:上架,2:下架不能为空"`
	Status         int    `p:"status" v:"required#状态#1:启用,2:禁用不能为空"`
	Tupian         string `p:"tupian" `
	Other          int    `p:"other" `
	Sort           int    `p:"sort" v:"required#排序不能为空"`
	OtherStr       string `p:"otherStr" v:"required#其他状态2#s1:已激活,s2a2:未激活#IsDictColumn,IsSearchParams不能为空"`
	AppGoodsInfoId int    `p:"appGoodsInfoId" v:"required#appGoodsInfoID不能为空"`
	TestName       string `p:"testName" `
}
type UpdateReqParams struct {
	Id             int    `p:"id" v:"required#ID不能为空"`
	Title          string `p:"title" `
	SubTitle       string `p:"subTitle" `
	Pic            string `p:"pic" `
	Types          int    `p:"types" `
	Status         int    `p:"status" `
	Tupian         string `p:"tupian" `
	Other          int    `p:"other" `
	Sort           int    `p:"sort" `
	OtherStr       string `p:"otherStr" `
	AppGoodsInfoId int    `p:"appGoodsInfoId" `
	TestName       string `p:"testName" `
}
type InfoReqParams struct {
	Id int `p:"id" v:"required#id不能为空"`
}
type ListReqParams struct {
	Title    string `p:"title" `
	SubTitle string `p:"subTitle" `
	Types    int    `p:"types"  default:"-1"`
	Status   int    `p:"status"  default:"-1"`
	Other    int    `p:"other"  default:"-1"`
	OtherStr string `p:"otherStr"  default:"-1"`
}
type DeleteReqParams struct {
	Ids []int `p:"ids" v:"required#ids不能为空"`
}
