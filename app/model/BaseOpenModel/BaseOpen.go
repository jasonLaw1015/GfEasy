// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package BaseOpenModel

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table baseOpen.
type Entity struct {
	Id         int         `orm:"id,primary" json:"id"`         // ID
	CreateTime *gtime.Time `orm:"createTime" json:"createTime"` // 创建时间
	UpdateTime *gtime.Time `orm:"updateTime" json:"updateTime"` // 更新时间
}

var (
	// Table is the table name of baseOpen.
	Table = "base_open"
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
}
type UpdateReqParams struct {
	Id int `p:"id" v:"required#ID不能为空"`
}
type InfoReqParams struct {
	Id int `p:"id" v:"required#id不能为空"`
}

type SysParamsReqParams struct {
	Name string `p:"name" v:"required#name不能为空"`
}
type ListReqParams struct {
}
type DeleteReqParams struct {
	Ids []int `p:"ids" v:"required#ids不能为空"`
}

type LoginReqParams struct {
	UserName   string `p:"username" v:"required#用户名不能为空"`
	Password   string `p:"password" v:"required#密码不能为空"`
	CaptchaId  string `p:"CaptchaId" v:"required#验证码ID不能为空"`
	VerifyCode string `p:"VerifyCode" v:"required#验证码不能为空"`
}
