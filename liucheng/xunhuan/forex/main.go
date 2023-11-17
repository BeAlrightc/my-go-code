package main

import(
	"fmt"
)
func main(){
	/*
	var ge int = 0
	var sum int = 0
	for i :=1;i<=100;i++{
		if i % 9 ==0{
			ge++
			sum = sum + i
		}
	}
	fmt.Printf("1~100之间是9的倍数额整数个数为%d,总和为%d",ge,sum)
    */
	//第二题
	// var n int =6
	// for i :=0;i<=n;i++{
	// 	fmt.Printf("%v + %v= %v\n",i,n-i,n)
	// }
	//使用while方式输出十句"hello world"
	//循环变量初始化
	var i int = 1
	// for {
	// 	if i>10 {
	// 		break  //跳出循环结束过程
	// 	}
	// 	fmt.Println("hello world",i)
	// 	i++  //循环变量的迭代
	// }
	//使用do while的形式进行实现
	for {
		fmt.Println("hello world" ,i)
		i++ //循环变量进行迭代
		if i>10{
			break //跳出for循环
		}
	}
}