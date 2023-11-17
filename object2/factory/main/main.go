package main
import (
	"fmt"
	"go_code/object2/factory/model"
)

func main(){
	//创建一个Student的实例
	// var stu = model.Student{
	// 	Name : "Tom",
	// 	Score : 78.9,
	// }
	//当student结构体首字母是小写的我们可以通过工厂模式解决
	var stu = model.NewStudent("tom~",88.8) //stu是一个结构体指针类型
	fmt.Println(*stu)
	fmt.Printf("name=%v score的值为=%v",stu.Name,stu.GetScore())
	//实际上底层会自动优化这样
	//fmt.Printf("name=%v score的值为=%v",stu.Name,(*stu).GetScore())
}
