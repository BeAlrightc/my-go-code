package main
import (
	"fmt"
	"time"
	"sync"
)
//需求:现在要计算1-200的各个数的阶乘，
// 并且把各个数的阶乘放入到map中，最后显示出来。要求使用goroutine完成

//思路
//1.编写一个函数，来计算各个数的阶乘，并放入到map中
//2.我们爱动的协程是多个，统计的结果放入到map中
//2.map应该做出一个全局的

var (
	myMap = make(map[int]int,10) 
	//声明一个全局的互斥锁
	//lock是一个全局的互斥锁
	//sync 是包：synchornized 同步
	//Mutex是互斥的意思
	lock sync.Mutex
)

//test函数就是计算n的阶乘，将这个结果放入到map中
func test(n int) {
	res := 1
	for i :=1; i <=n;i++ {
		res *= i
	}

	//这里我们将res放入到myMap中
	//加锁
	lock.Lock()
    myMap[n]= res//concurrent map writes
	//解锁
   lock.Unlock()
}

func main() {
	//我们这里开启多个协程完成这个任务[200个协程]
	for i :=1; i <=15; i++ {
		go test(i)
	}
	//休眠10秒
	time.Sleep(time.Second * 5)
	//输出结果，遍历结果
	lock.Lock()
	for i,v :=range myMap {
		fmt.Printf("map[%d]=%d\n",i,v)
	}
    lock.Unlock()
}