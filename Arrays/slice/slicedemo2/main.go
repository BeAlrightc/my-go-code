package main
import (
	"fmt"
)
func main(){
	//使用常规的for循环遍历切片
	var arr [5]int=[...]int{10,20,30,40,50}
	//slice :=arr[0:len(arr)]
	slice := arr[:]
	for i := 0;i<len(slice);i++{
		fmt.Printf("slice[%v]=%v ",i,slice[i])//slice[0]=20 slice[1]=30 slice[2]=40
	}
    //使用for -range 方式遍历切片
	for i,v :=range slice{
		fmt.Printf("slice[%v]=%v ",i,v)////slice[0]=20 slice[1]=30 slice[2]=40
	}

	slice2 := slice[1:2] //20
	fmt.Println("slice1=",slice2) //[20]

	//用append内置函数，可以对切片进行动态追加
	var slice3 []int = []int{100,200,300}
	//通过append直接对slice3追加具体的元素
	slice3 = append(slice3,400,500,600)
	
	fmt.Println("slice3=",slice3)

	//通过append将切片slice3追加到slice3
	slice3 = append(slice3,slice3...)
	fmt.Println("slice3=",slice3)

	//切片的拷贝操作
	//切片的copy内置函数完成拷贝，举例说明
	var slice4 []int =[]int{1,2,3,4,5}
	var slice5 = make([]int,10)
	copy(slice5,slice4)//将slice4拷贝给slice5
	fmt.Println("slice4=",slice4)//[1 2 3 4 5]
	fmt.Println("slice5=",slice5)//[1 2 3 4 5 0 0 0 0 0]
}