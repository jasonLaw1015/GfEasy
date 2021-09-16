// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package TaskInfoService

import (
	"gfEasy/app/model/TaskInfoModel"
	"gfEasy/app/model/TaskLogModel"
	"gfEasy/app/service/TaskLogService"
	Task "gfEasy/app/task"
	"gfEasy/library/utils"
	"gfEasy/library/utils/packObj"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var S = new(taskInfoService)

type taskInfoService struct {
}

// 查询单个信息
func (s *taskInfoService) Info(id int) (data *TaskInfoModel.Entity, err error) {

	err = TaskInfoModel.M.Where("id=?", id).Scan(&data)

	return
}

//添加操作
func (s *taskInfoService) Add(req *TaskInfoModel.AddReqParams) (insId int64, err error) {

	r, err := TaskInfoModel.M.OmitEmpty().Data(req).Insert()
	if err != nil {
		return 0, err
	}
	// 判断新增还是修改
	id, e := r.LastInsertId()
	s.modifyAfter(gconv.Int(id))
	return id, e
}

//修改操作
func (s *taskInfoService) Update(req *TaskInfoModel.UpdateReqParams) (Id int, err error) {

	_, err = TaskInfoModel.M.OmitEmpty().Data(req).Where("id=?", req.Id).Update()
	if err != nil {
		return 0, err
	}
	id := gconv.Int(req.Id)
	s.modifyAfter(id)

	return id, err
}

//分页查询
func (s *taskInfoService) Page(req *TaskInfoModel.PageReqParams) (total, page int, size int, list []*TaskInfoModel.Entity, err error) {
	page = req.Page
	size = req.Size

	M := TaskInfoModel.M

	if req.Status != -1 {
		M = M.Where("status=?", req.Status)
	}

	if req.Types != -1 {
		M = M.Where("types=?", req.Types)
	}

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
	list = make([]*TaskInfoModel.Entity, len(data))
	err = data.Structs(&list)

	return
}

//List
func (s *taskInfoService) List(condition g.Map) (list []*TaskInfoModel.Entity, err error) {

	err = TaskInfoModel.M.Where(condition).Scan(&list)
	return

}

//删除
func (s *taskInfoService) Delete(ids []int) (err error) {
	_, err = TaskInfoModel.M.Where("id IN(?)", ids).Delete()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除模型数据失败")
		return
	}
	return
}

//增加修改后，需要的操作
func (s *taskInfoService) modifyAfter(id int) (err error) {

	//发布redis 订阅的数据
	err = s.TaskPublish(g.Map{"id": id})
	return
}

//初始化所有开启的定时任务
func (s *taskInfoService) InitStartALLTaskList() (err error) {
	var validTask []*TaskInfoModel.Entity
	list, err := s.List(g.Map{"status": 1})
	nowUnix := gtime.Timestamp()
	for _, entity := range list {
		if entity.StartDate != nil && entity.EndDate != nil {
			//g.Dump(entity.StartDate.Unix(), entity.EndDate.Unix(), nowUnix)
			if entity.StartDate.Unix() <= nowUnix && entity.EndDate.Unix() >= nowUnix {
				validTask = append(validTask, entity)
			}
		} else {
			validTask = append(validTask, entity)
		}
	}

	for _, task := range validTask {
		//这里启动task错误，发生错误不处理。因为是启动任务boot使用
		err = s.Start(task)

		if err != nil {
			g.Log().Println(err.Error())
			continue
		}
	}

	return

}

//开启定时任务
//task是更改后状态的数据
func (s *taskInfoService) Start(task *TaskInfoModel.Entity) (err error) {
	//判断当前服务器的是否开启定时任务功能
	taskActive := g.Cfg().GetBool("server.taskActive", false)
	if taskActive == false {
		return gerror.New("当前节点-不支持定时任务")
	}
	status := 1 //1 成功 2失败
	detail := "任务执行成功"
	g.Log().Println("我开启定时任务：" + task.Cron + ":" + gconv.String(task.Id))

	//taskType状态 2:cron 1：时间间隔
	if task.TaskType == 2 {
		name := "task:cron:" + gconv.String(task.Id)
		//先判断存不存在，存在的话移除
		ent := gcron.Search(name)
		if ent != nil {
			gcron.Remove(name)
		}
		_, err = gcron.AddSingleton(gstr.Trim(task.Cron), func() {
			g.Log().Println("我是cron版定时任务：" + task.Cron + ":" + name)
			err = packObj.Executable.TaskRun(Task.Schedule, task.Service, nil)

			if err != nil {
				status = 2
				detail = err.Error()
			}

			log := TaskLogModel.AddReqParams{
				Status: status,
				TaskId: gconv.Int64(task.Id),
				Detail: detail,
			}

			TaskLogService.S.Add(&log)
		}, name)
	} else {
		name := "task:every:" + gconv.String(task.Id)
		//先判断存不存在，存在的话移除
		ent := gcron.Search(name)
		if ent != nil {
			gcron.Remove(name)
		}
		//设置时间间隔+次数的
		p := "@every " + gconv.String(task.Every) + "ms"
		_, err = gcron.AddTimes(p, task.Limit, func() {
			g.Log().Println("我是时间-间隔版定时任务：" + gconv.String(task.Every/1000) + ":" + gconv.String(task.Limit) + ":" + name)
			err = packObj.Executable.TaskRun(Task.Schedule, task.Service, nil)
			if err != nil {
				status = 2
				detail = err.Error()
			}

			log := TaskLogModel.AddReqParams{
				Status: status,
				TaskId: gconv.Int64(task.Id),
				Detail: detail,
			}

			TaskLogService.S.Add(&log)

		}, name)
	}
	return
}

//停止定时任务
//task是更改后状态的数据
func (s *taskInfoService) Stop(task *TaskInfoModel.Entity) (err error) {
	name := "task:every:" + gconv.String(task.Id)
	//taskType状态 2:cron 1：时间间隔
	if task.TaskType == 2 {
		name = "task:cron:" + gconv.String(task.Id)
	}

	gcron.Stop(name)
	return
}

//开启一次性定时任务
func (s *taskInfoService) Once(task *TaskInfoModel.Entity) (err error) {
	task.Limit = 1
	err = s.Start(task)
	return
}

//处理定时任务状态
func (s *taskInfoService) handleTask(id int) (err error) {
	//判断当前服务器的是否开启定时任务功能
	taskActive := g.Cfg().GetBool("server.taskActive", false)
	if taskActive == false {
		return gerror.New("当前节点-不支持定时任务")
	}
	var entity *TaskInfoModel.Entity
	err = TaskInfoModel.M.Where("id=?", id).Scan(&entity)
	g.Log().Println("handleTask", entity)
	if entity == nil || err != nil {
		return
	}
	//开始
	if entity.Status == 1 {
		s.Start(entity)
		//停止
	} else if entity.Status == 2 {
		s.Stop(entity)
	}
	return
}

/**
开启订阅-定时任务开启或关闭操作
redis订阅信息
*/
func (s *taskInfoService) AddTaskSubscribe() (err error) {
	conn := g.Redis().Conn()
	defer conn.Close()
	_, err = conn.Do("SUBSCRIBE", "TaskChannel")
	if err != nil {
		panic(err)
	}
	g.Log().Println("[redis] TaskChannel订阅成功")
	for {
		reply, err := conn.Receive()
		if err != nil {
			return err
		}
		//["message","TaskChannel","{\"Id\":1,\"status\":1}"]
		g.Log().Println(gconv.Strings(reply))
		arr := gconv.Strings(reply)
		type TaskParams struct {
			Id int `json:"id"`
		}
		var task TaskParams
		gconv.Struct(arr[2], &task)
		g.Log().Println(task)
		//处理定时任务状态
		s.handleTask(task.Id)
	}
	return
}

/**
	redis 发布信息
 	data:
		g.Map{
			"Id":     1,
			"status": 1,
		})
*/
func (s *taskInfoService) TaskPublish(data interface{}) (err error) {

	_, err = g.Redis().Do("PUBLISH", "TaskChannel", data)
	return
}

/**
获取app/task下的所有方法
*/
func (s *taskInfoService) GetTaskScheduleMethods() (methods []string, err error) {
	crMap := packObj.Executable.ReflectMethodMap(Task.Schedule)

	return utils.GetReflectkeys(crMap), err
}
