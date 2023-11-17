package main
import (
	"fmt"
)
func test(slice []int){
	slice[0]=100
}
func main(){
	//string底层是一个byte数组，因此string也可以进行切片处理
	str:= "hello mrliu"
	//使用切片获取mrliu
	slice :=str[6:]
	fmt.Println("slice",slice) //slice=mrliu
 //string是不可改变的，也就是说不能通过str[0]='z'方式来修改字符串
//  str[0] = 'z' //错误，编译不会通过string是不可变的

//如果需要修改字符串，可以先将string->[]byte / 或者 []rune ->修改->重写转成string
//"hello mrliu"=》改成"zello mrsliu"
// arr1 := []byte(str)
// arr1[0]='z'
// str = string(arr1)
// fmt.Println(str)//zello mrliu

//细节：我们转成[]byte后,可以处理英文和数字，但是没办法处理中文
//原因是 []byte字节来处理，而一个汉字是3个字节。因此就会出现erro
//解决办法是将string转成[]rune即可，因为[]rune是按照字符处理兼容汉字

arr1 := []rune(str)
arr1[0]='北' //utf-8 21271
str = string(arr1) 
fmt.Println(str)//北ello mrliu



}