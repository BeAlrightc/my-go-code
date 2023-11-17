package main
import (
	"fmt"
)
/*
案例2：请编写一个程序，要求如下

1. 声明一个结构体Circle，字段为radius
2. 声明一个方法area和CIrcle绑定，可以返回面积
3. 提示画出area执行过程+说明
*/
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
    return 3.14 * c.radius * c.radius
}

//为了提高效率，通常我们方法和结构体的指针类型进行绑定
func (c *Circle) area2() float64 {
	fmt.Printf("c是*cicle指向的地址=%p\n",c)
	c.radius = 10
	//因为c是一个指针，因此我们标准访问其字段的方式是(*c)
    return 3.14 * (*c).radius * (*c).radius
	//(*c).radius等价为c.radius
}

func main(){
	//创建一个结构体
	// var c Circle
	// c.radius = 4.0
	// area := c.area()
    // fmt.Printf("圆的面积为%v\n",area)
	//创建一个cicle变量
	var c Circle
	fmt.Printf("main的c结构体变量地址 =%p\n",&c)
	c.radius = 6.0
	//res2 := (&c).area2()
	//编译底层做了优化,(&c).area2()等价c.area2()
	//因为编译器会自动加上
	res2 := c.area2()
	fmt.Printf("圆的面积为%v\n",res2) //314
	fmt.Println("c.radius = ",c.radius) //10
	
	/*
    运行结果如下
	main的c结构体变量地址 =0xc0420120a0
    c是*cicle指向的地址=0xc0420120a0
    圆的面积为314
    c.radius =  10

	*/
}