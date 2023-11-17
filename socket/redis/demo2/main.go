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
	defer conn.Close() //关闭

	//2.通过go向redis写入数据string [key-val]
	//写入名字和年龄
	_, err=conn.Do("HMSet","user02","name","tom","age",19)
	if err != nil {
		fmt.Println("HMSet err=",err)
		return
	}

	//3.通过go向redis读入数据string [key-val]
	r, err :=redis.Strings(conn.Do("HMGet","user02","name","age"))
	if err != nil {
		fmt.Println("hGet err=",err)
		return
	}
	// fmt.Printf("r=%v\n",r) //r=[tom 19]
	for i,v := range r {
		fmt.Printf("r[%d]=%v\n",i,v)
	}

}