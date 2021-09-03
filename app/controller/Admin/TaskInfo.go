// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Admin

import (
	"goEasy/app/model/TaskInfoModel"
	"goEasy/app/model/TaskLogModel"
	"goEasy/app/service/TaskInfoService"
	"goEasy/app/service/TaskLogService"
	"goEasy/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var TaskInfo = new(taskInfo)

//控制器
type taskInfo struct{}

// @Summary 添加信息
// @Description 添加信息
// @Tags Admin-TaskInfo-admin定时任务表管理
// @Param data body TaskInfoModel.AddReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/taskInfo/add [post]
// @Security
func (c *taskInfo) Add(r *ghttp.Request) {
	var req *TaskInfoModel.AddReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := TaskInfoService.S.Add(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "新增成功", g.Map{"id": id})
}

// @Summary 修改信息
// @Description 修改信息
// @Tags Admin-TaskInfo-admin定时任务表管理
// @Param data body TaskInfoModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/taskInfo/update [post]
// @Security
func (c *taskInfo) Update(r *ghttp.Request) {
	var req *TaskInfoModel.UpdateReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	id, err := TaskInfoService.S.Update(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "更新成功", g.Map{"id": id})
}

// @Summary 删除信息
// @Description 删除信息
// @Tags Admin-TaskInfo-admin定时任务表管理
// @Param ids body TaskInfoModel.DeleteReqParams  true "ids[1,2,3..]"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/taskInfo/delete [post]
// @Security
func (c *taskInfo) Delete(r *ghttp.Request) {
	var req *TaskInfoModel.DeleteReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err := TaskInfoService.S.Delete(req.Ids)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "删除成功")
}

// @Summary 单个信息
// @Description 单个信息
// @Tags Admin-TaskInfo-admin定时任务表管理
// @Param data body TaskInfoModel.InfoReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/taskInfo/info [get]
// @Security
func (c *taskInfo) Info(r *ghttp.Request) {
	var req *TaskInfoModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	result, err := TaskInfoService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取信息成功", result)
}

// @Summary 信息列表
// @Description 信息列表
// @Tags Admin-TaskInfo-admin定时任务表管理
// @Param data body TaskInfoModel.ListReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/taskInfo/list [get]
// @Security
func (c *taskInfo) List(r *ghttp.Request) {
	var req *TaskInfoModel.ListReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	condition := g.Map{}

	if req.Name != "" {
		condition["name = ?"] = req.Name
	}

	if req.Status != -1 {
		condition["status = ?"] = req.Status
	}

	if req.Types != -1 {
		condition["types = ?"] = req.Types
	}

	result, err := TaskInfoService.S.List(condition)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "获取列表信息成功", result)
}

// @Summary 分页信息
// @Description 分页信息
// @Tags Admin-TaskInfo-admin定时任务表管理
// @Param data body TaskInfoModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/taskInfo/page [get]
// @Security
func (c *taskInfo) Page(r *ghttp.Request) {
	var req *TaskInfoModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := TaskInfoService.S.Page(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	pagination := g.Map{
		"page":  page,
		"size":  size,
		"total": total,
	}
	response.SusJson(true, r, "获取分页信息成功", g.Map{
		"total":      total,
		"list":       list,
		"pagination": pagination,
	})
}

// @Summary Log分页信息
// @Description Log分页信息
// @Tags Admin-TaskLog-Admin定时任务日志表管理
// @Param data body TaskLogModel.PageReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/taskInfo/log [get]
// @Security
func (c *taskInfo) Log(r *ghttp.Request) {
	var req *TaskLogModel.PageReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	total, page, size, list, err := TaskLogService.S.Page(req)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	pagination := g.Map{
		"page":  page,
		"size":  size,
		"total": total,
	}
	response.SusJson(true, r, "获取task日志信息成功", g.Map{
		"total":      total,
		"list":       list,
		"pagination": pagination,
	})
}

// @Summary 开始任务
// @Description 开始任务
// @Tags Admin-TaskInfo-admin定时任务表管理
// @Param data body TaskInfoModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/taskInfo/start [post]
// @Security
func (c *taskInfo) Start(r *ghttp.Request) {
	var req *TaskInfoModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}

	_, err := TaskInfoModel.M.Data(g.Map{
		"status": 1,
	}).Where("id=?", req.Id).Update()
	if err != nil {
		response.FailJson(true, r, err.Error())
	}

	entity, err := TaskInfoService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}

	err = TaskInfoService.S.Start(entity)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}

	response.SusJson(true, r, "成功", g.Map{"id": req.Id})
}

// @Summary 停止任务
// @Description 停止任务
// @Tags Admin-TaskInfo-admin定时任务表管理
// @Param data body TaskInfoModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/taskInfo/start [post]
// @Security
func (c *taskInfo) Stop(r *ghttp.Request) {
	var req *TaskInfoModel.InfoReqParams
	//获取参数&参数校验
	if err := r.Parse(&req); err != nil {
		response.FailJson(true, r, err.Error())
	}
	_, err := TaskInfoModel.M.Data(g.Map{
		"status": 2,
	}).Where("id=?", req.Id).Update()
	if err != nil {
		response.FailJson(true, r, err.Error())
	}

	entity, err := TaskInfoService.S.Info(req.Id)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}

	err = TaskInfoService.S.Stop(entity)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}

	response.SusJson(true, r, "成功", g.Map{"id": req.Id})
}

// @Summary 开始一次任务
// @Description 开始一次任务
// @Tags Admin-TaskInfo-admin定时任务表管理
// @Param data body TaskInfoModel.UpdateReqParams true "data"
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/taskInfo/start [post]
// @Security
//func (c *taskInfo) Once(r *ghttp.Request) {
//	var req *TaskInfoModel.InfoReqParams
//	//获取参数&参数校验
//	if err := r.Parse(&req); err != nil {
//		response.FailJson(true, r, err.Error())
//	}
//	entity, err := TaskInfoService.S.Info(req.Id)
//	if err != nil {
//		response.FailJson(true, r, err.Error())
//	}
//	entity.Status = 1
//	err = TaskInfoService.S.Once(entity)
//	if err != nil {
//		response.FailJson(true, r, err.Error())
//	}
//	response.SusJson(true, r, "成功", g.Map{"id": req.Id})
//}

// @Summary 获取定时任务的所有方法
// @Description 获取定时任务的所有方法
// @Tags Admin-TaskInfo-admin定时任务表管理
// @Success 1000 {object} response.Response "{"code": 1000, "data": [...]}"
// @Router /admin/taskInfo/getTaskScheduleMethods [get]
// @Security
func (c *taskInfo) GetTaskScheduleMethods(r *ghttp.Request) {

	methods, err := TaskInfoService.S.GetTaskScheduleMethods()
	if err != nil {
		response.FailJson(true, r, err.Error())
	}

	response.SusJson(true, r, "成功", methods)
}
