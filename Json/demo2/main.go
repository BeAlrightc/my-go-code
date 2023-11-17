package main
import (
	"fmt"
	"encoding/json"
)

//定义一个结构体
type Monster struct {
	Name string 
	Age int
	Birthday string
	Sal float64
	Skill string
}
//演示将json字符串。反序列化成struct
func umarshalstruct() {
	//说明str 在项目开发中，是通过网络传输获取到的...或者通过读取文件得到
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"
    //定义一个Monster实例
	var monster Monster

	err := json.Unmarshal([]byte(str),&monster)
    if err != nil {
		fmt.Printf("unmarshal err=%v\n",err)
	}
	fmt.Printf("反序列化后 monster=%v\n",monster)
	//单独取出结构体中的一个字段
	fmt.Printf("反序列化后 monster.Name=%v\n",monster.Name)
}

//演示将jason字符串反射成map
func unmarshalMap() {
	str := "{\"address\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"
    
	//定义一个map
	var a map[string]interface{}
	//反序列化就不需要进行make了因为他会自动进行make操作
    
	//反序列化
	err := json.Unmarshal([]byte(str),&a)
    if err != nil {
		fmt.Printf("unmarshal err=%v\n",err)
	}
	fmt.Printf("反序列化后 a=%v\n",a)
	//单独取出结构体中的一个字段
	// fmt.Printf("反序列化后 monster.Name=%v\n",monster.Name)


}
//演示将json串反序列化文slice
func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":30,\"name\":\"jack\"}," +
	"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":20,\"name\":\"tom\"}]"
	//定义一个切片
	var slice []map[string]interface{}

	//反序列化
	err := json.Unmarshal([]byte(str),&slice)
    if err != nil {
		fmt.Printf("unmarshal err=%v\n",err)
	}
	fmt.Printf("反序列化后 slice=%v\n",slice)

}




func main() {
	umarshalstruct()
//输出的结果为：反序列化后 monster={牛魔王 500 2011-11-11 8000 牛魔拳}
    unmarshalMap()
//输出结果为：反序列化后 a=map[address:洪崖洞 age:30 name:红孩儿]
	unmarshalSlice()
//反序列化后 slice=[map[address:北京 age:30 name:jack] map[address:[墨西哥 夏威夷] age:20 name:tom]]	

}