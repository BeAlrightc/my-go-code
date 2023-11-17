package main
import (
	"fmt"
)
type Point struct{
	x int
	y int
}
func main() {
   var a interface{}
   var point Point = Point{1,2}
   a = point //OK
   //如何将A赋给一个Point变量？
   var b Point
   //b = a //可以吗？ ==》erro
   b = a.(Point) //解决办法，类型断言
   fmt.Println(b) //1 2

//    //类型断言的其他案例
//    var x interface{}
//    var b2 float32 = 1.1
//    x = b2 //空接口可以接收任意类型
//    //x = >float32 【使用类型断言】
//    y := x.(float32)
//    fmt.Printf("y的类型是%T 值是%v",y,y) //y的类型是float32 值是1.1


//类型断言(带检测)
var x interface{}
var b2 float32 = 1.1
x = b2 //空接口可以接收任意类型
//x = >float32 【使用类型断言】
//y,ok := x.(float64)
//简写

if  y,ok := x.(float64);ok {
	fmt.Println("转换成功了")
    fmt.Printf("y的类型是%T 值是%v",y,y) //y的类型是float32 值是1.1
}else{
	fmt.Println("转换失败")
}
fmt.Println("检测失败了无妨，代码不要停")

}