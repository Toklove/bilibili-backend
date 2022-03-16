package main

import (
	"bilibili/application/Database"
	"bilibili/application/Router"
	"bilibili/core"
)

func main() {
	Database.ConnectDB() //加载数据库
	Router.Routes()      //加载路由
	core.RunApi()
}
