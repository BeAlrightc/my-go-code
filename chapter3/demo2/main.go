package main

import "fmt"
func main(){
	//1.golang的变量使用方式1
	//第一种，指定变量类型，声明后不赋值的话，使用默认值
	//int 的默认值为0
	var i int
	fmt.Println("i=",i)//结果为0
	//第二种，根据值自行判断变量的类型(类型推导)
	var num = 10.11
	fmt.Println("num=",num)
	//第三种：省略var，注意:=左侧的变量不应该是已经声明的，否则会导致编译报错
	//下面的方式等价于var name string name = "tom"
	name :="tom"//这个地方的:不能省略
	fmt.Println("name=",name)
	//4）多变量声明
    //在编程过程中，有时我们需要一次声明多个变量，golang也提供这样的语法
     

}