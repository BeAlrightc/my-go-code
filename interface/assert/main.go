package main
import (
	"fmt"
)

//声明定义一个接口
type Usb interface{
	Start()
	Stop()
}

type Phone struct {
    name string
}

type Camera struct {
	name string
}
//让phone实现usb接口的方法
func (p Phone) Start(){
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop(){
	fmt.Println("手机停止工作。。。")
}

func (p Phone) Call(){
	fmt.Println("手机开始打电话。。。")
}


func (c Camera) Start(){
	fmt.Println("相机开始工作。。。")
}
func (c Camera) Stop(){
	fmt.Println("相机停止工作。。。")
}
//计算机
type Computer struct{

}

//编写一个方法Working方法，接收一个Usb接口类型变量
//只要是实现了Usb接口（所谓实现Usb接口，就是指实现了Usb接口声明所有方法）
func (c Computer) Working (usb Usb){
	//通过usb接口变量来调用Start和Stop方法
     usb.Start()
	 //如果usb是指向Phone结构体变量，则还需要调用Call()方法
	 //类型断言
	 if phone, ok := usb.(Phone); ok {
		phone.Call()
	 }//否则断言失败还是继续执行下列方法
	 usb.Stop()
}

func main(){
   //定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
   //这里就体现出多态数组
   var usbArr  [3]Usb
   usbArr[0] = Phone{"iphone"}
   usbArr[1] = Phone{"小米"}
   usbArr[2] = Camera{"佳能"}
   //fmt.Println(usbArr) //[{iphone} {小米} {佳能}]
   //遍历
   var computer Computer
   for _,v := range usbArr{
          computer.Working(v)
		  fmt.Println()
   }
	
}
