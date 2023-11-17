package main
import (
	"fmt"
)
/*
Golang中的方法作用在**指定的数据类型**上的
（即：**和指定的数据类型绑定**）因此自定义类型，
都可以有方法。而不仅仅是struct,比如int,float32等都可以有方法
*/
type integer int
func (i integer) print() {
	fmt.Println("i=",i)
}
//编写一个方法修改i的值
func (i *integer) ch() {
	*i = *i +1
}

type Student struct {
	Name string
	Age int
}
//给Student实现方法String()
func (stu *Student) String() string {
   str := fmt.Sprintf("Name=[%v] Age=[%v]",stu.Name,stu.Age)
   return str
}
func main(){
    var i integer = 10
	i.print() //1 = 10
	i.ch()
	i.print() //i =11

	//定义一个Student变量
	stu := Student{
		Name : "tom",
		Age : 20,
	}
	//如果实现了*Student类型的String()方法会自动调用String()要传入&stu
	fmt.Println(&stu) 
}