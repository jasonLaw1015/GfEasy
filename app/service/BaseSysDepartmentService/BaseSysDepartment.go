// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package BaseSysDepartmentService

import (
	"gfEasy/app/model/BaseSysDepartmentModel"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var S = new(baseSysDepartmentService)

type baseSysDepartmentService struct {
}

// 查询单个信息
func (s *baseSysDepartmentService) Info(id int) (data *BaseSysDepartmentModel.Entity, err error) {

	err = BaseSysDepartmentModel.M.Where("id=?", id).Scan(&data)

	return
}

//添加操作
func (s *baseSysDepartmentService) Add(req *BaseSysDepartmentModel.AddReqParams) (insId int64, err error) {

	r, err := BaseSysDepartmentModel.M.OmitEmpty().Data(req).Insert()
	if err != nil {
		return 0, err
	}
	// 判断新增还是修改
	id, e := r.LastInsertId()
	return id, e
}

//修改操作
func (s *baseSysDepartmentService) Update(req *BaseSysDepartmentModel.UpdateReqParams) (Id int, err error) {

	_, err = BaseSysDepartmentModel.M.OmitEmpty().Data(req).Where("id=?", req.Id).Update()
	if err != nil {
		return 0, err
	}
	id := gconv.Int(req.Id)
	return id, err
}

//分页查询
func (s *baseSysDepartmentService) Page(req *BaseSysDepartmentModel.PageReqParams) (total, page int, size int, list []*BaseSysDepartmentModel.Entity, err error) {
	page = req.Page
	size = req.Size

	M := BaseSysDepartmentModel.M

	if req.KeyWord != "" {
		M = M.WhereLike("name", "%"+req.KeyWord+"%")

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
	list = make([]*BaseSysDepartmentModel.Entity, len(data))
	err = data.Structs(&list)

	return
}

//List
func (s *baseSysDepartmentService) List(condition g.Map) (list []*BaseSysDepartmentModel.Entity, err error) {

	err = BaseSysDepartmentModel.M.Where(condition).Scan(&list)
	return

}

//删除
func (s *baseSysDepartmentService) Delete(ids []int) (err error) {
	_, err = BaseSysDepartmentModel.M.Where("id IN(?)", ids).Delete()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除模型数据失败")
		return
	}
	return
}
