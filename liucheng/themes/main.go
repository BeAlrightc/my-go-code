package main

import (
	"fmt"
)
func main(){
	//第一题
	var a int32 =45
	var b int32 =78
    if a+b>=50 {
		fmt.Println("这个数的和大于50")
	}
   //第二题
	var c float64 =11.0
	var d float64 =17.0
	if c > 10.0 && d < 20.0 {
		fmt.Println("和=",(c+d)) //和=28
	}
	//第三题
	var num1 int32 = 3
	var num2 int32 = 12
	if (num1 + num2) % 3 ==0 && (num1 + num2) % 5 == 0 {
		fmt.Println("能被三整除又可以被5整除")
	}
	var year int
	//第四题求润年
	fmt.Println("请输入相应的年份：")
	fmt.Scanf("%d",&year)
	if (year% 4==0 && year % 100 !=0) || year % 400 ==0{
		fmt.Printf("%d是润年",year)
	}else{
		fmt.Printf("%d不是润年",year)
	}
}