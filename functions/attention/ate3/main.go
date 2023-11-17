package main

import (
	"fmt"
)

//编写一个sum函数，可以求出 1到多个int的和
func sum(n1 int,args... int) int{
	sum := n1
	//遍历arges
	for i :=0;i <len(args);i++{
		sum += args[i] //args[0]表示取出args切片的第一个元素的值，其他依次类推
	}
	return sum
}

func sum2(n1,n2 float32)float32{
	fmt.Printf("n1 type=%T\n",n1)
	return n1 +n2
}

func swap(n1 *int,n2 *int){
	//定义一个临时变量
	t :=*n1
	*n1=*n2
	*n2=t
} 
func main(){
//测试可变参数的使用
// res :=sum(10,0,-1,90,10)
// fmt.Println("res=",res)
fmt.Println("sum2=",sum2(1,2))

a :=10
b :=20
fmt.Printf("a=%v,b=%v",a,b)
swap(&a,&b)//传入的是地址
fmt.Printf("a=%v,b=%v",a,b)
}