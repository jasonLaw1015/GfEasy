// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package TaskLogService

import (
	"goEasy/app/model/TaskLogModel"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var S = new(taskLogService)

type taskLogService struct {
}

// 查询单个信息
func (s *taskLogService) Info(id int) (data *TaskLogModel.Entity, err error) {

	err = TaskLogModel.M.Where("id=?", id).Scan(&data)

	return
}

//添加操作
func (s *taskLogService) Add(req *TaskLogModel.AddReqParams) (insId int64, err error) {

	r, err := TaskLogModel.M.OmitEmpty().Data(req).Insert()
	if err != nil {
		return 0, err
	}
	// 判断新增还是修改
	id, e := r.LastInsertId()
	return id, e
}

//修改操作
func (s *taskLogService) Update(req *TaskLogModel.UpdateReqParams) (Id int, err error) {

	_, err = TaskLogModel.M.OmitEmpty().Data(req).Where("id=?", req.Id).Update()
	if err != nil {
		return 0, err
	}
	id := gconv.Int(req.Id)
	return id, err
}

//分页
func (s *taskLogService) Page(req *TaskLogModel.PageReqParams) (total, page int, size int, list []*TaskLogModel.LogPageResponse, err error) {
	page = req.Page
	size = req.Size

	M := g.DB("default").Model("task_log a").Safe()

	if req.KeyWord != "" {

	}
	if req.Id != -1 {
		M = M.Where("a.taskId", req.Id)
	}
	if req.Status != -1 {
		M = M.Where("a.status", req.Status)
	}
	if req.StartTime != "" {
		M = M.WhereGTE("a.createTime", req.StartTime)
	}
	if req.EndTime != "" {
		M = M.WhereLTE("a.createTime", req.EndTime)
	}
	if req.Order != "" && req.Sort != "" {
		M = M.Order("a." + req.Order + " " + req.Sort)
	}

	total, err = M.Count()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总行数失败")
		return
	}
	//指定filed
	M = M.LeftJoin("task_info b", "a.taskId=b.id").Fields("a.*,b.name AS taskName")
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
	list = make([]*TaskLogModel.LogPageResponse, 0)
	data, err := M.All()
	err = data.Structs(&list)

	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}
	return
}

//List
func (s *taskLogService) List(condition g.Map) (list []*TaskLogModel.Entity, err error) {

	err = TaskLogModel.M.Where(condition).Scan(&list)
	return

}

//删除
func (s *taskLogService) Delete(ids []int) (err error) {
	_, err = TaskLogModel.M.Where("id IN(?)", ids).Delete()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除模型数据失败")
		return
	}
	return
}
