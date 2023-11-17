package main
import (
	"fmt"
)
//冒泡排序
func BubbleSort(arr *[5]int){
	fmt.Println("排序前的ar=",(*arr))
	temp :=0//临时变量用来做交换的
	//完成第一轮排序（外层排序）
	for j :=0;j < 4;j++{
		if (*arr)[j] > (*arr)[j+1]{
			//交换
            temp = (*arr)[j]
			(*arr)[j]=(*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第一次排序过后arr=",(*arr))//[24 69 57 13 80]

	//完成第二轮排序（外层排序）
	for j :=0;j < 3;j++{
		if (*arr)[j] > (*arr)[j+1]{
			//交换
            temp = (*arr)[j]
			(*arr)[j]=(*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第三次排序过后arr=",(*arr))//[24 57 13 69 80]
    //完成第二轮排序（外层排序）
	for j :=0;j < 2;j++{
		if (*arr)[j] > (*arr)[j+1]{
			//交换
            temp = (*arr)[j]
			(*arr)[j]=(*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第三次排序过后arr=",(*arr))//arr= [24 13 57 69 80]

	//完成第四轮排序（外层排序）
	for j :=0;j < 1;j++{
		if (*arr)[j] > (*arr)[j+1]{
			//交换
            temp = (*arr)[j]
			(*arr)[j]=(*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第四次排序过后arr=",(*arr))////完成第二轮排序（外层排序）
	for j :=0;j < 2;j++{
		if (*arr)[j] > (*arr)[j+1]{
			//交换
            temp = (*arr)[j]
			(*arr)[j]=(*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第四次排序过后arr=",(*arr))//[13 24 57 69 80]

}

func BubbleSort2(arr *[5]int){
	fmt.Println("排序前的ar=",(*arr))
	temp :=0//临时变量用来做交换的
	len :=len((*arr))
	//完成第一轮排序（外层排序）
	for i :=0;i<len-1;i++{
	for j :=0;j < len-i-1;j++{
		if (*arr)[j] > (*arr)[j+1]{
			//交换
            temp = (*arr)[j]
			(*arr)[j]=(*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
}
	fmt.Println("排序过后arr=",(*arr))

}


	func main(){

	//定义一个数组
	arr := [5]int{24,69,80,57,13}
	//将数组传递给一个函数，完成排序
    //BubbleSort(&arr)
    BubbleSort2(&arr)

}