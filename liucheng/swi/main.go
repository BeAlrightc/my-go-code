package main

import (
	"fmt"
)

//写一个非常简单的函数
func test(b byte) byte {
	return b+1

}
func main(){
	/*
	请编写一个程序，该程序可以接收一个字符，
	比如：a,b,c,d,e,f a表示星期一，
	 b表示星期二...根据用户输入显示相应的信息，
	 要求使用switch语句完成
	 分析：
	 1，定义一个变量接受字符
	 2.使用switch完成
	
	var key byte
	fmt.Println("请输入一个字符  a b c d e f g")
	fmt.Scanf("%c",&key)

	switch test(key) { //调用、test函数
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
	case 'e':
		fmt.Println("周五")
	case 'f':
		fmt.Println("周六")
	case 'g':
		fmt.Println("周日")	
	default:
		fmt.Println("输入有误")						
	}
	*/

	var n1 int32 = 5
	var n2 int32 = 20 //改为int64是错误的

	switch n1{//n2的类型需要与n1保持一致
	case n2,10,5 : //case后面可以有多个表达值。匹配其中的一个就可以
		fmt.Println("ok") //输出了ok
	default:
		fmt.Println("输入有误")			
	}

	//switch后可以不带表达式，类似于if -else 分支来使用
	var age int = 10
	switch {
	case age ==10 :
		fmt.Println("age ==10")
    case age ==20 :
		fmt.Println("age ==20")
    case age ==30 :
		fmt.Println("age ==30")
	default :
	    fmt.Println("没有匹配到")				
	}
	//case中也可以就age的范围进行判断
	var score int = 90
	switch {
	case score > 90 :
		fmt.Println("成绩优秀")
	case score >=70 && score <=90 :
		fmt.Println("成绩优良")
	case score >=60 && score <70 :
		fmt.Println("成绩及格")
	default :
	fmt.Println("不及格")		
	}

	//switch后也可以直接声明/定义一个变量,分号结束，并不推荐
	switch grade :=90; {
	case grade > 90 :
		fmt.Println("成绩优秀_")
	case grade >=70 && grade <=90 :
		fmt.Println("成绩优良_")
	case grade >=60 && grade <70 :
		fmt.Println("成绩及格_")
	default :
	fmt.Println("不及格")		
	}

	//switch的穿透 fallthrought
	var num int = 10
	switch num{
	case 10 :
		fmt.Println("ok1")
		fallthrough  //默认只能穿透一层
    case 30 :
		fmt.Println("ok2")
		fallthrough
    case 20 :
		fmt.Println("ok3")		
	default :
		fmt.Println("没有匹配到")				
	}

	var x interface{}
	var y = 10.0
	x = y 
	switch i:= x.(type){
	case nil:
		fmt.Printf("x的类型是：%T",i)
    case int:
         fmt.Println("x是int型")
	case float64:
		fmt.Println("x是float64型")	 
    case func(int) float64:
		fmt.Println("x是func(int)型")
	case bool,string:
		fmt.Println("x是bool型或者string型")		 		
	default:
		fmt.Println("未知")
	}
}	
