package main

import (
	"fmt"
	//为了使用utils文件的变量或者是函数，我们需要首先引入这个包
    "go_code/zhizhen/model"
)


//演示golang中标识符的使用
func main() {
	//严格区分大小写
   var num int = 10
   var Num int = 20
   fmt.Printf("num=%v,Num=%v",num,Num) //10 20
   //标识符中不能含有空格
//    var ab c int = 30  //会报错

//_是空标识符，用于占用
// var _ int =40  // erro

//语法是可以通过的，但是不推荐使用
	var int int = 10
	fmt.Println(int)
	//使用utils.go的helloName 包名.标识符
	fmt.Println("model下的HelloName是：",model.HelloName)

} 