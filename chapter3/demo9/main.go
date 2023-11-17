package main

import (
	"fmt"
	// "unsafe"
)
func main(){
	var str string = "abc\nabc"
	var str2 string =`
	var b = false
	fmt.Println("b=",b)
    //注意事项
   //1.bool类型占用存储空间是1个字节
  // fmt.Println("b 的占用空间 =",unsafe.Sizeof(b) ) //b占用空间为1
   //2.bool类型只能取true或者false
   
//字符串

}
`
fmt.Println(str)
fmt.Println("====================")
fmt.Println(str2)

}