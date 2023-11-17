package main

import (
	"fmt"
)

func main() {
	//求两个数的最大值
	var n1 int =10
	var n2 int =40
	var max int
	if n1 > n2{
		max =n1
	}else{
		max =n2
	}
	fmt.Println("两个数的最大值为：",max)

	//求出三个数的最大值
	//先求出两个数的最大值然后让第三个数取比较

	var n3 int =45
    if n3 > max {
		max =n3
	}
	fmt.Println("三个数中的最大值是=",max)
	



	/*
//演示一把 &和*取值
a := 100
fmt.Println("a的地址是",&a)

//定义一个指针变量
var ptr *int = &a
fmt.Println("ptr指向的值是",*ptr)

var n int
var i int = 10
var j int = 12
//传统的三元运算符
//n = i >j ? i : j
//下面是go支持的
if i >j {
	n =i
}else{
	n =j
}
fmt.Println("n=",n) //12
*/

	// var a int =10
	// var b int = 12
	// var c int = 89
	// if a > b {
	// 	fmt.Println("两个数的最大值是：",a)
	// }else{
	// 	fmt.Println("两个数的最大值是：",b)
	// }

	// if a >b && a>c {
	// 	fmt.Println("三个数的最大值是：",a)
	// }
}