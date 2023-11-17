package main
import (
	"fmt"
)
var (
	Fun1 =func (n1 int,n2 int) int {
		return n1*n2
	}
)

func main(){
	//在定义匿名函数是就直接调用，这种方式匿名函数只能调用一次
   
	//案例演示，求两个数额和，使用匿名函数实现
	res1 :=func (n1 int,n2 int)int {
		return n1+n2
	}(10,20)

	fmt.Println("res1=",res1)  //30

	//第二种方法，将匿名函数func (n1 int,n2 int)int献给a变量
	//则a的数据类型就是函数类型,此时我们可以通过a完成调用
	a := func (n1 int,n2 int)int{
		return n1 -n2
	}
	res2 := a(10,30)
	
	fmt.Println("res2=",res2) //-20


	//全局匿名函数的使用
	res4 := Fun1(4,9)
    fmt.Println("res4=",res4)//36
}