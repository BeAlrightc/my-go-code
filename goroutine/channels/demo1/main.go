package main
import (
	"fmt"
)
func main() {
	//演示一下管道的使用
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int,3)

	//2.看看intChan是什么
	fmt.Printf("intchan的值是=%v intChan本身的地址=%p\n",intChan,&intChan)
    
	//3.像管道写入数据
	intChan<-10
	num := 211
	intChan<- num

	//注意点，当我们在给管道写入数据时，不能超过其容量
	intChan<- 50
	//intChan<- 98 //会报错
	//4.输出看看管道的长度和cap（容量）
	fmt.Printf("channel len =%v cap=%v\n",len(intChan),cap(intChan)) // 3,3
    
	//5.从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Printf("num2=%v\n",num2) //10
	fmt.Printf("channel len =%v cap=%v\n",len(intChan),cap(intChan))//2,3
	 
	//6.在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
    num3 := <-intChan
    num4 := <-intChan
    // num5 := <-intChan
    // fmt.Println("num3=",num3,"num4=",num4,"num5=",num5)//报错


	}   
