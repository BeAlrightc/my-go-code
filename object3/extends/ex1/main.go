package main
import (
	"fmt"
)

type A struct {
	Name string
	age int
}

func (a *A) SayOk(){
   fmt.Println(" A SayOk",a.Name)
}
func (a *A) hello(){
	fmt.Println(" A Hello",a.Name)
 }
type B struct {
	A
}
func (b *B)SayOk() {
	fmt.Println("B SayOK",b.Name)
}
func main(){
   var b B
//    b.A.Name = "tom"
//    b.A.age = 19
//    b.A.SayOk()
//    b.A.hello()
   b.Name = "jack"
   b.age  = 100
   b.SayOk()  //B SayOK jack
   b.hello() // A Hello jack


}