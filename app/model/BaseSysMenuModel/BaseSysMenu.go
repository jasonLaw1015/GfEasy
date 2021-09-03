// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package BaseSysMenuModel

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table baseSysMenu.
type Entity struct {
	Id         int         `orm:"id,primary" json:"id"`         // ID
	CreateTime *gtime.Time `orm:"createTime" json:"createTime"` // 创建时间
	UpdateTime *gtime.Time `orm:"updateTime" json:"updateTime"` // 更新时间
	ParentId   int64       `orm:"parentId" json:"parentId"`     // 父菜单ID
	Name       string      `orm:"name" json:"name"`             // 菜单名称
	Router     string      `orm:"router" json:"router"`         // 菜单地址
	Perms      string      `orm:"perms" json:"perms"`           // 权限标识
	Type       int         `orm:"type" json:"type"`             // 类型 0：目录 1：菜单 2：按钮
	Icon       string      `orm:"icon" json:"icon"`             // 图标
	OrderNum   int         `orm:"orderNum" json:"orderNum"`     // 排序
	ViewPath   string      `orm:"viewPath" json:"viewPath"`     // 视图地址
	KeepAlive  int         `orm:"keepAlive" json:"keepAlive"`   // 路由缓存
	IsShow     int         `orm:"isShow" json:"isShow"`         // 父菜单名称
}

var (
	// Table is the table name of baseSysMenu.
	Table = "base_sys_menu"
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
	ParentId  int64  `p:"parentId" `
	Name      string `p:"name" v:"required#菜单名称不能为空"`
	Router    string `p:"router" `
	Perms     string `p:"perms" `
	Type      int    `p:"type" v:"required#类型 0：目录 1：菜单 2：按钮不能为空"`
	Icon      string `p:"icon" `
	OrderNum  int    `p:"orderNum" v:"required#排序不能为空"`
	ViewPath  string `p:"viewPath" `
	KeepAlive int    `p:"keepAlive"`
	IsShow    int    `p:"isShow"`
}
type UpdateReqParams struct {
	Id        int    `p:"id" v:"required#ID不能为空"`
	ParentId  int64  `p:"parentId" `
	Name      string `p:"name" `
	Router    string `p:"router" `
	Perms     string `p:"perms" `
	Type      int    `p:"type" `
	Icon      string `p:"icon" `
	OrderNum  int    `p:"orderNum" `
	ViewPath  string `p:"viewPath" `
	KeepAlive int    `p:"keepAlive" `
	IsShow    int    `p:"isShow" `
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
