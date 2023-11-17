package main

import(
	"fmt"
	"strconv"
)

//golang基本数据类型转string

func main(){
	var num1 int = 99
	var num2 float64  = 23.456
	 var b bool = true
	 var mychar byte = 'h'
	var str string

	//是用第一种方式来转换 fmt.Sprint()
	str = fmt.Sprintf("%d",num1)  //num1转化为string
	fmt.Printf("str type %T str=%v\n",str,str) //结果为string 和99
    

	str = fmt.Sprintf("%f",num2)
	fmt.Printf("str type %T str=%v\n",str,str) //float64 23.456
    
	str = fmt.Sprintf("%t",b)
	fmt.Printf("str type %T str=%v\n",str,str)

	str = fmt.Sprintf("%c",mychar)
	fmt.Printf("str type %T str=%v",str,str)


	//第二种方式 是用 strconv函数的包
	var num3 int = 99
	 var num4 float64  = 23.456
	 var b2 bool = true
    //将int转换为字符串类型
	str = strconv.FormatInt(int64(num3),10)
	fmt.Printf("str type %T str=%q\n",str,str)

	//说明 f代表格式
	str = strconv.FormatFloat(num4,'f',10,64)
    fmt.Printf("str type %T str=%q\n",str,str)
	//将bool类型转换为字符串类型
	str =strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n",str,str)

    //strconv包中有一个函数Itoa
	var num5 int = 4567
	str = strconv.Itoa(num5)
	fmt.Printf("str type %T str=%q\n",str,str)


}