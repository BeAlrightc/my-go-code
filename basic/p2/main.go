package main

import(
	"fmt"
	"strconv"
)

//golang基本数据类型转string

func main(){
	var str string = "true"
	var b bool
	//b, _ = strconv.ParseBool(str)
	//说明
	//1.strconv,PareseBool(str) 函数会返回两个值（value bool,err erro）
	//2.因为我只想获取到value bool ，不想获取 err所以我是用下划线
	b , _ = strconv.ParseBool(str)
    fmt.Printf("b type %T b=%v\n",b,b)	

	var str2 string ="123456"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2,10,64)
	n2 = int(n1)  //将int64转化成int类型的
	fmt.Printf("n1 type %T n1=%v\n",n1,n1)
	fmt.Printf("n2 type %T n2=%v\n",n2,n2)	

    var str3 string = "123.456"
	var f1 float64 
	f1, _ = strconv.ParseFloat(str3, 64)
    fmt.Printf("f1 type %T f1=%v\n",f1,f1)

    //注意
	var str4 string = "hello"
	var n3 int64
	n3, _ =strconv.ParseInt(str4, 10,64)
	fmt.Printf("n3 type %T n3=%v\n",n3,n3) //如果没有转成功就会变成默认值为0的数

}