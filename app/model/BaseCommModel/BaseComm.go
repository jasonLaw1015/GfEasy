// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package BaseCommModel

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table baseComm.
type Entity struct {
	Id         int         `orm:"id,primary" json:"id"`         // ID
	CreateTime *gtime.Time `orm:"createTime" json:"createTime"` // 创建时间
	UpdateTime *gtime.Time `orm:"updateTime" json:"updateTime"` // 更新时间
}

var (
	// Table is the table name of baseComm.
	Table = "base_comm"
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
type PersonUpdateParams struct {
	NickName    string `p:"nickName" v:"required#昵称不能为空"`
	HeadImg     string `p:"headImg" v:"required#头像不能为空"`
	NowPassword string `p:"nowPassword" v:"required#当前密码不能为空"`
	NewPassword string `p:"newPassword" v:"required#新密码不能为空"`
}
type InfoReqParams struct {
	Id int `p:"id" v:"required#id不能为空"`
}
type ListReqParams struct {
}
type DeleteReqParams struct {
	Ids []int `p:"ids" v:"required#ids不能为空"`
}
