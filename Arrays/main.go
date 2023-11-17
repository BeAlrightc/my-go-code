package main
import (
	"fmt"
)
/*
一个养鸡场有6只鸡，他们的体重分别是3kg .5kg 1kg 3.4kg 
2kg 50kg请问这六只鸡的总体重是多少？平均体重是多少?
请你编写一个程序
*/
func main(){
 //使用数组来解决问题
 //1.定义一个数组
 var hens [6]float64
//2.给数组的每个元素赋值操作，元素下标从0开始
hens[0] =3.0 //hens数组的第一个元素赋值
hens[1] =5.0
hens[2] =1.0
hens[3] =3.4
hens[4] =2.0
hens[5] =50.0
//3.遍历数组求出总体重
totalweight :=0.0
for i :=0;i< len(hens);i++{
    totalweight += hens[i]
  }
  //4.求出总体重
//平均体重
avgweight := fmt.Sprintf("%.2f",totalweight/float64(len(hens)))
  fmt.Printf("鸡的总体重是：%v,平均体重是%v",totalweight,avgweight)
}