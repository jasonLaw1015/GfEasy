package main

import (
	_ "goEasy/boot"
	_ "goEasy/router"

	"github.com/gogf/gf/frame/g"
)

// @title       `goEasy`的API文档
// @version     1.0
// @description `GoFrame`基础开发框架示例服务API接口文档。
// @schemes     http
func main() {
	g.Server().Run()
}
