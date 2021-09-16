package boot

import (
	"gfEasy/app/service/TaskInfoService"
	_ "gfEasy/packed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

// 用于应用初始化。
func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	//监听是否有新的redis事件
	go TaskInfoService.S.AddTaskSubscribe()
	//定时任务注册
	TaskInfoService.S.InitStartALLTaskList()

}
