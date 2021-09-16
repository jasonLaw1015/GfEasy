// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package BaseSysLogService

import (
	"gfEasy/app/model/BaseSysConfModel"
	"gfEasy/app/model/BaseSysLogModel"
	"gfEasy/app/service/BaseSysConfService"
	ContextService "gfEasy/app/service/ContextService"
	"gfEasy/library/utils"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

var S = new(baseSysLogService)

type baseSysLogService struct {
}

// 查询单个信息
func (s *baseSysLogService) Info(id int) (data *BaseSysLogModel.Entity, err error) {

	err = BaseSysLogModel.M.Where("id=?", id).Scan(&data)

	return
}

//添加操作
func (s *baseSysLogService) Add(req *BaseSysLogModel.AddReqParams) (insId int64, err error) {

	r, err := BaseSysLogModel.M.OmitEmpty().Data(req).Insert()
	if err != nil {
		return 0, err
	}
	// 判断新增还是修改
	id, e := r.LastInsertId()
	return id, e
}

//修改操作
func (s *baseSysLogService) Update(req *BaseSysLogModel.UpdateReqParams) (Id int, err error) {

	_, err = BaseSysLogModel.M.OmitEmpty().Data(req).Where("id=?", req.Id).Update()
	if err != nil {
		return 0, err
	}
	id := gconv.Int(req.Id)
	return id, err
}

//分页查询
func (s *baseSysLogService) Page(req *BaseSysLogModel.PageReqParams) (total, page int, size int, list []*BaseSysLogModel.Entity, err error) {
	page = req.Page
	size = req.Size

	M := BaseSysLogModel.M

	if req.KeyWord != "" {
		M = M.Where("(action like ? OR ip like ? OR params like ? OR 1 like ?)", "%"+req.KeyWord+"%", "%"+req.KeyWord+"%", "%"+req.KeyWord+"%", "%"+req.KeyWord+"%")
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
	list = make([]*BaseSysLogModel.Entity, len(data))
	err = data.Structs(&list)

	return
}

//List
func (s *baseSysLogService) List(condition g.Map) (list []*BaseSysLogModel.Entity, err error) {

	err = BaseSysLogModel.M.Where(condition).Scan(&list)
	return

}

//删除
func (s *baseSysLogService) Delete(ids []int) (err error) {
	_, err = BaseSysLogModel.M.Where("id IN(?)", ids).Delete()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除模型数据失败")
		return
	}
	return
}

//记录日记
func (s *baseSysLogService) Record(r *ghttp.Request) (insId int64, err error) {
	var req *BaseSysLogModel.AddReqParams
	adminUser := ContextService.S.Get(r.GetCtx()).User
	var params string
	if r.Request.URL.Path != "/admin/baseComm/upload" {
		if r.Request.Method == "POST" {
			params = r.GetBodyString()
		} else {
			params, err = utils.Query2JsonString(r.Request.URL.RawQuery)
		}
		//对密码字段做下处理
		if gstr.Contains(params, "password") {
			paramsObj := gjson.New(params)
			paramsObj.Set("password", "******")
			params = paramsObj.MustToJsonString()
		}
	}

	data := g.Map{
		"adminUserId": adminUser.AdminUserId,
		"action":      r.Request.URL.Path,
		"ip":          r.GetClientIp(),
		"ipAddr":      r.RemoteAddr, //todo 以后再获取真实省份地址
		"params":      params,
	}
	gconv.Struct(data, &req)
	id, e := s.Add(req)
	return id, e
}

//清空日志
func (s *baseSysLogService) Clear(isAll bool) (err error) {
	if isAll {
		BaseSysLogModel.M.Where(1).Delete()
		return
	}
	conf, err := s.GetKeep("logKeep")
	if conf.CValue != "" {
		dayNum := gconv.Int(conf.CValue)

		//beforeDate := time.Now().AddDate(0, 0, -dayNum).Format("Y-m-d 00:00:00")
		beforeDate := gtime.Now().AddDate(0, 0, -dayNum).Format("Y-m-d 00:00:00")
		_, err1 := BaseSysLogModel.M.Where("createTime <=?", beforeDate).Delete()
		return err1
	}
	return nil
}

//设置日志保存时间
func (s *baseSysLogService) SetKeep(cKey string, cValue string) (err error) {
	err = BaseSysConfService.S.UpdateValue(cKey, cValue)
	if err != nil {
		g.Log().Error(err)
		return
	}
	return
}

//获得日志保存时间
func (s *baseSysLogService) GetKeep(key string) (data *BaseSysConfModel.Entity, err error) {
	data, err = BaseSysConfService.S.GetValue(key)

	if err != nil {
		g.Log().Error(err)
		return
	}
	return
}
