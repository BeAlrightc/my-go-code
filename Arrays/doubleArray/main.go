package main
import (
	"fmt"
)
/*
000000
001000
020300
000000
定义声明一个二维数组
*/
func demo1(){
	var arr [4][6]int
	 //赋初值
	 arr[1][2]=1
	 arr[2][1]=2
	 arr[2][3]=3
	 //遍历二维数组。按照要求输出图形
	 for i :=0; i< 4; i ++{
		for j :=0 ;j<6;j++{
			fmt.Print(arr[i][j]," ")
		}
		fmt.Println()
	 }	
}
func arrmemory(){
	var arr2 [2][3]int 
	arr2 [1][1]=10
	fmt.Println(arr2)

	fmt.Printf("arr2[0]的地址是%p\n",&arr2[0])
	//arr2[0]的地址是0xc04207a030 与arr2[1]相差24个字节
	fmt.Printf("arr2[1]的地址是%p\n",&arr2[1])
	//arr2[1]的地址是0xc04207a048

	fmt.Printf("arr2[0][0]的地址是%p\n",&arr2[0][0])
	//arr2[0][0]的地址是0xc04207a030
	fmt.Printf("arr2[1][0]的地址是%p\n",&arr2[1][0])
	// arr2[1][0]的地址是0xc04207a048
    
}
func demo3 (){
	var arr3 [2][3]int = [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println("arr3=",arr3)//arr3= [[1 2 3] [4 5 6]]
}

func bainli(){
	//演示二维数组的遍历
	var arr3 = [2][3]int{{1,2,3},{4,5,6}}

	//for循环来遍历
	for i :=0;i<len(arr3);i++{
		for j :=0; j < len(arr3[i]); j++{
			fmt.Printf("%v\t",arr3[i][j])
		}
		fmt.Println()
	}
	// for -range遍历
	for i,v := range arr3{
		for j , v2 := range v{
			fmt.Printf("arr3[%v][%v]=%v\t",i,j,v2)

		}
		fmt.Println()
		
	}
}
func main(){
 //demo1()
 //arrmemory()
 //demo3()
 bainli()
}