package main
import (
	"fmt"
	"encoding/json"
)

//定义一个结构体
type Monster struct {
	Name string `json:"monster_name"`//运用反射机制
	Age int `json:"monster_age"`
	Birthday string
	Sal float64
	Skill string
}

//将结构体序列化的演示
func testStruct() {
	//演示
	var monster = Monster{
		Name : "牛魔王",
		Age : 500,
		Birthday : "2011-11-11",
		Sal : 8000.0,
		Skill : "牛魔拳",
	}

	//将moster进行序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n",err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n",string(data))
}

//将Map序列化的演示
func testMap(){
	//定义一个Map
    var a map[string]interface{}
	//使用map，需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	//将a这个map进行序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n",err)
	}
	//输出序列化后的结果
	fmt.Printf("a map序列化后=%v\n",string(data))
}

//演示对切片进行序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	//使用map前，需要先make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 30
	m1["address"] = "北京"
	slice = append(slice,m1)

	var m2 map[string]interface{}
	//使用map前，需要先make
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 20
	m2["address"] = [2]string{"墨西哥","夏威夷"}
	slice = append(slice,m2)

	//将切片进行序列化操作
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n",err)
	}
	//输出序列化后的结果
	fmt.Printf("slice序列化后=%v\n",string(data))

}

//对基本数据类型进行序列化操作
func testFloat64() {
	var num1 float64 = 2345.67

	//对num1进行序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n",err)
	}
	//输出序列化后的结果
	fmt.Printf("num1序列化后=%v\n",string(data))
}
func main() {
	//演示将结构体，map,切片进行序列化
	testStruct()
//输出结果如下：monster序列化后={"Name":"牛魔王","Age":500,"Birthday":"2011-11-11","Sal":8000,"Skill":"牛魔拳"}	
	testMap()
//输出结果如下：a map序列化后={"address":"洪崖洞","age":30,"name":"红孩儿"}
    testSlice()
//输出结果如下：slice序列化后=[{"address":"北京","age":30,"name":"jack"},{"address":"墨西哥","age":20,"name":"tom"}]
    testFloat64() //num1序列化后=2345.67,将它变为字符串
	//将基本数据类型进行序列化意义不大
}