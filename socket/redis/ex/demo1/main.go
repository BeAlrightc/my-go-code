package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)


func main() {
	var names string
	var ages int
	var skills string


	//1.连接到redis
	conn,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=",err)
		return
	}
	defer conn.Close() //关闭

	fmt.Println("请输入names: ")
	fmt.Scan(&names)
	fmt.Println("请输入ages: ")
	fmt.Scan(&ages)
	fmt.Println("请输入skills: ")
	fmt.Scan(&skills)
	//go操作redis进行写的操作
	_, err=conn.Do("HMSet","monster","name",names,"age",ages,"skill",skills)
	if err != nil {
		fmt.Println("HMSet err=",err)
		return
	}
    //go操作redis进行读的操作
	r, err :=redis.Strings(conn.Do("HMGet","monster","name","age","skill"))
	if err != nil {
		fmt.Println("HMGet err=",err)
		return
	}

	for i,v := range r{
		fmt.Printf("r[%d]=%v\n",i,v)
	}

 fmt.Println("操作完成")

}