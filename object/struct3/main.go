package main
import (
	"fmt"
)

//结构体
type Point struct {
	 x int
	 y int
}
//结构体
type Rect struct {
	leftup,rightDown Point
}

type Rect2 struct {
	leftup,rightDown *Point
}

func main (){
    
	r1 := Rect{Point{1,2},Point{3,4}}

	//r1有4个int,在内存中连续分布
	//打印地址
	fmt.Printf("r1.leftup.x 地址=%p r1.leftup.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p\n",
    &r1.leftup.x,&r1.leftup.y,&r1.rightDown.x,&r1.rightDown.y)
	//输出结果如下
	//r1.leftup.x 地址=0xc04200a2c0 r1.leftup.y 地址=0xc04200a2c8 
	// r1.rightDown.x 地址=0xc04200a2d0 r1.rightDown.y 地址=0xc04200a2d8
    
	
	//r2有两个*Point类型，这两个*Point类型的本身地址也是连续的
	//但是他们指向的地址不一定是连续的
    
	r2 := Rect2{&Point{10,20}, &Point{30,40}}
	fmt.Printf("r2.leftup本身地址=%p  r2.rightDown 本身地址=%p \n",
    &r2.leftup,&r2.rightDown)
	//输出结果为:r2.leftup本身地址=0xc0420421c0  r2.rightDown 本身地址=0xc0420421c8
    
	//他们指向的地址不一定是连续的 这个要看编译器或系统运行时而定
	fmt.Printf("r2.leftup指向地址=%p  r2.rightDown 指向地址=%p \n",
    r2.leftup,r2.rightDown)
	//r2.leftup指向地址=0xc0420120b0  r2.rightDown 指向地址=0xc0420120c0


}