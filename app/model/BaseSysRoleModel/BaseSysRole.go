// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package BaseSysRoleModel

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table baseSysRole.
type Entity struct {
	Id          int         `orm:"id,primary" json:"id"`           // ID
	CreateTime  *gtime.Time `orm:"createTime" json:"createTime"`   // 创建时间
	UpdateTime  *gtime.Time `orm:"updateTime" json:"updateTime"`   // 更新时间
	AdminUserId int         `orm:"adminUserId" json:"adminUserId"` // 创建角色的admin用户ID
	Name        string      `orm:"name" json:"name"`               // 名称
	Label       string      `orm:"label" json:"label"`             // 角色标签
	Remark      string      `orm:"remark" json:"remark"`           // 备注
	Relevance   int         `orm:"relevance" json:"relevance"`     // 数据权限是否关联上下级
}

var (
	// Table is the table name of baseSysRole.
	Table = "base_sys_role"
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
	AdminUserId      int      `p:"adminUserId" v:"required#创建角色的admin用户ID不能为空"`
	Name             string   `p:"name" v:"required#名称不能为空"`
	Label            string   `p:"label" `
	Remark           string   `p:"remark" `
	Relevance        int      `p:"relevance" v:"required#数据权限是否关联上下级不能为空"`
	MenuIdList       []string `p:"menuIdList" `
	DepartmentIdList []string `p:"departmentIdList"`
}
type MenuAndDepartmentListParams struct {
}
type UpdateReqParams struct {
	Id               int      `p:"id" v:"required#ID不能为空"`
	AdminUserId      int      `p:"adminUserId" `
	Name             string   `p:"name" `
	Label            string   `p:"label" `
	Remark           string   `p:"remark" `
	Relevance        int      `p:"relevance" `
	MenuIdList       []string `p:"menuIdList" `
	DepartmentIdList []string `p:"departmentIdList"`
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

type InfoResponse struct {
	Id               int         `orm:"id,primary" json:"id"`                     // ID
	CreateTime       *gtime.Time `orm:"createTime" json:"createTime"`             // 创建时间
	UpdateTime       *gtime.Time `orm:"updateTime" json:"updateTime"`             // 更新时间
	AdminUserId      string      `orm:"adminUserId" json:"adminUserId"`           // 用户ID
	Name             string      `orm:"name" json:"name"`                         // 名称
	Label            string      `orm:"label" json:"label"`                       // 角色标签
	Remark           string      `orm:"remark" json:"remark"`                     // 备注
	Relevance        int         `orm:"relevance" json:"relevance"`               // 数据权限是否关联上下级
	MenuIdList       []int       `orm:"menuIdList" json:"menuIdList"`             //关联的部门id
	DepartmentIdList []int       `orm:"departmentIdList" json:"departmentIdList"` //关联的菜单id
}
