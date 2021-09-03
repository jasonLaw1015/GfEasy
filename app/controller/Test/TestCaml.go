package Test

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"goEasy/library/response"
	"goEasy/library/utils"
	"goEasy/library/utils/captcha"
)

var TestCaml = new(test)

type test struct{}

func (c *test) Test1(r *ghttp.Request) {
	str := "user_gDods_a"
	//小驼峰
	res := utils.Case2SmallCamel(str)

	//小驼峰2下划线
	caseStr := utils.Camel2Case(res)

	//小驼峰2大驼峰
	bigCamelStr := utils.Ucfirst(res)

	//大驼峰2小驼峰
	smallCamelStr := utils.Lcfirst(bigCamelStr)

	//下划线2大驼峰
	bigCamelStr2 := utils.Case2BigCamel(caseStr)

	////读取文件
	//routerStr := gfile.GetContents("./router/router.go")
	//adminObjStr := "group.ALL(\"/user\", Admin.User)"
	//apiObjStr := "group.ALL(\"/user\", Api.User)"
	////判断是否已有要加入的路由字符串
	//
	//pd := gstr.Contains(routerStr, adminObjStr)
	//if pd == false {
	//	router1 := gstr.StrTillEx(routerStr, "/**insert_router_end**/")
	//	router2 := gstr.StrEx(routerStr, "/**insert_router_end**/")
	//	//插入对应字符串
	//	routerStr = router1 + "\n\t\t" + adminObjStr + "\n\t\t/**insert_router_end**/\n\t\t" + router2
	//}
	//
	//pdApi := gstr.Contains(routerStr, apiObjStr)
	//if pdApi == false {
	//	routerApi1 := gstr.StrTillEx(routerStr, "/**api_insert_router_end**/")
	//	routerApi2 := gstr.StrEx(routerStr, "/**api_insert_router_end**/")
	//test代码 //	//插入对应字符串
	//	routerStr = routerApi1 + "\n\t\t" + adminObjStr + "\n\t\t/**insert_router_end**/\n\t\t" + routerApi2
	//}
	//保存回文件
	//g.Dump(routerStr)
	//utils.RegisterTableNameAllRouter("user", "./")
	//utils.RegisterTableNameAllRouter("userTableName", "./")

	k1, b1 := captcha.GetVerifyImgDigit()

	k2, b2, err := captcha.GetVerifyImgString()
	if err != nil {
		response.FailJson(true, r, "验证码错误", err.Error())
	}
	bl := captcha.VerifyString("odCGGV3QR7CS8uS9Emrg", "25470")
	response.SusJson(true, r, "处理成功", g.Map{
		"res": res, "caseStr": caseStr, "bigCamelStr": bigCamelStr, "smallCamelStr": smallCamelStr, "bigCamelStr2": bigCamelStr2,
		"k1": k1,
		"b1": b1,
		"k2": k2,
		"b2": b2,
		"bl": bl,
	})
}
