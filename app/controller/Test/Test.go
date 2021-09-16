package Test

import (
	"gfEasy/app/service/BaseSysLogService"
	"gfEasy/app/service/BaseSysRoleService"
	"gfEasy/app/service/TaskInfoService"
	Task "gfEasy/app/task"
	"gfEasy/library/response"
	"gfEasy/library/utils/packObj"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcron"
	"os/exec"
	"reflect"
)

var Tools = new(tools)

type tools struct{}

func (c *tools) HandleMenus2Camel(r *ghttp.Request) {
	BaseSysRoleService.S.HandleMenus2Camel()
	response.SusJson(true, r, "处理成功", nil)
}

func (c *tools) TestRecordLog(r *ghttp.Request) {
	BaseSysLogService.S.Record(r)
	response.SusJson(true, r, "处理成功", nil)
}

func (c *tools) TestReflect(r *ghttp.Request) {
	s1 := "Test"
	s3 := "Test1"
	s2 := "GoodsTask"

	packObj.Executable.TaskRun(Task.Schedule, s1, nil)
	packObj.Executable.TaskRun(Task.Schedule, s2, nil)
	packObj.Executable.TaskRun(Task.Schedule, s3, []reflect.Value{reflect.ValueOf("10s")})
	response.SusJson(true, r, "处理成功", nil)
}

func (c *tools) TestTaskPublish(r *ghttp.Request) {

	err := TaskInfoService.S.TaskPublish(nil)
	if err != nil {
		response.FailJson(true, r, "发布信息错误")
	}
	response.SusJson(true, r, "处理成功", nil)
}

func (c *tools) TestTaskInit(r *ghttp.Request) {

	err := TaskInfoService.S.InitStartALLTaskList()
	if err != nil {
		response.FailJson(true, r, "发布信息错误")
	}

	response.SusJson(true, r, "处理成功", nil)
}

func execCmd() (s string, err error) {
	ml := "ls"
	cmd := exec.Command("sh", "-c", ml)
	opBytes, err := cmd.Output()
	if err != nil {
		return
	}
	s = string(opBytes)
	return
}

func (c *tools) GetAllTaskEntries(r *ghttp.Request) {
	entries := gcron.Entries()
	g.Log().Println(entries)

	response.SusJson(true, r, "处理成功", entries)
}

func (c *tools) SelectActiveTask2UpdateStatus(r *ghttp.Request) {
	Task.Schedule.SelectActiveTask2UpdateStatus()

	response.SusJson(true, r, "处理成功", nil)
}
