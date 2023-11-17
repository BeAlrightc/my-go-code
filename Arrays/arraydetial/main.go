package main
import (
	"fmt"
)
//函数
func test01(arr [3]int){
  arr[0] = 88
  
}

//函数
func test02(arr *[3]int){
	(*arr)[0] = 88
	
  }

func main(){
	/*
   //数组创建后，如果没有赋值，有默认值(0值)
   //1.数值(整数系列，浮点数系列)=》0
   //2.字符串==》""
   //3.bool类型 ==》flase
  var arr01 [3]float32
  var arr02 [3]string
  var arr03 [3]bool
  fmt.Printf("arr01=%v arr02=%v arr03=%v",arr01,arr02,arr03)
//输出结果 arr01=[0 0 0] arr02=[  ] arr03=[false false false]

//数组的下标是从0开始
var arr04 [3]string //0-2
//fmt.Println(arr04[3])// 报错，原因是数组越界
*/
/*
arr := [3]int{11,22,33}
test01(arr)
fmt.Println("main里面的arr的值是",arr)//输出结果仍然是：[11 22 33]	
*/
arr := [3]int{11,22,33}
test02(&arr)
fmt.Println("main里面的arr的值是",arr)
//输出的内容是
// main里面的arr的值是 [88 22 33]

/*
	//1）数组是多个相同类型数据的组合，一个数组一旦
	// 声明/定义了，其长度是固定的，不能动态变化
	var arr01 [3]int
	arr01[0] =1
	arr01[1] =30
	//arr01[2] =1.1//这里会报错类型不一致
    arr01[2] =90
    // arr01[3] =890//数组会发生越界，超出指定范围长度

	fmt.Println(arr01)
	*/

}