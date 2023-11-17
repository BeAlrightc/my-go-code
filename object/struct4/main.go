package main
import (
	"fmt"
	"encoding/json"
)
type A struct{
    Num int
}
type B struct{
   Num int
}

type Student struct{
	Name string
	Age int
}

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

type Stu Student

func main(){
	
   var a A
   var b B
   fmt.Println(a,b)
   //a=b两个不同类型的结构体不能相互赋值操作
   //进行强制转换
   a = A(b) //成功，前提是两个结构体的字段名称
   //，字段类型是一模一样的，其中一个不一样都不可以
   fmt.Println(a,b)
   var stu1 Student
   var stu2 Stu
   //stu2 = stu1 //正确吗？会报错，可以这样修改 stu2=Stu(stu1)
   stu2=Stu(stu1)
   fmt.Println(stu1,stu2)

   //1.创建一个Monster变量
   monster :=Monster{"牛魔王",500,"芭蕉扇"}

   //2.将monster变量序列化为json格式的字符串
   jsonStr, err:= json.Marshal(monster)
   if err !=nil{
	fmt.Println("json处理错误",err)
   }
   fmt.Println("jsonStr",string(jsonStr))
   //输出如下：jsonStr {"Name":"牛魔王","Age":500,"Skill":"芭蕉扇"}
   //在字段后面加上json序列化后：
   //jsonStr {"name":"牛魔王","age":500,"skill":"芭蕉扇"}
}