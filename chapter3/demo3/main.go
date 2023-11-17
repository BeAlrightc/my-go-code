package main

import "fmt"
//定义三个全局变量
var n1 =100
var n2 = 200
var name = "jack"
//上面的声明也可以改成一次性声明
var (
	n3 = 300
	n4 = 200
	names = "jack"
)

func main(){
	
	//该案例演示golang如何一次性声明多个变量
	// var n1, n2, n3 int
	// fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)
	//一次性声明多个变量方式二
	// n1,name, n3 := 100, "tom_",888
	fmt.Println("n1=",n1,"name=",name,"n2=",n2)
	fmt.Println("n3=",n3,"names=",names,"n3=",n3)

	//同样使用类型推导
}