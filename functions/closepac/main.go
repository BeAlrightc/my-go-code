package main
import (
	"fmt"
	"strings"
)
//累加器
func Addupper() func (int) int {
	var n int = 10
	var str string ="hello"
	return func (x int)int{
		n = n+ x
		str += "x"
		fmt.Println("str=",str)
		return n
	}
}


// 1）编写一个函数makeSuffix(suffix string)可以接收一个文件后缀名（比如.jpg）.并返回一个闭包

// 2）调用闭包，可以传入一个文件名，如果改文件名没有指定的后缀(比如.jpg),则返回文件名.jpg.如果已有文件名就返回原文件名

// 3）要求使用闭包的方式完成

// 4）strings.HasSuffix:判断一个文件是否有后

func makeSuffix(suffix string) func(string) string{
    
	return func (name string) string{
		//如果name没有指定的后缀，则加上。否则就返回原来的名字
	   if !strings.HasSuffix(name,suffix){
		return name +suffix
	   }
	   return name
	}
}

func makeSuffix2(suffix string,name string) string{
    
		//如果name没有指定的后缀，则加上。否则就返回原来的名字
	   if !strings.HasSuffix(name,suffix){
		return name +suffix
	   }
	   return name
	}
func main(){
//测试makeSuffix使用
//返回一个闭包
f :=makeSuffix(".jpg")
fmt.Println("文件名处理后=",f("winter")) //输出winter.jpg
//传统方法就要传入两个参数
fmt.Println("文件名处理后=",makeSuffix2(".jpg","winter")) //输出bird.jpg
 


	// //使用前面的代码
	// f := Addupper()
	// fmt.Println(f(1)) //11
	// fmt.Println(f(2)) //13
	// fmt.Println(f(3)) //16
}