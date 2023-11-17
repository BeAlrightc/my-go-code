package main

import (
	"fmt"
)
func main(){
	//100以内的数求和，求出 当和第一次大于20的当前数
	sum := 0
	for i :=1;i <=100; i++ {
		sum +=i //求和
		if sum > 20 {
			fmt.Println("当sum大于20时,这个数是：",i)
			break
		}
	}

	//实现登录验证，有三次机会，如果用户名为“张无忌密码为888就提示登录成功
	//否则就提示还有几次机会
	var name string
	var passwd string
	var loginla int = 3
	for i:=1;i<=3;i++{
		fmt.Println("请输入姓名")
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		fmt.Scanln(&passwd)
		if name !="张无忌" || passwd != "888"{
			loginla --
			fmt.Printf("登录失败，你还有%v次机会\n",loginla)
			}else{
			fmt.Println("恭喜你登录成功")
			break
		}
	}
	
}