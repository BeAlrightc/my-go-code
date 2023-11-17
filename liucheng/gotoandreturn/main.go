package main

import (
	"fmt"
)
func main(){
	//演示goto的使用
	fmt.Println("ok1")
	goto lable1
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	lable1:
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
}