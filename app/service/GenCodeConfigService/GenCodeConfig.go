// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package GenCodeConfigService

import (
	"gfEasy/app/model/GenCodeConfigModel"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var S = new(genCodeConfigService)

type genCodeConfigService struct {
}

// 查询单个信息
func (s *genCodeConfigService) FirstInfo() (data *GenCodeConfigModel.Entity, err error) {

	err = GenCodeConfigModel.M.Order("id asc").Limit(1).Scan(&data)

	return
}

// 查询单个信息
func (s *genCodeConfigService) Info(id int) (data *GenCodeConfigModel.Entity, err error) {

	err = GenCodeConfigModel.M.Where("id=?", id).Scan(&data)

	return
}

//添加操作
func (s *genCodeConfigService) Add(req *GenCodeConfigModel.AddReqParams) (insId int64, err error) {

	r, err := GenCodeConfigModel.M.OmitEmpty().Data(req).Insert()
	if err != nil {
		return 0, err
	}
	// 判断新增还是修改
	id, e := r.LastInsertId()
	return id, e
}

//修改操作
func (s *genCodeConfigService) Update(req *GenCodeConfigModel.UpdateReqParams) (Id int, err error) {

	_, err = GenCodeConfigModel.M.OmitEmpty().Data(req).Where("id=?", req.Id).Update()
	if err != nil {
		return 0, err
	}
	id := gconv.Int(req.Id)
	return id, err
}

//修改操作
func (s *genCodeConfigService) Save(req *GenCodeConfigModel.SaveReqParams) (res interface{}, err error) {
	if req.Id == 0 {
		res, err = GenCodeConfigModel.M.Data(req).Insert()
	} else {
		res, err = GenCodeConfigModel.M.Data(req).Where("id =?", req.Id).Update()
	}

	if err != nil {
		return 0, err
	}

	return res, err
}

//分页查询
func (s *genCodeConfigService) Page(req *GenCodeConfigModel.PageReqParams) (total, page int, size int, list []*GenCodeConfigModel.Entity, err error) {
	page = req.Page
	size = req.Size

	M := GenCodeConfigModel.M

	if req.KeyWord != "" {
		M = M.Where("( 1 like ?)", "%"+req.KeyWord+"%")
	}
	if req.StartTime != "" {
		M = M.WhereGTE("createTime", req.StartTime)
	}
	if req.EndTime != "" {
		M = M.WhereLTE("createTime", req.EndTime)
	}
	if req.Order != "" && req.Sort != "" {
		M = M.Order(req.Order + " " + req.Sort)
	}

	total, err = M.Count()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总行数失败")
		return
	}
	//是否是导出excel，不用分页查询
	if req.IsExport {
		exportData, err1 := M.All()
		if err1 != nil {
			g.Log().Error(err1)
			err = gerror.New("获取导出数据失败")
			return
		}
		err = exportData.Structs(&list)
		return
	}

	M = M.Page(req.Page, req.Size)

	data, err := M.All()

	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}
	list = make([]*GenCodeConfigModel.Entity, len(data))
	err = data.Structs(&list)

	return
}

//List
func (s *genCodeConfigService) List(condition g.Map) (list []*GenCodeConfigModel.Entity, err error) {

	err = GenCodeConfigModel.M.Where(condition).Scan(&list)
	return

}

//删除
func (s *genCodeConfigService) Delete(ids []int) (err error) {
	_, err = GenCodeConfigModel.M.Where("id IN(?)", ids).Delete()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除模型数据失败")
		return
	}
	return
}
