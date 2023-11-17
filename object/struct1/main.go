package main
import (
	"fmt"
)
//指针，slice和map的零值都是nil即没有分配空间
//如果结构体的字段类型是这些，需要先make后才可以使用
type Person struct {
	Name string
	Age int
	scores [5]float64
	ptr *int //指针
    slice []int//切片
	map1 map[string]string//切片
}

type Monster struct {
	Name string
	Age int
}

func main() {

	//定义结构体变量
	var p1 Person
	fmt.Println(p1) //未赋值之前输出的是{ 0 [0 0 0 0 0] <nil> [] map[]}
    
	if p1.ptr == nil {
		fmt.Println("ok1")
	}

	if p1.slice == nil {
		fmt.Println("ok2")
	}

	if p1.map1 == nil {
		fmt.Println("ok3")
	}

	//使用slice,之前一定要make
	p1.slice = make([]int,10)
	p1.slice[0] = 100 

	//使用map也一定先要make操作
    p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom"
	fmt.Println(p1)

	//不同结构体变量的字段是独立的，互不影响，
	// 一个结构体变量字段的更改，不影响另一个。结构体是值类型
    var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 500

	monster2 := monster1//默认结构体是值类型，值拷贝
	monster2.Name = "青牛精"
	fmt.Println("monster1=",monster1)//monster1= {牛魔王 500}
	fmt.Println("monster2=",monster2)//monster2= {青牛精 500}
	

}