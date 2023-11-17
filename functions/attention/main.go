package main
import (
	"fmt"
)
//函数中的变量是局部的，函数外不生效
func test(){
	//n1是test函数的局部变量，只能在test函数中使用
	//var n1 int = 10

}
func test2(n1 int) {
	
   n1 = n1 + 10
   fmt.Println("test2(n1)=",n1)
}
//n1就是*int类型
func test3(n1 *int) {
	fmt.Println("n1的地址%v",&n1)
	*n1 = *n1 + 10
	fmt.Println("test03(n1)=",*n1)//30
 }
func main(){
	//这里不能使用n1，因为n1是test函数的局部变量
  // fmt.Println("n1=",n1)
//   n1 := 20
//   test2(n1)
//   fmt.Println("n1=",n1)
num := 20
fmt.Println("num的地址%v",&num)
test3(&num)
fmt.Println("main() num=",num)//30

}