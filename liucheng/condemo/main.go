package main

import (
	"fmt"
)
func main() {

	var posicount int
	var negacount int
	var num int
	for {
		fmt.Println("请输入一个整数")
		fmt.Scanln(&num)
		if num ==0 {
		break //终止for循环
		}
		if num >0 {
			posicount ++
			continue //结束本次循环，进入下一次循环
		}
			negacount ++
			
	}
	fmt.Printf("正数的个数是%v，负数的个数是%v",posicount,negacount)



/*
//1~100的奇数
for i:=1; i<=100; i++{
	if i%2!=0{
		fmt.Println(i)
	}else{
		continue
	}
}
*/

	/*
    continue案例
	这里演示一下指定标签的形式来使用
	*/
 //label2
//  for i := 0; i<4; i++{
//  //label1:设置一个标签
// for j :=0; j<10 ; j++{
// 	if j ==2 {
// 		continue
// 	}
// 	fmt.Println("j=",j)
//    }
  
//  }
}