package main

import (
	_ "gfEasy/boot"
	_ "gfEasy/router"

	"github.com/gogf/gf/frame/g"
)

// @title       `gfEasy`的API文档
// @version     1.0
// @description `GoFrame`基础开发框架示例服务API接口文档。
// @schemes     http
func main() {
	g.Server().Run()
}
