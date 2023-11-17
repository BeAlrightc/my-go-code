package main
import (
	"fmt"
)
func arry(){
	//从终端输入5个成绩，保存到float64数组，并输出
	var score[5]float64
	for i:=0;i<5;i++{
		fmt.Printf("请输入第%v个元素的值：",i)
		fmt.Scanln(&score[i])
	}
	// fmt.Println("score的值是",score)
	//遍历数组打印
	for i :=0;i< len(score);i++{
		fmt.Printf("score[%v]=%v\n",i,score[i])
	}
}
func main(){

// 四种初始化数组的方式
//way1
var numArr01 [3]int=[3]int{1,2,3}
fmt.Println("numArr01=",numArr01)
//输出结果为：numArr= [1 2 3]
//way2
var numArr02 =[3]int{5,6,7}
fmt.Println("numArr02=",numArr02)
//输出结果为：numArr02= [5 6 7]
//way3
var numArr03 =[...]int{8,9,10}
//这里的[...]是规定的写法
fmt.Println("numArr03=",numArr03)
//numArr03= [8 9 10]
// way4
var numArr04 =[...]int{1:800,0:900,2:999}
fmt.Println("numArr04=",numArr04)
//numArr04= [900 800 999]
//类型推导
numArr05 :=[...]string{1:"tom",0:"jfon",2:"feilipu"}
fmt.Println("numArr05=",numArr05)


	//arry()
	//var intArr [3]int//int占8个字节
	// //当我们定义完数组后，数组的各个元素有默认值0
	// fmt.Println(intArr)//[0 0 0]
	// fmt.Printf("数组的地址是：%p\n",&intArr)
	// //数组的地址是：0xc0420082c0
    // fmt.Printf("数组首地址是：%p，地址intArr[1]的地址是%p",&intArr[0],&intArr[1])
    // //数组首地址是：0xc04205c0a0，地址intArr[1]的地址是0xc04205c0a8

}