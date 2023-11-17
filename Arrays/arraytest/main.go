package main
import (
	"fmt"
	"math/rand"
	"time"
)
func th1(){//自己写的
	//声明一个数组
	var arr [26]byte
	arr[0]='A'
	for i :=1;i<26;i++{
		arr[i]=arr[i-1]+1
	}
	for a,v :=range arr{
		fmt.Printf("arr[%d]=%c ",a,v)
	}
}

func the1(){//老师写的
	var myChars [26]byte
	for i :=0; i < 26; i++ {
		myChars[i] ='A'+byte(i)//注意将i =>byte
	}
	for i :=0; i < 26; i++{
        fmt.Printf("%c  ",myChars[i])
	}
}
//请求出一个数组的最大值，并得到对应的下标
/*
1.声明一个数组[6]int{12,56,7,9,23,1}
2.假定第一个数为最大值，下标就为0
3，然后从第二个元素开始循环比较，如果发现有更大则交换
*/
func the2(){
  arr :=[6]int{12,56,7,90,23,1}
  var max int =arr[0]
  maxValIndex :=0
  for i :=1;i<len(arr);i++{
	if arr[i]>max{
		max=arr[i]
		maxValIndex=i
	}
  }
  fmt.Printf("max=%v,index=%v",max,maxValIndex)
  
}
//请求出一个数组的和以及平均值。for-range
func suma(){
	//1.声明一个数组arr :=[6]int{12,56,7,90,23,1}
	//2.求出sum
	//3.求出平均值
	arr :=[6]int{12,57,7,90,23,1}
	sum :=0
	for _,v :=range arr{
		//累积求和
		sum += v
	}
	fmt.Printf("数组的和是%v,数组的平均值是%.2f",sum,float64(sum)/float64(len(arr)))
}
/*
要求：随机生成5个数，并将其反转打印
思路
1.随机生成5个数，rand.Intn()函数
2，当我们得到随机数后就放到一个数组中 int数组
3.反转打印,交换的次数是len/2 2.倒数第一个和第一个交换倒数第二个与第二个进行交换
*/
func fanzhuan(){
	var intArr3 [5]int
	len :=len(intArr3)
	//为了每次生成的随机数都不一样，我们需要给一个seed值
	rand.Seed(time.Now().UnixNano())
for i := 0; i < len;i++{
	intArr3[i] =rand.Intn(100) //0<=n<=100
}
fmt.Println("交换前：",intArr3)
//3.反转打印,交换的次数是len/2 2.倒数第一个和第一个交换倒数第二个与第二个进行交换

temp :=0 //作为一个临时变量用于交换操作
for i := 0; i < len/2;i++{
	temp =intArr3[len-1 -i] //倒数第n个和第n个元素进行交换
	intArr3[len-1 -i]=intArr3[i]
	intArr3[i] = temp

}
fmt.Println(intArr3)

fmt.Println("交换后：",intArr3)


}


func main(){
  //th1()
  //the1()
  //the2()
//   suma()
  fanzhuan()


}