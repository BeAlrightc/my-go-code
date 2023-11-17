package main
import (
	"fmt"
)
func main() {
	intChan :=make(chan int,3)
	intChan<- 100
	intChan<- 200
	close(intChan) //close
	//这时不能够再写入到数channel
	//intChan<- 300 //panic: send on closed channel
	fmt.Println("okok~")
	//当管道关闭后，读取数据是可以的
	n1 := <-intChan
	fmt.Println("n1=",n1)

	//遍历管道
    intChan2 :=make(chan int,100)
	for i :=0; i < 100; i++ {
        intChan2 <- i *2 //放入100个数据进去管道之中
	}
	//遍历:这种遍历是错误的，因为遍历过程中管道的长度会变化
	// for i :=0; i < len(intChan2);++ {

	// }
	//在遍历时，如果channel没有关闭，则回出现deadlock的错误

	//在遍历时，如果cahnnel已经关闭，则会正常遍历数据，遍历完成后，就会退出遍历
	close(intChan2)
	for v := range intChan2 {
        fmt.Println("v=",v)
	}
}