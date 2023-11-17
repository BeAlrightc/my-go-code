package main

import (
	"fmt"
	//"math"
)
func main(){
	/*
	出票系统：根据淡旺季的月份和年龄。打印票价 【考虑学生先做】
	4~10旺季“
	  成人（18-60）:60
	  儿童（<18）:半价
	  老人(>60)：1/3

	 淡季：
	    成人：40
		其他：20 

		//分析思路
		1.month age两个变量 byte
		2.使用嵌套分支
	*/
	var month byte
	var age byte
	var price float64 =60.0
	fmt.Println("请输入您旅游的月份")
    fmt.Scanln(&month)
	fmt.Println("请输入您旅游的年龄")
    fmt.Scanln(&age)

	if month>=4 && month<=10{
		if age > 60{
			fmt.Printf("%v月份的票价是%v",month,price / 3)
		}else if age >=18{
            fmt.Printf("%v月份的票价是%v",month,price)
		}else{
			fmt.Printf("%v月份的票价： %v",month,price / 2)
		}
	}else{
		if age >=18 && age <60{
           fmt.Println("票价是",40)
		}else{
			fmt.Println("票价是",20)
		}
	}
}