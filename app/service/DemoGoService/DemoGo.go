// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package DemoGoService

import (
	"goEasy/app/model/DemoGoModel"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var S = new(demoGoService)

type demoGoService struct {
}

// 查询单个信息
func (s *demoGoService) Info(id int) (data *DemoGoModel.Entity, err error) {

	err = DemoGoModel.M.Where("id=?", id).Scan(&data)

	return
}

//添加操作
func (s *demoGoService) Add(req *DemoGoModel.AddReqParams) (insId int64, err error) {

	r, err := DemoGoModel.M.OmitEmpty().Data(req).Insert()
	if err != nil {
		return 0, err
	}
	// 判断新增还是修改
	id, e := r.LastInsertId()
	return id, e
}

//修改操作
func (s *demoGoService) Update(req *DemoGoModel.UpdateReqParams) (Id int, err error) {

	_, err = DemoGoModel.M.OmitEmpty().Data(req).Where("id=?", req.Id).Update()
	if err != nil {
		return 0, err
	}
	id := gconv.Int(req.Id)
	return id, err
}

//分页查询
func (s *demoGoService) Page(req *DemoGoModel.PageReqParams) (total, page int, size int, list []*DemoGoModel.Entity, err error) {
	page = req.Page
	size = req.Size

	M := DemoGoModel.M

	if req.Types != -1 {
		M = M.Where("types=?", req.Types)
	}

	if req.Status != -1 {
		M = M.Where("status=?", req.Status)
	}

	if req.Other != -1 {
		M = M.Where("other=?", req.Other)
	}

	if req.KeyWord != "" {
		M = M.Where("(title like ? OR  1 like ?)", "%"+req.KeyWord+"%", "%"+req.KeyWord+"%")
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
	list = make([]*DemoGoModel.Entity, len(data))
	err = data.Structs(&list)

	return
}

//List
func (s *demoGoService) List(condition g.Map) (list []*DemoGoModel.Entity, err error) {

	err = DemoGoModel.M.Where(condition).Scan(&list)
	return

}

//删除
func (s *demoGoService) Delete(ids []int) (err error) {
	_, err = DemoGoModel.M.Where("id IN(?)", ids).Delete()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除模型数据失败")
		return
	}
	return
}
