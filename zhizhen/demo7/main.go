package main

import (
	"fmt"
)
func main(){
	
	var i int =5;

	fmt.Printf("%d的二进制是%b\n",i,i) //输出的结果是：5的二进制是101
    //八进制：0~7，满8进1，以数字0开头进行表示
    var j int = 011
	fmt.Println("j=",j) //9
   //0~9及A-F以0x或0X表示
    var k int = 0x11
	fmt.Println("k=",k) //17

}