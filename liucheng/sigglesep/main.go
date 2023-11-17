package main

import (
	"fmt"
)
func main(){
	var x int = 1
	var y int = 1
	if(x>2){
		if(y>2){
			fmt.Println(x+y)
		}
		fmt.Println("hello")
	}else{
		fmt.Println("x is =",x) //输出x=1
	}
}