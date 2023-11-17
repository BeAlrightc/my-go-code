package main
import (
	"fmt"
)
//二分法查找的函数
/*
二分法查找的思路：比如我们要查找的数是findVal
1.arr是有一个有序数组，并且是从小到大排序
2.先找到中间的下标middle =(leftindex + rightindex)/2然后让中间的值和findval进行比较
逻辑：
2.1如果arr[middle]>findval,就应该问 leftindex----(middle -1)
2.1如果arr[middle]<findval,就应该问 middel+1----right
2.3如果Arr[middle]==findVal就找到
对上面的逻辑进行递归执行

递归退出条件
if lefetindex > rightindex
//找不到
return ..
思路---代码

*/

func BinaryFind(arr *[6]int,leftindex,rightindex,findVal int){
	//判断leftIndex是否大于rightindex
	if leftindex >rightindex{
		fmt.Println("没有找到")
		return
	}


	
	//先找到中间的下标
	middle :=(leftindex + rightindex) /2
	if (*arr)[middle] > findVal{
		//说明我们要查找的数，应该在 leftIndex ---middel-1之间
		BinaryFind(arr,leftindex,middle -1,findVal)
	}else if (*arr)[middle] < findVal{
		//说明我们要查找得数在middel + -----rightindex
		BinaryFind(arr,middle +1,rightindex,findVal)
	}else{
		//找到了
		fmt.Printf("找到了下标为%v\n",middle)
	}


}



func main(){
    arr := [6]int{1,8,10,89,1000,1234}
    BinaryFind(&arr,0,len(arr)-1,1000)//找到了下标为4
}