// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package TaskLogModel

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table taskLog.
type Entity struct {
	Id         int         `orm:"id,primary" json:"id"`         // ID
	CreateTime *gtime.Time `orm:"createTime" json:"createTime"` // 创建时间
	UpdateTime *gtime.Time `orm:"updateTime" json:"updateTime"` // 更新时间
	TaskId     int64       `orm:"taskId" json:"taskId"`         // 任务ID
	Status     int         `orm:"status" json:"status"`         // 状态 0:失败 1：成功
	Detail     string      `orm:"detail" json:"detail"`         // 详情描述
}

var (
	// Table is the table name of taskLog.
	Table = "task_log"
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
	Id        int    `p:"id" default:"-1"`
	KeyWord   string `p:"keyWord"`
	Page      int    `p:"page" default:"1"`
	Size      int    `p:"size" default:"20"`
	Status    int    `p:"status" default:"-1"`
	Order     string `p:"order"`
	Sort      string `p:"sort"`
	IsExport  bool   `p:"isExport"`
	StartTime string `p:"startTime"`
	EndTime   string `p:"endTime"`
}
type AddReqParams struct {
	TaskId int64  `p:"taskId" `
	Status int    `p:"status" v:"required#状态 2:失败 1：成功不能为空"`
	Detail string `p:"detail" `
}
type UpdateReqParams struct {
	Id     int    `p:"id" v:"required#ID不能为空"`
	TaskId int64  `p:"taskId" `
	Status int    `p:"status" `
	Detail string `p:"detail" `
}
type InfoReqParams struct {
	Id int `p:"id" v:"required#id不能为空"`
}
type ListReqParams struct {
	Status int `p:"status"`
}
type DeleteReqParams struct {
	Ids []int `p:"ids" v:"required#ids不能为空"`
}

type LogPageResponse struct {
	Id         int         `orm:"id,primary" json:"id"`         // ID
	CreateTime *gtime.Time `orm:"createTime" json:"createTime"` // 创建时间
	UpdateTime *gtime.Time `orm:"updateTime" json:"updateTime"` // 更新时间
	TaskId     int64       `orm:"taskId" json:"taskId"`         // 任务ID
	Status     int         `orm:"status" json:"status"`         // 状态 2:失败 1：成功
	Detail     string      `orm:"detail" json:"detail"`         // 详情描述
	TaskName   string      `orm:"taskName" json:"taskName"`     // 任务名
}
