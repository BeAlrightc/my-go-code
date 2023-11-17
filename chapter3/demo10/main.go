package main

import (
	"fmt"
)

//演示golang中字符串类型使用
func main(){
	//string的基本使用
	var address string = "北京长城 110 hello world"
  	fmt.Println(address)
    

	//字符串一旦赋值了，字符串就不能修改了，在Go中字符串是不可变的
	var str ="hello"
	//str[0] = 'a'  //这里不能去修改str的内容，即Go中的字符串是不可变的
    fmt.Println(str)

	str3 := "abc\nabc"
	fmt.Println(str3)

	//使用反引号``可以将源代码进行输出与
	str4 := `fmt.Println("hello world")`
	fmt.Println(str4)

	//字符串拼接方式
	var str5 = "hello" + "world"
	str5 += "haha"
    fmt.Println(str5)

	// var str6 = "hello" + "world" 
	// +"hello"+"hello"
	// +"world"
	// fmt.Println(str6) //会报错，+号要在上面

	var str6 = "hello" + "world" +
	"hello"+
	"hello"+
	"world"
	fmt.Println(str6) //正确输出
//查看这些基本数据类型的默认值
//%v表示的意思是按照变量的值输出
	var a int //0
	var b float32 //0
	var c float64 // 0 
	var isMarryied bool  //false
	var name string // " "
	fmt.Printf("a=%d,b=%f,c=%v,isMarryied=%v,name=%v",a,b,c,isMarryied,name)

}