package main
import (
	"fmt"
)
type Student struct{
   name string
   gender string
   age int
   id int
   score float64

}
func (stu *Student) say() string{
	//使用fmt.Sprintf()函数将值转换为字符串
    inforstr :=fmt.Sprintf("student的信息如下 name = [%v] gender=[%v] age = [%v] id=[%v] score=[%v]",stu.name,stu.gender,stu.age,stu.id,stu.score)
	 return inforstr
}

type Box struct{
	length float64
    width float64
	high  float64
}
func (b Box) getVolumn() float64{
    return b.length * b.width * b.high
}


type Vistor struct{
	Name string
	Age int
}
func (vistor Vistor) showPrice(){
	if vistor.Age >18{
		fmt.Println("门票价格为20元")
	}else{
		fmt.Println("免费")
	}

}

func main(){
var stu = Student{"小红","男",34,1001,99.5}
info :=stu.say()
fmt.Println(info)
//测试box
var box Box
box.length = 1.1
box.width = 2.0
box.high = 3.0
Volum := box.getVolumn()
fmt.Printf("这个黑子的体积是%.2f",Volum)

var vio Vistor

for {
	fmt.Println("请输入你的名字")
	fmt.Scanln(&vio.Name)
	if vio.Name == "n" {
		fmt.Println("退出程序")
	    break
	}
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&vio.Age)
	vio.showPrice()
}


}