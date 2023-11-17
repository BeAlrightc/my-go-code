package main
import (
	"fmt"
)

type Person struct{
	Name string
}

//给A类型绑定一个方法
func (p Person) test() {
	p.Name = "jack"
	fmt.Println("test()",p.Name)
}
func (p Person) speak(){
	fmt.Printf("%v是一个好人\n",p.Name)
}

func (p Person) jisuan() {
	var sum int = 0
	for i :=1;i <=1000;i++{
		sum += i
	}
	fmt.Printf("1+..+1000的结果是%v\n",sum)
}
//给Person结构体添加jisuan2方法,该方法可以接收一个数n，计算从1+...n的结果
func (p Person) jisuan2(n int) {
    var sum int = 0
	for i :=1;i <=n;i++{
		sum += i
	}
	fmt.Printf("1+..+%v的结果是%v\n",n,sum)
}
//给Person结构体添加getSum方法,可以计算两个数的和并返回结果。
func (p Person) getsum(n1 int,n2 int) int{
     return n1 + n2
}
func main(){
	//定义一个Person实例
	var p Person
	p.Name = "小红"
	//p.test() //调用方法
	//fmt.Println("main()",p.Name)//小红
	p.speak()
	p.jisuan()
	p.jisuan2(10000)
	res := p.getsum(10,20)
	fmt.Println("res=",res)

}