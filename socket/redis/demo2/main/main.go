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
	//写入名字
	_, err=conn.Do("HSet","user01","name","john")
	if err != nil {
		fmt.Println("hset err=",err)
		return
	}
	//写入age
	_, err=conn.Do("HSet","user01","age",10)
	if err != nil {
		fmt.Println("hset err=",err)
		return
	}

	//3.通过go向redis读入数据string [key-val]
	r1, err :=redis.String(conn.Do("HGet","user01","name"))
	if err != nil {
		fmt.Println("hGet err=",err)
		return
	}
    r2, err :=redis.Int(conn.Do("HGet","user01","age"))
	if err != nil {
		fmt.Println("hGet err=",err)
		return
	}
	



	fmt.Printf("操作成功! r1=%v,r2=%v\n",r1,r2) //操作成功! r1=john,r2=10
}