package main
import (
	"fmt"
	// "time"
	"errors"
)

func test(){
	//使用defer +recover来捕获处理异常
	defer func() {
		//err :=recover() //recover是内置函数，可以捕获到异常
		if err :=recover(); err != nil { //说明捕获到错误
			fmt.Println("err=",err)
			//这里就可以将错误发送给管理员...
			fmt.Println("发送邮件给@wew")
		}
	}()
	num1 :=10
	num2 :=0
	res :=num1/num2
	fmt.Println("res=",res)
}
//函数去读取以配置文件init.conf的信息
//如果文件名不正确，我们就返回一个自定义的错误

func readconf (name string) (err error){
	if name == "config.ini"{
		//读取
		return nil
	}else{
		//返回一个自定义的错误
		return errors.New("读取文件错误")
	}
}
func test02() {
   
	err :=readconf("config.ini")
	if err != nil{
		//如果读取文件发生错误就输出这个错误并终止程序
		panic(err)
	}
	fmt.Println("test02()继续执行")
}



func main(){
	//测试自定义错误
	test02()

// 	//测试
	
// for {
// 	test()
// 	time.Sleep(time.Second)
// 	}
 	fmt.Println("main()下面的代码...")
// }






}
	//num1 := 100
	//fmt.Printf("num1的类型是%T,num1的值是%v,num1的地址是%v\n",num1,num1,&num1)
    //输出结果如下
	//num1的类型是int,num1的值是100,num1的地址是0xc0420120a0
    
	
	//num2 := new(int)//*int
    //num2的类型%T ==> *int
	//num2的值 = 地址0xc0420120b8(这个地址是系统分配，不是固定的值)
	//num2的地址%v==地址0xc042004030（系统分配）
	// *num2 = 100
	// fmt.Printf("num2的类型是%T,num2存放的值是%v,num2指向的值是%v,num2的地址是%v",num2,num2,*num2,&num2)
    //输出结果如下所示：
	//num2的类型是*int,num2存放的值是0xc0420120b8,num2指向的值是0,num2的地址是0xc042004030

