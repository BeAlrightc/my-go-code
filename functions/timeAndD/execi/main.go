package main
import(
	"fmt"
	"time"
	"strconv"
)
//编写一个函数我们来算出他执行的时间
func test03(){
	str := ""
	for i :=0;i <100000;i++{
		str += "hello" +strconv.Itoa(i)
	}
}

func main(){
	//在执行test03前，先获取unix时间戳
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
    fmt.Printf("执行这个函数一共花了%v秒",end-start)
}