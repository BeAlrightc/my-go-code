package main

import (
	"fmt"
)
func modify(map1 map[int]int){
map1[10] = 900
}

//定义一个学生结构体
type Stu struct {
	Name string
	Age int
	Address string
}

func main() {
	//map是引用类型，
	// 遵守引用类型传递的机制，在一个函数接收map，修改后，会直接修改原来的map
  map1 := make(map[int]int)
  map1[1] = 90
  map1[2] = 88
  map1[10] = 1
  map1[20] = 2
  modify(map1)
  //观察结果如果map1[10]= 900说明map是引用类型
  fmt.Println(map1) //map[10:900 20:2 1:90 2:88]

  //map的容量达到后，再想map增加元素，会自动扩容，
// 并不会发生panic,也就是说map能动态的增长键值对(key -value)


//map的value也经常使用struct类型，更适合管理复杂的
// 数据(比前面value是一个map更好)，比如value为student结构体
//1.map的key为学生的学号是唯一的
//2.map的value为结构体，包含学生的名字，年龄，地址
students :=make(map[string]Stu, 10)
//创建2个学生
stu1 := Stu{"tom",18,"北京"}
stu2 := Stu{"jhon",19,"上海"}
students["no1"] = stu1
students["no2"] = stu2
fmt.Println(students) //map[no1:{tom 18 北京} no2:{jhon 19 上海}]

//遍历各个学生的信息
for k,v := range students {
	fmt.Printf("学生的编号是%v\n",k)
	fmt.Printf("学生的名字是%v\n",v.Name)
	fmt.Printf("学生的年龄是%v\n",v.Age)
	fmt.Printf("学生的住址是%v\n",v.Address)
	fmt.Println(" ")
}

}