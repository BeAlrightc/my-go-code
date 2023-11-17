package main
import (
	"fmt"
	"time"
)

func main() {
	start :=  time.Now().Unix()

	for num := 1; num <=100000; num++{
	   flag := true //假设是素数
		//判断num是不是素数
		for i :=2;i<num;i++{
			if num %i ==0 { //说明i不是素数
				flag = false
				break
			}
		}
		if flag {
			//将这个数就放入到primeChan中
			// primeChan<- num
		}
	}
	end :=  time.Now().Unix()
	fmt.Println("耗费的时间为=",end -start)
}