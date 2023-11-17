package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
	"unsafe"
)

//展示Golang整数类型的使用
func main(){
	
	// var i int =1
	// fmt.Println("i=",i)
	//测试一下int的范围
	//var j int8 = -129 //会报错
	//fmt.Println("j=",j)
	//测试一下
    var k uint8 = 255
	fmt.Println("k=",k)
	//int,uint,rune, byte的使用
	var a int =8900
	fmt.Println("a=",a)
	var b uint = 1
	var c byte = 3 //范围是0~255
	fmt.Println("b=",b,"c=",c)

	//整型的使用细节
	var n1 = 100 //n1是什么类型
     //这里我们给介绍一下如何查看某个变量的数据类型
	 //fmt.Printf()  可以用于做格式化的输出

	fmt.Printf("n1 的类型 %T ",n1)

	 //如何查看某个变量的占用字节大小和数据类型 （使用较多）
    var n2 int64 = 10
	//unsafe.Sizeof(n2)) 是unsafe包的一个函数，可以返回n1变量占用的字节数
	fmt.Printf("n2 的类型 %T n2占用的字节数为 %d ",n2,unsafe.Sizeof(n2))
    

	//golang程序中整型变量在使用时，遵守保大不保小的原则
	//即，在程序正常运行下，尽量使用占用空间小的数据类型
    //var age byte = 45


}

















































