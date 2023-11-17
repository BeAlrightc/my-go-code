package main
import (
	"fmt"
)

type AInterface interface {
	Say()

}

type Stu struct {
	Name string
}
func (stu Stu)Say(){
	fmt.Println("Stu Say()")
}

type integer int

func (i integer)Say() {
	fmt.Println("integer Say i = ",i)
}
type BInterface interface{
	Hello()
}
type Monster struct{

}
func (m Monster) Hello(){
	fmt.Println("Monster Hello()~~")
}
func (m Monster) Say(){
	fmt.Println("Monster Say()~~")
}
//此时刻Monster就实现了AInterface和BInterface

func main(){
   var stu Stu  //结构体变量，实现了Say() 实现了AInterface这个接口
   var a AInterface =stu
   a.Say() //Stu Say()
   var i integer = 10
   var b AInterface = i
   b.Say() //integer Say i =  10

   var monster Monster
   var a2 AInterface = monster //Monster Say()~~
   var b2 BInterface = monster //Monster Hello()~~
   a2.Say()
   b2.Hello()
}