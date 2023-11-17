 package main
 import (
	"fmt"
	"reflect"
 )

 //专门反射演示
 func reflectTest01(b interface{}) {
	//通过反射获取的传入的变量 type,kind,值
	//1.先获取到reflect.Type
	rTyp :=reflect.TypeOf(b)
	fmt.Println("rTyp=",rTyp) //int
	//2.获取到reflectValue
	rVal :=reflect.ValueOf(b) //100
	// n1 := 10
	// n2 := 2 + rVal
	// fmt.Println("n2=",n2)//erro 

	// n1 := 10
	n2 := 2 + rVal.Int()
	fmt.Println("n2=",n2) //102

	fmt.Printf("rVal=%v type=%T\n",rVal,rVal) //rVal=100 type=reflect.Value

	//下面我们将rVal转成interface{}
	iv := rVal.Interface()
	//将interface{}通过断言转成需要的类型
	num2 := iv.(int)
	fmt.Println("num2=",num2) //num2= 100
 }

 //专门演示对结构体的反射
 func reflectTest02(b interface{}) {
	//通过反射获取的传入的变量 type,kind,值
	//1.先获取到reflect.Type
	rTyp :=reflect.TypeOf(b)
	fmt.Println("rTyp=",rTyp) //rTyp= main.Student

	//2.获取到reflectValue
	rVal :=reflect.ValueOf(b)

	//3.获取变量对应的Kind
	//(1)rVal.Kind() ==>
	kind1 :=rVal.Kind()
	//(2)rTyp.Kind() ==>
	kind2 :=rTyp.Kind()
	fmt.Printf("kind =%v kind=%v\n",kind1,kind2)  //kind =struct kind=struct 两个都是一样的
	
	//下面我们将rVal转成interface{}
	iv := rVal.Interface()
	fmt.Printf("iv=%v ,iv type=%T\n",iv,iv) //iv={Tom 20} ,iv type=main.Student
	//将interface{}通过断言转成需要的类型
	// fmt.Printf("iv=%v ,iv type=%T name\n",iv,iv,iv.Name) //erro无法取出name的值
	//所以先要进行断言
	stu,ok :=iv.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n",stu.Name)//stu.Name=Tom
	}

	
 }

 type Student struct {
	Name string
	Age int
 }
 func main() {
	//-1.请编写一个函数，演示对（基本数据类型、interface{}、reflect.Value）进行反射的基本操作。
	
	//1、先定义一个int
	// var num int = 100
	// reflectTest01(num)

	//2.定义一个Student的实例
	stu := Student{
		Name : "Tom",
		Age : 20,
	}

	reflectTest02(stu)

 }