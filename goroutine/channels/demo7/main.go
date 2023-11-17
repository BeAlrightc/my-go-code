package  main
import (
	"fmt"
	"time"
)
func main() {

	//使用select可以解决从管道读取数据阻塞问题

	//1.先定义一个管道 10个数据 int
	intChan :=make(chan int, 10)
	for i := 0 ; i < 10 ;i ++{
		intChan<- i
	}
	//2.定义一个管道5个数据string
	stringChan :=make (chan string , 5)
	for i := 0; i < 5 ; i++ {
		stringChan <- "hello" +fmt.Sprintf("%d",i)
	}

	//传统方法遍历管道时，如果不关闭会阻塞而导致 deadlock

	//问题，在实际开发中，可能我们不好确定什么时候关闭该管道
	//可以使用select 方式解决
	label:
	for {
		select  {
		case v := <-intChan :  //注意：这里如果 intChan一直没有关闭，不会导致deadlocks,会自动到下一个case
			fmt.Printf("从intChan读取的数据%d\n",v)
			time.Sleep(time.Second)
		case v := <-stringChan :
			fmt.Printf("从stringChan读取的数据%s\n",v)	
			time.Sleep(time.Second)
		default :
			fmt.Println("都取不到，不玩了，你可以加入逻辑")	
			time.Sleep(time.Second)
			return
			break label
		}

	}

}