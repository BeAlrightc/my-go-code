package main

import (
	"fmt"
)
func main(){
	//for循环的第一种写法
	for i :=1;i<= 10;i++{
		fmt.Println("你好刘承",i)
	}
	//for循环的第二种写法
	j := 1 //循环变量初始化
	for j <= 10 { //循环条件
		fmt.Println("你好刘承传",j) //循环体
		j++  //循环变量的迭代
	}
	
	//for循环的第三种写法,这种写法通常会配合break进行使用
	k :=1
	for { //此处等价 for ; ; ;{}
		if k <= 10{
		fmt.Print("helllo\n")
		}else{ //k>10时就break跳出
			break //break就是跳出整个for循环
		}
		k++
	}
}