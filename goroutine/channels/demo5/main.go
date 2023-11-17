package main
import (
	"fmt"
	"time"
)
//向intChan放入 1-8000个数
func putNum(intChan chan int){
	for i := 0 ;i<20000; i++{
		intChan<- i
	}
	//关闭intChan
	close(intChan)
}

//从intchan中取出数据，并判断是否为素数,如果是就放入到primeChan
func primeNum(intChan chan int,primeChan chan int,exitChan chan bool){

	//使用for循环
	
	var flag bool
	for {
		//time.Sleep(time.Millisecond)
		num,ok := <-intChan
		if !ok { //intChan取不到的时候，就退出这个主for循环
			break
		}
		flag = true //假设是素数
		//判断num是不是素数
		for i :=2;i<num;i++{
			if num %i ==0 { //说明i不是素数
				flag = false
				break
			}
		}
		if flag {
			//将这个数就放入到primeChan中
			primeChan<- num
		}

	}
	fmt.Println("有一个协程因为取不到数据没退出！")
	//这里我们还不能关闭primeChan
	//向exitChan 写入true
	exitChan<- true
}
func main() {
	intChan :=make(chan int,1000)
	primeChan :=make(chan int,20000) //放入结果
	//标识退出的管道
	exitChan :=make(chan bool ,4) //4个
	

	start :=  time.Now().Unix()

	//开启一个协程，向intChan放入 1-8000个数
	go putNum(intChan)
	//开启4个协程，从intchan中取出数据，并判断是否为素数,如果是就放入到primeChan
	for i :=0;i<4; i++{
		go primeNum(intChan,primeChan,exitChan)
	}
	//这里我们主线程，进行处理
	go func() {
		for i :=0;i<4; i++{
			<-exitChan
		}

		end := time.Now().Unix()
		fmt.Println("使用协程耗费的时间=", end - start)
		//当我们从exitChan祛除了3个结果，就可以放心关闭primeChan
		close(primeChan)
	}()

	//遍历primeChan
	for {
		_,ok := <-primeChan
		if !ok {
			break
		}
		//将结果输出
		//   fmt.Printf("素数为=%d\n",res)
	}
	fmt.Println("main主线程退出")


}