package main

import "fmt"

//变量使用的注意事项
func main(){
	//该区域的数据值可以在同一类型范围内不断变化
	var i  int =10
	i = 30
	i = 50
	fmt.Println("i=",i)
	//i = 1.23 //汇报错，因为不是int,不能改变数据类型

	//变量在同一个作用域(在一个函数或者在代码块中)内不能重名
	//var i int = 50
	//i := 99  //也会报错。这样写就包含了定义变量
}