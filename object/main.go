package main
import (
	"fmt"
)
/*
张老太养了两只猫，一个名字叫小白，今年3岁，白色。
还有一只叫小花，今年100岁，花色。请编写一个程序。
当用户输入小猫的名字时，就显示该猫的名字，年龄，颜色。
如果用户输入的小猫名错误，则显示张老太没有这只小猫

*/

// func vars(){
// 	//1.使用变量的处理方式
// 	var cat1Name string = "小白"
// 	var cat1Age int = 3
// 	var cat1Color string = "白色"
    
// 	var cat2Name string = "小化"
// 	var cat2Age int = 100
// 	var cat2Color string = "花色"
// 	//2.使用数组来解决
// 	var catNames [2]string = [...]string{"小白","小花"}
//     var catAges [2]int = [...]int{3,100}
// 	var catColor [2]string = [...]string{"白色","花色"}
// 	//
// }

//定义一个cat结构体,将cat的各个字段/属性放入cat结构体进行管理
type Cat struct {
	Name string
	Age int
	Color string
	Hobby string
	Score [3] int //数组类型
	
}



func main() {
//创建一个cat结构体变量
var cat1  Cat  //类似于var a int
//输出cat1的结构体地址
fmt.Printf("cat1的地址是%p\n",&cat1) //cat1的地址是0xc0420460c0
cat1.Name = "小白"
cat1.Age = 3
cat1.Color = "白色"
cat1.Hobby = "吃<.))>>>"
fmt.Println("cat1=",cat1)

fmt.Println("猫猫的信息如下：")
fmt.Println("name=",cat1.Name,"Age=",cat1.Age,"Color=",cat1.Color)
fmt.Println("cat1.Hobby",cat1.Hobby)
//输出结果如下：
/*cat1= {小白 3 白色}
猫猫的信息如下：
name= 小白 Age= 3 Color= 白色
*/

}