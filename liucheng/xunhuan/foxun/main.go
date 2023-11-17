package main

import (
	"fmt"
)
func main(){
	//字符串传统遍历方式1
	var str string= "hello world!北京"
	/*for i :=0;i<len(str);i++{
		fmt.Printf("i=%d val=%c\n",i,str[i])
	}*/

	//字符串来遍历方式2 for range
	// for index, val := range str{
    //    fmt.Printf("index=%d,val=%c\n",index,val)
	// }

	//传统方式对字符串进行切片
	str2 := []rune(str) //把str转成了[]rune
	for i :=0;i<len(str2);i++{
		fmt.Printf("i=%d val=%c\n",i,str2[i])
	}

}