package main
import (
	"fmt"
)

//打印金字塔的案例
func jinzi(){
	var toleve int
	fmt.Println("请输入toleve值：")
	fmt.Scanln(&toleve)
	for i := 1;i<=toleve;i++{
	   //在打印星号前打印空格
	   for k :=1;k<=toleve-i;k++{
		   fmt.Print(" ")
	   }
	   for j :=1;j<=2*i-1;j++{
		 fmt.Printf("*")
	   }
	   fmt.Println(" ")
	}
   }	

//打印九九乘法表的函数
func chengfa(){
	var n int 
	fmt.Println("请输入n的值：")
	fmt.Scanln(&n)
	for i :=1; i<=n;i++{
       for j :=1;j<=i;j++{
		fmt.Printf("%v*%v=%v ",i,j,i*j)
	   }
	   fmt.Println(" ")
	}
}

func main(){
   chengfa()
   jinzi()
}