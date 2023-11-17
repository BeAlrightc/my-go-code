package main
import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
}

func main() {
	//定义一个存放任意数据类型的管道3个数据
	// var  allChan chan interface{}
    allChan := make(chan interface{},3)

	allChan<-10
	allChan<-"tom jack"
	cat :=Cat{"小花猫",4}
	allChan<- cat
 
//我们希望获得管道中的第三个元素，则先将前2个推出
<-allChan
<-allChan

newCat :=<-allChan //从管道中取出来的cat是什么
fmt.Printf("newCat=%T,newCat=%v\n",newCat,newCat)//newCat=main.Cat,newCat={小花猫 4}
//下面的写法是错误的,编译不通过，则使用类型断言就可以通过
// fmt.Printf("newCat.Name=%v",newCat.Name)
a :=newCat.(Cat)
fmt.Printf("newCat.Name=%v",a.Name)//newCat.Name=小花猫
}