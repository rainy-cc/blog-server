package main

import (
	"blog/initialize"
	"fmt"
)

func main() {
	// 初始化路由
	app := initialize.InitRoute()
	err := app.Run(":8080")
	if err != nil {
		fmt.Println("服务启动失败")
	}
}
