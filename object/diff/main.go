package main
import (
	"fmt"
)

type Person struct{
	Name string
}

//方法
func (p Person) test03(){
  fmt.Print("test03=",p.Name)
}
//函数
func test01(p Person){
    fmt.Println(p.Name)
}

func test02(p *Person){
    fmt.Println(p.Name)
}
func main(){
 p := Person{"tom"}
  //test01(p)
//  test01(&p)//不可以传地址
//test02(&p)

p.test03()

//(&p).test03()
}
