package main

import(
	"fmt"
)

func main() {
//重点讲解 /,%
//如果参与运算的数都是整数，那么除后，去掉小数部分，保留整数部分。不会四舍五入
fmt.Println(10/4) //是2不是2.5
var n1 float32 = 10 /4  //?结果也是2
fmt.Println(n1)

//如果我们需要保留小数部分，则需要有浮点数参与运算
var n2 float32 = 10.0 /4
fmt.Println(n2)
fmt.Println(-10 / -3)


//演示 % 的使用
//看一个公式 a%b=a - (a/b) *b
fmt.Println("10%3=",10 % 3) //结果余数是1
fmt.Println("-10%3=",-10 % 3) //结果余数是-1 : -10 - [(-10)/3]*3 =-10 -(-9)=-1
fmt.Println("10%-3=",10 % -3) //结果余数是 1
fmt.Println("-10%-3=",-10 % -3) //结果余数是 -1



//演示 ++ --
var i int = 10
i++  //等价 i = i+1
fmt. Println("i的值是",i)  //11
i-- //等价i = i -1
fmt. Println("i的值是",i)  //10
}