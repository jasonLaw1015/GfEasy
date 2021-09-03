// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package TaskInfoModel

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table taskInfo.
type Entity struct {
	Id          int         `orm:"id,primary" json:"id"`           // ID
	CreateTime  *gtime.Time `orm:"createTime" json:"createTime"`   // 创建时间
	UpdateTime  *gtime.Time `orm:"updateTime" json:"updateTime"`   // 更新时间
	JobId       string      `orm:"jobId" json:"jobId"`             // 任务ID
	RepeatConf  string      `orm:"repeatConf" json:"repeatConf"`   // 任务配置
	Name        string      `orm:"name" json:"name"`               // 名称
	Cron        string      `orm:"cron" json:"cron"`               // cron
	Limit       int         `orm:"limit" json:"limit"`             // 最大执行次数 不传为无限次
	Every       int         `orm:"every" json:"every"`             // 每间隔多少毫秒执行一次 如果cron设置了 这项设置就无效
	Remark      string      `orm:"remark" json:"remark"`           // 备注
	Status      int         `orm:"status" json:"status"`           // 状态 0:停止 1：运行
	StartDate   *gtime.Time `orm:"startDate" json:"startDate"`     // 开始时间
	EndDate     *gtime.Time `orm:"endDate" json:"endDate"`         // 结束时间
	Data        string      `orm:"data" json:"data"`               // 数据
	Service     string      `orm:"service" json:"service"`         // 执行的service实例ID
	Types       int         `orm:"types" json:"types"`             // 类型 0:系统 1：用户
	NextRunTime *gtime.Time `orm:"nextRunTime" json:"nextRunTime"` // 下一次执行时间
	TaskType    int         `orm:"taskType" json:"taskType"`       // 状态 0:cron 1：时间间隔
}

var (
	// Table is the table name of taskInfo.
	Table = "task_info"
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
	Status    int    `p:"status" default:"-1"`
	Types     int    `p:"types" default:"-1"`
	Order     string `p:"order"`
	Sort      string `p:"sort"`
	IsExport  bool   `p:"isExport"`
	StartTime string `p:"startTime"`
	EndTime   string `p:"endTime"`
}
type AddReqParams struct {
	JobId       string      `p:"jobId" `
	RepeatConf  string      `p:"repeatConf" `
	Name        string      `p:"name" v:"required#名称不能为空"`
	Cron        string      `p:"cron" `
	Limit       int         `p:"limit" `
	Every       int         `p:"every" `
	Remark      string      `p:"remark" `
	Status      int         `p:"status" v:"required#状态 2:停止 1：运行不能为空"`
	StartDate   *gtime.Time `p:"startDate" `
	EndDate     *gtime.Time `p:"endDate" `
	Data        string      `p:"data" `
	Service     string      `p:"service" `
	Types       int         `p:"types" v:"required#类型 2:系统 1：用户不能为空"`
	NextRunTime *gtime.Time `p:"nextRunTime" `
	TaskType    int         `p:"taskType" v:"required#状态 2:cron 1：时间间隔不能为空"`
}
type UpdateReqParams struct {
	Id          int         `p:"id" v:"required#ID不能为空"`
	JobId       string      `p:"jobId" `
	RepeatConf  string      `p:"repeatConf" `
	Name        string      `p:"name" `
	Cron        string      `p:"cron" `
	Limit       int         `p:"limit" `
	Every       int         `p:"every" `
	Remark      string      `p:"remark" `
	Status      int         `p:"status" `
	StartDate   *gtime.Time `p:"startDate" `
	EndDate     *gtime.Time `p:"endDate" `
	Data        string      `p:"data" `
	Service     string      `p:"service" `
	Types       int         `p:"types" `
	NextRunTime *gtime.Time `p:"nextRunTime" `
	TaskType    int         `p:"taskType" `
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
