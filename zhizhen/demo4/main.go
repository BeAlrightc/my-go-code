package main

import(
	"fmt"
)
func main(){
	//有两个变量a和b，要求将其进行交换，但是不允许使用中间变量，最终打印效果
	var a int = 10
	var b int = 20
	a = a + b  //30 
	b = a - b  // 10
	a = a -b  //30-10 =20

	fmt.Printf("a=%v和b=%v",a,b) //20 和 10
}