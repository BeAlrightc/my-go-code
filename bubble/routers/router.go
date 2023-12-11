package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//告诉gin框架模板文件引用的静态文件去哪里找,要使用绝对路径
	r.Static("/css", "D:/myfile/GO/project/qindanprojects/myproject/bubble/static/css")
	r.Static("/js", "D:/myfile/GO/project/qindanprojects/myproject/bubble/static/js")
	r.Static("/fonts", "D:/myfile/GO/project/qindanprojects/myproject/bubble/static/fonts")

	//告诉gin框架去哪里找模板
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	//v1
	v1Group := r.Group("v1")
	{
		//代办事项
		//添加
		v1Group.POST("/todo", controller.CreateTOdo)
		//查看所有的代办事项
		v1Group.GET("/todo", controller.GetTodoList)

		//修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATOdo)

	}
	return r
}
