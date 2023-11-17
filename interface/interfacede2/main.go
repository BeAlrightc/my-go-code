package main
import (
	"fmt"
)
type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}
type AInterface interface {
	BInterface
	CInterface
	test03()
}
//如果需要实现AIerface，就需要把BInterface和CInterface的方法都实现
type Stu struct{

}
//把所要实现的接口所有的方法都实现一下
func (stu Stu) test01(){
  fmt.Println("这是Test01")
}
func (stu Stu) test02(){
	fmt.Println("这是Test02")
}
func (stu Stu) test03(){
	fmt.Println("这是Test03")
}

type T interface {

}
func main(){
    //实践
	var stu Stu
	var a AInterface = stu
	a.test01()
	a.test02()
	a.test03()
	var t T =stu
	fmt.Println(t)
	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	fmt.Println(t2,test02)
}