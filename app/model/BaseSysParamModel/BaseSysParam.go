// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package BaseSysParamModel

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table baseSysParam.
type Entity struct {
	Id         int         `orm:"id,primary" json:"id"`         // ID
	CreateTime *gtime.Time `orm:"createTime" json:"createTime"` // 创建时间
	UpdateTime *gtime.Time `orm:"updateTime" json:"updateTime"` // 更新时间
	KeyName    string      `orm:"keyName" json:"keyName"`       // 键位
	Name       string      `orm:"name" json:"name"`             // 名称
	Data       string      `orm:"data" json:"data"`             // 数据
	DataType   int         `orm:"dataType" json:"dataType"`     // 数据类型 0:字符串 1：数组 2：键值对
	Remark     string      `orm:"remark" json:"remark"`         // 备注
}

var (
	// Table is the table name of baseSysParam.
	Table = "base_sys_param"
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
	KeyName  string `p:"keyName" v:"required#键位不能为空"`
	Name     string `p:"name" v:"required#名称不能为空"`
	Data     string `p:"data" v:"required#数据不能为空"`
	DataType int    `p:"dataType" v:"required#数据类型 0:字符串 1：数组 2：键值对不能为空"`
	Remark   string `p:"remark" `
}
type UpdateReqParams struct {
	Id       int    `p:"id" v:"required#ID不能为空"`
	KeyName  string `p:"keyName" `
	Name     string `p:"name" `
	Data     string `p:"data" `
	DataType int    `p:"dataType" `
	Remark   string `p:"remark" `
}
type InfoReqParams struct {
	Id int `p:"id" v:"required#id不能为空"`
}
type ListReqParams struct {
	Name string `p:"name"`
}
type DeleteReqParams struct {
	Ids []int `p:"ids" v:"required#ids不能为空"`
}
