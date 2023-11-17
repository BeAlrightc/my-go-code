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
	fmt.Scanln("请输入您旅游的月份")
    fmt.Scanln(&month)
	fmt.Scanln("请输入您旅游的年龄")
    fmt.Scanln(&age)

	if month>=4 && month<=10{
		if age > 60{
			fmt.Printf("票价是%v",price / 3)
		}else if age >=18{
            fmt.Printf("票价是%v",price)
		}else{
			fmt.Printf("票价： %v",price / 2)
		}
	}else{
		if age >=18 && age <60{
           fmt.Println("票价是",40)
		}else{
			fmt.Println("票价是",20)
		}
	}
	






/*
	var   second float64
	fmt.Println("请输入你的比赛成绩")
	fmt.Scanln(&second)
	//先判断是否进入决赛
	if second <=8{
		//进入决赛了
		var  gender string
		fmt.Println("请问您的性别是")
	    fmt.Scanln(&gender)
		if gender == "男"{
			fmt.Println("恭喜你进入男子组决赛")
		}else{
			fmt.Println("恭喜你进入女子组决赛")
		}

	}else{
		fmt.Println("对不起你被淘汰了")
	}
	*/




	
	/*
	//分析思路
	//设计三个变量，需要从控制台输入
	var height int32
	var money float32
	var handsome bool
	fmt.Println("请你输入身高(厘米)")
	fmt.Scanln(&height)
	fmt.Println("请你输入财富")
	fmt.Scanln(&money)
	fmt.Println("帅吗")
	fmt.Scanln(&handsome)
	if height>180 && money>10000000.0 && handsome==true{
		fmt.Println("我一定要嫁给他")
	}else if height>180 || money>10000000.0 || handsome==true{
        fmt.Println("嫁吧，比上不足，比下有余")
	}else{
		fmt.Println("不嫁")
	}
	*/








	//分析思路
	//1.a,b,c是三个float64
	//2.使用到给出的书序公式
	//3.使用到多分支
	//4.使用到math.squr方法
	//走代码
	/*
	var a float64 =3.0
	var b float64 =1.0
	var c float64 =10.0
	

	m := b*b -4 * a * c
	//多分支判断
	if m>0 {
		x1 :=-b + math.Sqrt(m) / 2 * a
		x2 :=-b - math.Sqrt(m) / 2 * a
		fmt.Printf("x1=%v,x2=%v",x1,x2)
	}else if m == 0 {
        x1 := (-b + math.Sqrt(m)) / 2 *a
		fmt.Printf("x1=%v",x1)
	}else{
		fmt.Println("无解...")
	}
	*/








	//多分支类型的理解
	/*岳小鹏参加golang考试，他和父亲岳不群达成承诺
	如果
	成绩100分奖励一台bmw
	成绩为(80,90)奖励一套iphone14
	当成绩（60,80）奖励一台ipad
	其他时，什么奖励都没有
	请从键盘输入岳小鹏的期末成绩，并加以判断*/
	/*var score int
	fmt.Println("请输入岳小鹏的成绩：")
	fmt.Scanf("%d",&score)
	if score==100{
		fmt.Println("奖励一台BMW")
	}else if score>80&&score<=99{
        fmt.Println("奖励一台iphone14")
	}else if score>=60&&score<=80{
		fmt.Println("奖励一个ipad")
	}else{
		fmt.Println("对不起，他什么都没有好好学习")
	}
	
	var b bool = true
	if b= false{ //在这里不能赋值只能写成==
		fmt.Println("a")
	}else if b {
		fmt.Println("b")
	}else if !b {
		fmt.Println("c")
	}else {
		fmt.Println("d")
	}
	*/
}
