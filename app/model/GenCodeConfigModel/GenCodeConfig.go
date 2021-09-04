// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package GenCodeConfigModel

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table genCodeConfig.
type Entity struct {
	Id                    int         `orm:"id,primary" json:"id"`                               // ID
	CreateTime            *gtime.Time `orm:"createTime" json:"createTime"`                       // 创建时间
	UpdateTime            *gtime.Time `orm:"updateTime" json:"updateTime"`                       // 更新时间
	DbHost                string      `orm:"dbHost" json:"dbHost"`                               // 数据库host
	DbPort                string      `orm:"dbPort" json:"dbPort"`                               // 数据库端口
	DbUser                string      `orm:"dbUser" json:"dbUser"`                               // 数据库用户
	DbPass                string      `orm:"dbPass" json:"dbPass"`                               // 数据库密码
	DbName                string      `orm:"dbName" json:"dbName"`                               // 数据库名
	DbType                string      `orm:"dbType" json:"dbType"`                               // 数据库类型，如mysql/pgsql/mssql/sqlite/oracle，默认mysql
	AdminGenPath          string      `orm:"adminGenPath" json:"adminGenPath"`                   // 生成代码前端目录
	ServerGenPath         string      `orm:"serverGenPath" json:"serverGenPath"`                 // 生成代码服务端目录
	DefaultCreatedAtLabel string      `orm:"defaultCreatedAtLabel" json:"defaultCreatedAtLabel"` // 默认的时间字段格式，如createTime
	DefaultUpdatedAtLabel string      `orm:"defaultUpdatedAtLabel" json:"defaultUpdatedAtLabel"` // 默认的时间字段格式，如updateTime
	IgnoreTablePrefix     string      `orm:"ignoreTablePrefix" json:"ignoreTablePrefix"`         // 忽略此前缀的表，逗号分隔
	TableNames            string      `orm:"tableNames" json:"tableNames"`                       // 表名，用逗号分隔
	ActiveCode            string      `orm:"activeCode" json:"activeCode"`                       // 激活码
}

var (
	// Table is the table name of genCodeConfig.
	Table = "gen_code_config"
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
	DbHost                string `p:"dbHost" v:"required#数据库host不能为空"`
	DbPort                string `p:"dbPort" v:"required#数据库端口不能为空"`
	DbUser                string `p:"dbUser" v:"required#数据库用户不能为空"`
	DbPass                string `p:"dbPass" v:"required#数据库密码不能为空"`
	DbName                string `p:"dbName" v:"required#数据库名不能为空"`
	DbType                string `p:"dbType" v:"required#数据库类型，如mysql/pgsql/mssql/sqlite/oracle，默认mysql不能为空"`
	AdminGenPath          string `p:"adminGenPath" v:"required#生成代码前端目录不能为空"`
	ServerGenPath         string `p:"ServerGenPath" v:"required#生成代码服务端目录不能为空"`
	DefaultCreatedAtLabel string `p:"defaultCreatedAtLabel" v:"required#默认的时间字段格式，如createTime不能为空"`
	DefaultUpdatedAtLabel string `p:"defaultUpdatedAtLabel" v:"required#默认的时间字段格式，如updateTime不能为空"`
	IgnoreTablePrefix     string `p:"ignoreTablePrefix" v:"required#忽略此前缀的表，逗号分隔不能为空"`
	TableNames            string `p:"tableNames"`
	ActiveCode            string `p:"activeCode" `
}
type UpdateReqParams struct {
	Id                    int    `p:"id" v:"required#ID不能为空"`
	DbHost                string `p:"dbHost" `
	DbPort                string `p:"dbPort" `
	DbUser                string `p:"dbUser" `
	DbPass                string `p:"dbPass" `
	DbName                string `p:"dbName" `
	DbType                string `p:"dbType" `
	AdminGenPath          string `p:"adminGenPath" `
	ServerGenPath         string `p:"ServerGenPath" `
	DefaultCreatedAtLabel string `p:"defaultCreatedAtLabel" `
	DefaultUpdatedAtLabel string `p:"defaultUpdatedAtLabel" `
	IgnoreTablePrefix     string `p:"ignoreTablePrefix" `
	TableNames            string `p:"tableNames" `
	ActiveCode            string `p:"activeCode" `
}

type SaveReqParams struct {
	Id                    int    `p:"id"`
	DbHost                string `p:"dbHost" v:"required#数据库host不能为空"`
	DbPort                string `p:"dbPort" v:"required#数据库端口不能为空"`
	DbUser                string `p:"dbUser" v:"required#数据库用户不能为空"`
	DbPass                string `p:"dbPass" v:"required#数据库密码不能为空"`
	DbName                string `p:"dbName" v:"required#数据库名不能为空"`
	DbType                string `p:"dbType" v:"required#数据库类型，如mysql/pgsql/mssql/sqlite/oracle，默认mysql不能为空"`
	AdminGenPath          string `p:"adminGenPath" v:"required#生成代码前端目录不能为空"`
	ServerGenPath         string `p:"ServerGenPath" v:"required#生成代码服务端目录不能为空"`
	DefaultCreatedAtLabel string `p:"defaultCreatedAtLabel" v:"required#默认的时间字段格式，如createTime不能为空"`
	DefaultUpdatedAtLabel string `p:"defaultUpdatedAtLabel" v:"required#默认的时间字段格式，如updateTime不能为空"`
	IgnoreTablePrefix     string `p:"ignoreTablePrefix" v:"required#忽略此前缀的表，逗号分隔不能为空"`
	ActiveCode            string `p:"activeCode" `
}
type InfoReqParams struct {
	Id int `p:"id" v:"required#id不能为空"`
}
type ListReqParams struct {
}
type DeleteReqParams struct {
	Ids []int `p:"ids" v:"required#ids不能为空"`
}
