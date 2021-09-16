package GenTableModel

import (
	"gfEasy/app/model/GenTableColumnModel"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table genTable.
type Entity struct {
	Id             int64       `orm:"id,primary" json:"id"`                  // 编号
	TableName      string      `orm:"tableName"       json:"tableName"`      // 表名称
	TableComment   string      `orm:"tableComment"    json:"tableComment"`   // 表描述
	ClassName      string      `orm:"className"       json:"className"`      // 实体类名称
	TplCategory    string      `orm:"tplCategory"     json:"tplCategory"`    // 使用的模板（crud单表操作 tree树表操作）
	PackageName    string      `orm:"packageName"     json:"packageName"`    // 生成包路径
	ModuleName     string      `orm:"moduleName"      json:"moduleName"`     // 生成模块名
	BusinessName   string      `orm:"businessName"    json:"businessName"`   // 生成业务名
	FunctionName   string      `orm:"functionName"    json:"functionName"`   // 生成功能名
	FunctionAuthor string      `orm:"functionAuthor"  json:"functionAuthor"` // 生成功能作者
	Options        string      `orm:"options"          json:"options"`       // 其它生成选项
	CreateBy       string      `orm:"createBy"        json:"createBy"`       // 创建者
	CreateTime     *gtime.Time `orm:"createTime"      json:"createTime"`     // 创建时间
	UpdateBy       string      `orm:"updateBy"        json:"updateBy"`       // 更新者
	UpdateTime     *gtime.Time `orm:"updateTime"      json:"updateTime"`     // 更新时间
	Remark         string      `orm:"remark"           json:"remark"`        // 备注
}

//实体扩展
type EntityExtend struct {
	Id             int64                            `orm:"tableId,primary" json:"tableId"`        // 编号
	TableName      string                           `orm:"tableName"       json:"tableName"`      // 表名称
	TableComment   string                           `orm:"tableComment"    json:"tableComment"`   // 表描述
	ClassName      string                           `orm:"className"       json:"className"`      // 实体类名称
	TplCategory    string                           `orm:"tplCategory"     json:"tplCategory"`    // 使用的模板（crud单表操作 tree树表操作）
	PackageName    string                           `orm:"packageName"     json:"packageName"`    // 生成包路径
	ModuleName     string                           `orm:"moduleName"      json:"moduleName"`     // 生成模块名
	BusinessName   string                           `orm:"businessName"    json:"businessName"`   // 生成业务名
	FunctionName   string                           `orm:"functionName"    json:"functionName"`   // 生成功能名
	FunctionAuthor string                           `orm:"functionAuthor"  json:"functionAuthor"` // 生成功能作者
	Options        string                           `orm:"options"         json:"options"`        // 其它生成选项
	CreateBy       string                           `orm:"createBy"        json:"createBy"`       // 创建者
	CreateTime     *gtime.Time                      `orm:"createTime"      json:"createTime"`     // 创建时间
	UpdateBy       string                           `orm:"updateBy"        json:"updateBy"`       // 更新者
	UpdateTime     *gtime.Time                      `orm:"updateTime"      json:"updateTime"`     // 更新时间
	Remark         string                           `orm:"remark"           json:"remark"`        // 备注
	TreeCode       string                           `json:"treeCode"`                             // 树编码字段
	TreeParentCode string                           `json:"treeParentCode"`                       // 树父编码字段
	TreeName       string                           `json:"treeName"`                             // 树名称字段
	Columns        []*GenTableColumnModel.ResEntity `json:"columns"`                              // 表列信息
	PkColumn       *GenTableColumnModel.ResEntity   `json:"pkColumn"`                             // 表列信息
}

//数据表需要的信息
type TableStructEntity struct {
	TableName    string                           `json:"tableName"`    // 表名称
	TableComment string                           `json:"tableComment"` // 表描述
	Columns      []*GenTableColumnModel.ResEntity `json:"Columns"`
}

var (
	// Table is the table name of goods.
	Table = "gen_table"
	M     = g.DB("default").Model(Table).Safe()
)

//修改页面请求参数
type EditReqParams struct {
	Id             int64  `p:"tableId" v:"required#主键ID不能为空"`
	TableName      string `p:"tableName"  v:"required#表名称不能为空"`
	TableComment   string `p:"tableComment"  v:"required#表描述不能为空"`
	ClassName      string `p:"className" v:"required#实体类名称不能为空"`
	FunctionAuthor string `p:"functionAuthor"  v:"required#作者不能为空"`
	TplCategory    string `p:"tplCategory"`
	PackageName    string `p:"packageName" v:"required#生成包路径不能为空"`
	ModuleName     string `p:"moduleName" v:"required#生成模块名不能为空"`
	BusinessName   string `p:"businessName" v:"required#生成业务名不能为空"`
	FunctionName   string `p:"functionName" v:"required#生成功能名不能为空"`
	Remark         string `p:"remark"`
	Params         string `p:"params"`
	Columns        string `p:"columns"`
	TreeCode       string `p:"treeCode"`
	TreeParentCode string `p:"treeParentCode"`
	TreeName       string `p:"treeName"`
	UserName       string
}

type GenCodeReq struct {
	TableNames            string `p:"tableNames" v:"required#请选择要生成的表"`
	ServerGenPath         string `p:"serverGenPath" default:"./" v:"required#服务端所在文件夹地址，如/Users/yons/Documents/develop/go/src/gitee.com/jasonLaw1015/gfEasy/app/controller/Common/GenTable.go"`
	AdminGenPath          string `p:"adminGenPath" default:"./" v:"required#admin文件夹地址"`
	DbHost                string `p:"dbHost" v:"required#需要数据库host"`
	DbPort                string `p:"dbPort" v:"required#需要数据库port"`
	DbUser                string `p:"dbUser" v:"required#需要数据库user"`
	DbPass                string `p:"dbPass" v:"required#需要数据库user"`
	DbName                string `p:"dbName" v:"required#需要数据库名"`
	DbType                string `p:"dbType" default:"mysql" v:"required#需要数据库类型，如mysql"`
	DefaultCreatedAtLabel string `p:"defaultCreatedAtLabel" v:"required#创建时间的命名"`
	DefaultUpdatedAtLabel string `p:"defaultUpdatedAtLabel" v:"required#更新时间的命名"`
	IgnoreTablePrefix     string `p:"ignoreTablePrefix" v:"required#更新时间的命名" default:"base_,gen_,task_"`
}

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
	Page    int    `p:"page"`
	Size    int    `p:"size"`
	Sort    string `p:"sort"`
	Types   string `p:"types"`
	Order   string `p:"order"`
}

type AddReqParams struct {
	Name      string `p:"name" v:"required#商品名称不能为空"`
	Inventory int    `p:"inventory" v:"required#库存不能为空"`
}

type UpdateReqParams struct {
	Id        int    `p:"id" v:"required#id不能为空"`
	Name      string `p:"name" v:"required#商品名称不能为空"`
	Inventory int    `p:"inventory" v:"required#库存不能为空"`
}

type InfoReqParams struct {
	Id int `p:"id" v:"required#id不能为空"`
}

/**
condition := g.Map{
    "title like ?"         : "%九寨%",
    "online"               : 1,
    "hits between ? and ?" : g.Slice{1, 10},
    "exp > 0"              : nil,
    "category"             : g.Slice{100, 200},
}
// SELECT * FROM article WHERE title like '%九寨%' AND online=1 AND hits between 1 and 10 AND exp > 0 AND category IN(100,200)
db.Model("article").Where(condition).All()
*/
type ListReqParams struct {
	Types int    `p:"types"`
	Name  string `p:"name"`
}

type DeleteReqParams struct {
	Ids []int `p:"ids" v:"required#ids不能为空"`
}
