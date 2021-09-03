package packObj

import (
	"github.com/gogf/gf/errors/gerror"
	"reflect"
	"sync"
)

type MapsType map[string]reflect.Value

type singleton struct{}

var ins *singleton
var once sync.Once

var Executable = GetIns()

func GetIns() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})

	return ins
}

//
//  ReflectMethodMap
//  @Description: 把taskObj的方法转成map返回。 然后调用crMap["Test"].Call(nil)调用
// taskObj
//  @return MapsType
//
func (l *singleton) ReflectMethodMap(taskObj interface{}) MapsType {
	crMap := make(MapsType, 0)
	vf := reflect.ValueOf(taskObj)

	vft := vf.Type()
	//读取方法数量
	mNum := vf.NumMethod()
	//g.Log().Println("NumMethod:", mNum)
	//遍历路由器的方法，并将其存入控制器映射变量中
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		//fmt.Println("index:", i, " MethodName:", mName)
		crMap[mName] = vf.Method(i) //<<<
	}
	return crMap
}

//
//  TaskRun 传入定时任务的方法的字符串函数名。运行该方法的代码
//  @Description: 运行Task.Schedule的方法，可传参数。没参数时，传nil
// taskMethod 需要运行的方法
// paramList paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)};参考http://c.biancheng.net/view/118.html
//  @return err
//
func (l *singleton) TaskRun(pack interface{}, taskMethod string, paramList []reflect.Value) (err error) {
	rMap := l.ReflectMethodMap(pack)
	if _, ok := rMap[taskMethod]; ok != true {
		return gerror.New("没有该方法:" + taskMethod)
	}
	rMap[taskMethod].Call(paramList)
	return
}
