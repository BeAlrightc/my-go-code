package main

import (
	"fmt"
)

//演示golang中字符类型使用
func main(){
	
	var c1 byte = 'a'
	var c2 byte = '0'
	//当我们输出byte值时，就直接输出了对应的字符ASCI码值
    fmt.Println("c1=",c1,"c2=",c2) //c1=92 c2=48
	//如果我们希望输出对应的字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n",c1,c2) //输出结果为 c1=a c2=0

	//var c3 byte = '北'  //overflow 溢出
	var c3 int = '北'  //汉字的ASCI码值比较大
	fmt.Printf("c3=%c c3对应的码值为=%d\n",c3,c3) //北 c3对应的码值为21271
    
	var c4 int = 22269 // 22269的ASCI码是国字
	fmt.Printf("c4=%c",c4) //输出的结果就是国字

   //字符类型是可以进行运算的，相当于于一个整数，运算时是按照码值进行运算的
   var num1 = 10 + 'a' //10 +97 =107
   fmt.Println("num1=",num1)
}