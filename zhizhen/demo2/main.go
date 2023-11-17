package main
import (
	"fmt"
)

//声明一个测试函数(测试)
func test() bool {
	fmt.Println("test....")
return true
}

func main() {
	var i  int = 10
	//短路与的
	//说明：i<9 这个为false 第二条件根本不会去判断不会调用test函数
   if i < 9 && test(){
	fmt.Println("ok")
   }
   //短路或
   //说明：i>9 这个为true 第二条件根本不会去判断不会调用test函数
   if i > 9 || test(){
	fmt.Println("ok")
   }


 /*
	//演示如何使用逻辑运算符
  var age int =40
  if  age > 30 && age < 50{
	fmt.Println("ok1")
  }

  if  age > 30 && age < 40{
	fmt.Println("ok2")
  }

  //演示逻辑或的使用  ||
  if  age > 30 || age < 50{
	fmt.Println("ok3")
  }

  if  age > 30 || age < 40{
	fmt.Println("ok4")
  }
 //演示逻辑非的使用 !
 if  !(age > 30) {
	fmt.Println("ok5") //不输出
  }
  */

}