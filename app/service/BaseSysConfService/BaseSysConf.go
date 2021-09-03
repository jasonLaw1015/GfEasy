// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package BaseSysConfService

import (
	"goEasy/app/model/BaseSysConfModel"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var S = new(baseSysConfService)

type baseSysConfService struct {
}

// 查询单个信息
func (s *baseSysConfService) Info(id int) (data *BaseSysConfModel.Entity, err error) {

	err = BaseSysConfModel.M.Where("id=?", id).Scan(&data)

	return
}

//添加操作
func (s *baseSysConfService) Add(req *BaseSysConfModel.AddReqParams) (insId int64, err error) {

	r, err := BaseSysConfModel.M.OmitEmpty().Data(req).Insert()
	if err != nil {
		return 0, err
	}
	// 判断新增还是修改
	id, e := r.LastInsertId()
	return id, e
}

//修改操作
func (s *baseSysConfService) Update(req *BaseSysConfModel.UpdateReqParams) (Id int, err error) {

	_, err = BaseSysConfModel.M.OmitEmpty().Data(req).Where("id=?", req.Id).Update()
	if err != nil {
		return 0, err
	}
	id := gconv.Int(req.Id)
	return id, err
}

//分页查询
func (s *baseSysConfService) Page(req *BaseSysConfModel.PageReqParams) (total, page int, size int, list []*BaseSysConfModel.Entity, err error) {
	page = req.Page
	size = req.Size

	M := BaseSysConfModel.M

	if req.KeyWord != "" {

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
	list = make([]*BaseSysConfModel.Entity, len(data))
	err = data.Structs(&list)

	return
}

//List
func (s *baseSysConfService) List(condition g.Map) (list []*BaseSysConfModel.Entity, err error) {

	err = BaseSysConfModel.M.Where(condition).Scan(&list)
	return

}

//删除
func (s *baseSysConfService) Delete(ids []int) (err error) {
	_, err = BaseSysConfModel.M.Where("id IN(?)", ids).Delete()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除模型数据失败")
		return
	}
	return
}

//获得配置参数值
func (s *baseSysConfService) GetValue(key string) (data *BaseSysConfModel.Entity, err error) {

	err = BaseSysConfModel.M.Where("cKey=?", key).Scan(&data)

	return
}

/**
 * 更新配置参数
 * @param cKey
 * @param cValue
 */
func (s *baseSysConfService) UpdateValue(cKey string, cValue string) (err error) {
	data := g.Map{
		"cKey":   cKey,
		"cValue": cValue,
	}
	_, err = BaseSysConfModel.M.OmitEmpty().Data(data).Save()
	return
}
