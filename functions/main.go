package main

import (
	"fmt"
	ut "go_code/functions/utils"
)
//将计算的功能放到一个函数中，在需要的时候进行调用
// func cal (n1 float64,n2 float64,operator byte) float64{
// 	var res float64
// 	switch operator {
// 	case '+':
//           res = n1 + n2
// 	case '-':
// 		  res = n1 - n2
// 	case '*':
// 		  res = n1 * n2
// 	case '/':
// 		  res = n1 / n2	  
// 	default:
// 		 fmt.Println("操作符错误")  
// 	}
// 	return res
// }

func main(){
	/*
    请大家完成这样一个需求：
	输入两个数，在输入一个运算符(+-*   / ,得到结果)*/
	// var n1 float64 =1.2
	// var n2 float64 =2.3
	// var operator byte ='/'
	result := ut.Cal(3.2,2.1,'-')
	fmt.Println("result=..",result)
	// var res float64
	// switch operator {
	// case '+':
    //       res = n1 + n2
	// case '-':
	// 	  res = n1 - n2
	// case '*':
	// 	  res = n1 * n2
	// case '/':
	// 	  res = n1 / n2	  
	// default:
	// 	 fmt.Println("操作符错误")  
	// }
	// fmt.Println("res=",res)

	//使用函数去解决这个问题：

}