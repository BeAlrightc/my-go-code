package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

func main() {
	//创建数据库
	//sql：CREATE DATABASE bubble;
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err) //数据库都连接不上的话就没有必要往下走了
	}
	//延迟注册关闭
	defer dao.Close() //程序退出关闭数据库连接
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{}) //todos
	r := routers.SetupRouter()

	r.Run(":9090")
}
