package main
import (
	"fmt"
)
func main(){
	//演示切片的基本使用
	var intArr [5]int = [...]int{1,22,33,66,99}
	//声明定义一个切片
	/*
     1.slice 就是切片的名称
	 2.intArr[1:3] 表示slice 引用到intArr这个数组
	 3.应用inArr数组的起始下标为1终止下标为3不包含3
	*/
    slice := intArr[1:3] 
    fmt.Println("intArr=",intArr) // intArr= [1 22 33 66 99]
    fmt.Println("slice 的元素是=",slice)//slice 的元素是= [22 33]
    fmt.Println("slice 的元素的个数是=",len(slice))//slice 的元素的个数是= 2
	fmt.Println("slice 的容量是=",cap(slice))//slice 的容量是= 4
	//切片的容量是可以动态变化的 cability
	fmt.Printf("intArr[1]的地址=%p\n",&intArr[1])
	fmt.Printf("slice[0]的地址=%p slice[0]=%v\n",&slice[0],slice[0])
	slice[0]=34 //相当于*intArr[1]=34
	fmt.Println("intArr=",intArr)//intArr= [1 34 33 66 99]
}
