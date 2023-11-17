package main
import (
	"fmt"
)
//,Monkey结构体
type Monkey struct {
	Name string
}

//声明接口
type BirdAble interface {
	Flying()
}

//声明接口
type FishAble interface {
	Swimming()
}
func (this *Monkey) climbing() {
	fmt.Println(this.Name,"生来会爬树")
}

//创建LittleMonkey结构体
type LittleMonkey struct {
	Monkey //继承
}

//让littleMonkey实现BirdAble的Flying()方法
func (this *LittleMonkey) Flying(){
	fmt.Println(this.Name,"通过学习会飞翔")
}
//littleMonkey实现FishAble的Swimming()方法
func (this *LittleMonkey) Swimming(){
	fmt.Println(this.Name,"通过学习会游泳")
}
func main(){
    
	//创建LittleMonkey实例
	monkey := LittleMonkey{
		Monkey {
			Name : "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
  
}