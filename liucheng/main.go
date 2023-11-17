package main

import (
	"fmt"
)
//编写一个程序可以输入一个人的年龄，
// 如果同志的年龄大于18岁则输出你的年龄大于18岁
//要对自己的行为负责

//分析：年龄 ==》var age int
//2.从控制台接收输入

func main(){
var age int
fmt.Println("请输入你的年龄:")
fmt.Scanln(&age)
if age >=18 {
	fmt.Println("你的年龄大于18岁，要为自己的行为所负责！")
      }else{
		fmt.Println("你的年龄不大，这次放过你了")
	  }

}