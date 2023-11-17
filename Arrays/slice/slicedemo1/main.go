package main
import (
	"fmt"
)
func main(){
	//演示切片的使用make
	var slice []float64 = make([]float64,5,10)
	slice[1] = 10
	slice[3] = 20
    //对于切片，必须使用make
	fmt.Println(slice)
	fmt.Println("slice的size=",len(slice))
	fmt.Println("slice的cap=",cap(slice))


	//第3中方式：定义一个切片，直接就指定具体数组，使用原理类似make方式

	var strSlice []string = []string {"tom","jack","mary"}
	fmt.Println(strSlice)
	fmt.Println("strSlice的size=",len(strSlice))//3
	fmt.Println("strSlice的cap=",cap(strSlice))//3





	/*
    var slice []int =make([]int,4,10)
    fmt.Println(slice)//默认值为0 [0 0 0 0]
    fmt.Println("slice len=",len(slice),"slice cap=",cap(slice))
    //slice len= 4 slice cap= 10
	slice[0]=100
    slice[2]=100
    fmt.Println(slice)//[100 0 100 0]
	*/
    
}