package main

import (
	"fmt"
)
//编写一个函数 test
func test(n1 int){
	n1 = n1 + 1
    fmt.Println("n1=",n1)
}
func getSum(n1 int,n2 int) int {
	return n1 + n2
}

func getSumAndSub(n1 int,n2 int) (int,int) {
	sum := n1 +n2
	sub := n1 -n2
	return sum,sub
}
func main(){
	n1 :=10
  
  //调用test
  test(n1)
  sum := getSum(1,3)
  fmt.Println("getSum得到的值是：",sum)

  //调用getSumAndSub()函数
  res1,res2 := getSumAndSub(1,2)
  fmt.Printf("res1=%v,res2=%v\n",res1,res2)//res1=3 res2=1
//希望忽略某个返回值则使用_符号表示占位忽略
_, res3 := getSumAndSub(1,2)
fmt.Println("res3=",res3) //1-2=-1


}