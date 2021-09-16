// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Task

import (
	"gfEasy/app/model/AppGoodsInfoModel"
	"gfEasy/app/model/TaskInfoModel"
	"gfEasy/app/service/AppGoodsInfoService"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

var Schedule = new(schedule)

type schedule struct {
}

func (s *schedule) Test() (err error) {
	AppGoodsInfoService.S.Info(3)

	goods, err := AppGoodsInfoModel.M.All()

	g.Dump("goods", goods)
	g.Log().Println("test1")
	return nil
}

func (s *schedule) Test1() (err error) {

	g.Log().Println("test1==>")
	return nil
}

//查询gcron有激活的定时任务，定期更新所有任务状态
func (s *schedule) SelectActiveTask2UpdateStatus() error {
	g.Log().Println("SelectActiveTask2UpdateStatus")
	entries := gcron.Entries()
	g.Log().Println(entries)
	activeArr := garray.NewArray()
	//var activeArr []int
	for _, entry := range entries {
		arr := gstr.Split(entry.Name, ":")
		id := arr[len(arr)-1]
		activeArr.Append(gconv.Int(id))
	}
	//更新没激活的任务为停止
	TaskInfoModel.M.Data(g.Map{
		"status": 2,
	}).WhereNotIn("id", activeArr.Slice()).Update()
	//更新没激活的任务为开始
	TaskInfoModel.M.Data(g.Map{
		"status": 1,
	}).WhereIn("id", activeArr).Update()
	return nil
}
