package main
import (
	"fmt"
)
var name string="tom" //全局变量
func test01(){
	fmt.Println(name)
}

func test02(){//编译器采用就近原则
	name ="jack"
	fmt.Println(name) //jack
}
var age int= 50
var Name string= "jhon"
//函数
func test(){
	//age和Name作用域就只在test函数内部
	 age := 10
	 Name := "tom"
	fmt.Println("age=",age)
	fmt.Println("Name=",Name)
}
func main(){
fmt.Println(name) //tom
test01()//tom
test02()//jack
test01()//jack

	//test() //10 tom
//   fmt.Println("age=",age) //会报错
//   fmt.Println("age=",age)//50 
//   fmt.Println("Name=",Name)//jhon

//   for i :=0 ;i<=10;i++{
// 	fmt.Println("i=",i)
//   }
//  fmt.Println("i=",i)会报错

}