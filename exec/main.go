package main

import(
	"fmt"
)

func main(){
	//练习题
var n1 int32 = 12
var n2 int64
var n3 int8
var n4 int8

n2 = int64(n1)+ 20  
n3 =int8(n1) + 128 // 编译不通过   
n4 = int8(n1) +127 //编译不通过 但是结果不是127+12
fmt.Println("n2=",n2,"n3",n3)


}