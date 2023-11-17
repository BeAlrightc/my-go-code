package main
import (
	"fmt"
	"strconv"
   "strings"
)
func main(){
	//统计字符串的长度，按字节len(str)
    str := "hello"
    str2 := "hello我"//go的编码统一utf8(asciii的字符（字母和数字）占一个字节，一个汉字占三个字节)
	fmt.Println("str len =",len(str)) //str len =5
	fmt.Println("str2 len =",len(str2)) //str len =8（5个字母一个汉字）
   //题r :=[]rune(str)


   str3 := "hello北京"
   //字符串遍历，同时处理有中文的问题r :=[]rune(str)
   r := []rune(str3)
   for i :=0;i< len(r); i++{
	fmt.Printf("字符=%c\n",r[i])
   }

   //字符串转整数：n,err :=strconv.Atoi("12")
   n,err :=strconv.Atoi("hello")
   if err != nil{
	fmt.Println("转换错误",err)//转换错误 strconv.Atoi: parsing "hello": invalid syntax
   }else{
	fmt.Println("转成的结果是",n) //转成的结果是123
   }

   //整数转字符串：str = strconv.Itoa(12345)
   str = strconv.Itoa(12345)
   fmt.Printf("str=%v,str=%T\n",str,str)

   //字符串转[]byte: var bytes =[]byte("hello go")
   var bytes = []byte("hello go")
   fmt.Printf("bytes=%v\n",bytes)

   //[]byte转字符串：str=string([]byte{97,98,99})
   str = string([]byte{97,98,99})
   fmt.Printf("str=%v\n",str)

   //10进制转2,8,16进制：str=strconv.FormatInt(123,2) //2>8,16
   str=strconv.FormatInt(123,2)
   fmt.Printf("123对应的二进制是%v\n",str)//1111011
   
   str=strconv.FormatInt(123,16)
   fmt.Printf("123对应的十六进制是%v\n",str)//7b

   //查找子串是否在指定的字符串中：string.Contains("sefood","food") //trues
  b := strings.Contains("seafood","foo")
  fmt.Printf("b=%v\n",b)//true
  //统计一个字符串有几个指定的子串: 
 //strings.Count("ceheese","e")//4
 c :=strings.Count("ceheese","e")
 fmt.Printf("c=%v\n",c) //4

 //不想区分大小写的字符串比较(==是区分字母大小写的):
//fmt.Println(strings.EqualFold("abc","Abc")) //true
fmt.Println(strings.EqualFold("abc","Abc"))//true

fmt.Println("abc"=="Abc")//false

//返回子串在字符串第一次出现的index值
// ，如果没有返回-1:strings.Index("NLT_abc","abc") //4
e := strings.Index("NLT_abc","abc")
fmt.Printf("e=%v\n",e) //e=4

//返回子串在字符串最后一次出现的index，如没有就返回-1：strings.LastIndex("go "）
index :=strings.LastIndex("go golang","go")
fmt.Printf("index=%v\n",index)//3

//将指定的子串替换成另外一个
// 子串strings.Peplace("go go hello","go","go语言",n)
// n可以指定你希望替换几个，如果n=1表示全部替换
str4 :="go go hello"
str = strings.Replace(str4,"go","go语言",-1)
fmt.Printf("str=%v\n",str) //str=go语言 go语言 hello

//按照指定的某个字符，为分割标识，
// 将一个字符串拆分成字符串数组：strings.Split("hello,world,ok",",")
strArr :=strings.Split("hello,world,ok",",")
for i :=0;i<len(strArr);i++{
   fmt.Printf("str[%v]=%v\n",i,strArr[i])
}
fmt.Printf("strAr=%v\n",strArr)//strAr= [hello world ok]

//将字符串的字母进行大小写转换
// ：strings.ToLower("GO") //go strings.ToUpper("GO") //GO
str = "goLang Hello"
str= strings.ToLower(str)
fmt.Printf("str=%v\n",str)//str=golang hello
//将字符串的字母全部转换成大写
str= strings.ToUpper(str)
fmt.Printf("str=%v\n",str)//str=GOLANG HELLO

//将字符串左右两边的空格去掉：
// strings.TrimSpace(" tn a lone gropher ntrn  ")
str=strings.TrimSpace(" tn a lone gropher ntrn  ")
fmt.Printf("str=%v\n",str)//str=tn a lone gropher ntrn

//将字符串左右两边指定的字符去掉
// ：strings.Trim("!hello!","!")//["hello"]//将左右两边!和""去掉
str =strings.Trim("!hello!","!")
fmt.Printf("str=%v\n",str)//str=hello

//将字符串左边指定的字符去掉;
// strings.TrimLeft("! hello!","!")//["hello"]//将左边!和“”去掉
str=strings.TrimLeft("! hello!","!")
fmt.Printf("str=%v\n",str)//str= hello!

//判断字符串是否以指定的字符串开头
// ：strings.HasPrefix("ftp://192.168.10.1","ftp")//true
boo := strings.HasPrefix("ftp://192.168.10.1","ftp")
fmt.Printf("boo=%v\n",boo) //boo=true

//判断字符串是否以指定的字符结束
// ：strings.HasSuffix("NLT_abc.jpg","abc")//false
boo=strings.HasSuffix("NLT_abc.jpg","abc")
fmt.Printf("boo=%v\n",boo)//false
}