package main
import (
	"fmt"
	//引入包
	"go_code/functions/funint/utils"
)

var age = test()

//为了看到全局变量是先被初始化的，我们这里先写函数
func test()int{  //1
	fmt.Println("test")
	return 90
}

//init函数，通常可以在init函数中完成初始调用
func init(){//先执行 2
	fmt.Println("init()....")
}

func main(){ //3
  fmt.Println("main()...age=",age)
  fmt.Println("Age=",utils.Age,"Name",utils.Name)
}