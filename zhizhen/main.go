package main

import(
	"fmt"
)
//演示golang中的指针类型
func main(){


	//基本数据在内存的布局
	var i int = 10 
	// i的地址是多少 &i
	fmt.Println("i的地址=",&i)  //结果为：i的地址= 0xc04205e058
	//1.a是一个指针变量
	//2. a的类型是 *int
	//3.a 本身就是 &i
	var a *int =&i
	fmt.Println("i的值是",i) //10
	fmt.Println("a存放的值是",a)//0xc04205e058
	fmt.Println("a指向的值是",*a)//10
	fmt.Println("指针本身的地址是",&a)//0xc04207c020


}