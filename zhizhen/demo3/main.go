	package main
	
	import (
		"fmt"
	)

	func main() {
		//赋值运算符的使用演示
		// var i int 
		// i = 10   //基本赋值

		//有两个变量，a和b要求将其进行交换，最终打印结果
		//a =9 b=2 ==>a =2 b =9
		a := 9
		b := 2
        //定义一个临时变量
		t :=a
		a =b
		b =t
    fmt.Printf("交换后的情况如下 a=%v,b=%v",a,b)
 
	//复和赋值的操作
	a +=7  //等价 a = a+7
	fmt.Println("a+=7的值为",a)


	}