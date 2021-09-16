// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package GenCodeLogService

import (
	"gfEasy/app/model/GenCodeLogModel"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var S = new(genCodeLogService)

type genCodeLogService struct {
}

// 查询单个信息
func (s *genCodeLogService) Info(id int) (data *GenCodeLogModel.Entity, err error) {

	err = GenCodeLogModel.M.Where("id=?", id).Scan(&data)

	return
}

//添加操作
func (s *genCodeLogService) Add(req *GenCodeLogModel.AddReqParams) (insId int64, err error) {

	r, err := GenCodeLogModel.M.OmitEmpty().Data(req).Insert()
	if err != nil {
		return 0, err
	}
	// 判断新增还是修改
	id, e := r.LastInsertId()
	return id, e
}

//修改操作
func (s *genCodeLogService) Update(req *GenCodeLogModel.UpdateReqParams) (Id int, err error) {

	_, err = GenCodeLogModel.M.OmitEmpty().Data(req).Where("id=?", req.Id).Update()
	if err != nil {
		return 0, err
	}
	id := gconv.Int(req.Id)
	return id, err
}

//分页查询
func (s *genCodeLogService) Page(req *GenCodeLogModel.PageReqParams) (total, page int, size int, list []*GenCodeLogModel.Entity, err error) {
	page = req.Page
	size = req.Size

	M := GenCodeLogModel.M

	if req.KeyWord != "" {
		M = M.Where("(title like ? OR description like ? OR  1 like ?)", "%"+req.KeyWord+"%", "%"+req.KeyWord+"%", "%"+req.KeyWord+"%")
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
	M = M.Order("createTime desc")

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
	list = make([]*GenCodeLogModel.Entity, len(data))
	err = data.Structs(&list)

	return
}

//List
func (s *genCodeLogService) List(condition g.Map) (list []*GenCodeLogModel.Entity, err error) {

	err = GenCodeLogModel.M.Where(condition).Scan(&list)
	return

}

//删除
func (s *genCodeLogService) Delete(ids []int) (err error) {
	_, err = GenCodeLogModel.M.Where("id IN(?)", ids).Delete()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除模型数据失败")
		return
	}
	return
}
