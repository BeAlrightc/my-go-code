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
	_, err=conn.Do("Set","name","tomjerry毛毛和老鼠")
	if err != nil {
		fmt.Println("set err=",err)
		return
	}

	//3.通过go向redis读入数据string [key-val]
	r, err :=redis.String(conn.Do("Get","name"))
	if err != nil {
		fmt.Println("Get err=",err)
		return
	}

	//返回的r是interfacce{}
	//因为name对应的值是字符串，因此我们需要转换[在一开始就用redis提供的方法进行转换]
	//nameString := r.(string) erro


	fmt.Println("操作成功!",r) //操作成功! tomjerry毛毛和老鼠
}