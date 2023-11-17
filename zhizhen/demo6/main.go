package main

import (

	"fmt"
)
func main() {
	//要求：可以从控制台接收用户信息 【姓名，年龄，薪水，是否通过考试】
   //方式1：fmt.Scanln()
   //1.先声明需要的变量
   var name string
   var age byte
   var sal float32
   var isPass bool

   /*
   fmt.Println("请输入姓名")
   //当程序执行到fmt.Scanln(name)是，程序会停止运行，并等待用户的输入
   fmt.Scanln(&name)

   fmt.Println("请输入年龄")
   fmt.Scanln(&age)

   fmt.Println("请输入薪水")
   fmt.Scanln(&sal)

   fmt.Println("是否通过考试")
   fmt.Scanln(&isPass)

fmt.Printf("名字是%v \n 年龄是%v \n 薪水是 %v \n 是否通过考试 %v ",name,age,sal,isPass)

*/

//方式2.一次性输入这些信息 fmt.Scanf()可以按指定的格式输入信息
fmt.Println("姓名，年龄，薪水，是否通过考试,输入时采用空格隔开")
fmt.Scanf("%s %d %f %t ",&name,&age,&sal,&isPass)
fmt.Printf("名字是%v \n 年龄是%v \n 薪水是 %v \n 是否通过考试 %v ",name,age,sal,isPass)

}