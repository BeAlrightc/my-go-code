package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go向redis写入数据和读取数据
	//1.连接到redis
	conn,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=",err)
		return
	}
	fmt.Println("conn succ...",conn)
}