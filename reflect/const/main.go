package main
import (
	"fmt"
)
func main() {
	var num int
	num = 90 
	//常量声明的时候必须赋值
	const tax int = 90
	// tax = 10 //常量是不能修改的
	fmt.Println(num,tax)
}