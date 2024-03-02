# go语言学习笔记

## 一、golang初体验:

1.简单体验案例：

```go
package main{ //把这个test.go归属到main
import "fmt"  //引入一个包 
func main(){
	//输出hello
	fmt.Println("hello world")

}
}
```

(1)go文件的后缀是.go

(2)package main

表示该hello.go文件所在的包是main,在go中每一个文件都要归属于一个包

(3)import “fmt” 表示引入一个包，包名为fmt引入该包后就可以这个包里面的函数

(4)func main(){

}func 表示一个函数main表示一个主函数，go函数执行的入口

(5)fmt.Println("hello world")

表示调用fmt里面的一个包的函数println进行输出hello world

通过go build 命令 对该go文件进行编译，生成 exe文件

最后执行该go文件的话就是进入cmd窗口 进行编译操作

![](D:\myfile\后端\go语言学习\pic\执行文件.jpg)

运行hello.exe文件即可

![](D:\myfile\后端\go语言学习\pic\执行exe.jpg)

注意：通过go run 命令可以直接运行hello,go程序（类似于执行一个脚本的文件的形式）在真实的生产环境会先编译再去运行go程序

linux下如何开发go环境,与windows开发基本一样，只是运行在可执行文件时使用./文件名方式运行

演示

![](D:\myfile\后端\go语言学习\pic\linux编译和运行.jpg)



## 二、golang执行流程分析

如果说是对源码编译后，再执行，Go的执行流程如下图所示：

![](D:\myfile\后端\go语言学习\pic\golng源码执行流程.jpg)

go run 文件名

![](D:\myfile\后端\go语言学习\pic\gorun方式执行.jpg)

通过运行进行对比前者执行速度更快，后者执行速度更慢，因为前面是直接运行二进制语言，后面还要先进行编译成二进制的语言再执行

实际的开发中还是先进行编译成二进制文件

**两种执行流程的区别**（面试重点）

1）如果我们先编译生成了可执行文件，那么我们可以将该可执行文件拷贝到没有go开发环境的机器上，仍然可以运行

2）如果我们是直接go run  go源代码，那么要在另一台机器上运行，也需要go开发环境，否则无法运行。

3)在编译时，编译器将程序运行依赖的库文件包含在可执行文件中，所以，可执行文件变大了很多

编译和运行的注意事项

![](D:\myfile\后端\go语言学习\pic\什么是编译.jpg)

![](D:\myfile\后端\go语言学习\pic\指定编译运行.jpg)

## 三、go的语法

#### 1.go语言开发注意事项

![](D:\myfile\后端\go语言学习\pic\go语法\pic1.jpg)

1）go源文件以"go"为扩展名

2）go应用程序的执行入口是main函数。

```go
package main //引入main包
func main(){ //调用main函数

}
```

3)go语言严格区分大小写

4）go方法、一由条条语句构成，每个语句后不需要分号（go语言会在每行后自动加上分号），这也体现出go语言的简洁性

```go
fmt.Println("hello world")  //后面不需要加上分号
```

5）go编译是一行一行进行编译的，因此我们一行就写一条语句，不能把多条语句写在同一行，否则会报错

```go
package main  //把这个test.go归属到main
import "fmt"  //引入一个包 
func main() {
	//输出hello
	fmt.Println("hello world")
	fmt.Println("hello world")
	fmt.Println("hello world")
	fmt.Println("hello world")
	fmt.Println("hello world")
	fmt.Println("hello world")
	fmt.Println("hello world")
	//下面是错的
	fmt.Println("hello world") fmt.Println("hello world") 

}

```

6）go语言定义的变量或者import的包如果没有使用到，代码不能编译通过

7）大括号都是成对出现的，缺一不可

2.golang常用的转义字符(escape char)

```
1) \t  一个制表位，实现对齐的功能
2) \n  换行符
3) \\ 一个\
3) \" 一个"
5) \r 一个回车 fmt.println("天龙八部雪山飞龙\r张飞")

```

```go
package main

import "fmt" //fmt包主要提供格式化和输入的函数
func main() {
	//演示转义字符的使用

	fmt.Println("tom\tjack")
	fmt.Println("D:\\myfile\\GO\\project\\src\\go_code\\escapechar")
	fmt.Println("tom说\"i love you\"")
	//\r表示回车，从当前行的最前面开始输出，覆盖掉以前的内容
	fmt.Println("天龙八部雪上飞狼\r张飞")
}

//输出的结果
D:\myfile\GO\project\src\go_code\escapechar>go run main.go
tom     jack
D:\myfile\GO\project\src\go_code\escapechar
tom说"i love you"
张飞八部雪上飞狼


练习：
fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
//结果如下：
姓名    年龄    籍贯    住址
john    12      河北    北京
```

3.golang开发常用的问题

![](D:\myfile\后端\go语言学习\pic\go语法\pic2.jpg)

小结与提示：

**学习编程最容易犯的错就语法错误，go要求你必须按照语法规则编写代码，如果你的程序违反了语法规则。例如：忘记了大括号、引导、或者拼错了单词go编译器都会报错，尝试去看懂编译器会报告的错误信息。**

4.go语言注释类型

1.go支持C语言风格的/* */块注释，也支持C++风格的//注释，行注释更通用，块注释主要针对包的详情说明或者屏蔽大块的代码

1）行注释

```go
package main

import "fmt" //fmt包主要提供格式化和输入的函数
func main() {
	
    
	//fmt.Println("hello world")
	//fmt.Println("i love you")
	//fmt.Println("i hate you")
	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
}
```

2）块注释(多行注释)

```go
package main

import "fmt" //fmt包主要提供格式化和输入的函数
func main() {
    //格式
    /*注释内容*/
	
	/*fmt.Println("hello world")
	fmt.Println("i love you")
	fmt.Println("i hate you")
	*/
	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
}
```

使用细节：

1）.对于行注释和块注释，被注释的文字，不会被go编辑器执行

2）块注释里面不允许有块注释嵌套

2.规范的代码风格

Go官方推荐使用行注释来注释整个方法语句。

要有正确的缩进和空白对程序编写代码

在dos下将go语言的代码显示出来

![](D:\myfile\后端\go语言学习\pic\go语法\pic3.jpg)

运算符的两边习惯性各加一个空格。比如: 2 + 4 * 5.

```go
var num = 2 + 4 * 5
```

go语言的代码风格

{  不可以单独占一行

```go
ackage main

import "fmt" //fmt包主要提供格式化和输入的函数
func main() {
	
	
	/*fmt.Println("hello world")
	fmt.Println("i love you")
	fmt.Println("i hate you")
	*/
	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
	var num = 2 + 4 * 5
	fmt.Println(num)
}

ackage main

以下这种方式是错误的
import "fmt" //fmt包主要提供格式化和输入的函数
func main() 
{//不可以单独占据一行
	/*fmt.Println("hello world")
	fmt.Println("i love you")
	fmt.Println("i hate you")
	*/
	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
	var num = 2 + 4 * 5
	fmt.Println(num)
}
```

一行不能超过80个字符

golang官网网站 http://golang.org#

熟练使用官网熟练其api的使用



中文网:http://studygolang.com



golang的包和源文件和函数的关系简图

![](D:\myfile\后端\go语言学习\pic\go语法\pic4.jpg)

要想学好go就多看官网并学习这个API

5.Dos的常用命令（针对windows系统）

```
目录的操作
md test //创建test目录
md test01 test02 //一下子创建两个目录
dir     //查看当前目录
cd /d f:  //切换到f盘 切换到某个盘符
cd  d:\test //切换到d盘的test100目录  ---绝对路径
cd test //切换到在当前目录的test目录   ---相对路径
cd ..  //切换到上一级目录
cd \  //切换到根目录
rd test01 //删除一个空目录
rd /q/s test01 //删除非空目录，将这个目录下的所有文件都删除 /q表示不用询问
rd /s test01 //以询问的方式删除这个目录

文件的操作
echo hello > abc.txt 将hello追加到一个新的abc.txt文件当中
copy abc.txt d:\test\test03 将abc.txt文件复制到test03目录下面
copy abc.txt d:\test\test03\sd.txt  //复制文件后再重新命名
move abc.txt d:\test\test02  //移动abc.txt文件
del abc.txt //删除一个文件
del *.txt //将所有txt类型的文件
其他指令
cls   //清屏
exit //退出终端

```

练习：

1）独立编写hello golang的程序

```go
package main

import "fmt" //fmt包主要提供格式化和输入的函数
func main() {
	fmt.Println("hello golang")
}
```

2）将个人基本信息(姓名，性别，籍贯，住址)打印倒终端上进行输出，每个信息各站一行

```go
package main

import "fmt" //fmt包主要提供格式化和输入的函数
func main() {
	fmt.Println("hello golang")
	//fmt.Println("")
    fmt.Println("刘承传\n男\n中国\n五指峰")
}
```

本章知识回顾

**go语言的SDK是什么？**

SDK就是啊软件开发工具包，我们做go开发，首先需要安装并配置好‘go

**Golang的环境变量配置及其作用**

GROOT:指定go sdk的安装目录。

Path:指令 sdk:bin目录 go.exe godoc.exe gofmt.exe

GOPATH 就是go项目的工作目录,所有项目的源码都放在这个目录下面

**golang程序的编写、编译、运行的步骤是什么**

编写：就是写代码

编译：go build 源码 =》生成一个二进制可执行文件

运行 : 1.对可执行文件运行 2.go run 源码

 四、变量的学习

变量是程序的基本组成单位

![](D:\myfile\后端\go语言学习\pic\go语法\pic5.jpg)



1.变量的使用步骤

1）声明变量（也叫定义变量）

2）给变量赋值

3）使用变量

```go
package main

import "fmt"
func main(){
	//定义变量/声明变量
	var i int
	//给i赋值
	i = 9
	//使用变量
	fmt.Println("i=",i)
}
```

2.变量的注意事项

1)变量表示内存中的一个存储区域

2)该区域有自己的名称(变量名)和类型(数据类型)

示意图：

![](D:\myfile\后端\go语言学习\pic\go语法\pic6.jpg)

3.变量使用的三种方式：

1）第一种：指定变量类型，声明后若不赋值，使用默认值

2）第二种，根据值自行判断变量的类型(类型推导)

3）第三种：省略var，注意:=左侧的变量不应该是已经声明的，否则会导致编译报错

```go
package main

import "fmt"
func main(){
	//1.golang的变量使用方式1
	//第一种，指定变量类型，声明后不赋值的话，使用默认值
	//int 的默认值为0
	var i int
	fmt.Println("i=",i)//结果为0
	//第二种，根据值自行判断变量的类型(类型推导)
	var num = 10.11
	fmt.Println("num=",num)
	//第三种：省略var，注意:=左侧的变量不应该是已经声明的，否则会导致编译报错
	//下面的方式等价于var name string name = "tom"
	name :="tom"//这个地方的:不能省略
	fmt.Println("name=",name)
	//4）多变量声明
    //在编程过程中，有时我们需要一次声明多个变量，golang也提供这样的语法
     

}
```

4）多变量声明

在编程过程中，有时我们需要一次声明多个变量，golang也提供这样的语法

```go
package main

import "fmt"
func main(){
	
	//该案例演示golang如何一次性声明多个变量
	// var n1, n2, n3 int
	// fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)
	//一次性声明多个变量方式二
	 n1,name, n3 := 100, "tom_",888
	fmt.Println("n1=",n1,"name=",name,"n3=",n3)

	//同样使用类型推导
}
```

如何一次性声明多个全局变量[在go中函数的外部定义变量就是全局变量]

5）该区域的数据值可以在同一类型范围内不断变化

```go
package main

import "fmt"

//变量使用的注意事项
func main(){
	//该区域的数据值可以在同一类型范围内不断变化
	var i  int =10
	i = 30
	i = 50
	fmt.Println("i=",i)
	//i = 1.23 //汇报错，因为不是int,不能改变数据类型
}
```

6）变量在同一个作用域(一个函数或者一个代码块中)内不能重名

```go
package main

import "fmt"

//变量使用的注意事项
func main(){
	//该区域的数据值可以在同一类型范围内不断变化
	var i  int =10
	i = 30
	i = 50
	fmt.Println("i=",i)
	//i = 1.23 //汇报错，因为不是int,不能改变数据类型

	//变量在同一个作用域(在一个函数或者在代码块中)内不能重名
	//var i int = 50
	//i := 99  //也会报错。这样写就包含了定义变量
}
```

7）变量=变量名 + 值 +数据类型，这一点大家要注意，变量的三要素

8）goglang的变量如果没有赋初值，编译器会使用默认值，比如int会默认值是0,string的默认值是空串

4.程序中 + 号的使用

1）当左右两边都是数值类型的时候，则做加法运算

2）当左右两边都是字符串，则做字符串拼接

```go
package main

import "fmt"

//展示Golang的使用
func main(){
	
	var i =1
	var j =2
	var r = i+j //做加法运算
	fmt.Println("r=",r) //result is 5

	var str1 = "hello"
	var str2 = "world"
	var res = str1 + str2
	fmt.Println("res=",res) //result is helloworld

}
```

5.变量的数据类型

![](D:\myfile\后端\go语言学习\pic\go语法\pic7.jpg)

1）int数据类型

![](D:\myfile\后端\go语言学习\pic\go语法\pic8.jpg)

```go
package main

import "fmt"

//展示Golang整数类型的使用
func main(){
	
	var i int =1
	fmt.Println("i=",i)
	//测试一下int的范围
	var j int8 = -129 //会报错 因为其范围是-128 ~ 128
	fmt.Println("j=",j)

}
```

无符号int整型介绍：

![](D:\myfile\后端\go语言学习\pic\go语法\pic9.jpg)

```go
	//测试一下
    var k uint8 = -1  //会报错
	fmt.Println("k=",k)
```

整型的类型

![](D:\myfile\后端\go语言学习\pic\go语法\pic10.jpg)

```go
//int,uint,rune, byte的使用
	var a int =8900
	fmt.Println("a=",a)
	var b uint = 1
	var c byte = 3 //范围是0~255
	fmt.Println("b=",b,"c=",c)
```

整型的使用细节

1）Golang各个整数类型分为：有符号和无符号 int uint的大小和系统有关系

2)Golang的整型默认为int型

3)如何在程序中查看某个变量的字节大小和数据类型

```go
package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
	"unsafe"
)

//展示Golang整数类型的使用
func main(){
	
	// var i int =1
	// fmt.Println("i=",i)
	//测试一下int的范围
	//var j int8 = -129 //会报错
	//fmt.Println("j=",j)
	//测试一下
    var k uint8 = 255
	fmt.Println("k=",k)
	//int,uint,rune, byte的使用
	var a int =8900
	fmt.Println("a=",a)
	var b uint = 1
	var c byte = 3 //范围是0~255
	fmt.Println("b=",b,"c=",c)

	//整型的使用细节
	var n1 = 100 //n1是什么类型
     //这里我们给介绍一下如何查看某个变量的数据类型
	 //fmt.Printf()  可以用于做格式化的输出

	fmt.Printf("n1 的类型 %T ",n1)

	 //如何查看某个变量的占用字节大小和数据类型 （使用较多）
    var n2 int64 = 10
	//unsafe.Sizeof(n2)) 是unsafe包的一个函数，可以返回n1变量占用的字节数
	fmt.Printf("n2 的类型 %T n2占用的字节数为 %d ",n2,unsafe.Sizeof(n2))


}
```

4)Golang程序中整型变量在使用时，遵守保小不保大的原则，即在保证程序正确运行下，尽量使用占用空间小的数据类型。【如年龄】

```go
//golang程序中整型变量在使用时，遵守保大不保小的原则
	//即，在程序正常运行下，尽量使用占用空间小的数据类型
    var age byte = 45
```

5)bit:计算机中最小的存储单位。byte：计算机中基本存储单元.



2）小数类型浮点型

基本介绍

小数类型就是用于存放小数的，比如 1.2 0.23 1.911

案例演示：

```go
func main(){

	var price float32 = 89.12
	fmt.Println("price",price)
	

}
```

浮点类型的使用

![](D:\myfile\后端\go语言学习\pic\go语法\pic11.jpg)

```go
func main(){

	var price float32 = 89.12
	fmt.Println("price",price)
	var num1 float32 = -0.00089
	var num2 float64 = -780956.09
	fmt.Println("num1",num1,"num2",num2)

}
```

```go
//位数部分可能丢失，造成精度损失。 -123.00000901

	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
    fmt.Println("num3",num3,"num4",num4)
```

说明float64 的精度比float32的更高

//保存精度更高的数应该用float64

![](D:\myfile\后端\go语言学习\pic\go语法\pic12.jpg)

```go
package main

import "fmt"

//演示golang中小数类型的使用
func main(){

	var price float32 = 89.12
	fmt.Println("price",price)
	var num1 float32 = -0.00089
	var num2 float64 = -780956.09
	fmt.Println("num1",num1,"num2\n",num2)
    
	//位数部分可能丢失，造成精度损失。 -123.00000901

	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
    fmt.Println("num3",num3,"num4",num4)

	//golang的浮点类型默认为声明为float64 类型
	var num5 = 1.1
	fmt.Printf("num5的数据类型是 %T\n",num5) //输出结果为float64

	//十进制整数形式.如：5.12    .512(必须有小数点)
	num6 := 5.12
	num7 := .123 //=> 0.123
	fmt.Println("num6",num6,"num7\n",num7)

	//科学计数法模式
	num8 := 5.1234e2 // ?5.1234 *10的2次方
	num9 := 5.1234E2 // ?5.1234 *10的2次方
	num10 := 5.1234E-2 // ?5.1234 *10的-2次方
	fmt.Println("num8=",num8,"num9=",num9,"num10",num10)
}
```

**字符类型**

Golang中没有专门的字符类型，如果要存储单个字符(字母)，一般使用byte来保存

**字符串就是一串固定长度的字符连接起来的字符序列**，Go的字符是由单个字节连接起来的。也就是说对于传统的字符串是由字符组成，而**go语言的字符串不同**，他是由**字节**组成的

[官方string、归属为https://tour,go-zh.org/basic/11]

```go
package main

import (
	"fmt"
)

//演示golang中字符类型使用
func main(){
	
	var c1 byte = 'a'
	var c2 byte = '0'
	//当我们输出byte值时，就直接输出了对应的字符ASCI码值
    fmt.Println("c1=",c1,"c2=",c2) //c1=92 c2=48
	//如果我们希望输出对应的字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n",c1,c2) //输出结果为 c1=a c2=0

	//var c3 byte = '北'  //overflow 溢出
	var c3 int = '北'  //汉字的ASCI码值比较大
	fmt.Printf("c3=%c c3对应的码值为=%d",c3,c3) //北 c3对应的码值为21271
}
```

对上面代码说明

1)如果我们保存的字符在ASCII表，比如0-1,a-z,A-Z.直接可以保存到byte

2)如果我们保存的字符对应的码值大于255,这是我们考虑int类型

3)如果我们需要按照字符的方式输出，这时我们只需要格式化输出printf("%c")

字符类型使用细节

![](D:\myfile\后端\go语言学习\pic\go语法\pic13.jpg)

```go
package main

import (
	"fmt"
)

//演示golang中字符类型使用
func main(){
	
	var c1 byte = 'a'
	var c2 byte = '0'
	//当我们输出byte值时，就直接输出了对应的字符ASCI码值
    fmt.Println("c1=",c1,"c2=",c2) //c1=92 c2=48
	//如果我们希望输出对应的字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n",c1,c2) //输出结果为 c1=a c2=0

	//var c3 byte = '北'  //overflow 溢出
	var c3 int = '北'  //汉字的ASCI码值比较大
	fmt.Printf("c3=%c c3对应的码值为=%d\n",c3,c3) //北 c3对应的码值为21271
    
	var c4 int = 22269 // 22269的ASCI码是国字
	fmt.Printf("c4=%c",c4) //输出的结果就是国字

   //字符类型是可以进行运算的，相当于于一个整数，运算时是按照码值进行运算的
   var num1 = 10 + 'a' //10 +97 =107
   fmt.Println("num1=",num1)
}
```

字符类型的本质探讨

1）字符类型存储到计算机中，需要将字符对应的码值（整数）找出来

存储：字符=》对应值--》二进制--》存储

读取：二进制---》码值----》字符---》读取

2）字符和码值的对应关系是通过字符编码表决定的（是规定好的）

3）Go语言的编码都统一成了utf-8，和其他的编程语言来说，非常的方便，很统一，再也没有编码的困扰了

布尔类型

基本介绍

1）布尔类型也是bool类型，bool类型数据只允许取值true和false

2）bool类型占一个字节

3）boolean类型适用于逻辑运算，一般用于程序流程控制：这个后面会详细介绍

```go
package main

import (
	"fmt"
	"unsafe"
)
func main(){
	var b = false
	fmt.Println("b=",b)
    //注意事项
   //1.bool类型占用存储空间是1个字节
   fmt.Println("b 的占用空间 =",unsafe.Sizeof(b) ) //b占用空间为1
   //2.bool类型只能取true或者false
   

}
```

if条件控制语句

for循环控制语句

**字符串（String）类型**

基本介绍

字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的Go语言的字符串的字节使用UTF-8编码标识Unicode文本

```go
package main

import (
	"fmt"
)

//演示golang中字符串类型使用
func main(){
	//string的基本使用
	var address string = "北京长城 110 hello world"
  	fmt.Println(address)

	
}
```

**String 使用的注意事项和细节**

1）Go语言的字符串的字节使用UTF-8编码标识Unicode文本，这样Golang统一使用UTF-8编码，乱码问题不会困扰程序员

2）字符串一旦赋值，字符串就不能更改了，在go中字符串是不可变的

```go
//字符串一旦赋值了，字符串就不能修改了，在Go中字符串是不可变的
	var str ="hello"
	//str[0] = 'a'  //这里不能去修改str的内容，即Go中的字符串是不可变的
    fmt.Println(str)
```

3）字符串的两种表示形式

（1）双引号，会识别转义字符

（2）反引号，以字符串的原生形式输出，包括换行和特殊字符，可以防止攻击、输出源代码等效果

4）字符串的拼接方式

```go
//字符串拼接方式
	var str5 = "hello" + "world"
	str5 += "haha"
    fmt.Println(str5)
```

5）当一行字符串太长时，需要使用到多行字符串，可以如下处理

```go
// var str6 = "hello" + "world" 
	// +"hello"+"hello"
	// +"world"
	// fmt.Println(str6) //会报错，+号要在上面

	var str6 = "hello" + "world" +
	"hello"+
	"hello"+
	"world"
	fmt.Println(str6) //正确输出
```



****

**基本数据类型默认值**

基本介绍：

在go中所有的数据类型都有一个默认值，当程序员没有赋值时就会保留默认值，在go中，默认值又叫0值，以下就是这个默认值：

基本数据类型一览表

![](D:\myfile\数据库\数据库资料\Mysql笔记（黑马）\结构图\新建文件夹\pic12.jpg)

```go
//查看这些基本数据类型的默认值
//%v表示的意思是按照变量的值输出
	var a int //0
	var b float32 //0
	var c float64 // 0 
	var isMarryied bool  //false
	var name string // " "
	fmt.Printf("a=%d,b=%f,c=%v,isMarryied=%v,name=%v",a,b,c,isMarryied,name)
```



**Golang基本数据类型的转换**

介绍

Golang和java/c不同，Go在不同类型的变量之间赋值时需要显示转换，也就是说Golang中数据类型不能自动转换

基本语法

表达式T(v)将值V转换为类型T

```go
T:就是数据类型，比如int32,int64,float32等等
V:就是需要转换的变量
package main

import (
	"fmt"
	// "unsafe"
)
//演示golang中基本数据类型的转换
func main(){
	
var i int = 100
//希望将i => float
var n1 float32 =float32(i)
var n2 int8 = int8(i)
var n3 int64 = int64(i) //低精度转换为高精度
fmt.Printf("i=%v n1=%v n2=%v n3=%v",i,n1,n2,n3) 

}


```

细节说明

1）G哦，数据类型的转换可以是从表示范围小-->表示范围大，也可以是 范围大--》范围小

2）被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化

```go
/被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化
fmt.Printf("i type is %T",i)  //输出结果为int32
}
```

3）在转换中，比如将int64转成int8[-128~127]，编译时不会报错，只是转换的结果是按溢出处理，和我们希望的结果不一样。因此在转换时需要考虑范围。

```go
var nu1 int64 =999999
var nu2 int8 = int8(nu1) //结果会按照溢出处理
fmt.Println("nu2=",nu2)  //输出结果是63
```

练习题：

```go
func main(){
	//练习题
var n1 int32 = 12
var n2 int64
var n3 int8

n2 = n1+ 20  //int32交给int64 错误
n3 =n1 + 20   //int32交给int8 错误

```

```
func main(){
	//练习题
var n1 int32 = 12
var n2 int64
var n3 int8
var n4 int8

n2 = int64(n1)+ 20  
n3 =int8(n1) + 128 // 编译不通过   
n4 = int8(n1) +127 //编译不通过 但是结果不是127+12
fmt.Println("n2=",n2,"n3",n3)


}
```



基本数据类型和string的转换

基本介绍

在程序开发中，我们经常将基本数据类型转换成string或者将string转换为基本数据类型

基本类型转为string类型

方式1:fmt,Sprintf("%参数"，表达式)

```go
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

}
```

函数的介绍

func Sprintf

```
func Sprintf(format string, a ...interface{}) string
Sprint根据format参数生成格式化的字符串并返回该字符串
```

参数需要和表达式的数据类型相匹配

fmt.Sprintf()会返回转换后的字符串

方式2：fmt.Sprintf()..会返回转换后的字符串

```go
func FormatBool(b bool) string
func FormatFloat(f float64,fmt byte,prec bitSize int) string
func FormatInt(int64,base int) string
func FormatUnit(i int64,base int) string

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

```

案例演示：

```go
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
```

String类型转基本数据类型

1)使用

![](D:\myfile\数据库\数据库资料\Mysql笔记（黑马）\结构图\新建文件夹\pic13.jpg)

```go
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

}
```

注意事项

在将string类型转成基本数据类型时，要确保string类型能够转换成有效的数据，比如我们可以把“123” 转成一个整数，但是不能把"hello" 转成一个整数，如果这样做，Golang直接将其转成0.

```go
//注意
	var str4 string = "hello"
	var n3 int64
	n3, _ =strconv.ParseInt(str4, 10,64)
	fmt.Printf("n3 type %T n3=%v\n",n3,n3) //输出是0
```

## 四、指针，以及标识符

### 1.指针

1.基本介绍

1）基本数据类型，变量存的值，也叫值类型

2）获取变量的地址用&，比如 var num int ,获取num的地址：&num

3)指针类型，变量存的是一个地址，这个地址指向的空间存的才是真正值，比如： var ptr *int =&num

4)获取指针类型所指的值，使用*,比如：var *ptr int,使用 *ptr获取p指向的值

5）举例说明

```go
package main

import(
	"fmt"
)
//演示golang中的指针类型
func main(){


	//基本数据在内存的布局
	var i int = 10 
	// i的地址是多少 &i
	fmt.Println("i的地址=",&i)  //结果为：i的地址= 0xc04205e058
	//1.a是一个指针变量
	//2. a的类型是 *int
	//3.a 本身就是 &i
	var a *int =&i
	fmt.Println("i的值是",i) //10
	fmt.Println("a存放的值是",a)//0xc04205e058
	fmt.Println("a指向的值是",*a)//10
    fmt.Println("指针本身的地址是",&a)//0xc04207c020
    或者可以使用printf()进行格式化输出

}
```

内存布局图：

![](D:\myfile\后端\go语言学习\pic\指针.jpg)

关于指针的内存布局操作

![](D:\myfile\后端\go语言学习\pic\指针内存布局.jpg)

练习，通过指针修改所存变量地址中存放的值

```go
package main

import(
	"fmt"
)

func main(){
var a int =999
fmt.Println("a的地址是",&a)
fmt.Println("指针修改之前a的值是",a)
var prt *int = &a
fmt.Println("prt的值是",*prt)
*prt = 23  //通过指针去修改a变量的值
fmt.Println("prt的值是",*prt)
fmt.Println("通过指针修改后这个a的值是",a)

}
```

判断对错

```go
func main(){
var a int = 300
var ptr *int =a //错误应该是 &a
}

func main(){
var a int = 300
var prt *float32 = &a //错误，类型不一致，不应该是float
}


func main(){
var a int =300
var b int =400

var ptr *int =&a
*ptr = 100 //a=100
ptr = &b  //存放b的地址
*ptr = 200  //b=200
fmt.Printf("a=%d, b=%d,*ptr=%d",a,b,*ptr)


}
```

指针的细节说明

1)值类型，都有对应的指针类型，形式为："数据类型"，比如int 的对应的指针就是 * *int,而float32对应的指针类型就是*  *float32 ,以此类推

2)值类型：基本数据类型int系列，float系列，bool string \数组和结构体struct

### 2.值类型和引用类型

1说明：

1)值类型:基本数据类型int系列，float系列，bootl,string、数组和结构体struct

2)引用类型：指针，slice切片、map、管道chan,interface等都是引用类型

2值类型和引用类型的区别

1）值类型：变量直接存储，内存通常在栈中分配

![](D:\myfile\后端\go语言学习\pic\值类型内存.jpg)

2）引用类型：变量存储的是一个地址，这个地址对应的空间才真正存储数据（值），内存通常在堆中分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾由GC来回收

![](D:\myfile\后端\go语言学习\pic\引用类型内存.jpg)

### 3.标识符的命名规范

#### 1.标识符概念

1）golang对各种变量、方法等命名时使用的字符序列称为标识符

2）凡是自己可以起名字的地方都叫标识符

#### 2.标识符的命名规则(6点)

1）由26个英文字母大小写，0~9，下划线—组成

2）数字不可以开头

3）Golang中严格区分大小写

4）标识符不能包含空格

5）下划线"—"本身在go中是一个特殊字符，称为空标识符，可以代表任何其他的标识符，但是它对应的值会忽略（比如，忽略某个返回值），所以仅能被作为占位符使用不能作为标识符使用

6）不能以系统保留关键字作为标识符，比如break、if等等...

```go
package main

import "fmt"

//演示golang中标识符的使用
func main() {
	//严格区分大小写
   var num int = 10
   var Num int = 20
   fmt.Printf("num=%v,Num=%v",num,Num) //10 20
   //标识符中不能含有空格
//    var ab c int = 30  //会报错

//_是空标识符，用于占用
// var _ int =40  // erro

} 
```

go中的保留字

![](D:\myfile\后端\go语言学习\pic\关键字.jpg)

标识符的案例

```go
hello  //OK
hello12 //OK
1hello //erro 不能以数字开头
h-b  //erro 不能使用-
x h   //erro 不能有空格
h_4   //ok
_ab   //ok
int    //ok 不推荐，甚至不要去用,语法通过不推荐使用
var int int =10
fmt.Println("int的值",int)//10success
float32  //ok
_  //erro
Abc  //ok


```

标识符命名注意事项

1）包名：**保持package的名字和目录保持一致**，尽量采取有意义的包名，简短，有意义不要和标准库冲突

2）变量名、函数名、常量名，采用驼峰法

```
var stuName string ="Tom"
var goodPrice int32 =124//第一个单词字母小写往后的单词采用大写的形式
```

3）如果使用变量名、函数名、常量名首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用（可以简单理解成，**首字母大写是公有的，首字母小写是私有的**）

![](D:\myfile\后端\go语言学习\pic\导包.jpg)

预定标识符

![](D:\myfile\后端\go语言学习\pic\预定标识符.jpg)

保留关键字的介绍：

![](D:\myfile\后端\go语言学习\pic\保留关键字.jpg)

运算符：

运算符是一种特殊的符号，用以表示数据的运算，赋值和比较等

1）算术运算符

主要是对数值型的变量进行运算，比如加减乘除，在Go程序中使用最多

![](D:\myfile\后端\go语言学习\pic\算术运算符.jpg)

```go
package main

import(
	"fmt"
)

func main() {
//重点讲解 /,%
//如果参与运算的数都是整数，那么除后，去掉小数部分，保留整数部分。不会四舍五入
fmt.Println(10/4) //是2不是2.5
var n1 float32 = 10 /4  //?结果也是2
fmt.Println(n1)

//如果我们需要保留小数部分，则需要有浮点数参与运算
var n2 float32 = 10.0 /4
fmt.Println(n2) //2.5
}
//演示%的用法
//演示 % 的使用
//看一个公式 a%b=a - (a/b) *b
fmt.Println("10%3=",10 % 3) //结果余数是1
fmt.Println("-10%3=",-10 % 3) //结果余数是-1 : -10 - [(-10)/3]*3 =-10 -(-9)=-1
fmt.Println("10%-3=",10 % -3) //结果余数是 1
fmt.Println("-10%-3=",-10 % -3) //结果余数是 -1
//演示 ++ --
var i int = 10
i++  //等价 i = i+1
fmt. Println("i的值是",i)  //11
i-- //等价i = i -1
fmt. Println("i的值是",i)  //10
```

细节说明：

1）对于除号 ”/“，他的整数和小数除是有区别的，整数之间做除法时，只保留整数部分而舍弃小数部分，例如 x := 19/5 结果是3

2）当对一个数取模式时，可以等价 a%b =a- a/b *b,这样我们就可以看到取模的本质运算

3）Golang的自增自减只能作为一个独立语言时，不能这样使用 b :=a++或者b :=a--

4）Golang的++和--只能写在变量的后面，不能写在变量的前面，既：a++ a--没有++a 或者--a

5）Golang的设计者去掉c/java中的自增自减的容易混淆的写法让Golang的语法更加简洁统一

![](D:\myfile\后端\go语言学习\pic\课堂练习.jpg)

练习：

（1）假如还有97天放假 问 xx个星期零xx天

（2）定义一个变量保存华氏温度，华氏温度转换为摄氏温度的公式是：5/9*(华氏温度-100)求出华氏温度对应的摄氏温度

```go
package main

import (
	"fmt"
)

func main(){
	//题1
 var day int  =97
 var week int = day/7
 var mo int = day%7
fmt.Printf("还有%d天放假,是%d个星期%d天",day,week,mo)
 
var huashi float32 = 134.2
var sheshi float32 = 5.0 / 9 * (huashi - 100)
fmt.Printf("%v对应的摄氏温度=%v",huashi,sheshi) 19 
}
```

2）赋值运算符

介绍

赋值运算符就是将某个运算后的值，赋给指定的变量

赋值运算符的分类：

![](D:\myfile\后端\go语言学习\pic\赋值运算符.jpg)



![](D:\myfile\后端\go语言学习\pic\赋值运算符位运算.jpg)

案例演示

（1）赋值基本案例

（2）有两个变量，a和b ，要求将其交换，最终打印结果

（3）+=的使用案例

```go
package main
	
	import (
		"fmt"
	)

	func main() {
		//赋值运算符的使用演示
		// var i int 
		// i = 10   //基本赋值

		//有两个变量，a和b要求将其进行交换，最终打印结果
		//a =9 b=2 ==>a =2 b =9
		a := 9
		b := 2
        //定义一个临时变量
		t :=a
		a =b
		b =t
    fmt.Printf("交换后的情况如下 a=%v,b=%v",a,b)
 
	//复和赋值的操作
	a +=7  //等价 a = a+7
	fmt.Println("a+=7的值为",a)


	}
```



赋值运算符的特点

（1）运算顺序从右往左

（2）赋值运算符的左边只能是变量 ，右边 可以是变量、表达式、常量值

（3）复合运算符等价于下面的效果

a+=3  a=a+3

3）比较运算符/关系运算符

(1)关系运算符的结果都是bool型，也就是要么为true,要么为false

(2)关系表达式经常用字在if结构中或者循环条件中

关系运算符一览表

![](D:\myfile\后端\go语言学习\pic\关系运算符一览表.jpg)

细节说明

- 关系运算符的结果都是bool型，也就是要么为true,要么为false
- 关系运算符组成的表达式，我们称为关系表达式 a>b
- 比较运算符”==“不能写成 "="

面试题

```go
package main

import(
	"fmt"
)
func main(){
	//有两个变量a和b，要求将其进行交换，但是不允许使用中间变量，最终打印效果
	var a int = 10
	var b int = 20
	a = a + b  //30 
	b = a - b  // 10
	a = a -b  //30-10 =20

	fmt.Printf("a=%v和b=%v",a,b) //20 和 10
}
```



4）逻辑运算符

介绍

用于连接多个条件（一般来讲就是关系表达式），最终的结果也是一个bool值

逻辑运算一览表：

![](D:\myfile\后端\go语言学习\pic\逻辑运算一览表.jpg)

案例演示

```go
package main
import (
	"fmt"
)

func main() {
  //演示如何使用逻辑运算符
  var age int =40
  if  age > 30 && age < 50{
	fmt.Println("ok1")
  }

  if  age > 30 && age < 40{
	fmt.Println("ok2")
  }

  //演示逻辑或的使用  ||
  if  age > 30 || age < 50{
	fmt.Println("ok3")
  }

  if  age > 30 || age < 40{
	fmt.Println("ok4")
  }
 //演示逻辑非的使用 !
 if  !(age > 30) {
	fmt.Println("ok5") //不输出
  }
}
```

注意事项和使用细节

- &&也叫短路与，如果第一个条件为false，则第二个条件不会判断，最终结果为false
- ||也叫短路或，如果第一个条件为true，则第二个条件不会判断，最终结果为true

案例演示：

```go
//声明一个测试函数(测试)
func test() bool {
	fmt.Println("test....")
return true
}

func main() {
	var i  int = 10
	//短路与的
	//说明：i<9 这个为false 第二条件根本不会去判断不会调用test函数
   if i < 9 && test(){
	fmt.Println("ok")
   }
   //短路或
   //说明：i>9 这个为true 第二条件根本不会去判断不会调用test函数
   if i > 9 || test(){
	fmt.Println("ok")
   }

    
```

运算符的优先级

（1）运算符有着不同的优先级，所谓优先级就是表达式运算中的运算顺序，如下表，上一行运算符总优先于下一行

![](D:\myfile\后端\go语言学习\pic\运算符优先级、.jpg)

（2）只有单目运算符、赋值运算符是从右往左运算的

（3）大致的顺序整理

1. 括号，++.--
2. 单目运算
3. 算术运算符
4. 移位运算
5. 关系运算
6. 位运算
7. 逻辑运算符
8. 赋值运算符
9. 逗号

6）其他运算符

位运算符

![](D:\myfile\后端\go语言学习\pic\位运算符.jpg)

![](D:\myfile\后端\go语言学习\pic\其他运算符.jpg)

课堂练习

（1）两个数的最大值

（2）求三个数的最大值

```go
func main() {
	//求两个数的最大值
	var n1 int =10
	var n2 int =40
	var max int
	if n1 > n2{
		max =n1
	}else{
		max =n2
	}
	fmt.Println("两个数的最大值为：",max)

	//求出三个数的最大值
	//先求出两个数的最大值然后让第三个数取比较

	var n3 int =45
    if n3 > max {
		max =n3
	}
	fmt.Println("三个数中的最大值是=",max)
}
```



键盘输入语句

介绍

在编程中，需要接收用户输入的数据，就可以使用键盘输入语句来获取’

步骤：

1）导入fmt包

2）调用fmt包的fmt.Scanln()或者fmt.Scanf()

案例演示：

要求：可以从控制台上接受用户信息。[姓名，年龄，薪水,是否通过考试]

1）使用fmt.Scanln()获取

```go
fmt.Println("请输入姓名")
   //当程序执行到fmt.Scanln(name)是，程序会停止运行，并等待用户的输入
   fmt.Scanln(&name)

   fmt.Println("请输入年龄")
   fmt.Scanln(&age)

   fmt.Println("请输入薪水")
   fmt.Scanln(&sal)

   fmt.Println("是否通过考试")
   fmt.Scanln(&isPass)

fmt.Printf("名字是%v \n 年龄是%v \n 薪水是 %v \n 是否通过考试 %v ",name,age,sal,isPass)
```

2）使用fmt.Scanf()获取

```go
/方式2.一次性输入这些信息 fmt.Scanf()可以按指定的格式输入信息
fmt.Println("姓名，年龄，薪水，是否通过考试,输入时采用空格隔开")
fmt.Scanf("%s %d %f %t ",&name,&age,&sal,&isPass)
fmt.Printf("名字是%v \n 年龄是%v \n 薪水是 %v \n 是否通过考试 %v ",name,age,sal,isPass)
```

### 4.go语言机制转换

1）进制介绍

- 二进制 : 0,1 满2进1在golang中不能直接使用二进制表示一个整数，因为他沿袭了C语言的风格要输出就%b

- 十进制:0~9满10进1

- 八进制:0~7满8进1，以数字0开头表示

- 十六进制：0~9A-F满16进1以0x或0x开头表示

  此处A~F不区分大小写

  如：0x21AF+1=02x21B0

2)案例使用

```go
func main(){
	var i int =5;

	fmt.Printf("%d的二进制是%b\n",i,i) //输出的结果是：5的二进制是101
    //八进制：0~7，满8进1，以数字0开头进行表示
    var j int = 011
	fmt.Println("j=",j) //9
   //0~9及A-F以0x或0X表示
    var k int = 0x11
	fmt.Println("k=",k) //17

}
```



进制转换的介绍

第一组（其他进制转换为十进制）

![](D:\myfile\后端\go语言学习\pic\其他进制转十进制.jpg)

1）二进制转十进制

规则：从最低位开始（右边的），将每个位上的数提取出来，乘以2的（位数-1）次方然后求和

```
案例：请将二进制：1011转成十进制的数
1011=1*1+1*2+0*2*2+1*2*2*2=1+2+0+8=11


```

2）八进制转十进制

规则：从最低位开始（右边的），将每个位上的数提取出来，乘以8的（位数-1）次方然后求和

```
请将0123转成十进制的数
3+2*8+1*8*8=83
```

3）十六进制转十进制

规则：从最低位开始（右边的），将每个位上的数提取出来，乘以8的（位数-1）次方然后求和

```
请将0x34A转成十进制的数
10+4*16+3*16*16=842

```

第二组（十进制转其他进制）

1）十进制转二进制

规则：将该数不断除于2，直到商为0为止，然后将得到的余数倒过来就是对应的二进制

案例；请将56转成2进制

![](D:\myfile\后端\go语言学习\pic\十进制转2进制.jpg)



2）十进制转八进制

规则：将该数不断除于8，直到商为0为止，然后将得到的余数倒过来就是对应的二进制

3）十进制转十六进制

规则：将该数不断除于16，直到商为0为止，然后将得到的余数倒过来就是对应的二进制

第三组

1）二进制转八进制

规则：将二进制数**每三位一组**（从低位开始组合），转成对应的八进制数即可

案例将二进制 11010101转成八进制

11010101=0325

2）二进制转十六进制

规则：将二进制**每四位数一组**（从低位开始组合），转换对应的十六进制数即可

案例：请将二进制：11010101转成16进制数为：0XD5

11100101 转化为八进制：0345

1110010110转化为十六进制：0x396

第四组

1）将八进制数每1位转成对应的一个3位的二进制数即可

案例：请将0237转成二进制

0237=10011111

2)将十六进制每1位转成1个4位的二进制数

0x3D=0x111101

原码 补码反码

![](D:\myfile\后端\go语言学习\pic\原码补码反码.jpg)

![](D:\myfile\后端\go语言学习\pic\原细节.jpg)

golang有三个位运算

![](D:\myfile\后端\go语言学习\pic\golang的位运算.jpg)

2的补码：00000010

3的补码    00000011

2&3=00000010 =2 

2|3=00000011=3

2^3=00000001

-2^2=-4

-2的原码为：10000010

-2的反码为：11111101

-2的补码：11111110

2的补码     00000010

故-2^2=11111100 的原码是先求出他的反码然后进行取反操作

反码=补码-1：11111011==》原码（符号位不变其他取反）：10000100=-4

-1：10000001=》反码=》11111110补码：11111111 +00000001=全是0

```
func main(){
	
	//位运算的知识演示
	fmt.Println(2&3) //2
	fmt.Println(2|3)//3
	fmt.Println(2^3)//1
	//下面是带有负数的二进制
	fmt.Println(-2^2)//-4

}
```

golang有两个移位运算符

```
>> 右移、 <<左移
右移运算符 >>:低位溢出，符号位不，并用符号位补溢出的高位
左移运算符<< 符号位不变，低位补0
```

案例演示：

```
都要将原来的数变成补码进行运算
a:=1>>2//00000001=>00000000=0 右移即补码左边补0
c:=1<<2 //00000001=>00000100=4左移即补码右边补0
fmt.Println(1>>2)//0
fmt.Println(1<<2)//4

```

## 五、流程控制语句

介绍：在程序中，程序运行的流程控制决定程序如何执行的，是我们应该掌握的，主要有三大流程控制语句

### 1.顺序控制

介绍：程序从上到下逐行地执行，中间没有任何判断效果

一个案例说明，必须下面的代码，没有判断，也没有跳转。因此程序按照默认的流程执行。即顺序控制

```go
var days int  = 97
var week int =days / 7
var day int = days % 7
fmt.Println("%d个星期零%d天",week,day)
```

流程图

![](D:\myfile\后端\go语言学习\pic\顺序控制.jpg)

顺序控制举例和注意事项

Golang中定义变量时采用合法额前向引用。如

```go
func main(){
var num1 int = 0
var num2 int = num1 + 20
fmt.Println(num2)
}
错误形式：不可以先使用后声明
func main(){
var num2 int = num1 + 20 //错误，一定要注意先后顺序
var num1 int = 10
fmt,Println(num2)
}
```

### 2.分支控制

分支控制if-else介绍

让程序有选择的执行，分支控制有三种。

#### 2.1单分支

基本语法

```
if 条件表达式{
   执行代码块
}
说明：当条件表达式为true时，就会执行{}的代码。
注意：{}是必须有的，就算你只写一行代码
```

应用案例

```go
package main

import (
	"fmt"
)
//编写一个程序可以输入一个人的年龄，
// 如果同志的年龄大于18岁则输出你的年龄大于18岁
//要对自己的行为负责

//分析：年龄 ==》var age int
//2.从控制台接收输入

func main(){
var age int
fmt.Println("请输入你的年龄:")
fmt.Scanln(&age)
if age >=18 {
	fmt.Println("你的年龄大于18岁，要为自己的行为所负责！")
      }
}
```

单分支细节说明

Go的if还有一个强大的地方就是条件判断语句里允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了

```golang
//golang支持在if中，直接定义一个变量，比如下面的
if age := 20; age > 18{
     fmt.Println("你的年龄大于18岁，要为自己的行为所负责！")
}
```



#### 2.2双分支

基本语法

```
if 条件表达式{
    执行代码块1
}else{
    执行代码块2
}
说明：当条件表达式成立，即执行代码块1，否则执行代码块2{}也是必须有的
```

```go
func main(){
var age int
fmt.Println("请输入你的年龄:")
fmt.Scanln(&age)
if age >=18 {
	fmt.Println("你的年龄大于18岁，要为自己的行为所负责！")
      }else{//强制编程语法。else一定在这个地方
		fmt.Println("你的年龄不大，这次放过你了")
	  }

}
```

双分支细节

双分支只会执行其中的一个分支。

```go
package main

import (
	"fmt"
)
func main(){
	var x int = 1
	var y int = 1
	if(x>2){//golang中可以写小括号，但是官方不推荐就是
		if(y>2){
			fmt.Println(x+y)
		}
		fmt.Println("hello")
	}else{
		fmt.Println("x is =",x) //输出x=1
	}
}
```

```go
var x int = 4
if > 2  //编译都不会通过，因为没有大括号
  fmt.Println("ok")
else
  fmt.Println("helo")
//没有大括号就是错的。编译错误没有什么输出

判断下列代码是否正确
var x int = 4
if x >2{
    fmt.Println("ok")
}//也是错的因为else要在这个括号的后面
else{
    fmt.Println("hello")
}
//编译错误，原因是else不能换行输出
```

#### 单分支和双分支的四个题目

1）编写程序，声明2 个 int32型变量并赋值。判断两个数之和。如果大于等于50，打印“hello world"

```go
var a int32 =45
	var b int32 =78
    if a+b>=50 {
		fmt.Println("这个数的和大于50")
	}
```

2)编写程序，声明2个float64型变量并赋值，判断第一个数大于10.0且第二个数小于20.0打印这两个数之和

```
 //第二题
	var c float64 =11.0
	var d float64 =17.0
	if c > 10.0 && d < 20.0 {
		fmt.Println("和=",(c+d)) //和=28
	}
```

3）定义两个变量int32 判断二者之和，是否能被3又被5整除，打印提示信息

```
//第三题
	var num1 int32 = 3
	var num2 int32 = 12
	if (num1 + num2) % 3 ==0 && (num1 + num2) % 5 == 0 {
		fmt.Println("能被三整除又可以被5整除")
	}
```

4）判断一个年份是否为闰年 ，闰年的条件时符合二者条件之一：（1）年份能被4整除但不能被100整除；（2）能被400整除

```go
//第四题求润年
	fmt.Println("请输入相应的年份：")
	fmt.Scanf("%d",&year)
	if (year% 4==0 && year % 100 !=0) || year % 400 ==0{
		fmt.Printf("%d是润年",year)
	}else{
		fmt.Printf("%d不是润年",year)
	}
```



2.3多分支

基本语法:

```
if 条件表达式1{
执行代码块1
}else if 条件表达式2 {
       执行代码块2
}
...
else{
  执行代码块n
}

```

对上面基本语法的说明

1）多分支的判断流程如下

（1）先判断条件表达式1是否成立，如果为真，就执行代码块1

  （2）如果条件表达式1为假就继续判断条件表达式2是否成立，如果条件表达式2为真就执行代码块2，

  （3）以此类推

  （4）如果所有条件表达式不成立，则执行else的语句块。

2）else不是必须的

3）多分支最终只能有一个执行入口

流程图（更加清晰的说清楚）：

![](D:\myfile\后端\go语言学习\pic\多分支流程图.jpg)

```go
func main(){
	//多分支类型的理解
	/*岳小鹏参加golang考试，他和父亲岳不群达成承诺
	如果
	成绩100分奖励一台bmw
	成绩为(80,90)奖励一套iphone14
	当成绩（60,80）奖励一台ipad
	其他时，什么奖励都没有
	请从键盘输入岳小鹏的期末成绩，并加以判断*/
	var score int
	fmt.Println("请输入岳小鹏的成绩：")
	fmt.Scanf("%d",&score)
	if score==100{
		fmt.Println("奖励一台BMW")
	}else if score>80&&score<=99{
        fmt.Println("奖励一台iphone14")
	}else if score>=60&&score<=80{
		fmt.Println("奖励一个ipad")
	}else{
		fmt.Println("对不起，他什么都没有好好学习")
	}
}

```

分支控制if-else案例介绍

![](D:\myfile\后端\go语言学习\pic\多分支演示.jpg)

```go
var b bool = true
	if b= false{ //在这里不能赋值，只能写成==
		fmt.Println("a")
	}else if b {
		fmt.Println("b")
	}else if !b {
		fmt.Println("c")
	}else {
		fmt.Println("d")
	}
```

![](D:\myfile\后端\go语言学习\pic\多分支案例演示3.jpg)

```go
package main

import (
	"fmt"
	"math"
)
func main(){
	//分析思路
	//1.a,b,c是三个float64
	//2.使用到给出的书序公式
	//3.使用到多分支
	//4.使用到math.squr方法
	//走代码
	var a float64 =3.0
	var b float64 =1.0
	var c float64 =10.0
	

	m := b*b -4 * a * c
	//多分支判断
	if m>0 {
		x1 :=-b + math.Sqrt(m) / 2 * a
		x2 :=-b - math.Sqrt(m) / 2 * a
		fmt.Printf("x1=%v,x2=%v",x1,x2)
	}else if m == 0 {
        x1 := (-b + math.Sqrt(m)) / 2 *a
		fmt.Printf("x1=%v",x1)
	}else{
		fmt.Println("无解...")
	}
```

![](D:\myfile\后端\go语言学习\pic\多分支案例4.jpg)

```go
//分析思路
	//设计三个变量，需要从控制台输入
	var height int32
	var money float32
	var handsome bool
	fmt.Println("请你输入身高(厘米)")
	fmt.Scanln(&height)
	fmt.Println("请你输入财富")
	fmt.Scanln(&money)
	fmt.Println("帅吗")
	fmt.Scanln(&handsome)
	if height>180 && money>10000000.0 && handsome==true{
		fmt.Println("我一定要嫁给他")
	}else if height>180 || money>10000000.0 || handsome==true{
        fmt.Println("嫁吧，比上不足，比下有余")
	}else{
		fmt.Println("不嫁")
	}
```

嵌套分支

基本介绍：

在一个分支结构中又完整嵌套了另一个完整的分支结构，里面的分支结构称为内层分支外面的称为外层分支

基本语法

```
if条件表达式{
  if条件表达式{
  
  }else{
  
  }
}
```

注意：嵌套分支不宜过多，建议控制在3层内

应用案例1：

参加白米运动会，如果用时8秒以内进入决赛，否则提示淘汰。并且根据特别提示进入男子组或女子组，输入成绩和性别，进行判断

float64 second | string gender

```go
var   second float64
	fmt.Println("请输入你的比赛成绩")
	fmt.Scanln(&second)
	//先判断是否进入决赛
	if second <=8{
		//进入决赛了
		var  gender string
		fmt.Println("请问您的性别是")
	    fmt.Scanln(&gender)
		if gender == "男"{
			fmt.Println("恭喜你进入男子组决赛")
		}else{
			fmt.Println("恭喜你进入女子组决赛")
		}

	}else{
		fmt.Println("对不起你被淘汰了")
	}
```

```go
//案例2：出票系统
func main(){
	/*
	出票系统：根据淡旺季的月份和年龄。打印票价 【考虑学生先做】
	4~10旺季“
	  成人（18-60）:60
	  儿童（<18）:半价
	  老人(>60)：1/3

	 淡季：
	    成人：40
		其他：20 

		//分析思路
		1.month age两个变量 byte
		2.使用嵌套分支
	*/
	var month byte
	var age byte
	var price float64 =60.0
	fmt.Println("请输入您旅游的月份")
    fmt.Scanln(&month)
	fmt.Println("请输入您旅游的年龄")
    fmt.Scanln(&age)

	if month>=4 && month<=10{
		if age > 60{
			fmt.Printf("票价是%v",price / 3)
		}else if age >=18{
            fmt.Printf("票价是%v",price)
		}else{
			fmt.Printf("票价： %v",price / 2)
		}
	}else{
		if age >=18 && age <60{
           fmt.Println("票价是",40)
		}else{
			fmt.Println("票价是",20)
		}
	}
```

#### switch分支结构

1）switch语句用于基于不同条件执行不同动作，。**每一个case分支都是唯一的，从上到下逐一测试**，直到匹配为止

2）**匹配项后面也不需要再加break**

基本语法

```go
switch 表达式{
case 表达式1,表达式2,...:
     语句块1
case 表达式3,表达式4,...:
     语句块2
...
default:
     语句块
}
```

switch流程图

![](D:\myfile\后端\go语言学习\pic\switch流程图.jpg)

对上图的说明和总结

1）switch的执行的流程是，先执行表达式,得到值，然后和case表达式进行比较，如果相等，就匹配到，然后执行对应的case语句块，然后退出switch的控制

2）如果switch的表达式的值没有和任何表达式的值匹配成功，执行后退出switch机制

3)golang的表达式可以有多个，并且使用逗号间隔

4）golang中的case语句块不需要写break，因为默认就会有break.当程序执行完语句块之后，就直接退出该switch控制结构

案例：

请编写一个程序，该程序可以接收一个字符，比如：a,b,c,d,e,f a表示星期一， b表示星期二...根据用户输入显示相应的信息，要求使用switch语句完成

```go
func main(){
	/*
	请编写一个程序，该程序可以接收一个字符，
	比如：a,b,c,d,e,f a表示星期一，
	 b表示星期二...根据用户输入显示相应的信息，
	 要求使用switch语句完成
	 分析：
	 1，定义一个变量接受字符
	 2.使用switch完成
	*/
	var key byte
	fmt.Println("请输入一个字符  a b c d e f g")
	fmt.Scanf("%c",&key)

	switch key {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
	case 'e':
		fmt.Println("周五")
	case 'f':
		fmt.Println("周六")
	case 'g':
		fmt.Println("周日")	
    default:
        fmt.Println("输入有误")						

	}

```

​	关于switch的细节讨论

1）case后是一个表达式（即：常量值、变量、一个有返回值的函数等都可以）

```
//写一个非常简单的函数
func test(b byte) byte {
	return b+1

}
func main(){	
	var key byte
	fmt.Println("请输入一个字符  a b c d e f g")
	fmt.Scanf("%c",&key)

	switch test(key) { //调用、test函数
```

2）case后的各个表达式的值的数据类型，必须和switch的表达式数据类型一致

```go
var n1 int32 = 20
	var n2 int32 = 20 //改为int64是错误的

	switch n1{//n2的类型需要与n1保持一致
	case n2 :
		fmt.Println("ok") //输出了ok
	default:
		fmt.Println("输入有误")			
	}
```

3）case后面可以带多个表达式，使用逗号间隔。比如case表达式1，表达式2

```go
var n1 int32 = 5
	var n2 int32 = 20 //改为int64是错误的

	switch n1{//n2的类型需要与n1保持一致
	case n2,10,5 : //case后面可以有多个表达值。匹配其中的一个就可以
		fmt.Println("ok") //输出了ok
	default:
		fmt.Println("输入有误")			
	}
```

4)case后面的表达式如果是常量值（字面量）则要求不能重复

```
var n1 int32 = 20
	var n2 int32 = 20 //改为int64是错误的
    var n3 int32 =5
	switch n1{//n2的类型需要与n1保持一致
	case n2,10,5 : //case后面可以有多个表达值。匹配其中的一个就可以
		fmt.Println("ok") //输出了ok
    case 5://错误上面已经有一个匹配值了，我们不可以重复case
         fmt.Print("5")
    case n3: //这种情况是可以正常进行编译操作的的可以说是骗过了编译器
         fmt.Println("n3")
	default:
		fmt.Println("输入有误")			
	}
```

5）case后面不需要带break程序匹配到一个case后就会执行对应的代码块，然后退出switch,如果一个斗匹配不到则执行default



6）default语句不是必须的

我们不写default也可以

7）switch后也可以不带表达式，类似于if-else分支来使用

```go
//switch后可以不带表达式，类似于if -else 分支来使用
	var age int = 10
	switch {
	case age ==10 :
		fmt.Println("age ==10")
    case age ==20 :
		fmt.Println("age ==20")
    case age ==30 :
		fmt.Println("age ==30")
	default :
	    fmt.Println("没有匹配到")				
	}
//case中也可以就age的范围进行判断
	var score int = 90
	switch {
	case score > 90 :
		fmt.Println("成绩优秀")
	case score >=70 && score <=90 :
		fmt.Println("成绩优良")
	case score >=60 && score <70 :
		fmt.Println("成绩及格")
	default :
	fmt.Println("不及格")		
	}
```

8）switch后也可以直接声明/定义一个变量，分号结束，不推荐

```go
//switch后也可以直接声明/定义一个变量,分号结束，并不推荐
	switch grade :=90; {
	case grade > 90 :
		fmt.Println("成绩优秀")
	case grade >=70 && grade <=90 :
		fmt.Println("成绩优良")
	case grade >=60 && grade <70 :
		fmt.Println("成绩及格")
	default :
	fmt.Println("不及格")		
	}
```

9）switch穿透-fallthough，如果在case语句块后增加fallthrough,则会继续执行下一个case也叫switch穿透

```go
	//switch的穿透 fallthrought
	var num int = 10
	switch num{
	case 10 :
		fmt.Println("ok1")
		fallthrough  //默认只能穿透一层
    case 30 :
		fmt.Println("ok2")
    case 20 :
		fmt.Println("ok3")		
	default :
		fmt.Println("没有匹配到")				
	}
```

10)Type Switch ：switch语句还可以被用于 typ-switch来判断某个interface变量中实际指向的变量类型 xtype()类似于多态操作

![](D:\myfile\后端\go语言学习\pic\type-sw.jpg)

```go
	var x interface{}
	var y = 10.0
	x = y 
	switch i:= x.(type){
	case nil:
		fmt.Printf("x的类型是：%T",i)
    case int:
         fmt.Println("x是int型")
	case float64:
		fmt.Println("x是float64型")//会输出	  
    case func(int) float64:
		fmt.Println("x是func(int)型")
	case bool,string:
		fmt.Println("x是bool型或者string型")		 		
	default:
		fmt.Println("未知")
	}
```

案例:

1).使用switch把小写类型的char型转为大写（键盘输入）。只转换a,b,c,d,e,其他的输出“Other”

```go
func main() {
	//1，使用switch将小写字符改为大写字母
	var char byte
	fmt.Println("请输入相应的字符")
	fmt.Scanf("%c",&char)
	switch char{
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
    case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default :
	fmt.Println("other")						
	}
```

2)对学生成绩大于60分的，输出”合格"。低于60的输出“不合格”。（注：输入的成绩不能大于100）

```go
 var score float64
    fmt.Println("请输入成绩")
	fmt.Scanln(&score)
	switch int(score / 60){
	    case 1:
			fmt.Println("及格")	
		case 0:
			fmt.Println("不及格")	
	    default:
			fmt.Println("输入有误...")
```



3）根据用户指定月份，打印该月份所属的季节。3,4,5.春季 6,7,8夏季 9,10,11秋季 12,1,冬季

```go
//3,根据用户指定月份，打印该月份所属的季节。
	// 3,4,5.春季 6,7,8夏季 9,10,11秋季 12,1,
	// 冬季
	var month byte
	fmt.Println("请输入月份")
	fmt.Scanln(&month)
	switch month {
	case 3,4,5:
		fmt.Println("春季")
	case 6,7,8:
		fmt.Println("夏季")	
	case 9,10,11:
		fmt.Println("秋季")	
	case 12,1,2:
		fmt.Println("冬季")	
	default:
		fmt.Println("您输错了月份")	
	}
```

switch和if的比较

1）如果判断的具体数值不多，而且符合整数、浮点数、字符、字符串这几种类型。建议使用switch语句简洁高效

2）其他情况：对区间判断和结果为bool类型的判断，用if if的使用范围更广



### 3.循环控制

#### for循环控制

基本语法

```
for循环变量初始化;循环条件;循环变量迭代{
      循环操作（语句）
}
```

对上面的语法格式说明

1)对for循环来说，有四个要素

2)循环变量初始化

3)循环条件

4)循环操作（语句），有人也叫循环体。

5)循环变量的迭代

6)for循环执行的顺序

for循环执行的顺序说明

1）执行循环变量的初始化，比如 i :=1

2)执行循环条件，比如 i<=n

3)如果循环条件为真就执行循环操作

4)执行循环变量迭代

5)反复执行 2 3 4 步骤，直到循环条件为false就退出for循环



案例

编写一个程序循环十次"你好刘承传"

```go
package main

import (
	"fmt"
)
func main(){
	for i :=1;i<= 10;i++{
		fmt.Println("你好刘承传",i)
	}
}
```



![](D:\myfile\后端\go语言学习\pic\for循环的执行内存图.jpg)

for循环的注意事项和细节说明

1）循环条件时返回一个布尔值的表达式

2）for循环的第二种使用方式

```
for 循环判断条件 {
   //循环执行语句
}
将变量初始化和变量迭代写到其他的位置
案例：
//for循环的第二种写法
	j := 1 //循环变量初始化
	for j <= 10 { //循环条件
		fmt.Println("你好刘承传",j) //循环体
		j++  //循环变量的迭代
	}
```

3）for循环的第三种使用方法

```go
for {
     //循环执行语句
}
上面的写法等价for ; ; {}是一个无限循环，通常需要配合break语句使用
//for循环的第三种写法,这种写法通常会配合break进行使用
	k :=1
	for { //此处等价 for ; ; ;{}
		if k <= 10{
		fmt.Print("helllo\n")
		}else{ //k>10时就break跳出
			break //break就是跳出整个for循环
		}
		k++
	}
```

for的注意事项和细节说明

4）Golang提供for-range的方式，可以方便遍历字符串和数组（注：数组的遍历）

```go
way1传统方式进行遍历	
str := "hello world!"
for i :=0;i<len(str);i++{
fmt.Printf("i=%d val=%c\n",i,str[i])
}

way2for-range方式进行遍历
func main(){
str :="hello,world!北京"
for index,val :=range str {
fmt.Printf("index=%d val=%c\n",index,val)
}

注意：for-range在遍历字符串时，是按照字符来遍历的，而不是按照字节来的，请注意这点
}
```

上面的细节讨论：

如果我们的字符串含有中文，那么传统的遍历字符串方式，就是错误。会出现乱码，原因是对字符串的遍历是按照字节进行遍历的，而一个汉字在utf8编码对应的是3个字节

如何解决，需要将str []str切片

```go
//传统方式对字符串进行切片
	str2 := []rune(str) //把str转成了[]rune
	for i :=0;i<len(str2);i++{
		fmt.Printf("i=%d val=%c\n",i,str2[i])
	}
```

对应for-range遍历字符串时，是按照字符串来遍历的，而不是按照字节来的。请注意这一点

for循环练习题

1）打印1~100之间所有是9的倍数的整数的个数及总和

```go
func main(){
	var ge int = 0
	var sum int = 0
	for i :=1;i<=100;i++{
		if i % 9 ==0{
			ge++
			sum = sum + i
		}
	}
	fmt.Printf("1~100之间是9的倍数额整数个数为%d,总和为%d",ge,sum)
}
```

2）完成下面的表达式输出

```go
//第二题
	var n int =6
	for i :=0;i<=n;i++{
		fmt.Printf("%v + %v= %v\n",i,n-i,n)
	}
```

#### while 和do...while的实现

Go语言没有while和do ..while语法，我们可以使用for循环来实现其使用效果

1）for循环实现while效果：先进行判断条件是否符合在进行循环输出

```go
循环变量初始化
for {
    if 循环条件表达式{
     break  //跳出for循环
    }
    循环操作（语句）
    循环变量迭代
}
说明：

```

1. for循环是一个无限循环
2. break语句就是跳出循环

使用上面的while循环输出十句“hello world”

```go
//使用while方式输出十句"hello world"
	//循环变量初始化
	var i int = 1
	for {
		if i>10 {
			break  //跳出循环结束过程
		}
		fmt.Println("hello world",i)
		i++  //循环变量的迭代
	}
```

使用do..while实现:先进行循环当条件不符合时就跳出循环

语法

```go
循环变量初始化
for {
   循环操作(语句)
   循环变量迭代
   if 循环条件表达式{
     break //跳出for循环
   }
}
```

对上图的说明

1. 上面的循环是先执行，再判断，因此至少执行一次
2. 当循环条件成立后，就会执行break,break就是跳出for循环，结束for循环

案例演示使用上面的dowhile实现输出十句helloworld

```
//使用do while的形式进行实现
var i int  = 1
	for {
		fmt.Println("hello world" ,i)
		i++ //循环变量进行迭代
		if i>10{
			break //跳出for循环
		}
	}
```

多重循环的控制

介绍：

1）将一个循环放在另一个循环体内，就形成了嵌套循环，在外边的for称为外层循环在里面的for循环称为内层循环（建议一般使用两层，最多不得超过三层）

2）实质上，嵌套循环就是把内层循环当作外层循环的循环体。当只有内层循环的循环条件为false时，才会完全跳出内层循环，才可结束外层循环的当次循环，开始下一次循环

3）设外层循环次数为m次，内层为n次，则内层循环体实际上需要执行m*n=nm次

应用实例

1》统计3个班级成绩情况，每个班有5名同学，求出各个班的平均分和所有班级的平均分（学生成绩从键盘输入）

```go
//案例1
	/*
	 1)统计3个班的成绩情况，每个班5名同学，
	 求出各个班的平均分和所有班级的平均分（学生的成绩从键盘输入
	 ）
	 分析实现思路
	 1.统计一个班成绩情况，每个班有5名同学，求出这个班的平均分
	 2.学生数就5个 
     3.声明一个变量sum统计班级的总分
	 4.定义一个变量统计总成绩得出所有学生的平均分

	 分析思路2
	 要统计三个班，走三次
	 分析思路3
	 1,将代码激活
	 2.定义两个变量表示班级的个数和班级的人数

	 代码实现
	*/
	var classNum int =2
	var stuNum int = 5
	var totalSum float64 = 0.0
	for j :=1;j <=classNum;j++{
	sum := 0.0
	for i := 1;i<=stuNum;i++{
		var score float64
		fmt.Printf("请输入第%d个班的第%d个学生成绩\n",j,i)
		fmt.Scanln(&score)
        //累计总分
		sum += score
	}
	fmt.Printf("第%d个班级的平均分是%v\n",j,sum / float64(stuNum))
	//将总分进行累加
	totalSum += sum
	fmt.Println("  ")
 }
 fmt.Printf("各个班级的总成绩是%v各个班级的平均分是%v",totalSum,totalSum/ float64(classNum * stuNum))
}
```

2》统计三个班及格人数，每个班有5名同学

```go
func main(){
	//案例1
	/*
	 1)统计3个班的成绩情况，每个班5名同学，
	 求出各个班的平均分和所有班级的平均分（学生的成绩从键盘输入
	 ）
	 分析实现思路
	 1.统计一个班成绩情况，每个班有5名同学，求出这个班的平均分
	 2.学生数就5个 
     3.声明一个变量sum统计班级的总分
	 4.定义一个变量统计总成绩得出所有学生的平均分

	 分析思路2
	 要统计三个班，走三次
	 分析思路3
	 1,将代码激活
	 2.定义两个变量表示班级的个数和班级的人数

	 //分析思路3统计班级及格人数
	 1,声明一个变量用于保存及格人数


	 代码实现
	*/
	var passc int = 0
	var classNum int =2
	var stuNum int = 5
	var totalSum float64 = 0.0
	for j :=1;j <=classNum;j++{
	sum := 0.0
	for i := 1;i<=stuNum;i++{
		var score float64
		fmt.Printf("请输入第%d个班的第%d个学生成绩\n",j,i)
		fmt.Scanln(&score)
        //累计总分
		sum += score
		//判断分数是否及格
		if score >=60{
          passc ++
		}
	}
	fmt.Printf("第%d个班级的平均分是%v\n",j,sum / float64(stuNum))
	//将总分进行累加
	totalSum += sum
	fmt.Println("  ")
 }
 fmt.Printf("各个班级的总成绩是%v各个班级的平均分是%v及格人数是%v",totalSum,totalSum/ float64(classNum * stuNum),passc)
}
```

打印一个矩形

```go
//矩形
for i := 1;i<=5;i++{
	for j :=1;j<=5;j++{
      fmt.Printf("*")
	}
	fmt.Println(" ")
 }
 //左侧直角三角形
 for i := 1;i<=5;i++{
	for j :=1;j<=i;j++{
      fmt.Printf("*")
	}
	fmt.Println(" ")
 }
//打印金字塔

案例2打印金字塔经典案例
案例分析
代码思路
1.打印一个金字塔
 */
 var toleve int =9
 for i := 1;i<=toleve;i++{
	//在打印星号前打印空格
	for k :=1;k<=toleve-i;k++{
		fmt.Print(" ")
	}
	for j :=1;j<=2*i-1;j++{
      fmt.Printf("*")
	}
	fmt.Println(" ")
 }
}
```

![](D:\myfile\后端\go语言学习\pic\打印金字塔.jpg)

打印空心金字塔

```go
//打印空心金字塔
 /*
   *
  * *
  ****
  分析：在我们给每行打印*号时需要考虑一个问题是打印*还是打印空格
  每层的第一个和最后一个就是打印星星，其他的都打印空格
  分析到例外情况，最底层是全部打印星星
 */
 var toleve int =9
 for i := 1;i<=toleve;i++{
	//在打印星号前打印空格
	for k :=1;k<=toleve-i;k++{
		fmt.Print(" ")
	}
	//在空格后面打印星星
	for j :=1;j<=2*i-1;j++{
      if j ==1 || j==2 * i -1 || i==toleve{
		fmt.Printf("*")
	  }else{
		fmt.Print(" ")
	  }
	}
	fmt.Println(" ")
 }
```

![](D:\myfile\后端\go语言学习\pic\空心金字塔.jpg)

打印九九乘法表

```go
/*
 打印九九乘法表

 
 */
 for i:=1;i<=9;i++{
	for j:=1;j<=i;j++{
		fmt.Printf("%v * %v= %v  ",i,j,i*j)
	}
	fmt.Println(" ")
 }
}
```

### 4.跳转控制语句

#### break

看需求：

随机生成1-100的一个数，直到生成了99这个数，看看你一共用了几次

分析：编写一个无限循环的控制，然后不停地随机生成数，当生成了99时，就退出这个无限循环==》break

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {

    //为了生成一个随机数，还需要个rand设置一个种子
	//time.Now().Unix():返回一个从1970:01:01的0分0秒到现在的秒数
	//rand.Seed(time.Now().Unix())
	//如何生成1-100的整数
	//n := rand.Intn(100) +1 //[0-100]故加一变成那个[1 100]
   // fmt.Println(n)

	/*
   思路：
   编写一个无限循环控制，然后不停的随机生成数，当生成了99时
   就退出这个无限循环==>break
	*/
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) +1 
		fmt.Println("n=",n)
		count ++
		if(n ==99){
			break //跳出for循环
		}
	}
	fmt.Printf("生成99一共使用了%v次",count)
     
}

```

基本介绍:

break语句用于终止某个语句块的执行，用于中断当前for循环或跳出switch语句

基本语法

```
( ......
break;
  ......
)
```

注意事项或细节说明

1）break语句出现在多层嵌套的语句块中时，可以通过标签指明要终止的是那一层语句块

2）标签的基本使用

```
label1:{...
label2:  {...
label3:    {...
           break label2;
           ....

          }
        }
      }
```

案例代码:

```
//这里演示一下指定标签的形式使用 break
	label2:
    for i := 0; i < 4; i++{
		//label1:  //设置标签
		for j := 0; j < 10; j++{
			if j ==2 {
				//break //break默认会跳出最近的for循环
			
			    //break label1
			    break label2
			}
			fmt.Println("j=",j) //j=0 j=1
		}
	}
```

对案例的说明：

- break默认跳出最近的for循环
- break后面可以指定标签，跳出标签对应的for循环

break语句练习题

1)100以内的数求和，求出当和第一次大于20的当前数。

```go
func main(){
	sum := 0
	for i :=1;i <=100; i++ {
		sum +=i //求和
		if sum > 20 {
			fmt.Println("当sum大于20时,这个数是：",i)
			break
		}
	}
```

2）实现登录验证，有三次机会，如果用户名为“张无忌”，密码"888"提示登录成功，否则提示还有几次机会

```go
//实现登录验证，有三次机会，如果用户名为“张无忌密码为888就提示登录成功
	//否则就提示还有几次机会
	var name string
	var passwd string
	var loginla int = 3
	for i:=1;i<=3;i++{
		fmt.Println("请输入姓名")
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		fmt.Scanln(&passwd)
		if name !="张无忌" || passwd != "888"{
			loginla --
			fmt.Printf("登录失败，你还有%v次机会\n",loginla)
			}else{
			fmt.Println("恭喜你登录成功")
			break
		}
	}
```

#### continue

基本介绍：

1）continue语句用于结束本次循环，继续执行下一次循环

2）continue语句出现在多层嵌套循环语句体中时，可以通过标签指明要跳过的是哪一层循环，这个和前面的break标签的使用的规则一样

基本语法

```go
{
 ...
 continue;
 ...
}
```

案例分析continue的使用

```go
func main() {
	/*
    continue案例
	这里演示一下指定标签的形式来使用
	*/
 //label2
 for i := 0; i<4; i++{
 //label1:设置一个标签
for j :=0; j<10 ; j++{
	if j ==2 {
		continue
	}
	fmt.Println("j=",j)
   }
  
 }
}
```

输出4次134..9，但是每次都没有2在其中。

continue输出1~100的奇数

```go
//1~100的奇数
for i:=1; i<=100; i++{
	if i%2!=0{
		fmt.Println(i)
	}else{
		continue
	}
}
```

从键盘读入个数不确定的整数，并判断读入额正数和负数额个数，输入为0时结束程序for循环break continue完成

```go
func main() {

	var posicount int
	var negacount int
	var num int
	for {
		fmt.Println("请输入一个整数")
		fmt.Scanln(&num)
		if num ==0 {
		break //终止for循环
		}
		if num >0 {
			posicount ++
			continue //结束本次循环，进入下一次循环
		}
			negacount ++
			
	}
	fmt.Printf("正数的个数是%v，负数的个数是%v",posicount,negacount)


```

#### goto

介绍：

1）Go语言的goto语句可以无条件地转移到程序中指定的行

2）goto语句通常与条件语句配合使用。可用来实现条件转移，跳出循环体等功能

3）在Go程序设计中一般不主张使用goto语句，以免造成程序流程的混乱，使理解与调试程序都产生困难

基本语法

```go
goto label
...
label:statement
```

```go
func main(){
	//演示goto的使用
	fmt.Println("ok1")
	goto lable1
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	lable1:
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
}
//运行结果为：
ok1
ok5
ok6
ok7
```

#### return

return使用在方法，表示跳出所在的函数或方法

```go
func main(){
for i :=1; i<=10; i++{
if i ==3{
return
}
fmt.Println(i)
}
fmt.Println("helloworld")
}
```

说明：

1）如果return 是在普通的函数，则表示跳出这个函数，即不再执行这个函数

2）如果return在main函数，表示终止main函数，也就是说终止程序

## 六、函数

#### 1.函数入门

为什么需要函数？

完成一个需求：输入两个数，再输入一个运算符(+-*/),得到结果

使用传统方法：

```go
func main(){
	/*
    请大家完成这样一个需求：
	输入两个数，在输入一个运算符(+-*   / ,得到结果)*/
	var n1 float64 =1.2
	var n2 float64 =2.3
	var operator byte ='/'
	var res float64
	switch operator {
	case '+':
          res = n1 + n2
	case '-':
		  res = n1 - n2
	case '*':
		  res = n1 * n2
	case '/':
		  res = n1 / n2	  
	default:
		 fmt.Println("操作符错误")  
	}
	fmt.Println("res=",res)
}
```

分析代码上的问题：

- 代码冗余太长
- 不好改动，不利于代码进行维护
- 函数可以解决这个问题

什么是函数：

为完成某一个功能的程序指令（语句）的集合，称为函数

在Go中，函数分为：自定义函数、系统函数（查看GO编程手册）

函数基本语法：

```
func 函数名 (形象列表) (返回值列表){
   执行语句...
   return 返回值列表
}
```

1）形参列表:表示函数的输入

2）函数中的语句：表示为了实现某一个功能的代码块

3）函数可以有返回值，也可以没有

案例入门：

```go
//将计算的功能放到一个函数中，在需要的时候进行调用
func cal (n1 float64,n2 float64,operator byte) float64{
	var res float64
	switch operator {
	case '+':
          res = n1 + n2
	case '-':
		  res = n1 - n2
	case '*':
		  res = n1 * n2
	case '/':
		  res = n1 / n2	 
    case '%':
		  res = n1 % n2    
	default:
		 fmt.Println("操作符错误")  
	}
	return res
}
//调用函数:
func main(){
result := cal(3.2,2.1,'-') //输入参数就可以对这个函数调用
fmt.Println("result=",result) //1.1
}
```

#### 2.包

为什么要用包

1）在实际的开发中，我们往往需要在不同的文件中，去调用其他文件的定义的函数，比如main.go中，去使用utils.go文件中的函数，如何实现？

2）现在有两个程序员共同开发一个go项目，两个程序员都想定义定义个交cal名字的函数怎么办

包的原理图

包的本质实际上就是创建不同的文件夹，来存储程序文件

此图用来说明一下包的原理图：

![](D:\myfile\后端\go语言学习\pic\包的原理图.jpg)

包的示意图

![](D:\myfile\后端\go语言学习\pic\包的示意图.jpg)

包的介绍

包的基本概念：

说明：go的每一个文件都是属于一个包的，也就是说go是以包的形式来管理文件和项目目录结构的

包的三大作用

- 区分相同名字的函数、变量等标识符
- 当程序文件很多时，可以很好的管理项目
- 控制函数、变量等访问访问范围，即作用域

打包的基本语法

```
pacakage 包名
```

引入包的基本语法

```
import "包的路径"
```

包使用的快速入门案例

![](D:\myfile\后端\go语言学习\pic\包utils.jpg)

```
utils包：
package utils

import (
	"fmt"
)

//将计算的功能放入一个函数中，然后在需要的时候进行使用
//为了让马哥其他的包文件使用Cal函数，需要将c大写，类似java中public
func Cal (n1 float64,n2 float64,operator byte) float64{
	
	var res float64
	switch operator {
	case '+':
          res = n1 + n2
	case '-':
		  res = n1 - n2
	case '*':
		  res = n1 * n2
	case '/':
		  res = n1 / n2	  
	default:
		 fmt.Println("操作符错误")  
	}
	return res
}

引入utils包
import (
	"fmt"
	"go_code/functions/utils"
)
调用
func main(){
result := utils.Cal(3.2,2.1,'-')
	fmt.Println("result=..",result)
}
//这样就完事了
```

包的注意事项以及细节

1）在给一个文件打包时，该包对应一个文件夹，比如在这里的utils文件夹对应的包名就是utils，文件的包名通常和文件夹名一致，一般为小写字母

2).当一个文件要使用其他的包函数或变量是需要先引入对应的包

引入方式1: import “包名”

引入方式2：

```
import(
    "name"
    "name"

)
```

3)pacakage指令在文件第一行，然后是、import指令

4)在import包是路径从$GOPATH的src下开始，不用带src,编译器会自动从src下开始引入

5）为了让其他包的文件，可以访问到本包的函数，则该函数名的首字母需要大写

6）在访问其他包的函数名时，其语法是包.函数，

```
utils.Cal(1,2)
```

7)如果包名较长，Go支持给包取别名，注意细节：取别名后，原来的包名就不能使用了

```go
import (
	"fmt"
	//前面的是别名
	ut "go_code/functions/utils"
)
//别名调用，原来的包名就会失效，就必须使用别名调用
result := ut.Cal(3.2,2.1,'-')
	fmt.Println("result=..",result)
```

8）在同一个包下，不能有相同的函数名（也不能有相同的全局变量名），否则报重复定义的错误

9）如果你要编译成一个可执行文件，就需要将这个包声明为main,即pacakage main这就是一个语法规范，如果你写一个库，包名可以自定义

![](D:\myfile\后端\go语言学习\pic\编译为可执行.jpg)

注意：-o表示输出 bin\my.exe表示在当前目录中的bin下面



#### 3.函数的调用机制

##### 通俗理解

![](D:\myfile\后端\go语言学习\pic\调用机制图.jpg)

##### 调用过程：

![](D:\myfile\后端\go语言学习\pic\diaoyong2.jpg)

对上图的说明：

（1）在调用一个函数时，会给该函数分配一个新的空间，编译器会通过自身的处理让这个新的空间和其他栈的空间区分开来

（2）在每个函数对应的栈中，数据空间都是独立的，不会混淆

（3）当一个函数调用完毕后，程序会销毁这个函数的栈空间

案例理解：

1）传入一个数进行加1的操作

```go
package main

import (
	"fmt"
)
//编写一个函数 test
func test(n1 int){
	n1 = n1 + 1
    fmt.Println("n1=",n1)
}
func main(){
	n1 :=10
  
  //调用test
  test(n1)

}
```

2）计算两个数，并返回 getSum

```go
func getSum(n1 int,n2 int) int {
//当函数有、return语句时，就是将结果返回给调用者
//即谁调用我，我就返回给谁
	return n1 + n2
}
//调用
func main(){
	n1 :=10
  
  //调用test
  test(n1)
  sum := getSum(1,3)
  fmt.Println("getSum得到的值是：",sum)//4

}
```

##### return语句

基本语法

Go函数支持返回多个值，这一点是其他编程语言没有的

```
func 函数名 (形参列表) (返回值类型列表){
    语句...
    return 返回值列表
}
```

注意：

（1）如果返回多个值时，在接收时，希望忽略某个返回值则使用_符号表示占位忽略

（2）如果返回值只有一个，（返回值类型列表） 可以不写()

案例演示

请编写函数，可以计算两个数的和和差，并返回结果

```go
func getSumAndSub(n1 int,n2 int) (int,int) {
	sum := n1 +n2
	sub := n1 -n2
	return sum,sub
}
调用
func main(){
  //调用getSumAndSub()函数
  res1,res2 := getSumAndSub(1,2)
  fmt.Printf("res1=%v,res2=%v",res1,res2)//res1=3 res2=1
}
```

一个细节说明：希望忽略某个返回值则使用_符号表示占位忽略

```go
//希望忽略某个返回值则使用_符号表示占位忽略
_, res3 := getSumAndSub(1,2)
fmt.Println("res3=",res3) //1-2=-1

```

##### 递归调用

基本介绍：

**一个函数在函数体内又调用了本身，我们称为递归调用**

案例入门：

```
func test(n int){
   if n > 2 {
       n--
       test(n)
   }
   fmt.Println("n=",n) //n=2 n=2 n=3
   //这里是if里面执行完了就会执行下面的输出
}
func test2 (n int) { //这里走了if就不会走else,
    if n > 2{
      n-- //递归必须向退出递归逼近
      test2(n)
    }else{
       fmt.Println("n=",n) //n=2 
    }
}
以上代码传入test(4)分别输出什么

```

![](D:\myfile\后端\go语言学习\pic\递归示意图.jpg)

函数递归需要遵守的重要原则

1）执行一个函数时，就创建一个新的受保护的独立空间（新函数栈）

2）函数额局部变量是独立的，不会相互影响

3）递归必须向退出递归的条件逼近，否则就是无限递归了，死龟了

4）当一个函数执行完毕，或遇到return就会返回遵守递归调用，就将结果返回给谁，当函数执行完毕或者返回时，该函数也会被销毁

117集

练习题：

题1：斐波那契数列

请使用递归的方式，求出斐波那契数1,1,2,3,5,8,13...给你一个整数n求出他的值是多少

```go
//斐波那契数
 func feibo(n int) int{
	if n==1 || n==2{
		return 1
	}else{
	return feibo(n-1)+feibo(n-2)
	}
 }
```

题2：求函数值

已知f(1)=3; f(n)= 2*f(n-1)+1

请使用递归的思想编程，求出f(n)的值

```go
//求函数的值 当f(1)=3,f(n)=2*f(n-1)+1
func han(n int) int {
	if n==1{
		return 3
	}else{
		return 2*han(n-1)+1
	}
}
```

题3猴子吃桃问题

有一堆桃子，猴子第一天吃了其中的一半，并且多吃了一个：以后每天猴子都吃了其中的一半，然后再多吃一个，想再吃时（还没吃），发现只有1个桃子。问题：最初共有多少个桃子

思路分析：

1. 第10天只有1个桃子
2. 第九天有几个桃子=(第10天桃子数 + 1)*2
3. 规律第n天桃子数peach(n)=(peach(n+1)+1)*2

```go
//猴子吃桃问题
func MonkeyPeach(n int) int {
	if n>10 || n<0{
		fmt.Println("输入的天数不对")
        return 0
	}
	if n==10 {
		return 1
	}else{
		return (MonkeyPeach(n+1)+1)*2
  }
}

//反过来
func maonkey2(n int) int {
	if n==1 {
		return 1
	}else{
		return (maonkey2(n-1)+1)*2
	}
}
```

#### 4.函数注意事项和细节讨论

1）函数的形参列表可以是多个。返回值列表也可以是多个

2）形参列表和返回值列表的数据类型可以是值类型和应用类型

3）函数的命名遵循标识符命名规范，首字母不能是数字，**首字母大写该函数可以被本包文件和其他包文件使用，类似于public，首字母小写，只能被本包文件使用，其他包文件不能使用，类似于private**

4）函数中的变量是局部的，函数外不生效

```go
package main
import (
	"fmt"
)
//函数中的变量是局部的，函数外不生效
func test(){
	//n1是test函数的局部变量，只能在test函数中使用
	var n1 int = 10

}
func main(){
	//这里不能使用n1，因为n1是test函数的局部变量
   fmt.Println("n1=",n1)
}
```

5)基本数据类型和数组默认都是值传递，即进行值拷贝。在函数内修改，不会影响到原来的值。

```go
func test2(n1 int) {
   n1 = n1 + 10
   fmt.Println("test2(n1)=",n1)
}
func main(){
	//这里不能使用n1，因为n1是test函数的局部变量
  // fmt.Println("n1=",n1)
  n1 := 20
  test2(n1)//30 不会影响到外面的值
  fmt.Println("n1=",n1) //20


}
```

6）如果希望函数内的变量能修改函数外的变量(指的是默认以值传递的方式的数据类型)，可以传入变量的地址&，函数内以指针的方式操作变量。**从效果上来看类似引用**

```go
//n1就是*int类型
func test3(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("test03(n1)=",*n1)//30
 }
func main(){
	
num := 20
test3(&num)//传的是地址，函数内部通过指针修改变量
fmt.Println("main() num=",num)//30

}
```

![](D:\myfile\后端\go语言学习\pic\指针传递修改函数外变量的值.jpg)

7）Go函数不支持重载

```
func test2(n1 int) {
	
   n1 = n1 + 10
   fmt.Println("test2(n1)=",n1)
}
func test2(n1 int，n2 int) {//错误不支持重载会报重复定义的错
	
}
```

8）在Go中，**函数也是一种数据类型，**可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数进行调用。

```go
package main
import (
	"fmt"
)
//在go中，函数也是一种数据类型
//可以赋给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数进行调用
func getSum(n1 int,n2 int) int{
	return n1+n2
}
func main(){
   a :=getSum
   fmt.Printf("a的类型%T,getSum类型是%T\n",a,getSum)

   res := a(10,40) //等价 res =:= getSum(10,40)
   fmt.Println("res=",res)
}
```

9）函数既然是一种数据类型，因此在go中，函数可以作为形参，并且调用

```go
package main
import (
	"fmt"
)
//在go中，函数也是一种数据类型
//可以赋给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数进行调用
func getSum(n1 int,n2 int) int{
	return n1+n2
}

//函数既然可以作为一种数据类型，因此在Go中，函数可以作为形参，并且调用
func myFunc(funvar func(int,int)int,num1 int,num2 int) int{
	return funvar(num1,num2)
}
func main(){
   a :=getSum
   fmt.Printf("a的类型%T,getSum类型是%T\n",a,getSum)

   res := a(10,40) //等价 res =:= getSum(10,40)
   fmt.Println("res=",res)

   //看案例
   res2 :=myFunc(getSum,50,60)
   fmt.Println("res2=",res2) //110

}
```

10）为了简化数据类型定义，Go支持定义数据类型

```go
基本语法：type 自定义数据类型名 数据类型 //理解：相当于一个别名
案例：type myInt int //这时myInt就等价int来使用了
-----
 type myInt int //给int取了别名 在go中myInt和int虽然都是
   //int类型，但是在go中还是认为myInt 和int是不同的数据类型

   var num1 myInt
   var num2 int
   num1 = 40
  // num2= num1 //报错不是同一个类型
  num2 = int(num1) //这样进行转换就可以
   fmt.Println("num1=",num1) //40
---------------------------------------   
  
案例：type mySum func(int,int)int //这时mySum就等价一个函数类型 func(int,int)int
 -----  
    //在举一个案例
 type myFunt func(int,int)int //这时myFun就是 func(int,int)int类型

 func myFunc2(funvar myFunt,num1 int,num2 int) int{
	return funvar(num1,num2)
}
//main函数进行调用
//案例结果
   res3 :=myFunc2(getSum,500,600)
   fmt.Println("res2=",res3) //1100
}
-------------------------------------
```

11)支持对函数返回值命名

```go
func cal(n1 int,n2 int) (sum int,sub int){
sum = n1 +n2
sub = n1 - n2
return  //这样就不用这样return sum sub之类额
}

//main中调用
 a,b := cal(1,2)
   fmt.Printf("a=%v,b=%v",a,b) //3,-1,a是sum b是sub
}
```

12）使用_标识符，忽略返回值

```go
func cal(n1 int,n2 int) (sum int,sub int){
sum = n1 +n2
sub = n1 - n2
return 
}
fun main(){
res,_:=cal(10,20)
fmt.Printf("res=%d",res1) //30
}
```

13)Go支持可变参数

```
//支持0到多个参数
func sum(arg.. int)sum int{
}
//支持1到多个参数
func sum(n1 int,args.. int)sum int{
}
```

说明

(1)args是slice,通过args[index]可以访问到各个值

(2)案例演示:编写一个函数sum,可以求出1到多个int的和

(3)如果一个函数的形参列表中有可变参数，则可变参数需要放在形参列表最后

```go
package main

import (
	"fmt"
)

//编写一个sum函数，可以求出 1到多个int的和
func sum(n1 int,args... int) int{
	sum := n1
	//遍历arges
	for i :=0;i <len(args);i++{
		sum += args[i] //args[0]表示取出args切片的第一个元素的值，其他依次类推
	}
	return sum
}
func main(){
//测试可变参数的使用
res :=sum(10,0,-1,90,10)
fmt.Println("res=",res)
}
```

函数练习题

```
func sum2(n1,n2 float32)float32{
	fmt.Printf("n1 type=%T\n",n1) //n1 type=float32
	return n1 +n2
}
func main(){
fmt.Println("sum2=",sum2(1,2)) //3


}
```

![](D:\myfile\后端\go语言学习\pic\函数练习.jpg)

最后一句会报错因为myfunc就接收两个int的参数但是b是三个int的参数，类型不匹配。



练习题三

请编写一个函数swap(n1 int,n2 int)可以交换n1和n2的值

```go
func swap(n1 *int,n2 *int){
	//定义一个临时变量
	t :=*n1
	*n1=*n2
	*n2=t
} 
func main(){
a :=10
b :=20
fmt.Printf("a=%v,b=%v",a,b)//10,20
swap(&a,&b)//传入的是地址
fmt.Printf("a=%v,b=%v",a,b) //20,10
}
```

#### 5.init函数

基本介绍

每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被Go运行框架调用，也就是说init会在main函数前被调用

案例说明

```go
package main
import (
	"fmt"
)

//init函数，通常可以在init函数中完成初始调用
func init(){//先执行
	fmt.Println("init()....")
}

func main(){
   fmt.Println("main()....")
}
//执行结果
D:\myfile\GO\project\src\go_code\functions\funint>go run main.go
init()....
main()....
```

init函数的注意事项和注意细节

1）如果一个文件同时包含变量定义，init函数和main函数，则执行的流程是先init函数后main函数

```go
package main
import (
	"fmt"
)

var age = test()

//为了看到全局变量是先被初始化的，我们这里先写函数
func test()int{  //1
	fmt.Println("test")
	return 90
}

//init函数，通常可以在init函数中完成初始调用
func init(){//先执行 2
	fmt.Println("init()....")
}

func main(){ //3
  fmt.Println("main()...age=",age)
}
```

2)init函数最主要的作用，就是完成一些初始化的工作

```go
utils.go
package utils
import (
	"fmt"
)
var Age int
var Name string
//Aage 和name全局变量我们需要在main.go中使用
//但是我们需要初始化Age和Name

//init函数完成初始化
func init(){
	fmt.Println("init函数执行")
  Age = 100
  Name= "初始化"
}
//调用
package main
import (
	"fmt"
	//引入包
	"go_code/functions/funint/utils"
)

var age = test()

//为了看到全局变量是先被初始化的，我们这里先写函数
func test()int{  //1
	fmt.Println("test")
	return 90
}

//init函数，通常可以在init函数中完成初始调用
func init(){//先执行 2
	fmt.Println("init()....")
}

func main(){ //3
  fmt.Println("main()...age=",age)
  fmt.Println("Age=",utils.Age,"Name",utils.Name)
}
```

3)面试题：案例如果main.go和utils.go都含有变量定义，init函数时，执行的流程又是怎么样

![](D:\myfile\后端\go语言学习\pic\执行顺序.jpg)

#### 6.匿名函数

介绍：

Go支持匿名函数，如果我们某个函数只是希望使用一次，可以考虑使用匿名函数，匿名函数也可以实现多次使用

匿名函数使用方式1

在定义匿名函数是就直接调用，这种方式匿名函数只能调用一次

```go
func main(){
	//在定义匿名函数是就直接调用，这种方式匿名函数只能调用一次
   
	//案例演示，求两个数额和，使用匿名函数实现
	res1 :=func (n1 int,n2 int)int {
		return n1+n2
	}(10,20)

	fmt.Println("res1=",res1)  //30
}
```

匿名函数使用方式2

将匿名函数赋给一个变量(函数变量)，再通过该变量来调用匿名函数

```go
//第二种方法，将匿名函数func (n1 int,n2 int)int献给a变量
	//则a的数据类型就是函数类型,此时我们可以通过a完成调用
	a := func (n1 int,n2 int)int{
		return n1 -n2
	}
	res2 := a(10,30)
	
	fmt.Println("res2=",res2) //-20

```



全局匿名函数

如果将匿名函数赋给一个全局变量，那么这个匿名函数，就成为一个全局匿名函数，可以在程序有效

```go
package main
import (
	"fmt"
)
var (
	Fun1 =func (n1 int,n2 int) int {
		return n1*n2
	}
)

func main(){
	//全局匿名函数的使用
	res4 := Fun1(4,9)
    fmt.Println("res4=",res4)//36
}
```

#### 7.闭包

介绍：

闭包就是一个函数与其相关的引用环境组合成一个整体（实体）

案例演示：

```go
package main
import (
	"fmt"
)
//累加器
func Addupper() func (int) int {
	var n int = 10
	return func (x int)int{
		n = n+ x
		return n
	}
}
func main(){
	//使用前面的代码
	f := Addupper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16
}
```

对上面代码的说明和总结

1)AddUpper是一个函数，返回额度数据类型是fun(int)int

2)闭包的说明：

```go
var n int = 10
	return func (x int)int{
		n = n+ x
		return n
	}
```

返回的是一个匿名函数，但是这个匿名函数引用到函数外的n，因此这个匿名函数就和n形成了一个整体，构成闭包

3)可以这样理解：闭包是类，函数是操作，n是字段

4）当我们反复调用f函数时，因为n是初始化1次，因此每调用一次就进行累加

5）我们要搞清楚闭包的关键，就是要分析出返回的函数他使用（引用）到哪些变量，因为函数和它引用的变量是一个整体

```go
//累加器
func Addupper() func (int) int {
	var n int = 10
	var str string ="hello"
	return func (x int)int{
		n = n+ x
		str += "x"
		fmt.Println("str=",str)
		return n
	}
}
func main(){
	//使用前面的代码
	f := Addupper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16
}
//运行结果;
str= hellox
11
str= helloxx
13
str= helloxxx
16
```

闭包的最佳实践

请编写一个程序，具体要求如下

1）编写一个函数makeSuffix(suffix string)可以接收一个文件后缀名（比如.jpg）.并返回一个闭包

2）调用闭包，可以传入一个文件名，如果改文件名没有指定的后缀(比如.jpg),则返回文件名.jpg.如果已有文件名就返回原文件名

3）要求使用闭包的方式完成

4）strings.HasSuffix:判断一个文件是否有后缀

```go
func makeSuffix(suffix string) func(string) string{
    
	return func (name string) string{
		//如果name没有指定的后缀，则加上。否则就返回原来的名字
	   if !strings.HasSuffix(name,suffix){
		return name +suffix
	   }
	   return name
	}
}
测试：
func main(){
//测试makeSuffix使用
//返回一个闭包
f :=makeSuffix(".jpg")
fmt.Println("文件名处理后=",f("winter")) //输出winter.jpg
fmt.Println("文件名处理后=",f("bird.jpg")) //输出bird.jpg
}
```

代码说明：

1）返回的函数和makeSuffix（suffix string）的suffiix变量和返回的函数组合成一个闭包，因为返回的函数引用到suffix这个变量

2）我们体会一下闭包的好处，如果使用传统的方法，也可以实现这个轻松的功能，但是传统方法需要每次都传入后缀名，比如.jpg,而闭包可以保留上次引用的某个值，所以我们传入一次就可以反复使用，仔细体会把

```go
//传统方法
func makeSuffix2(suffix string,name string) string{
    
		//如果name没有指定的后缀，则加上。否则就返回原来的名字
	   if !strings.HasSuffix(name,suffix){
		return name +suffix
	   }
	   return name
	}
	func main(){
//传统方法就要传入两个参数
fmt.Println("文件名处理后=",makeSuffix2(".jpg","winter")) //输出winter.jpg

}
```

#### 8.defer

为什么需要defer

在函数中，程序员经常需要创建资源(比如，数据库连接、文件句柄、锁等)，为了在函数执行完毕后，及时的释放资源，Go的设计者提供defer(延时机制)

快速入门案例

```go
package main
import (
	"fmt"
)
func sum(n1 int,n2 int)int{
    //当执行到defer时，暂时不会执行会将defer后面的语句压入到独立的栈中(defer栈)
    //当函数执行完毕后，再从defer栈，按陷入后出的方式中出栈，执行
	defer fmt.Println("ok1 n1=",n1)
	defer fmt.Println("ok2 n2=",n2)
	res := n1 + n2
	fmt.Println("ok3 res=",res)
	return res

}
func main(){
 sum(10,20)
 //输出：
//  ok3 res= 30
// ok2 n2= 20
// ok1 n1= 10
// }
```

defer的注意事项和最佳实践

1)当go执行到了一个defer时，不会立即执行defer后的语句，而是将defer后的语句压入到一个栈中，然后继续执行函数的下一个语句

2）当函数执行完毕后，从defer栈中，依次从栈顶取出语句执行（遵从栈 先入后出的机制）

3）在defer将语句放入栈中也会将相关的值拷贝同时入栈

```go
func sum(n1 int,n2 int)int{
	defer fmt.Println("ok1 n1=",n1)
	defer fmt.Println("ok2 n2=",n2)
	n1++ //n1=11
    n2++ //n2=21
    res := n1 + n2
	fmt.Println("ok3 res=",res)
	return res

}
func main(){
 sum(10,20)
 }
//输出结果仍然不变
ok3 res= 32
ok2 n2= 20
ok1 n1= 10
```

defer主要的价值是在：当函数执行完毕后，可以及时的释放函数创建的资源

```go
func test(){
//关闭文件资源
file = openfile(文件名)
defer file.close()
}

func test2(){
//释放数据库资源
connect = openDatabase()
defer connect.close()
//其他代码
}
```

1）在golang编程中的通常做法是，创建资源后，比如（打开了文件，获取了数据库的连接，或者是锁资源，可以理解为defer file.Close() defer connect.Close()）

2）在defer后，可以继续创建资源

3）当函数完毕后，系统依次从defer栈中，取出语句，关闭资源

4）这种机制非常简洁，程序员们不用再为什么时机关闭资源而烦心。

#### 9.函数参数的传递方式

基本介绍

我们在讲解函数注意事项和使用细节时，已经讲过值类型和引用类型了，这里我们在系统的总结一下，因为这是重难点，值类型参数默认就是值传递了，而引用类型参数默认就是引用传递

两种传递方式

1）值传递

2）引用传递

其实，不管是值传递还是引用传递，传递给函数的都是**变量的副本**，不同的是，**值传递**的是**拷贝**，**引用传递**的是**地址的拷贝**，一般来说，地址拷贝的效率高，因为数据量小，而值拷贝决定拷贝的数据大小，数据越大，效率越低

值类型和引用类型

1）值类型：基本数据类型int系列，float系列,bool，string、数组和结构体struct

2）引用类型：指针、slice切片、map、管道cha、interface等都是引用类型

值传递和引用传递的使用特点

1）值类型默认值传递：变量直接存储值，内存通常在栈中分配

2）引用类型默认是引用传递：变量存储的是一个地址，这个地址对应的空间才真正存储数据（值），内存通常在堆中分配，当任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，由GC回收

变量作用域

说明

1）函数内部声明/定义的变量叫局部变量，作用域仅限于函数内部

```go
package main
import (
	"fmt"
)
var name string="tom" //全局变量
func test01(){
	fmt.Println(name)
}

func test02(){//编译器采用就近原则
	name ="jack"
	fmt.Println(name) //jack
}

func main(){
fmt.Println(name) //tom
test01()//tom
test02()//jack
test01()//jack

```

2）函数外部声明/定义的变量叫全局变量，作用域在整个包都有效，如果首字母为大写，则作用域在整个程序有效

```go
package main
import (
	"fmt"
)

var age int= 50
var Name string= "jhon"
//函数
func test(){
	//age和Name作用域就只在test函数内部
	 age := 10
	 Name := "tom"
	fmt.Println("age=",age)
	fmt.Println("Name=",Name)
}
func main(){
	test() //10 tom
//   fmt.Println("age=",age) //会报错
  fmt.Println("age=",age)//50 
  fmt.Println("Name=",Name)//jhon
}
```

3）如果变量是在一个代码块，比如for/if中，那么这个变量的作用域就在该代码块

```go
for i :=0 ;i<=10;i++{
	fmt.Println("i=",i)
  }
//  fmt.Println("i=",i)会报错

```

下列代码是否报错

```go
var Age int = 20
Name :="tom" //相当于 var Name string Name="tom"在函数外不能够这样写
func main(){
fmt.Println("name",name1)
}
```

函数的课堂练习

1）函数可以没有返回值，编写一个函数，从终端输入一个整数，打印对应的金字塔

```go
//打印金字塔的案例
func jinzi(){
	var toleve int
	fmt.Println("请输入toleve值：")
	fmt.Scanln(&toleve)
	for i := 1;i<=toleve;i++{ //行
	   //在打印星号前打印空格
	   for k :=1;k<=toleve-i;k++{
		   fmt.Print(" ")
	   }
	   for j :=1;j<=2*i-1;j++{
		 fmt.Printf("*")
	   }
	   fmt.Println(" ")
	}
   }	
```

2）编写一个函数从终端输入一个整数（1-9），打印出对应的九九乘法表

```go
//打印九九乘法表的函数
func chengfa(){
	var n int 
	fmt.Println("请输入n的值：")
	fmt.Scanln(&n)
	for i :=1; i<=n;i++{
       for j :=1;j<=i;j++{
		fmt.Printf("%v*%v=%v ",i,j,i*j)
	   }
	   fmt.Println(" ")
	}
```

#### 10.字符串中常用的函数

说明：字符串在我们程序开发中，使用的是非常多的，常用的函数需要参考官方编程手册

1）统计字符串的长度，按字节len(str)

```go
func main(){
	//统计字符串的长度，按字节len(str)
    str := "hello"
    str2 := "hello我"//go的编码统一utf8(asciii的字符（字母和数字）占一个字节，一个汉字占三个字节)
	fmt.Println("str len =",len(str)) //str len =5
	fmt.Println("str2 len =",len(str2)) //str len =8（5个字母一个汉字）
}
```

2）字符串遍历，同时处理有中文的问题r :=[]rune(str)

```go
str3 := "hello北京"
   //字符串遍历，同时处理有中文的问题r :=[]rune(str)
   r := []rune(str3)
   for i :=0;i< len(r); i++{
	fmt.Printf("字符=%c\n",r[i])
   }
```

3）字符串转整数：n,err :=strconv.Atoi("12")

```go
//字符串转整数：n,err :=strconv.Atoi("12")
   n,err :=strconv.Atoi("123")
   if err != nil{
	fmt.Println("转换错误",err)
   }else{
	fmt.Println("转成的结果是",n) //转成的结果是123
   }
//如果传入hello
/字符串转整数：n,err :=strconv.Atoi("12")
   n,err :=strconv.Atoi("hello")
   if err != nil{
	fmt.Println("转换错误",err)//转换错误 strconv.Atoi: parsing "hello": invalid syntax
   }else{
	fmt.Println("转成的结果是",n) 
   }
```

4）整数转字符串：str := strconv.ltoa(12345)

```go
/整数转字符串：str = strconv.Itoa(12345)
   str = strconv.Itoa(12345)
   fmt.Printf("str=%v,str=%T",str,str)//str=12345,str=string
```

5）字符串转[]byte: var bytes =[]byte("hello go")

```go
//字符串转[]byte: var bytes =[]byte("hello go")
   var bytes = []byte("hello go")
   fmt.Printf("bytes=%v",bytes)
   //输出结果为：
  bytes=[104 101 108 108 111 32 103 111]
```

6）[]byte转字符串：str=string([]byte{97,98,99})

```go
//[]byte转字符串：str=string([]byte{97,98,99})
   str = string([]byte{97,98,99})
   fmt.Printf("str=%v",str)//str=abc

```

7）10进制转2,8,16进制：str=strconv.FormatInt(123,2) //2>8,16

```go
//10进制转2,8,16进制：str=strconv.FormatInt(123,2) //2>8,16
   str=strconv.FormatInt(123,2)
   fmt.Printf("123对应的二进制是%v\n",str)//1111011
   
   str=strconv.FormatInt(123,16)
   fmt.Printf("123对应的十六进制是%v\n",str)//7b
```

8）查找子串是否在指定的字符串中：string.Contains("sefood","food") //true

```go
//查找子串是否在指定的字符串中：string.Contains("sefood","food") //trues
  b := strings.Contains("seafood","foo")
  fmt.Printf("b=%v",b)//true
```

9)统计一个字符串有几个指定的子串: strings.Count("ceheese","e")//4

```go
//统计一个字符串有几个指定的子串: 
 //strings.Count("ceheese","e")//4
 c :=strings.Count("ceheese","e")
 fmt.Printf("c=%v\n",c) //4
```

10）不区分大小写的字符串比较(==是区分字母大小写的):fmt.Println(strings.EqualFold("abc","Abc")) //true

```go
//不想区分大小写的字符串比较(==是区分字母大小写的):
//fmt.Println(strings.EqualFold("abc","Abc")) //true
fmt.Println(strings.EqualFold("abc","Abc"))//true
fmt.Println("abc"=="Abc")//false ==是区分大小写的
```

11）返回子串在字符串第一次出现的index值，如果没有返回-1:strings.Index("NLT_abc","abc") //4

```go
//返回子串在字符串第一次出现的index值
// ，如果没有返回-1:strings.Index("NLT_abc","abc") //4
e := strings.Index("NLT_abc","abc")
fmt.Printf("e=%v",e) //e=4  如果返回的是-1那说明就没有包含
```

12）返回子串在字符串最后一次出现的index，如没有就返回-1：strings.LastIndex("go "）

```go
//返回子串在字符串最后一次出现的index，如没有就返回-1：strings.LastIndex("go "）
index :=strings.LastIndex("go golang","go")
fmt.Printf("index=%v\n",index)//3
```

13）将指定的子串替换成另外一个子串strings.Peplace("go go hello","go","go语言",n)n可以指定你希望替换几个，如果n=1表示全部替换

```go
str4 :="go go hello"
str = strings.Replace(str4,"go","go语言",1)
fmt.Printf("str=%v\n",str) //str=go语言 go hello
//将n变成-1时，表示全部替换
str = strings.Replace(str4,"go","go语言",-1)
fmt.Printf("str=%v\n",str) //str=go语言 go语言 hello

```

14）按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组：strings.Split("hello,world,ok",",")

```go
//按照指定的某个字符，为分割标识，
// 将一个字符串拆分成字符串数组：strArr :=strings.Split("hello,world,ok",",")
for i :=0;i<len(strArr);i++{
   fmt.Printf("str[%v]=%v\n",i,strArr[i])
}
fmt.Printf("strAr=%v\n",strArr)//strAr= [hello world ok]
}
//输出结果为：
str[0]=hello
str[1]=world
str[2]=ok
strAr=[hello world ok]
```

15）将字符串的字母进行大小写转换：strings.ToLower("GO") //go strings.ToUpper("GO") //GO

```go
str = "goLang Hello"
str= strings.ToLower(str)
fmt.Printf("str=%v\n",str)//str=golang hello
//将字符串的字母全部转换成大写
str= strings.ToUpper(str)
fmt.Printf("str=%v\n",str)//str=GOLANG HELLO
```

16）将字符串左右两边的空格去掉：strings.TrimSpace(" tn a lone gropher ntrn  ")

```go
//将字符串左右两边的空格去掉：
// strings.TrimSpace(" tn a lone gropher ntrn  ")
str=strings.TrimSpace(" tn a lone gropher ntrn  ")
fmt.Printf("str=%v\n",str)//str=tn a lone gropher ntrn
```

17）将字符串左右两边指定的字符去掉：strings.Trim("!hello!","!")//["hello"]//将左右两边!和""去掉

```go
//将字符串左右两边指定的字符去掉
// ：strings.Trim("!hello!","!")//["hello"]//将左右两边!和""去掉
str =strings.Trim("!hello!","!")
fmt.Printf("str=%v\n",str)//str=hello
```

18）将字符串左边指定的字符去掉;strings.TrimLeft("! hello!","!")//["hello"]//将左边!和“”去掉

```go
//将字符串左边指定的字符去掉;
// strings.TrimLeft("! hello!","!")//["hello"]//将左边!和“”去掉
str=strings.TrimLeft("! hello!","!")
fmt.Printf("str=%v\n",str)//str= hello!
```

19）将字符串右边指定的字符串去掉：strings.TrimRight("! hello!","!")//["hello"]//将右边的！和""去掉

```go
str=strings.TrimRight("! hello!","!")
fmt.Printf("str=%v\n",str)//str= !hello
```

20）判断字符串是否以指定的字符串开头：strings.HasPrefix("ftp://192.168.10.1","ftp")//true

```go
//判断字符串是否以指定的字符串开头
// ：strings.HasPrefix("ftp://192.168.10.1","ftp")//true
boo := strings.HasPrefix("ftp://192.168.10.1","ftp")
fmt.Printf("boo=%v\n",boo) //boo=true

```

21）判断字符串是否以指定的字符结束：strings.HasSuffix("NLT_abc.jpg","abc")//false

```go
//判断字符串是否以指定的字符结束
// ：strings.HasSuffix("NLT_abc.jpg","abc")//false
boo=strings.HasSuffix("NLT_abc.jpg","abc")
fmt.Printf("boo=%v\n",boo)//false
```

#### 11.时间和日期相关的函数

说明：在编程中，程序员经常使用到日期相关的函数，比如：统计某段代码执行花费的时间

1）时间和日期相关的函数，需要导入time包

![](D:\myfile\后端\go语言学习\pic\time包.jpg)

2）time.Time类型，用于表示时间

```go
//看看日期和时间相关的函数和方法使用
   //1.获取当前时间
   now := time.Now()
   fmt.Printf("now=%v now type=%T",now,now)
   //输出结果如下
//    now=2023-08-23 18:37:45.9279802 +0800 CST m=+0.008001001 now type=time.Time
```

3）获取当前时间的方法

```
now :=time.Now()//now的类型就是time.Time
```

4）如何获取其他的日期

```go
//2.通过now可以获取到年，月，日，时分秒
fmt.Printf("年=%v\n",now.Year())
fmt.Printf("月=%v\n",now.Month())
fmt.Printf("日=%v\n",now.Day())
fmt.Printf("时=%v\n",now.Hour())
fmt.Printf("分=%v\n",now.Minute())
fmt.Printf("秒=%v\n",now.Second())

fmt.Printf("月=%v\n",int(now.Month()))//转成我们喜欢的数字
```

5）格式化日期时间

 方式1：就是使用Printf 或者SPrintf

```go
//格式化日期和时间
//way1
fmt.Printf("当前年月日 %02d-%02d-%02d %02d:%02d:%02d \n",
now.Year(),now.Month(),now.Day(),
now.Hour(),now.Minute(),now.Second())//当前年月日 2023-08-23 19:17:28

dateStr := fmt.Sprintf("当前年月日 %d-%d-%d-%d-%d-%d \n",now.Year(),
now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())

fmt.Printf("dateStr=%v",dateStr)//dateStr=当前年月日 2023-8-23-19-21-36
```

方式2

```go
fmt.Printf(now.Format("2006/01/02 15:04:05"))
fmt.Println()
fmt.Printf(now.Format("2006-01-02"))
fmt.Println()
fmt.Printf(now.Format("15:04:05"))
fmt.Println()
//输出结果如下所示
2023/08/23 19:26:21
2023-08-23
19:26:21
```

说明："2006/01/02 15:04:05" 这个字符串的各个数字是固定的，必须这样写。

“2006/01/02 15:04:05”这个字符串各个数字可以自由组合，这样可以按照程序需求来返回时间和日期

6）时间的常量

```go
const{
  Nanosecond Duration =1 //纳秒
  Microsecond =1000 * Nanosecond //微秒
  Millisecond =1000 * Microsecond //毫秒
  Second         =1000 * Millisecond //秒
  Minute         = 60 *Second//分钟
  Hour           = 60 *Minute//小时
}
```

**常量的作用**：在程序中可用于获取指定时间单位的时间，比如想得到的100毫秒 100 * time.Millisecond

7）休眠

func Sleep(d Duration)

案例：

```go
time.Sleep(100 * time.Millisecond)//休眠100毫秒
//每隔0.1秒就打出一个数字，打印到100时就退出
i := 0
for {
	i++
	fmt.Println(i)
	//休眠
	time.Sleep(time.Millisecond * 100)
	if i== 100{
		break
	}
}
```

8）获取当前Uix时间戳和unixnano时间戳（作用是可以获取随机数字）

unix时间戳

unixnano时间戳

![](D:\myfile\后端\go语言学习\pic\unix时间戳.jpg)

```go
//unix和unixnano的使用
fmt.Printf("unix时间戳=%v unixnano时间戳=%v",now.Unix(),now.UnixNano())
//输出有以下结果：
//unix时间戳=1692792690 unixnano时间戳=1692792690199200800

```

实践案例：

1）编写一段代码来统计函数test3()执行的时间

```go
package main
import(
	"fmt"
	"time"
	"strconv"
)
//编写一个函数我们来算出他执行的时间
func test03(){
	str := ""
	for i :=0;i <100000;i++{
		str += "hello" +strconv.Itoa(i)
	}
}

func main(){
	//在执行test03前，先获取unix时间戳
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
    fmt.Printf("执行这个函数一共花了%v秒",end-start)
}
```

#### 12.内置函数

说明：

Golang设计者为了编程方便，提供了一些函数，这些函数可以直接使用，我们称为go的内置函数文档:https://studygolang.com/pkgdoc ->builtin

1）len:用来求长度，比如string、array、map、channel

2）new:用来分配内存，主要用来配置类型，比如int,float32,struct...返回的是指针

```go
func main(){
	num1 := 100
	fmt.Printf("num1的类型是%T,num1的值是%v,num1的地址是%v\n",num1,num1,&num1)
    //输出结果如下
	//num1的类型是int,num1的值是100,num1的地址是0xc0420120a0
    
	
	num2 := new(int)//*int
    //num2的类型%T ==> *int
	//num2的值 = 地址0xc0420120b8(这个地址是系统分配，不是固定的值)
	//num2的地址%v==地址0xc042004030（系统分配）
	*num2 = 100
	fmt.Printf("num2的类型是%T,num2存放的值是%v,num2指向的值是%v,num2的地址是%v",num2,num2,*num2,&num2)
    //输出结果如下所示：
	//num2的类型是*int,num2存放的值是0xc0420120b8,num2指向的值是0,num2的地址是0xc042004030

}
```

![](D:\myfile\后端\go语言学习\pic\new的内存图.jpg)

3）make:用来分配内存，主要用来分配引用类型，比如chan、map、slice这个我们后面讲解

#### 13.go的错误处理机制

先看示例代码，看看输出了什么

```go
func test(){
	num1 :=10
	num2 :=0
	res :=num1/num2
	fmt.Println("res=",res)
}
func main(){
	//测试
	test()

	fmt.Println("main()下面的代码...")
}

```

对上面的代码进行总结

1)在默认情况下，当发生错误后(panic)，程序就会退出(崩溃)

2)如果我们希望。当发生错误后，可以捕获到错误，并进行处理，保证代码可以继续执行。还可以在捕获到错误后给管理员一个提示（邮件或短信）

3)这里引出我们对错误进行处理

**错误处理**

1）Go语言追求简洁优雅，所以Go语言不支持传统的try...catch...finally这种处理。

2）Go中引入的处理方式是: defer,panic,recover

3）这几个异常的使用场景可以简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理

```go
package main
import (
	"fmt"
)

func test(){
	//使用defer +recover来捕获处理异常
	defer func() {
		err :=recover() //recover是内置函数，可以捕获到异常
		if err != nil { //说明捕获到错误
			fmt.Println("err=",err)
		}
	}()
	num1 :=10
	num2 :=0
	res :=num1/num2
	fmt.Println("res=",res)
}
func main(){
	//测试
	test()

	fmt.Println("main()下面的代码...")
    //执行结果如下：
err= runtime error: integer divide by zero
main()下面的代码...
```

错误处理的好处

**进行错误处理后，程序不会轻易挂掉。如果加入预警代码，就可以让程序更加健壮**，看一个案例演示：

```go
func test(){
	//使用defer +recover来捕获处理异常
	defer func() {
		//err :=recover() //recover是内置函数，可以捕获到异常
		if err :=recover(); err != nil { //说明捕获到错误
			fmt.Println("err=",err)
			//这里就可以将错误发送给管理员...
			fmt.Println("发送邮件给@wew")
		}
	}()
	num1 :=10
	num2 :=0
	res :=num1/num2
	fmt.Println("res=",res)
}
func main(){
	//测试
	
for {
	test()
	time.Sleep(time.Second)
	}
	fmt.Println("main()下面的代码...")
}
```

自定义错误

Go程序中，也支持自定义错误，使用errros.New和panic内置函数

1）errors.New("错误说明")，会返回一个erro类型的值，表示一个错误

2）panic内置函数，接收一个interface{}类型的值（也就是任何值了）作为参数。可以接收erro类型的变量，**输出错误信息，并退出程序**



案例演示：

```go
//函数去读取以配置文件init.conf的信息
//如果文件名不正确，我们就返回一个自定义的错误

func readconf (name string) (err error){
	if name == "config.ini"{
		//读取
		return nil
	}else{
		//返回一个自定义的错误
		return errors.New("读取文件错误")
	}
}
func test02() {
   
	err :=readconf("config.ini")
	if err != nil{
		//如果读取文件发生错误就输出这个错误并终止程序
		panic(err)
	}
	fmt.Println("test02()继续执行")
}
```

## 七、数组

### 1.为什么需要数组

问题：

一个养鸡场有6只鸡，他们的体重分别是3kg .5kg 1kg 3.4kg 2kg 50kg请问这六只鸡的总体重是多少？平均体重是多少?请你编写一个程序

传统方法：定义六个变量进行求值即可

问题：不利于数据的管理和维护，不够灵活，我们需要使用新的数据类型数组

数组介绍

**数组可以存放多个同一类型的数据。数组也是一种数据类型，在Go中数组就是值类型**

### 2.数组快速入门

```go
/*
一个养鸡场有6只鸡，他们的体重分别是3kg .5kg 1kg 3.4kg 
2kg 50kg请问这六只鸡的总体重是多少？平均体重是多少?
请你编写一个程序
*/
func main(){
 //使用数组来解决问题
 //1.定义一个数组
 var hens [6]float64
//2.给数组的每个元素赋值操作，元素下标从0开始
hens[0] =3.0 //hens数组的第一个元素赋值
hens[1] =5.0
hens[2] =1.0
hens[3] =3.4
hens[4] =2.0
hens[5] =50.0
//3.遍历数组求出总体重
totalweight :=0.0
for i :=0;i< len(hens);i++{
    totalweight += hens[i]
  }
  //4.求出总体重
//平均体重
avgweight := fmt.Sprintf("%.2f",totalweight/float64(len(hens)))
  fmt.Printf("鸡的总体重是：%v,平均体重是%v",totalweight,avgweight)
}

```

对上面代码总结

1）使用数组解决问题，程序的可维护性增加

2）而且的方法代码更加清晰，也容易扩展

### 3、数组的定义和内存布局

数组的定义

```
var 数组名[数组大小] 数据类型
var a[5]int
赋初值 a[0]=1 a[1]=30 ...
```

数组内存（重要）

```go
func main(){
	var intArr [3]int
	//当我们定义完数组后，数组的各个元素有默认值0
	fmt.Println(intArr)//[0 0 0]
	fmt.Printf("数组的地址是：%p",&intArr)//数组的地址是：0xc0420082c0
      fmt.Printf("数组首地址是：%p",&intArr[0])
    ////数组的首地址是：0xc0420082c0

}
```

1. 数组的地址可以通过数组名来获取：&intArr
2. 数组的第一个元素的地址就是数组的首地址

```go
func main(){
	var intArr [3]int//int占8个字节 如果是int32就是4个字节
	//当我们定义完数组后，数组的各个元素有默认值0
	fmt.Println(intArr)//[0 0 0]
	fmt.Printf("数组的地址是：%p\n",&intArr)
	//数组的地址是：0xc0420082c0
    fmt.Printf("数组首地址是：%p，地址intArr[1]的地址是%p",&intArr[0],&intArr[1])
    //数组首地址是：0xc04205c0a0，地址intArr[1]的地址是0xc04205c0a8

}
```

#### 数组的使用

访问数组元素

数组名[下标]比如：你要使用a数组的第三个元素 a[2]

案例：

循环输入5个成绩，保存到float64数组,并输出

```go
func arry(){
	//从终端输入5个成绩，保存到float64数组，并输出
	var score[5]float64
	for i:=0;i<5;i++{
		fmt.Printf("请输入第%v个元素的值：",i)
		fmt.Scanln(&score[i])
	}
	// fmt.Println("score的值是",score)
	//遍历数组打印
	for i :=0;i< len(score);i++{
		fmt.Printf("score[%v]=%v\n",i,score[i])
	}
}
```

四种初始化数组的方式：

```go
func main(){

// 四种初始化数组的方式
//way1
var numArr01 [3]int=[3]int{1,2,3}
fmt.Println("numArr01=",numArr01)
//输出结果为：numArr= [1 2 3]
//way2
var numArr02 =[3]int{5,6,7}
fmt.Println("numArr02=",numArr02)
//输出结果为：numArr02= [5 6 7]
//way3
var numArr03 =[...]int{8,9,10}
//这里的[...]是规定的写法
fmt.Println("numArr03=",numArr03)
//numArr03= [8 9 10]
// way4
var numArr04 =[...]int{1:800,0:900,2:999}
fmt.Println("numArr04=",numArr04)
//numArr04= [900 800 999]
 //类型推导   
 numArr05 :=[...]string{1:"tom",0:"jfon",2:"feilipu"}
 fmt.Println("numArr05=",numArr05)
//numArr05= [jfon tom feilipu]
}
```

#### 数组的遍历

for -range结构遍历

这是Go语言一种独有的遍历，可以用来遍历访问数组元素

基本语法

```
for index,value := range array01{
...
}
```

说明

- 第一个返回值index是数组的下标
- 第二个value是在该下标位置的值
- 他们都是仅在for循环内部可见的局部变量
- 遍历数组元素的时候，如果不想使用下标index，可以直接把下标index标为下划线
- index和value的名称不是固定的，即程序员可以自行指定，一般命名为index和value

案例演示：

```go
func main(){
	//演示for -range遍历数组
	heroes :=[...]string{"刘备","张飞","关羽"}
	fmt.Println(heroes)
	//for -range遍历
	for i,v :=range heroes{
		fmt.Printf("i=%v,v=%v",i,v)//i=0,v=刘备i=1,v=张飞i=2,v=关羽s
	//除此之外这样遍历也可以
        fmt.Printf("heroes[%d]=%v\n",i,heroes[i])
    }
    //不要元素的下标只要元素的值可以这样写：
    
    for _,v :=range heroes{
		fmt.Printf("元素的值=%v\n",v)
	}
}
}
```

#### 数组的注意事项和细节

1）数组是多个相同类型数据的组合，一个数组一旦声明/定义了，其长度是固定的，不能动态变化

```go
//1）数组是多个相同类型数据的组合，一个数组一旦
	// 声明/定义了，其长度是固定的，不能动态变化
	var arr01 [3]int
	arr01[0] =1
	arr01[1] =30
	//arr01[2] =1.1//这里会报错类型不一致
    arr01[2] =90
    // arr01[3] =890//数组会发生越界，超出指定范围长度

	fmt.Println(arr01)

```

2）var arr []int这时arr是一个slice切片

数组需要写大小 var arr[3]int这样的写法才是数组

3）数组中的元素可以是任何数据类型，包括值类型和引用类型，但是不能混用



4）数据创建后，如果没有赋值，有默认值

```go
数组类型数组 默认值为0
字符串数组，默认值""
bool数组，默认值为false

//数组创建后，如果没有赋值，有默认值(0值)
   //1.数值(整数系列，浮点数系列)=》0
   //2.字符串==》""
   //3.bool类型 ==》flase
  var arr01 [3]float32
  var arr02 [3]string
  var arr03 [3]bool
  fmt.Printf("arr01=%v arr02=%v arr03=%v",arr01,arr02,arr03)
//输出结果 arr01=[0 0 0] arr02=[  ] arr03=[false false false]

```

5）使用数组的步骤：1.声明数组并开辟空间 2给数组各个元素赋值 3使用数组

6）数组的下标是从0开始的

```go
//数组的下标是从0开始
var arr04 [3]string //0-2
fmt.Println(arr04[3])// 报错，原因是数组越界

```

7）数组下标必须在指定范围内使用，否则报panic,数组越界比如：var arr[5]int 则下标为0~4

8）Go的数组属于值类型，在默认情况下不是值传递，因此会进行值拷贝。数组间不会相互影响

![](D:\myfile\后端\go语言学习\pic\数组是值传递.jpg)

```go
//函数
func test01(arr [3]int){
  arr[0] = 88
  
}
main中进行调用
arr := [3]int{11,22,33}
test01(arr)
fmt.Println(arr)//输出结果仍然是：[11 22 33]	无影响
```

9）如想在其它函数中，去修改原来的数组，可以使用引用传递【指针方式】

![](D:\myfile\后端\go语言学习\pic\数组地址传递.jpg)

```go
//函数
func test02(arr *[3]int){
	(*arr)[0] = 88
	
  }
  main中进行调用
  arr := [3]int{11,22,33}
test02(&arr)
fmt.Println("main里面的arr的值是",arr)
//输出的内容是
// main里面的arr的值是 [88 22 33]
```

10）长度是数组类型的一部分，在传递函数参数时，需要考虑数组的长度看案例

```go
题1 默认值拷贝
func modify(arr []int){ //编译就直接报错因为没有指定长度
arr[0] = 100
fmt.Println("modify的值arr",arr)
}
func main(){
var arr = [...]int{1,2,3}
modify(arr)
}

题2 默认值拷贝
func modify(arr [4]int){ 
arr[0] = 100
fmt.Println("modify的值arr",arr)
}
func main(){
var arr = [...]int{1,2,3}
modify(arr)
}

题2 默认值拷贝
func modify(arr [4]int){ 
arr[0] = 100
fmt.Println("modify的值arr",arr)
}
func main(){
var arr = [...]int{1,2,3}
modify(arr)
}
//编译错误，长度是数据类型的一部分

题2 默认值拷贝
func modify(arr [3]int){ 
arr[0] = 100
fmt.Println("modify的值arr",arr) 100 2 3
}
func main(){
var arr = [...]int{1,2,3}
modify(arr) // 1 2 3
}
//这个正确，但是不能修改成功

```

#### 数组的应用案例

1）创建一个byte类型的26个元素的数组，分别放置'A'-'Z',使用for循环访问所有元素并打印出来，提示：字符数据运算'A'+1->'B'

```go
func th1(){//自己写的
	//声明一个数组
	var arr [26]byte
	arr[0]='A'
	for i :=1;i<26;i++{
		arr[i]=arr[i-1]+1
	}
	for a,v :=range arr{
		fmt.Printf("arr[%d]=%c ",a,v)
	}
}

func the1(){//老师写的
	var myChars [26]byte
	for i :=0; i < 26; i++ {
		myChars[i] ='A'+byte(i)//注意将i =>byte
	}
	for i :=0; i < 26; i++{
        fmt.Printf("%c  ",myChars[i])
	}
}
func main(){
  //th1()
  the1()

}
```

2）请求出一个数组的最大值，并得到对应的下标

```go
//请求出一个数组的最大值，并得到对应的下标
/*
1.声明一个数组[6]int{12,56,7,9,23,1}
2.假定第一个数为最大值，下标就为0
3，然后从第二个元素开始循环比较，如果发现有更大则交换
*/
func the2(){
  arr :=[6]int{12,56,7,90,23,1}
  var max int =arr[0]
  maxValIndex :=0
  for i :=1;i<len(arr);i++{
	if arr[i]>max{
		max=arr[i]
		maxValIndex=i
	}
  }
  fmt.Printf("max=%v,index=%v",max,maxValIndex) 
}
```

3）请求出一个数组的和以及他的平均值 。for-range

```go
//请求出一个数组的和以及平均值。for-range
func suma(){
	//1.声明一个数组arr :=[6]int{12,56,7,90,23,1}
	//2.求出sum
	//3.求出平均值
	arr :=[6]int{12,57,7,90,23,1}
	sum :=0
	for _,v :=range arr{
		//累积求和
		sum += v
	}
	fmt.Printf("数组的和是%v,数组的平均值是%.2f",sum,float64(sum)/float64(len(arr)))
}


func main(){
  suma()

}
```

数组的复杂使用---数组反转

要求：随机生成5个数，并将其反转打印

```go
/*
要求：随机生成5个数，并将其反转打印
思路
1.随机生成5个数，rand.Intn()函数
2，当我们得到随机数后就放到一个数组中 int数组
3.反转打印
*/
func fanzhuan(){
	var intArr3 [5]int
    len :=len(intArr3) //先算出数组的长度，避免反复调用
	//为了每次生成的随机数都不一样，我们需要给一个seed值
	rand.Seed(time.Now().UnixNano())
for i := 0; i < len);i++{
	intArr3[i] =rand.Intn(100) //0<=n<=100
}
    
   fmt.Println("交换前：",intArr3)
//3.反转打印,交换的次数是len/2 2.倒数第一个和第一个交换倒数第二个与第二个进行交换

temp :=0 //作为一个临时变量用于交换操作
for i := 0; i < len/2;i++{
	temp =intArr3[len-1 -i] //倒数第n个和第n个元素进行交换
	intArr3[len-1 -i]=intArr3[i]
	intArr3[i] = temp

}
fmt.Println(intArr3)

fmt.Println("交换后：",intArr3)


最后main中调用即可
```

### 4.slice切片

**为什么需要切片？**

先看一个需求：我们需要一个数组用于保存学生的成绩，但是学生的个数是不确定的，请问怎么办，解决方案：使用切片

#### 1.基本介绍

1）切片的英文slice

2）切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制

3）切片的使用和数组类似，遍历切片、访问切片的元素和求切片长度len(slice)都一样

4）切片的长度是可以变化的，因此切片是一个可以动态变化数组

5）切片定义的基本语法

```go
var 变量名 []类型
```

比如: var a []int

入门案例

```go
func main(){
	//演示切片的基本使用
	var intArr [5]int = [...]int{1,22,33,66,99}
	//声明定义一个切片
	/*
     1.slice 就是切片的名称
	 2.intArr[1:3] 表示slice 引用到intArr这个数组
	 3.应用inArr数组的起始下标为1终止下标为3不包含3
	*/
    slice := intArr[1:3] 
    fmt.Println("intArr=",intArr) // intArr= [1 22 33 66 99]
    fmt.Println("slice 的元素是=",slice)//slice 的元素是= [22 33]
    fmt.Println("slice 的元素的个数是=",len(slice))//slice 的元素的个数是= 2
	fmt.Println("slice 的容量是=",cap(slice))//slice 的容量是= 4
	//切片的容量是可以动态变化的 cability
    fmt.Printf("intArr[1]的地址=%p\n",&intArr[1])
	fmt.Printf("slice[0]的地址=%p slice[0]=%v\n",&slice[0],slice[0])
	slice[0]=34 //相当于*intArr[1]=34
	fmt.Println("intArr=",intArr)//intArr= [1 34 33 66 99]
}

```

2.切片在内存中的形式

为了让大家更加深入的理解切片，我们画图分析一下切片在内存中是如何布局的

![](D:\myfile\后端\go语言学习\pic\slice内存图.jpg)

这是一个非常重要的知识点

1）以前面的案例来分析切片在内存中的布局

2）切片底层的数据结构可以理解成一个结构体struct

3）输出切片和切片的引用地址

#### 2.切片使用的三种方式

##### way1

第一种方式：定义一个切片，然后让切片去引用一个已经创建好的数组像前面的案例

```go
func main(){
	//演示切片的基本使用
	var intArr [5]int = [...]int{1,22,33,66,99}
	//声明定义一个切片
	/*
     1.slice 就是切片的名称
	 2.intArr[1:3] 表示slice 引用到intArr这个数组
	 3.应用inArr数组的起始下标为1终止下标为3不包含3
	*/
    slice := intArr[1:3] 
    fmt.Println("intArr=",intArr) // intArr= [1 22 33 66 99]
    fmt.Println("slice 的元素是=",slice)//slice 的元素是= [22 33]
    fmt.Println("slice 的元素的个数是=",len(slice))//slice 的元素的个数是= 2
	fmt.Println("slice 的容量是=",cap(slice))//slice 的容量是= 4
```

##### way2

第二种方式：通过make来创建切片

基本语法：

```
var 切片名 []type = make([].len,[cap])
参数说明：type就是数据类型 len:大小 cap指定切片的容量可选
```

案例演示

```go
func main(){
    var slice []int =make([]int,4,10)
    fmt.Println(slice)//默认值为0
    fmt.Println("slice len=",len(slice),"slice cap=",cap(slice))
    slice[0]=100
    slice[2]=100
    fmt.Println(slice)
    
}

//演示切片的使用make
	var slice []float64 = make([]float64,5,10)
	slice[1] = 10
	slice[3] = 20
    //对于切片，必须使用make
	fmt.Println(slice)
	fmt.Println("slice的size=",len(slice))
	fmt.Println("slice的cap=",cap(slice))

```

![](D:\myfile\后端\go语言学习\pic\makeslice.jpg)

内部有个数组是不可见的

**对上面代码的小结：(面试重点)**

1）通过make方式创建切片可以指定切片的大小和容量

2）如果没有给切片的各个元素赋值，那么就会使用默认值[int,float=>0 string=>" " bool=>false]

3）通过make方式创建的切片对应的数组是由make底层维护，对外不可见，即只能通过slice方式去访问

##### way3

第3中方式：定义一个切片，直接就指定具体数组，使用原理类似make方式

```go

	//第3中方式：定义一个切片，直接就指定具体数组，使用原理类似make方式

	var strSlice []string = []string {"tom","jack","mary"}
	fmt.Println(strSlice)
	fmt.Println("strSlice的size=",len(strSlice))//3
	fmt.Println("strSlice的cap=",cap(strSlice))//3
```

3.切片的遍历

切片的遍历和数组一样，也有两种方式‘

1）for循环常规遍历方式

案例演示

```go

func main(){
	//使用常规的for循环遍历切片
	var arr [5]int=[...]int{10,20,30,40,50}
	slice :=arr[1:4]//20 30 40
	for i := 0;i<len(slice);i++{
		fmt.Printf("slice[%v]=%v ",i,slice[i])//slice[0]=20 slice[1]=30 slice[2]=40
	}

```

2）for-range结构遍历切片

案例演示：

```go
//使用for -range 方式遍历切片
	for i,v :=range slice{
		fmt.Printf("slice[%v]=%v ",i,v)////slice[0]=20 slice[1]=30 slice[2]=40
	}
```

#### 3.切片的注意事项

1）切片初始化时 var slice = arr[start index:endindex]

说明：从arr数组下标为startindex,取到下标为endindex的元素（不含 arr[endindex]）

2）切片初始化时，仍然不能越界。范围在[0-len(arr)]之间，但是可以动态增长

1）var slice = arr[0:end]可以简写 var slice = arr[:end]

2）var slice = arr[start:len(arr)]可以简写：var slice = arr[start:]

3）var slice = arr[0:len(arr)]可以简写：var slice = arr[:]

- cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素

- 切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组或者make一个空间供切片来使用

- 切片可以继续切片

  ```go
      slice2 := slice[1:2] //20
  	fmt.Println("slice1=",slice2) //[20]
  ```

  当一个切片元素发生变化，其关联的数组和其他切片也会变化。因为切片是引用数据类型

  3）用append内置函数，可以对切片进行动态追加

  ```go
  //用append内置函数，可以对切片进行动态追加
  	var slice3 []int = []int{100,200,300}
  	//通过append直接对slice3追加具体的元素
  	slice3 = append(slice3,400,500,600)
  	
  	fmt.Println("slice3=",slice3)
  
  	//通过append将切片slice3追加到slice3
  	slice3 = append(slice3,slice3...)
  	fmt.Println("slice3=",slice3)
  }
  ```

  

4）切片append操作的底层原理分析

![](D:\myfile\后端\go语言学习\pic\slice append底层2.jpg)

- 切片append操作的本质就是对数组的扩容
- go底层会创建一个新的数组newArr（安装扩容后的大小）
- 将slice原来包含的元素拷贝到新的数组newArr
- slice重新引用到newArr
- 注意newArr是在底层来维护的，程序员是不可见的
- 案例演示

 5)切片的拷贝操作

切片使用copy内置函数完成拷贝

```go
//切片的拷贝操作
	//切片的copy内置函数完成拷贝，举例说明
	var slice4 []int =[]int{1,2,3,4,5}
	var slice5 = make([]int,10)
	copy(slice5,slice4)//将slice4拷贝给slice5
	fmt.Println("slice4=",slice4)//[1 2 3 4 5]
	fmt.Println("slice5=",slice5)//[1 2 3 4 5 0 0 0 0 0]
```

(1)说明：copy(para1,para2):para1和para2都是切片类型

(2)按照上面的代码来看，slice4和slice5的数据空间是独立的，相互不影响，也就是说slice[0]=9999,slice5[0]仍然是1不会受到影响

思考题

```go
var a []int =[]int{1,2,3,4,5}
	var slice5 = make([]int,1)
	copy(slice5,a)//ok只拷贝一个元素
fmt.Println(slice5) //[1]
```

上面的代码没有问题

切片式引用类型，所以在传递时，遵守**引用传递机制**

```go
func main(){
var slice[]int
var arr[5]int = [...]int{1,2,3,4,5}
slice = arr[:]
var slice2 = slice
slice2[0] =10
fmt.Println("slice2",slice2) //[10,2,3,4,5]
fmt.Println("slice",slice)//[10,2,3,4,5]
fmt.Println("arr",arr) //[10,2,3,4,5]
}
```

```go
func test(slice []int){
	slice[0]=100
}
func main(){
	var slice = []int{1,2,3,4}
	fmt.Println("slice=",slice)//[1,2,3,4]
	test(slice)
	fmt.Println("slice=",slice)[100,2,3,4]
	}
```

#### 4.string和slice

1)string底层是一个byte数组，因此string也可以进行切片处理

案例演示：

```
func main(){
	//string底层是一个byte数组，因此string也可以进行切片处理
	str:= "hello mrliu"
	//使用切片获取mrliu
	slice :=str[6:]
	fmt.Println("slice",slice) //slice=mrliu
}
```

2)string和切片在内存种的形式，”abcd“画出内存示意图

![](D:\myfile\后端\go语言学习\pic\string切片示意图.jpg)

3）string是不可改变的，也就是说不能通过str[0]='z'方式来修改字符串

```go
	//string底层是一个byte数组，因此string也可以进行切片处理
	str:= "hello mrliu"
	//使用切片获取mrliu
	slice :=str[6:]
	fmt.Println("slice",slice) //slice=mrliu
 //string是不可改变的，也就是说不能通过str[0]='z'方式来修改字符串
//  str[0] = 'z' //错误，编译不会通过string是不可变的
```

4）如果需要修改字符串，可以先将string->[]byte / 或者 []rune ->修改->重写转成string

```go
//如果需要修改字符串，可以先将string->[]byte / 或者 []rune ->修改->重写转成string
"hello mrliu"=》改成"zello mrsliu"
 arr1 := []byte(str)
 arr1[0]='z'
 str = string(arr1)
 fmt.Println(str)//zello mrliu

//细节：我们转成[]byte后,可以处理英文和数字，但是没办法处理中文
//原因是 []byte字节来处理，而一个汉字是3个字节。因此就会出现erro
//解决办法是将string转成[]rune即可，因为[]rune是按照字符处理兼容汉字

arr1 := []rune(str)
arr1[0]='北' //utf-8 21271
str = string(arr1) 
fmt.Println(str)//北ello mrliu
```

4.切片的课堂练习

说明：编写一个函数fbn(n int),要求完成

1）可以接收一个n int

2）能够将斐波那契的数列放到切片中

3）提示，斐波那契的数列形式

arr[0] =1;arr[1]=1;arr[2]=2;arr[3]=3;arr[4]=5;arr[5]=8

```go
func fbn (n int)([]uint64){
	//声明一个切片，切片大小n
	fbnSlice :=make([]uint64,n)
	//第一个数和第二数为1
	fbnSlice[0]=1
	fbnSlice[1]=1
	//使用for循环来存放斐波那契的数列
	for i := 2; i < n ;i++{
		fbnSlice[i]=fbnSlice[i - 1] + fbnSlice [i - 2]
	}
	return fbnSlice
	
}
func main(){
	/*
说明：编写一个函数fbn(n int),要求完成

1）可以接收一个n int

2）能够将斐波那契的数列放到切片中

3）提示，斐波那契的数列形式
arr[0] =1;arr[1]=1;arr[2]=2;arr[3]=3;arr[4]=5;arr[5]=8
	
思路：
1.声明一个函数fbn(n int)([]uint64)
2.编写fbn(n int)进行for循环来存放斐波那契数列
*/
fnbSlice :=fbn(10)
fmt.Println(fnbSlice) //[1 1 2 3 5 8 13 21 34 55]

}
```

### 5.二维数组

（2）快速入门案例

请使用二维数组输出如下图形

```
000000
001000
020300
000000
```

使用方式：先声明/定义再赋值

实现：

1）语法：

```
var 数组名[大小][大小]类型
var arr [2][3]int,再赋值
```

2）代码演示

```go
/*
000000
001000
020300
000000
定义声明一个二维数组
*/
func demo1(){
	var arr [4][6]int
	 //赋初值
	 arr[1][2]=1
	 arr[2][1]=2
	 arr[2][3]=3
	 //遍历二维数组。按照要求输出图形
	 for i :=0; i< 4; i ++{
		for j :=0 ;j<6;j++{
			fmt.Print(arr[i][j]," ")
		}
		fmt.Println()
	 }	
}
func main(){
 demo1()
}
```

4)二维数组在内存中的存在形式（重点）

![](D:\myfile\后端\go语言学习\pic\二维数组在内存布局.jpg)

```go
func arrmemory(){
	var arr2 [2][3]int 
	arr2 [1][1]=10
	fmt.Println(arr2)

	fmt.Printf("arr2[0]的地址是%p\n",&arr2[0])
	//arr2[0]的地址是0xc04207a030 与arr2[1]相差24个字节
	fmt.Printf("arr2[1]的地址是%p\n",&arr2[1])
	//arr2[1]的地址是0xc04207a048

	fmt.Printf("arr2[0][0]的地址是%p\n",&arr2[0][0])
	//arr2[0][0]的地址是0xc04207a030
	fmt.Printf("arr2[1][0]的地址是%p\n",&arr2[1][0])
	// arr2[1][0]的地址是0xc04207a048
    
}
```



（2）使用方式2：直接初始化

1)声明 :

```
var 数组名 [大小][大小]类型 =[大小][大小]类型{{初值..},{初值...}}
```

2)赋值（有默认值，比如 int 类型就是0）

3）使用演示

```go
func demo3 (){
	var arr3 [2][3]int = [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println("arr3=",arr3)//arr3= [[1 2 3] [4 5 6]]
}
```

4）说明：二维数组在声明/定义时也应有四种写法【和一维数组类似】

```
var 数组名 [大小] [大小]类型 = [大小][大小]类型{{初值...},{初值...}}
var 数组名 [大小] [大小]类型 = [...][大小]类型{{初值...},{初值...}}
var 数组名 =[大小][大小]类型{{初值...},{初值...}}
var 数组名 =[...][大小]类型{{初值...},{初值...}}
```

(3)二维数组的遍历

1）双层for循环完成遍历

案例：

```go
func bainli(){
	//演示二维数组的遍历
	var arr3 = [2][3]int{{1,2,3},{4,5,6}}

	//for循环来遍历
	for i :=0;i<len(arr3);i++{
		for j :=0; j < len(arr3[i]); j++{
			fmt.Printf("%v\t",arr3[i][j])
		}
		fmt.Println()
	}
}
```

2）for-range方式完成遍历

案例演示：

```go
// for -range遍历
	for i,v := range arr3{
		for j , v2 := range v{
			fmt.Printf("arr3[%v][%v]=%v\t",i,j,v2)

		}
		fmt.Println()
		
	}
```

(3)二维数组的应用案例

定义一个二维数组，用于保存三个班，每个班五名同学成绩，并求出每个班级的平均分、以及所有班级的平均分

```go
package main
import (
	"fmt"
)
/*
定义一个二维数组，用于保存三个班，每个班五名同学成绩，
并求出每个班级的平均分、以及所有班级的平均分
*/
func classf(){
	//定义一个二维数组
	var scores [3][5]float64
	//2循环的输入数据
	for i :=0; i < len(scores); i++{
		for j :=0; j < len(scores[i]);j++{
			fmt.Printf("请输入第%d班的第%d个学生的成绩\n",i+1,j+1)
		    fmt.Scanln(&scores[i][j])
		}
	}
	fmt.Println(scores)

	//遍历输出成绩后的二维数组，统计平均分
	totalSum  := 0.0 //定义一个变量用于统计所有班级的分数
	for i :=0; i < len(scores); i++{
		sum := 0.0 //定义一个变量，用于累计各个班级的成绩
		for j :=0; j < len(scores[i]);j++{
			sum += scores[i][j]
		}
		totalSum += sum
		fmt.Printf("第%d班级的总分%v,平均分为%v\n",i+1,sum,
		sum/float64(len(scores[i])))
	}
	fmt.Printf("所有班级额度总分是%v,所有班级的平均分是%v",
	totalSum,totalSum/15)


}

func main(){
classf()
}
```



### 1.排序

#### 1）排序的基本介绍

排序就是将一组数据，依指定的顺序进行排列的过程

#### 2）排序的分类：

（1）内部排序：

指将需要处理的所有数据都加载到内部存储器中进行排序。

包括（**交换式排序法**、**选择式排序法和****插入排序法**）



（2）外部排序法

数据量过大，无法全部加载到内存中，需要借助外部存储进行排序。包括（**合并排序法**和**直接合并排序法**）



#### 3）交换式排序

交换式排序属于内部排序法，是运用数据值比较后，依判断规则对数据位置进行交换，以达到排序的目的

交换式排序又分为两种

1）冒泡排序法(Bubble sort)



2)快速排序法(Quick sort)

#### 4）交换式排序法--冒泡排序

（1）基本思想

冒泡排序（Bubble sort0）的基本思想是：通过对待排序序列从后向前（从下标较大的元素开始），依次比较相邻元素的排序码，若发现逆序交换，使排序码较小的元素逐渐从后部移向前部（从下标较大的单元移向下标较小的单元），就像水底下的气泡一样逐渐向上冒。

因为排序的过程中，各元素不断接近自己的位置，如果一趟比较下来没有进行过交换，就说明序列有序，因此要在排序过程中设置一个标志flag判断元素是否进行过交换，从而减少不必要的比较

（2）案例

我们将5个数字：24,69,80,57,13使用冒泡排序将其排成一个从小到大的有序数列

![](D:\myfile\后端\go语言学习\pic\冒泡排序比较图2.jpg)

```go
package main
import (
	"fmt"
)
//冒泡排序
func BubbleSort(arr *[5]int){
	fmt.Println("排序前的ar=",(*arr))
	temp :=0//临时变量用来做交换的
	//完成第一轮排序（外层排序）
	for j :=0;j < 4;j++{
		if (*arr)[j] > (*arr)[j+1]{
			//交换
            temp = (*arr)[j]
			(*arr)[j]=(*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第一次排序过后arr=",(*arr))//[24 69 57 13 80]

	//完成第二轮排序（外层排序）
	for j :=0;j < 3;j++{
		if (*arr)[j] > (*arr)[j+1]{
			//交换
            temp = (*arr)[j]
			(*arr)[j]=(*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第三次排序过后arr=",(*arr))//[24 57 13 69 80]
    //完成第二轮排序（外层排序）
	for j :=0;j < 2;j++{
		if (*arr)[j] > (*arr)[j+1]{
			//交换
            temp = (*arr)[j]
			(*arr)[j]=(*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第三次排序过后arr=",(*arr))//arr= [24 13 57 69 80]

	//完成第四轮排序（外层排序）
	for j :=0;j < 1;j++{
		if (*arr)[j] > (*arr)[j+1]{
			//交换
            temp = (*arr)[j]
			(*arr)[j]=(*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第四次排序过后arr=",(*arr))////完成第二轮排序（外层排序）
	for j :=0;j < 2;j++{
		if (*arr)[j] > (*arr)[j+1]{
			//交换
            temp = (*arr)[j]
			(*arr)[j]=(*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
	fmt.Println("第四次排序过后arr=",(*arr))//[13 24 57 69 80]

}

func BubbleSort2(arr *[5]int){
	fmt.Println("排序前的ar=",(*arr))
	temp :=0//临时变量用来做交换的
	len :=len((*arr))
	//完成第一轮排序（外层排序）
	for i :=0;i<len-1;i++{
	for j :=0;j < len-i-1;j++{
		if (*arr)[j] > (*arr)[j+1]{
			//交换
            temp = (*arr)[j]
			(*arr)[j]=(*arr)[j+1]
			(*arr)[j+1] = temp
		}
	}
}
	fmt.Println("排序过后arr=",(*arr))

}


	func main(){

	//定义一个数组
	arr := [5]int{24,69,80,57,13}
	//将数组传递给一个函数，完成排序
    //BubbleSort(&arr)
    BubbleSort2(&arr) //调用改写之后的冒泡排序

}
```

### 2.查找

1）介绍

在Golang中，我们常用的查找有两种

（2） 顺序查找

（2）二分查找



2）案例演示

（1）有一个数列：白眉鹰王、金毛狮王、紫衫龙王、青翼斧王

猜数游戏：从键盘任意输入一个名称，判断数列中是否包含此名称

```go
package main
import (
	"fmt"
)
/*
（1）有一个数列：白眉鹰王、
金毛狮王、紫衫龙王、青翼斧王

猜数游戏：从键盘任意输入一个名称，
判断数列中是否包含此名称
思路：
1.定义一个数组：白眉鹰王、金毛狮王、紫衫龙王、青翼斧王
2.从控制台接收一个名词，依次比较如果发现有就提示
*/
func find1(){
	names := [4]string{"白眉鹰王","金毛狮王","紫衫龙王","青翼斧王"}
	var heroName = " "
	fmt.Println("请输入要查找的人名...")
	fmt.Scanln(&heroName)
	//顺序查找第一种方式
	for i := 0; i < len(names); i++{
		if heroName == names[i]{
			fmt.Printf("找到了%v，下标%v",heroName,i)
			break
		} else if i ==(len(names)-1){//判断当i为最后一个下标
			fmt.Printf("没有找到%v",heroName)
		}
	}
}

//顺序查找第二种方式(推荐)
func find2(){
	names := [4]string{"白眉鹰王","金毛狮王","紫衫龙王","青翼斧王"}
	var heroName = " "
	fmt.Println("请输入要查找的人名...")
	fmt.Scanln(&heroName)

	index := -1
	for i :=0; i < len(names); i++{
		if heroName == names[i]{
			index =i
			break
		}
	}
	if index != -1{
		fmt.Printf("找到了%v，下标%v",heroName,index)
	}else{
		fmt.Printf("没有找到%v",heroName)
	}

}

func main(){
	// find1()
	find2()
  
}
```

 （2）请对一个有序数列进行二分查找{1,8,10,89,1000,1234}.输入一个数看看该数组是否存在此数，并且求出下标，如果没有就提示“没有这个数”会使用到递归

二分法查找的思路分析

```
arr = [1,8,10,89,1000,1234] 8
二分法查找的思路：比如我们要查找的数是findVal
1.arr是有一个有序数组，并且是从小到大排序
2.先找到中间的下标middle =(leftindex + rightindex)/2然后让中间的值和findval进行比较
逻辑：
2.1如果arr[middle]>findval,就应该问 leftindex----(middle -1)
2.1如果arr[middle]<findval,就应该问 middel+1----right
2.3如果Arr[middle]==findVal就找到
对上面的逻辑进行递归执行

递归退出条件
if lefetindex > rightindex
//找不到
return ..
思路---代码
```

```go
func BinaryFind(arr *[6]int,leftindex,rightindex,findVal int){
	//判断leftIndex是否大于rightindex
	if leftindex >rightindex{
		fmt.Println("没有找到")
		return
	}


	
	//先找到中间的下标
	middle :=(leftindex + rightindex) /2
	if (*arr)[middle] > findVal{
		//说明我们要查找的数，应该在 leftIndex ---middel-1之间
		BinaryFind(arr,leftindex,middle -1,findVal)
	}else if (*arr)[middle] < findVal{
		//说明我们要查找得数在middel + -----rightindex
		BinaryFind(arr,middle +1,rightindex,findVal)
	}else{ //就是当Arr[middle]==findVal的时候
		//找到了
		fmt.Printf("找到了下标为%v\n",middle)
	}


}

func main(){
    arr := [6]int{1,8,10,89,1000,1234}
    BinaryFind(&arr,0,len(arr)-1,1000)//找到了下标为4
}
```

## 八、map

### 1.map的基本介绍

map是key -value数据结构，又称为字段或者关联数组。类似其他编程语言的集合，在编程中经常使用到

### 2.map的声明

基本语法

```go
var map 变量名 map[keptype]valuetype
```

key可以是什么类型

golang中的map，的key可以是很多种类型，比如bool，数字，string,指针，channel,还可以是只包含前面几个类型的 接口，结构体，数组

**通常key为int 、string**

**注意：slice，map还有function不可以，因为这几个没法用 ==来判断**

valuetype可以是什么类型

valuetype的类型和key基本一样

**通常为：数字(整数，浮点数)string,map,struct**

map声明的举例

```go
var a map[string]string
var a map[string]int
var a map[int]string
var a map[string]map[string]string
```

**注意：声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用**

案例演示：

```go
func main(){
	//map的声明和注意事项，map是无序的数据结构
	var a map[string]string
	//使用、map前，需要先make,make的作用就是给map分配数据空间
	a  = make(map[string]string,10) //最大可以放10对
	a["ao1"]="宋江"
	a["ao2"]="吴用"
	a["ao3"]="李逵"
	a["ao4"]="林冲"
	a["ao5"]="吴用"
	a["ao5"]="无名" //会覆盖同key的值
    fmt.Println(a)
}
```

对上面代码的说明：

1）map在使用前一定要make

2）map的key是不能重复，如果重复了，则以最后这个key-value为准

3）map的value是可以相同的

4）map的key-value是无序的

### 3.map的使用方式

1）方式1

```go
func main(){
	//map的声明和注意事项，map是无序的数据结构
	var a map[string]string
	//使用、map前，需要先make,make的作用就是给map分配数据空间
	a  = make(map[string]string,10) //最大可以放10对
	a["ao1"]="宋江"
	a["ao2"]="吴用"
	a["ao3"]="李逵"
	a["ao4"]="林冲"
	a["ao5"]="吴用"
	a["ao5"]="无名" //会覆盖同key的值
    fmt.Println(a)
}
```

2）方式2

```go
//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)//map[no1:北京 no2:天津 no3:上海]

```

3）方式3

```go
//第三种方式
	heroes := map[string]string{
		"heroe1" : "宋江",
		"heroe2" : "林冲",
	}
herroes["hertoe3"] = "张顺"
fmt.Println(heroes)//map[heroe2:林冲 heroe1:宋江 heroes:张顺]
```

练习：演示一个key-value的value是map的案例

比如：我们要存放3个学生的信息，每个学生有name,sex,adress信息

思路：map[string]map[string]string

代码：

```go
package main
import(
	"fmt"
)
func main(){
	studentsMap := make(map[string]map[string]string) 
	studentsMap["no1"] = make(map[string]string,3)
	studentsMap["no1"]["name"] = "tom"
	studentsMap["no1"]["sex"] = "男"
	studentsMap["no1"]["adrress"] = "北京"
 //第二个学生
    studentsMap["no2"] = make(map[string]string,3)
	studentsMap["no2"]["name"] = "jhon"
	studentsMap["no2"]["sex"] = "男"
	studentsMap["no2"]["adrress"] = "上海"

	fmt.Println(studentsMap["no1"])
	fmt.Println(studentsMap["no2"])
	//打印结果如下
	/*
    map[sex:男 adrress:北京 name:tom]
    map[name:jhon sex:男 adrress:上海]
	*/
}
```

### 4.map的增删改查操作

#### 1）map增加和更新

map["key"] =value //如果key还没有，就是增加，如果key存在就是修改

```go
func main(){
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)
	//因为no3这个key已经存在，因此下面的这句话就是修改
	cities["no3"] = "深圳"
    fmt.Println(cities)


}
```

#### 2）map删除

说明：

delete(map,"key"),delete是一个内置函数，如果key存在，就删除该key-value.如果key不存在，不操作，但是也不会报错

func delete

```
func delete(m map[Type]Type1,key Type)
```

内建函数delete按照指定的键将元素从个映射中删除，若m为nil或无此元素，delete将不进行操作

案例演示：

```go
func main(){
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)
	//因为no3这个key已经存在，因此下面的这句话就是修改
	cities["no3"] = "深圳"
    fmt.Println(cities)
    //演示删除
	delete(cities,"no1")
	fmt.Println(cities)//map[no2:天津 no3:深圳]没有no1
    //当delete指定的key不存在时，删除不会操作，也不会报错
	delete(cities,"no5")
	fmt.Println(cities)////map[no2:天津 no3:深圳]
    
    //如果希望一次性删除所有的key
	//1.遍历啊所有的key,如何逐一删除[遍历]
	//2.直接make一个新空间
	cities = make (map[string]string)//效率较高
	fmt.Println(cities)
}
```

#### 3）map的查找

案例演示

```go
    cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
//演示map的查找
	val, ok := cities["no1"]
	if ok{
		fmt.Printf("有no1的key,值为：%v\n",val)
	}else{
		fmt.Println("没有no1的key，不存在这个值")
	}
```

注意：如果cities这个map中存在“no1",那么findRes就会返回true，否则返回false

4）map的遍历:

案例演示 相对复杂的map遍历，该map的value又是一个map

**说明：map的遍历使用for-range的结构遍历**

```go
//使用for-range遍历map
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
    
	for k,v := range cities {
		fmt.Printf("k=%v v=%v\n",k,v)

	}
	//输出结果如下
	//k=no1 v=北京k=no2 v=天津k=no3 v=上海
	
	//使用for-range遍历比较复杂的map
	studentsMap := make(map[string]map[string]string) 
	studentsMap["no1"] = make(map[string]string,3)
	studentsMap["no1"]["name"] = "tom"
	studentsMap["no1"]["sex"] = "男"
	studentsMap["no1"]["adrress"] = "北京"
 //第二个学生
    studentsMap["no2"] = make(map[string]string,3)
	studentsMap["no2"]["name"] = "jhon"
	studentsMap["no2"]["sex"] = "男"
	studentsMap["no2"]["adrress"] = "上海"
 
	for k1,v1 :=range studentsMap{
        fmt.Println("k1=",k1)
		for k2,v2 := range v1 {
			fmt.Printf("\t k2=%v v2 = %v\n",k2,v2)
		}
	}
	/*
    输出结果如下
	k1= no1
         k2=name v2 = tom
         k2=sex v2 = 男
         k2=adrress v2 = 北京
k1= no2
         k2=sex v2 = 男
         k2=adrress v2 = 上海
         k2=name v2 = jhon
     
```

map的长度

func len

```
func len(v Type)int
```

内置函数的len返回v的长度，这取决于具体类型：

```
数组：v中元素的数量
数组指针：*v中元素的数量(v为nil时panic)
切片\映射：v中元素的数量：若v为nil,len(v)为0
字符串：v中字节的数量
通道：通道缓存中队列(未读取)元素的数量nil,len(v)即为0
```

案例演示

```go
fmt.Printf("cities有%v 对key-value \n",len(cities)) //3对
```

### 5.map切片

基本介绍

切片的数据类型如果是map,则我们称为silice of map ,ma切片，这样使用规则map个数就可以动态变化了

案例演示

要求：使用一个map来记录monster的信息name和age，也就是说一个monster对应一个map，并且妖怪的个数可以动态的增加=>map切片

```go
func main(){
	//演示map切片的使用
	//1.声明一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string,2)//准备放入两个妖怪
	//2.增加一个妖怪的信息
	if monsters[0] ==nil {
		monsters[0] = make(map[string]string,2)
		monsters[0]["name"]="牛魔王"
		monsters[0]["age"]="500"
	}

	if monsters[1] ==nil {
		monsters[1] = make(map[string]string,2)
		monsters[1]["name"]="玉兔精"
		monsters[1]["age"]="400"
	}

	//下面这个写法越界.
	// if monsters[2] ==nil {
	// 	monsters[2] = make(map[string]string,2)
	// 	monsters[2]["name"]="狐狸精"
	// 	monsters[2]["age"]="300"
	// }
	//这里我们需要使用到切片的append函数，可以动态的增加monster
	//1.先定义一个monster信息
	newMonster := map[string]string{
		"name" : "新的妖怪 -火云邪神",
		"age" : "200",
	}
	monsters = append(monsters,newMonster)

	fmt.Println(monsters)
}
```

### 6.map排序

基本介绍

1）golang中没有一个专门的方法针对map的key进行排序

2）golang中的map默认是无序的，注意也不是按照添加的顺序存放的，你每次遍历，得到的输出可能也不一样

3）golang中map的排序，是先将key进行排序，然后根据key值遍历输出即可

案例演示

```go
func main(){
	//map排序
	map1 :=make(map[int]int,10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	fmt.Println(map1)//map[10:100 1:13 4:56 8:90]

	//如何按照map的key的顺序进行排序输出
	//1.先将map的key放入到切片中
	//2.对切片排序、
	//3.遍历切片，然后按照key来输出map的值
	var keys []int
	for k,_:=range map1{
		keys = append(keys,k)
	}
	//排序
	sort.Ints(keys)
	fmt.Println(keys)

	//遍历切片，然后按照key来输出map的值
	for _, k := range keys{
		fmt.Printf("map[%v]=%v\n",k,map1[k])
	}
}
```

### 7.map的使用细节

1）map是引用类型，遵守引用类型传递的机制，在一个函数接收map，修改后，会直接修改原来的map[案例演示]

```go
func main() {
	//map是引用类型，
	// 遵守引用类型传递的机制，在一个函数接收map，修改后，会直接修改原来的map
  map1 := make(map[int]int)
  map1[1] = 90
  map1[2] = 88
  map1[10] = 1
  map1[20] = 2
  modify(map1)
  //观察结果如果map1[10]= 900说明map是引用类型
  fmt.Println(map1) //map[10:900 20:2 1:90 2:88]

}
```

2）map的容量达到后，再想map增加元素，会自动扩容，并不会发生panic,也就是说map能动态的增长键值对(key -value)

3）map的value也经常使用struct类型，更适合管理复杂的数据(比前面value是一个map更好)，比如value为student结构体

```go
//定义一个学生结构体
type Stu struct {
	Name string
	Age int
	Address string
}
func main() {

//map的value也经常使用struct类型，更适合管理复杂的
// 数据(比前面value是一个map更好)，比如value为student结构体
//1.map的key为学生的学号是唯一的
//2.map的value为结构体，包含学生的名字，年龄，地址
students :=make(map[string]Stu, 10)
//创建2个学生
stu1 := Stu{"tom",18,"北京"}
stu2 := Stu{"jhon",19,"上海"}
students["no1"] = stu1
students["no2"] = stu2
fmt.Println(students) //map[no1:{tom 18 北京} no2:{jhon 19 上海}]

//遍历各个学生的信息
for k,v := range students {
	fmt.Printf("学生的编号是%v\n",k)
	fmt.Printf("学生的名字是%v\n",v.Name)
	fmt.Printf("学生的年龄是%v\n",v.Age)
	fmt.Printf("学生的住址是%v\n",v.Address)
	fmt.Println(" ")
}

}
//输出结果如下
map[no1:{tom 18 北京} no2:{jhon 19 上海}]
学生的编号是no2
学生的名字是jhon
学生的年龄是19
学生的住址是上海

学生的编号是no1
学生的名字是tom
学生的年龄是18
学生的住址是北京


D:\myfile\GO\project\src\go_code\map\mapdetails>
```

8.综合练习题：

1）使用map[string]map[string]string的map类型

2）key:表示用户名，是唯一的，不可以重复

3）如果某个用户名存在，就将其密码修改"8888",如果不存在就增加这个用户信息，（包括nickname和密码pwd）

4）编写一个函数modifyUser(users map[string]map[string] string,name string)完成上述功能

```go
package main
import(
	"fmt"
)
func modifyUsers(users map[string]map[string]string,name string){
  //判断users中是否有name
//   v , ok := users[name]
   if users[name] != nil {
	    //有这个用户
		users[name]["pws"] = "8888"
   }else {
	//没有这个用户
	users[name] = make(map[string]string,2)
	users[name]["pws"] = "8888"
	users[name]["nicname"] = "昵称" + name //示意
   }
}
func main(){
	users :=make(map[string]map[string]string)
	users["smith"] =make(map[string]string,2)
	users["smith"]["pwd"] = "999999"
	users["smith"]["nickname"] = "小花猫"
    modifyUsers(users,"tom")
    modifyUsers(users,"mary")
    modifyUsers(users,"smith")


	fmt.Println(users)
	//输出结果为：map[mary:map[pws:8888 nicname:
	//昵称mary] smith:map[nickname:小花猫 pws:8888 pwd:999999] tom:map[pws:8888 nicname:昵称tom]]
}
```

## 九、面向对象编程(上)

### 1.问题与解决思路

问题：

张老太养了两只猫，一个名字叫小白，今年3岁，白色。还有一只叫小花，今年100岁，花色。请编写一个程序。当用户输入小猫的名字时，就显示该猫的名字，年龄，颜色。如果用户输入的小猫名错误，则显示张老太没有这只小猫

使用现有的技术来解决这个问题

1）单独的定义变量解决

代码演示

```go
//1.使用变量的处理方式
	var cat1Name string = "小白"
	var cat1Age int = 3
	var cat1Color string = "白色"
    
	var cat2Name string = "小化"
	var cat2Age int = 100
	var cat2Color string = "花色"
```

2）使用数组解决

代码演示

```go
//2.使用数组来解决
	var catNames [2]string = [...]string{"小白","小花"}
    var catAges [2]int = [...]int{3,100}
	var catColor [2]string = [...]string{"白色","花色"}
	//
```

现有的技术解决的缺点分析

- 使用变量或者数组来解决养猫的问题，不利于数据的管理和维护。因为名字，年龄，颜色都是属于一只猫，但是这里分开保存
- 如果我们希望对一只猫的属性（名字、年龄、颜色）绑定方法也不好处理
- 引出讲解的知识点结构体



### 2.结构体

**一个程序就是一个世界，有很多对象（变量）**

#### 1》Golang语言面向对象编程说明

1）Golang也支持面向对象编程(OOP),但是和传统的面向对象编程有区别，并不是**纯粹的面向对象语言**。所以我们说**Golang支持面向对象编程的特性**是比较准确的

2）Golang没有类（class），go语言的结构体(struct)和其他编程语言的类(class)有同等的地位，你可以理解Golang是基于struct来实现OOP特性的

3）Golang面向对象编程非常简洁，去掉了传统OOP语言的继承，方法重载、构造函数和析构函数、隐藏的this指针等等

4）Golang仍然有面向对象编程的继承，封装和多态的特性，只是实现的方式和其他的OOP语言不一样，比如继承：Golang没有extends关键字，继承是通过匿名字段来实现

5）Golang面向对象(OOP)很优雅，OOP本身就是语言类型系统(type system)的一部分，通过接口(interface)关联，**耦合性低**，也非常灵活。后面学习就可以体会到了，也就是说在Golang中**面向接口编程**是非常重要的特性

#### 2》结构体与结构体变量(实例/对象)的关系的示意图

![](D:\myfile\后端\go语言学习\pic\面向对象\pic1.jpg)

简单来说就是类似与java的类和对象

对上图的说明

1）将一类事物的特性提取出来（比如猫类），形成一个新的数据类型，就是结构体

2）通过这个结构体，我们可以创建多个变量（实例对象）

3）事物可以是猫类，也可以是Person,Fish或是某个工具类



特征对象抽取图：

![](D:\myfile\后端\go语言学习\pic\面向对象\pic2.jpg)

#### 3》入门案例(using struct to solve the problem of cat growing)

```go
package main
import (
	"fmt"
)
/*
张老太养了两只猫，一个名字叫小白，今年3岁，白色。
还有一只叫小花，今年100岁，花色。请编写一个程序。
当用户输入小猫的名字时，就显示该猫的名字，年龄，颜色。
如果用户输入的小猫名错误，则显示张老太没有这只小猫

*/

// func vars(){
// 	//1.使用变量的处理方式
// 	var cat1Name string = "小白"
// 	var cat1Age int = 3
// 	var cat1Color string = "白色"
    
// 	var cat2Name string = "小化"
// 	var cat2Age int = 100
// 	var cat2Color string = "花色"
// 	//2.使用数组来解决
// 	var catNames [2]string = [...]string{"小白","小花"}
//     var catAges [2]int = [...]int{3,100}
// 	var catColor [2]string = [...]string{"白色","花色"}
// 	//
// }

//定义一个cat结构体,将cat的各个字段/属性放入cat结构体进行管理
type Cat struct {
	Name string
	Age int
	Color string
	Hobby string
}



func main() {
//创建一个cat结构体变量
var cat1  Cat  //类似于var a int
cat1.Name = "小白"
cat1.Age = 3
cat1.Color = "白色"
cat1.Hobby = "吃<.))>>>"
fmt.Println("cat1=",cat1)

fmt.Println("猫猫的信息如下：")
fmt.Println("name=",cat1.Name,"Age=",cat1.Age,"Color=",cat1.Color)
fmt.Println("cat1.Hobby",cat1.Hobby)
//输出结果如下：
/*cat1= {小白 3 白色}
猫猫的信息如下：
name= 小白 Age= 3 Color= 白色
*/
}
```



**4》结构体和结构体变量(实例)的区别和联系**

通过上面的案例和讲解我们可以看出

1）结构体是**自定义的数据类型**。代表一类事物

2）结构体变量(实例)是具体的，实际的，代表一个具体变量

**5》结构体变量（实例）在内存中存在的布局**

```go
var cat1 Cat
cat1.Age = 3
cat1.Color = "白色"

```

画出示意图P185(还没有看完因为其他事)

![](D:\myfile\后端\go语言学习\pic\面向对象\pic3.jpg)

### 3.结构体的具体应用

1)结构体的声明

```go
type 结构体名称 struct {
   filed1 type
   filed2 type
}

//举例
//如果结构体的名称首字母为大大写，那就意味着该结构体可以被其他包使用
type Student struct {
    Name string
    Age int
    Socre float32
}
```



2）字段/属性

基本介绍：

（1）从概念或叫法上看，结构体字段=属性=filed（集授课中，统一叫字段）

（2）字段是结构体的一一个组成部分，一般是基本数据类型、数组。比如我们前面定义猫结构体的Name string就是属性



3）字段和属性的细节说明

1、字段声明语法同变量，示例：字段名 字段类型

2、字段的类型可以为：基本类型、数组或引用类型

3、在创建一个结构体变量后，如果没有给字段赋值，都对应一个零值（默认值），规则同前面讲的一样

布尔类型 false,整型是0 ，字符是 " "

数组类型的默认值和它的元素类型相关,比如score [3]int的默认值是[0 0 0]

指针，slice和map的零值都是nil即没有分配空间（需要用make进行分配空间的操作）

```go
package main
import (
	"fmt"
)
//指针，slice和map的零值都是nil即没有分配空间
//如果结构体的字段类型是这些，需要先make后才可以使用
type Person struct {
	Name string
	Age int
	scores [5]float64
	ptr *int //指针
    slice []int//切片
	map1 map[string]string//切片
}

func main() {

	//定义结构体变量
	var p1 Person
	fmt.Println(p1) //未赋值之前输出的是{ 0 [0 0 0 0 0] <nil> [] map[]}
    
	if p1.ptr == nil {
		fmt.Println("ok1")
	}

	if p1.slice == nil {
		fmt.Println("ok2")
	}

	if p1.map1 == nil {
		fmt.Println("ok3")
	}

	//使用slice,之前一定要make
	p1.slice = make([]int,10)
	p1.slice[0] = 100 

	//使用map也一定先要make操作
    p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom"
	fmt.Println(p1)
}
```

4）不同结构体变量的字段是独立的，互不影响，一个结构体变量字段的更改，不影响另一个。结构体是值类型

```go
//不同结构体变量的字段是独立的，互不影响，
	// 一个结构体变量字段的更改，不影响另一个。结构体是值类型
    var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 500

	monster2 := monster1//默认结构体是值类型，值拷贝
	monster2.Name = "青牛精"
	fmt.Println("monster1=",monster1)//monster1= {牛魔王 500}
	fmt.Println("monster2=",monster2)//monster2= {青牛精 500}
	

}
```

示意图

![](D:\myfile\后端\go语言学习\pic\面向对象\pic4.jpg)

### 4.创建结构体变量和访问结构体字段

1）方式1-直接声明

```go
案例演示：var person Person
```

2)方式2-{}

```go
案例演示：var person Person = Person{}

//方式2
	p2 := Person{"mary",20}
	// p2.Name = "Tom"
	// p2.Age = 18
	fmt.Println(p2)
```

3）方式3

```
案例：var person *Person = new (Person)
```



```gos
//方式3-8
	//案例：var person *Person = new (Person)

	var p3 *Person= new(Person)
	//因为p3是一个指针，因此标准的给字段赋值方式
	//(*p3).Name = "smith"也可以这样写 p3.Name = "smith"
	//原因是go的设计者为了程序员使用方便，在底层会对p3.Name = "smith"进行优化
	//底层会给p3加上取值运算 (*p3).Name = "smith"
	(*p3).Name = "smith"
	p3.Name = "jhon"
	(*p3).Age = 30
	fmt.Println(*p3)//{jhon 30}
```

4）方式4-{}

```go
案例：var person *Person = &Person{}
//var person *Person = &Person{"jerry",60}写成这种形式也是对的

//方式4-{}案例：var person *Person = &Person{}
	var person *Person = &Person{}
	//因为person是一个指针，因此标准的访问其字段的方法是
    //(*peron).Name = "scott"
	//go的设计者为了方便,也可以person.Name = "scott"
	//底层会进行优化操作
	(*person).Name = "scott"
	(*person).Age= 30
	fmt.Println(*person)

```

说明

1. 第三种和第四种方式返回的是结构体指针
2. 结构体指针访问字段的标准方式应该是：(*结构体指针).字段名，比如(*Person).Name = "tom"
3. 但go做了一个简化，也支持 结构体指针.字段名,比如1perosn.Name = "tom".更加符合程序员使用的习惯，go编译器底层对person.Name做了转化(*Person).Name

### 5.struct类型的内存分配机制

看看结构体在内存中的存在形式

```go
var p4 Person
	p4.Name= "小明"
	p4.Age = 23
	var p5 *Person = &p4
    //fmt.Println(*p5.Age)这种用法是错误的
    //需要将*号包起来，因为.的优先级比*的优先级更高
	fmt.Println((*p5).Age)//10
	fmt.Println(p5.Age) //10

	p5.Name = "tom~"
	fmt.Printf("p5.Name=%v p4.Name=%v \n",p5.Name,p4.Name)
	fmt.Printf("p5.Name=%v p4.Name=%v \n",(*p5).Name,p4.Name)
	/*
	输出结果如下所示
	p5.Name=tom~ p4.Name=tom~
    p5.Name=tom~ p4.Name=tom~

	*/
fmt.Printf("p4的地址是%p\n",&p4)
	//p4的地址是0xc042056420
	
	fmt.Printf("p5的地址是%p p5的值是%p\n",&p5,p5)
	//p5的地址是0xc04207a020 p5的值是0xc042056420
```

其内存图分析如下

![](D:\myfile\后端\go语言学习\pic\面向对象\pic5.jpg)

### 6.结构体的注意事项和使用细节

#### 1）结构体的所有字段在内存中是连续的

```go
package main
import (
	"fmt"
)

//结构体
type Point struct {
	 x int
	 y int
}
//结构体
type Rect struct {
	leftup,rightDown Point
}

func main (){
    
	r1 := Rect{Point{1,2},Point{3,4}}

	//r1有4个int,在内存中连续分布
	//打印地址
	fmt.Printf("r1.leftup.x 地址=%p r1.leftup.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p\n",
    &r1.leftup.x,&r1.leftup.y,&r1.rightDown.x,&r1.rightDown.y)
	//输出结果如下
	//r1.leftup.x 地址=0xc04200a2c0 r1.leftup.y 地址=0xc04200a2c8 
	// r1.rightDown.x 地址=0xc04200a2d0 r1.rightDown.y 地址=0xc04200a2d8

    //r2有两个*Point类型，这两个*Point类型的本身地址也是连续的
	//但是他们指向的地址不一定是连续的
    
	r2 := Rect2{&Point{10,20}, &Point{30,40}}
	fmt.Printf("r2.leftup本身地址=%p  r2.rightDown 本身地址=%p \n",
    &r2.leftup,&r2.rightDown)
	//输出结果为:r2.leftup本身地址=0xc0420421c0  r2.rightDown 本身地址=0xc0420421c8
    
	//他们指向的地址不一定是连续的 这个要看编译器或系统运行时而定
	fmt.Printf("r2.leftup指向地址=%p  r2.rightDown 指向地址=%p \n",
    r2.leftup,r2.rightDown)
	//r2.leftup指向地址=0xc0420120b0  r2.rightDown 指向地址=0xc0420120c0


}
```

内存图

![](D:\myfile\后端\go语言学习\pic\面向对象\pic6.jpg)

![](D:\myfile\后端\go语言学习\pic\面向对象\pic7.jpg)

2）结构体是用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段（名字，个数和类型要完全一致）

```go
type A struct{
    Num int
}
type B struct{
   Num int
}
func main(){
   var a A
   var b B
  fmt.Println(a,b)
   //a=b两个不同类型的结构体不能相互赋值操作
   //进行强制转换
   a = A(b) //成功，前提是两个结构体的字段名称
   //，字段类型是一模一样的，其中一个不一样都不可以
   fmt.Println(a,b)
}

```

3）结构体进行type重新定义（相当于取别名），**Golang认为是新的数据类型，但是相互可以强转**

```go
type Student struct{
	Name string
	Age int
}
type Stu Student

func main(){
 
   var stu1 Student
   var stu2 Stu
 //stu2 = stu1 //正确吗？会报错，可以这样修改 stu2=Stu(stu1)
   stu2=Stu(stu1)
   fmt.Println(stu1,stu2)
}
```

```go
type integer int
func main(){
var i integer = 10
var j int = 20
//j = i //正确吗？ 也不可以的如果要就必须强制进行转换
j = int(i)
fmt.Println(i,j)
}
```

4）struct的每个字段上，可以写一个tag,该tag可以通过反射机制获取，常见的使用场景就是序列号和反序列化

举例：

```go
package main
import (
	"fmt"
	"encoding/json"
)
type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}
func main(){
//1.创建一个Monster变量
   monster :=Monster{"牛魔王",500,"芭蕉扇"}

   //2.将monster变量序列化为json格式的字符串
    //json.Marshal函数中使用反射
   jsonStr, err:= json.Marshal(monster)
   if err !=nil{
	fmt.Println("json处理错误",err)
   }
   fmt.Println("jsonStr",string(jsonStr))
   //输出如下：jsonStr {"Name":"牛魔王","Age":500,"Skill":"芭蕉扇"}
   //在字段后面加上json序列化后：
   //jsonStr {"name":"牛魔王","age":500,"skill":"芭蕉扇"}
}
```

序列化的应用场景：

![](D:\myfile\后端\go语言学习\pic\面向对象\pic8.jpg)

### 7.方法

#### 1)基本介绍：

在某些情况下，我们需要声明(定义)方法，比如Person结构体，除了有一些字段外(年龄，姓名..)Person结构体还有一些行为比如：可以说话、跑步。。通过学习，还可以做算术题，这时要**使用方法才可以完成**

Golang中的**方法是作用在指定的数据类型上**的（即。和指定的数据类型绑定）因此**自定义类型都可以有方法，而不仅仅是struct**

#### 2）方法的声明和调用

```go
type A struct{
    Num int
}
func (a A)test(){
    fmt.Println(a.Num)
}
```

对上面的语法的说明

- func (a A) test(){} 表示A结构体有一方法，方法名为test
- (a A ) 体现test方法是和A类型绑定的

举例说明：

```go
package main
import (
	"fmt"
)

type Person struct{
	Name string
}

//给A类型绑定一个方法
func (p Person) test() {
    p.Name = "jack"
	fmt.Println("test()",p.Name) //jack
}

func main(){
	//定义一个Person实例
	var p Person
	p.Name = "小红"
	p.test() //调用方法
    fmt.Println("main()",p.Name) //小红，这里不会因为test()方法中的操作而改变，因为struct不是引用传递而是值传递还是小红
}
```

#### 3）对上面的代码总结和说明

1. test方法和Person类型绑定

2. test方法只能通过Person类型的变量来调用，而不能直接调用，也不能使用其他类型变量来调用

   ```go
   下面的使用方式都是错误的
   var dog Dog
   dog.test()
   test()
   ```

   

3. func (p Person) test() {}...p表示哪个Person变量调用，这个p就是它的副本，这点和函数传参非常相似

4. p这个·名字，是可以有程序员指定，不是固定的，比如修改成person也是可以的

#### 4）方法的快速入门

（1）给Person结构体添加speak方法，输出xxx是一个好人

```go
func (p Person) speak(){
	fmt.Printf("%v是一个好人\n",p.Name)
}

```

（2）给Person结构体添加jisuan方法，可以计算从1+..+1000的结果

```go
func (p Person) jisuan() {
	var sum int = 0
	for i :=1;i <=1000;i++{
		sum += i
	}
	fmt.Printf("1+..+1000的结果是%v\n",sum)
}
```

（3）给Person结构体添加jisuan2方法,该方法可以接收一个数n，计算从1+...n的结果

```go
func (p Person) jisuan2(n int) {
    var sum int = 0
	for i :=1;i <=n;i++{
		sum += i
	}
	fmt.Printf("1+..+%v的结果是%v\n",n,sum)
}
```

（4）给Person结构体添加getSum方法,可以计算两个数的和并返回结果。

```go
func (p Person) getsum(n1 int,n2 int) int{
     return n1 + n2
}
```

（5）方法的调用

```go
    p.speak()
	p.jisuan()
	p.jisuan2(10000)
	res := p.getsum(10,20)
	fmt.Println("res=",res)
```

#### 6）方法的调用和传参机制原理(vital)

说明：方法的调用和传参机制和函数基本不一样，不一样的地方是方法调用时，会将调用方法的变量，当做实参也传递给方法。

案例1：画出前面getSum()方法的执行过程+说明

![](D:\myfile\后端\go语言学习\pic\面向对象\pic9.jpg)

**说明：**

1. 在通过一个变量去调用方法时，其调用机制和函数一样
2. 不一样的地方是，变量调用方法时，该变量本身也会作为一个参数传递到方法(如果变量是值类型，则进行值拷贝，如果变量是引用类型则进行地址拷贝)





案例2：请编写一个程序，要求如下

1. 声明一个结构体Circle，字段为radius
2. 声明一个方法area和CIrcle绑定，可以返回面积
3. 提示画出area执行过程+说明

```go
package main
import (
	"fmt"
)
/*
案例2：请编写一个程序，要求如下

1. 声明一个结构体Circle，字段为radius
2. 声明一个方法area和CIrcle绑定，可以返回面积
3. 提示画出area执行过程+说明
*/
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
    return 3.14 * c.radius * c.radius
}

func main(){
	//创建一个结构体
	var c Circle
	c.radius = 4.0
	area := c.area()
    fmt.Printf("圆的面积为%v\n",area)
}
```

执行内存图：

![](D:\myfile\后端\go语言学习\pic\面向对象\pic10.jpg)

其实这样的效率非常慢，会一对一进行值拷贝，用指针可以直接修改它就更快

#### 7）方法的声明(定义)

```go
func (recevier type) methodName (参数列表) (返回值列表){
         方法体
         return 返回值
}
```

1. 参数列表，表示方法输入
2. recevier type：表示这个方法和type这个类型进行绑定，或者说该方法作用于type类型
3. recevier type：type可以是结构体，也可以是其他自定义类型
4. recevier:就是type类型的一个变量（实例）比如：Person结构体的一个变量（实例）
5. 参数列表：表示方法输入
6. 返回值列表：表示返回的值，可以多个
7. 方法主体。表示为了实现某一个功能代码块
8. return语句不是必须的。

#### 8）方法注意事项和细节讨论

1. 结构体类型是值类型，在方法调用中，遵守值类型的传递机制，是值拷贝传递方式

   ```go
   //为了提高效率，通常我们方法和结构体的指针类型进行绑定
   func (c *Circle) area2() float64 {
   	//因为c是一个指针，因此我们标准访问其字段的方式是(*c)
       return 3.14 * (*c).radius * (*c).radius
   	//(*c).radius等价为c.radius
   }
   
   func main(){
   	//创建一个cicle变量
   	var c Circle
   	c.radius = 5.0
   	//res2 := (&c).area2()
   	//编译底层做了优化,(&c).area2()等价c.area2()
   	//因为编译器会自动加上
       res2 := c.area2()
   	fmt.Printf("圆的面积为%v\n",res2) //面积为78.5
   	
   }
   ```

   

2. 如果程序员希望在方法中，修改结构体变量的值，可以通过结构体指针的方式来处理

   ```go
   func (c *Circle) area2() float64 {
   	c.radius = 10
   	//因为c是一个指针，因此我们标准访问其字段的方式是(*c)
       return 3.14 * (*c).radius * (*c).radius
   	//(*c).radius等价为c.radius
   }
   
   func main(){
   	
   	//创建一个cicle变量
   	var c Circle
   	c.radius = 6.0
   	//res2 := (&c).area2()
   	//编译底层做了优化,(&c).area2()等价c.area2()
   	//因为编译器会自动加上
        res2 := c.area2()
   	fmt.Printf("圆的面积为%v\n",res2) //314
   	fmt.Println("c.radius = ",c.radius) //10
   	
   }
   ```

   ```go
   //为了提高效率，通常我们方法和结构体的指针类型进行绑定
   func (c *Circle) area2() float64 {
   	fmt.Printf("c是*cicle指向的地址=%p\n",c)
   	c.radius = 10
   	//因为c是一个指针，因此我们标准访问其字段的方式是(*c)
       return 3.14 * (*c).radius * (*c).radius
   	//(*c).radius等价为c.radius
   }
   
   func main(){
   	//创建一个结构体
   	// var c Circle
   	// c.radius = 4.0
   	// area := c.area()
       // fmt.Printf("圆的面积为%v\n",area)
   	//创建一个cicle变量
   	var c Circle
   	fmt.Printf("main的c结构体变量地址 =%p\n",&c)
   	c.radius = 6.0
   	//res2 := (&c).area2()
   	//编译底层做了优化,(&c).area2()等价c.area2()
   	//因为编译器会自动加上
   	res2 := c.area2()
   	fmt.Printf("圆的面积为%v\n",res2) //314
   	fmt.Println("c.radius = ",c.radius) //10
   	
   	/*
       运行结果如下
   	main的c结构体变量地址 =0xc0420120a0
       c是*cicle指向的地址=0xc0420120a0
       圆的面积为314
       c.radius =  10
   
   	*/
   }
   ```

   ![](D:\myfile\后端\go语言学习\pic\面向对象\pic11.jpg)

3. Golang中的方法作用在**指定的数据类型**上的（即：**和指定的数据类型绑定**）因此自定义类型，都可以有方法。而不仅仅是struct,比如int,float32等都可以有方法

   ```go
   package main
   import (
   	"fmt"
   )
   /*
   Golang中的方法作用在**指定的数据类型**上的
   （即：**和指定的数据类型绑定**）因此自定义类型，
   都可以有方法。而不仅仅是struct,比如int,float32等都可以有方法
   */
   type integer int
   func (i integer) print() {
   	fmt.Println("i=",i)
   }
   //编写一个方法修改i的值
   func (i *integer) ch() {
   	*i = *i +1
   }
   func main(){
       var i integer = 10
   	i.print() //1 = 10
   	i.ch()
   	i.print() //i =11
   }
   ```

   

4. 方法的访问范围控制的规则，和函数一样，方法名首字母小写，只能在本包访问，方法首字母大写，可以在本包和其他包访问

5. 如果一个类型实现了String()这个方法，那么fmt.Println默认会调用这个变量的String（）进行输出

   ```go
   //定义一个Student结构体
   type Student struct {
   	Name string
   	Age int
   }
   //给Student实现方法String()
   func (stu *Student) String() string {
      str := fmt.Sprintf("Name=[%v] Age=[%v]",stu.Name,stu.Age)
      return str
   }
   func main(){
   	//定义一个Student变量
   	stu := Student{
   		Name : "tom",
   		Age : 20,
   	}
   	//如果实现了*Student类型的String()方法会自动调用String()要传入&stu
   	fmt.Println(&stu) 
   }
   ```


#### 9）有关方法的课外习题

1. 编写结构体(MethodUtils),编程一个方法，方法不需要参数，在方法中打印一个10 *8 的矩形，在main方法中调用该方法

   ```go
   package main
   import (
   	"fmt"
   )
   
   type MethodUtils struct{
       //字段...
   }
   //给MethodUtils编写方法
   func (m MethodUtils) print(){
   	for i :=1; i<=10;i++{
   		for j :=1;j<=8;j++{
   			fmt.Print("* ")
   		}
   		fmt.Println()
   	}
   }
   func main(){
   	/*
      编写结构体(MethodUtils),编程一个方法，
      方法不需要参数，在方法中打印一个10 *8 的矩形，
      在main方法中调用该方法
   	*/
   	var m MethodUtils
   	m.print()
   }
   ```

   

2. 编写一个方法，提供m和n两个参数，方法中打印m*n的矩形

   ```go
   //编写一个方法，提供m和n两个参数，方法中打印m*n的矩形
   func (mu MethodUtils) print2(m int, n int){
   	for i :=1; i<=m;i++{
   		for j :=1;j<=n;j++{
   			fmt.Print("* ")
   		}
   		fmt.Println()
   	}
   } 
   func main(){
   	/*
      编写结构体(MethodUtils),编程一个方法，
      方法不需要参数，在方法中打印一个10 *8 的矩形，
      在main方法中调用该方法
   	*/
   	var m MethodUtils
   	// m.print()
       m.print2(3,4)
   }
   ```

   

3. 编写一个方法算该矩形的面积(可以接收长len和宽width),将其作为方法返回值。在main方法中调用该方法，接收返回的面积值并打印

   ```go
   func  (mu MethodUtils) area(length float32,width float32)float32{
         return length * width
   }
   func main(){
   	res :=m.area(4.0,5.6)
   	fmt.Println("该矩形的面积为：",res)
   }
   ```

   

4. 编写方法，判断一个数是奇数还是偶数

   ```go
   func (mu *MethodUtils) qi(m int) {
   	if m %2==0{
   		fmt.Printf("%v是偶数",m)
   	}else{
           fmt.Printf("%v是奇数",m)
   	}
   }
   ```

   

5. 根据行、列、字符打印对应行数和列数的字符，比如：行：3，列：2，字符 * ，则打印相应的效果

   ```go
   func (mu *MethodUtils) printe(n int, m int,key string) {
   	for i :=1;i<= n;i++{
   		for j :=1;j<=m;j++{
   			fmt.Print(key)
   		}
   		fmt.Println()
   	}
   }
   ```

   

6. 定义小小计算器结构体(Calcuator),实现加减乘除四个功能

   实现形式1：分4个方法完成

   ```go
   //实现形式1：
   type Calcuator struct{
   	Num1 float64
   	Num2 float64
   }
   //求和
   func (calcuator *Calcuator) getSum() float64{
      
   	return calcuator.Num1 +calcuator.Num2
   }
   //求差
   func (calcuator *Calcuator) getJian() float64{
      
   	return calcuator.Num1 -calcuator.Num2
   }
   //求积
   func (calcuator *Calcuator) getJi() float64{
      
   	return calcuator.Num1 *calcuator.Num2
   }
   //求商
   func (calcuator *Calcuator) getShang() float64{
      
   	return calcuator.Num1 /calcuator.Num2
   }
   func main(){
   	 n := Calcuator{24.0,8.0}
   	he :=n.getSum()
   	cha :=n.getJian()
   	ji :=n.getJi()
   	shangha :=n.getShang()
   	fmt.Printf("和为%v 差为%v 积为%v 商为%v",he,cha,ji,shangha)
   }
   ```

   实现形式2：用一个方法搞定

   ```go
   //实现形式2
   func (calcuator *Calcuator) getRes(operator byte) float64 {
      res :=0.0
      switch operator {
      case '+':
   	   res = calcuator.Num1 +  calcuator.Num2
      case '-':
   	   res = calcuator.Num1 -  calcuator.Num2
      case '*':
   	   res = calcuator.Num1 *  calcuator.Num2
      case '/':
   	   res = calcuator.Num1 / calcuator.Num2
      default:
   	   fmt.Println("运算符输入有误")
   	}
   	return res
   }
   func main(){
   s := Calcuator{32.0,8.0}
   	res :=s.getRes('*')
   	fmt.Println("运算的结果为：",res)
   }
   ```

   #### 10）方法和函数的区别

   1. 调用方式不一样

      函数的调用方式 ： 函数名（实参列表）

      方法的调用方式：变量.方法名（实参列表）

   2. 对于普通函数，接收者为值类型的时候，不能将指针类型的数据直接传递，反之亦然

      ```go
      type Person struct{
      	Name string
      }
      
      //函数
      func test01(p Person){
          fmt.Println(p.Name)
      }
      
      func test02(p *Person){
          fmt.Println(p.Name)
      }
      func main(){
       p := Person{"tom"}
        test01(p)
      //  test01(&p)//不可以传地址
      test02(&p)
      
      }
      
      ```

      

   3. 对于方法如(struct的方法)，接受者为值类型的时候，可以直接用指针类型的变量调用方法，反过来同样也可以

      ```go
      //方法
      func (p Person) test03(){
        fmt.Print("test03=",p.Name)
      }
      
      func main(){
       p := Person{"tom"}
        
      p.test03() //可以
      (&p).test03()//指针传入地址也可以，从形式上是传入了地址但是本质上仍然是值拷贝
      }
      
      ```

      总结：

      1. 不管调用形式如何，真正决定是值拷贝还是地址拷贝，看这个方法是和哪个类型绑定
      2. 如果是和值类型，比如(p Person),则是值拷贝，如果是指针类型，是(p *Person)则是地址拷贝

### 8.面向对象编程的应用实例

1）步骤

1. 声明(定义)结构体，确定结构体名
2. 编写结构体的字段
3. 编写结构体的方法

2）案例

1. 编写一个Student结构体，包含name、gender、age、id、score字段，分别是string、string、int、int、float54类型

2. 结构体中声明一个say()方法，返回string类型，方法返回信息包含所有字段值

3. 在main1方法中，创建一个stsudent结构体实例(变量),并访问say方法，并将调用结构打印输出

   ```go
   package main
   import (
   	"fmt"
   )
   type Student struct{
      name string
      gender string
      age int
      id int
      score float64
   
   }
   func (stu *Student) say() string{
   	//使用fmt.Sprintf()函数将值转换为字符串
       inforstr :=fmt.Sprintf("student的信息如下 name = [%v] gender=[%v] age = [%v] id=[%v] score=[%v]",stu.name,stu.gender,stu.age,stu.id,stu.score)
   	 return inforstr
   }
   func main(){
   var stu = Student{"小红","男",34,1001,99.5}
   info :=stu.say()
   fmt.Println(info)
   }
   ```

3)案例2：盒子案例

1. 编程创建一个Box结构体，在其中声明三个字段表示一个立方体的长、宽、高，长宽高要从终端获取

2. 声明一个方法获取立方体的体积

3. 创建一个Box结构体变量，打印给定尺寸的立方体的体积

   ```go
   type Box struct{
   	length float64
       width float64
   	high  float64
   }
   func (b Box) getVolumn() float64{
       return b.length * b.width * b.high
   }
   
   func main(){
   //测试box
   var box Box
   box.length = 1.1
   box.width = 2.0
   box.high = 3.0
   Volum := box.getVolumn()
   fmt.Printf("这个黑子的体积是%.2f",Volum)
   }
   ```

3)案例3：景区门票案例

1. 一个景区根据游人的年龄收取不同价格的民票，比如年龄大于18，收费20元，其他情况门票免费

2. 请编写Visitor结构体，根据年龄段决定能够购买的门票价格并输出

   ```go
   type Vistor struct{
   	Name string
   	Age int
   }
   func (vistor Vistor) showPrice(){
   	if vistor.Age >18{
   		fmt.Println("门票价格为20元")
   	}else{
   		fmt.Println("免费")
   	}
   
   }
   func main(){
       var vio Vistor
   
   for {
   	fmt.Println("请输入你的名字")
   	fmt.Scanln(&vio.Name)
   	if vio.Name == "n" {
   		fmt.Println("退出程序")
   	    break
   	}
   	fmt.Println("请输入你的年龄")
   	fmt.Scanln(&vio.Age)
   	vio.showPrice()
   }
   }
   ```

   

## 十、面向对象编程（下）

### 1.创建结构体变量时指定字段值

1）说明：Golang在创建结构体实例（变量）时，可以直接指定字段值

2）创建结构体变量时指定字段值方式

1. way1

   ```go
   //在创建结构体变量时，就直接指定字段的值
   	var stu1 = Stu{"tom",23}
   	stu2 :=Stu{"晓明",20}
   
   //创建结构体变量时。把字段名和字段值写在一起
   	var stu3 = Stu{ //这种写法更加稳健不依赖于字段的顺序
   		Name : "jack",
   		Age : 20,
   	}
   	//可以类型颠倒
   	var stu3 = Stu{ //这种写法更加稳健不依赖于字段的顺序
   		Age : 20,
   		Name : "jack",
   	}
   ```

   

2. way2

   ```go
   //方式2.返回结构体的指针类型(!!!)
   var stu5 *(Stu)= &Stu{"小王",29} //stu2--->地址 ---》结构体数据[xxxx,xxxx]
   	stu6 :=&Stu{"小王",39}
   
   	//在创建结构体指针变量是，把字段名和字段值写在一起，这种写法，就不依赖于字段的定义顺序
   	var stu7 = &Stu{
   		Name : "小李",
   		Age : 49,
   	}
   
   	var stu8 = &Stu{
   		Age :59,
   		Name : "小李~",
   
   	}
   	fmt.Println(*stu5,*stu6,*stu7,*stu8)//
   }
   ```

   

### 2.工厂模式

1）说明

Golang的额结构体没有构造函数，通常可以使用工厂模式来解决这个问题

看一个需求

一个结构体的声明是这样的

```go
package model
type Student struct{
Name string ....
}
```

因为这里的Student的首字母S是大写的，如果我们想在其它包创建Student的实例(比如main包)，引入model包后，就可以直接创建Student结构体的变量(实例)

**但是问题来了，如果首字母是小写的，比如是type student struct {,,,}就不行了，怎么办---》工厂模式来解决**

2）工厂模式来解决问题

使用工程模式实现挎包创建结构体实例(变量)的案例

1》如果model包的结构体变量首字母大写，引入后，直接使用没有任何问题

```go
package model
//定义一个结构体
type Student struct{
	Name string
	Score float64
}
package main
import (
	"fmt"
	"go_code/object2/factory/model"
)

func main(){
	//创建一个Student的实例
	var stu = model.Student{
		Name : "Tom",
		Score : 78.9,
	}
	fmt.Println(stu)
}

```

2》如果model包的结构体变量首字母小写，引入后，不能直接使用，可以使用工厂模式解决

```go
package model
//定义一个结构体
type student struct{
	Name string
	Score float64
}
//因为student结构体首字母是小写，因此是只能在
//model包使用，通过工厂模式解决
func NewStudent(n string,s float64) *student{
	return &student{
		Name : n,
		Score : s,
	}
}

在main包下的main,go中使用
package main
import (
	"fmt"
	"go_code/object2/factory/model"
)

func main(){
	//当student结构体首字母是小写的我们可以通过工厂模式解决
	var stu = model.NewStudent("tom~",88.8)
	fmt.Println(*stu)
}

```

问题：

思考一下，如果model包的Student的结构体字段Score改为score,我们还可以正常访问吗？当然不可以正常访问，我们要使用一个方法去访问

```go
package model
//定义一个结构体
type student struct{
	Name string
	score float64
}

//因为student结构体首字母是小写，因此是只能在
//model包使用，通过工厂模式解决
//返回这个结构体变量的指针
func NewStudent(n string,s float64) *student{
	return &student{
		Name : n,
		score : s,
	}
}

//如果score字段首字母小写，则在其他包不可以访问，我们可以提供方法
//就是用结构体指针来访问结构体字段
func (s *student) GetScore() float64{
    return s.score //（底层会优化为//return (*s).score）
}


package main
import (
	"fmt"
	"go_code/object2/factory/model"
)

func main(){
	
	//当student结构体首字母是小写的我们可以通过工厂模式解决
	var stu = model.NewStudent("tom~",88.8)
	fmt.Println(*stu)
    //注意了，这里我们使用getScore方法去访问stu,score
	fmt.Printf("name=%v Score=%v",stu.Name,stu.GetScore())
    //fmt.Printf("name=%v score的值为=%v",stu.Name,(*stu).GetScore())
}

```

### 3.vscode快捷键

1）删除当前行 ctrl +shift +k [也可自定义]

2）向上/向下复制当前行Shfit + Alt + 下箭头/上箭头

3）补全代码 alt +/

4）添加注释和取消注释 ctrl + /

5）快速修复 alt + /

6）快速格式化代码 shift + alt +f

特别说明：

1）VSCODE的快捷键不要和输入法冲突，否则不会生效

2）一些快捷键需要安装Go插件后才能生效

### 4.面向对象编程的三大特性

#### 1）基本介绍

Golang仍然有面向对象编程的继承，封装和多态的特性，只是实现的方式和其他的OOP语言不一样，下面就一一来讲解吧

#### 2）面向对象编程思想--抽象

如何理解抽象？

我们在定义一个结构体的时候，实际上就是把一类事物的共有的属性和行为提取出来，形成一个**物理模型**（结构体），这种研究问题的方法称为抽象。

![](D:\myfile\后端\go语言学习\pic\面向对象\pic12.jpg)

代码实现：

```go
package main
import (
	"fmt"
)
//定义一个结构体
type Account struct{
	AccountNo string
	Pwd string
	Balance float64
}

//方法
//1.存款
func (account *Account) Deposite(money float64,pwd string){
	//看看输入的密码是否正确
	if pwd !=account.Pwd{
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance += money
	fmt.Println("存款成功")
}
//取款
func (account *Account) WithDraw(money float64,pwd string){
	//看看输入的密码是否正确
	if pwd !=account.Pwd{
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看存款金额是否正确
	if money <= 0 || money >account.Balance{
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance -= money
	fmt.Println("取款成功")
}

//查询余额

func (account *Account)Query(pwd string){
	//看看输入的密码是否正确
	if pwd !=account.Pwd{
		fmt.Println("你输入的密码不正确")
		return
	}
	fmt.Printf("你的账号为=%v 余额=%v \n",account.AccountNo,account.Balance)
}


func main(){
	//测试一下
	account := Account{
		AccountNo : "gs111111",
		Pwd : "666666",
		Balance  : 1000.0,
	}
	for{
		//这里可以做的更加灵活，就是让用户通过控制台来输入命令
		//菜单
	account.Query("666666")
	account.Deposite(200,"666666")
	account.Query("666666")
	account.WithDraw(150.0,"666666")
	account.Query("666666")
	}

}
```

#### 3）面向对象编程思想-封装

##### &1介绍

封装就是把抽象出来的字段和对字段的操作封装在一起，数据被保护在内部，程序的其他包只有通过被授权的操作（方法），才能对字段进行操作

##### &封装的理解和好处

1. 隐藏实现细节

2. 可以对数据进行验证，保证安全合理

   ```go
   type Person struct{
   Age int
   }
   ```

   

##### &如何实现封装

1. 对结构体中的属性进行封装(字段名小写)
2. 通过方法，包 实现封装

4）封装的实现步骤

1. 将结构体、字段（属性）的首字母小写（不能导出了，其他包不能使用，类似private）

2. 给结构体所在包提供了一个工厂模式的函数，首字母大写。类似一个构造函数

3. 提供一个首字母大写的Set方法（类似其他语言的public）,用于对属性判断并赋值

   ```go
   func (var 结构体类型名)SetXxx(参数列表)(返回值列表){
       //加入数据验证的业务逻辑
       var.字段 = 参数
   }
   ```

   

4. 提供一个首字母大写的Get方法（类似其他语言的public）,用于获取属性的值

   ```
   func (var 结构体类型名)GetXxx(){
   return var.字段
   }
   ```

   

   特别说明：在Golang开发中并没有特别强调封装，这点并不像java，所以学过java的朋友不能总是用java语法特征来看待Glang,Golang本身对向对象的特性做了简化。

##### &案例入门

有一个程序，不能随便查看人的年龄，工资等隐私，并对输入的年龄进行合理的验证.设计：model包(person go) main包(main.go调用Person结构体)

```go
person.go
package model
import (
	"fmt"
)
type person struct{
	Name string
	age int //字母小写，其他包不能访问
	sal float64
}

//写一个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person{
	return &person{
		Name : name,
	}
}
//为了访问age和sal我们编写一个SertXxx的方法和GetXxx方法
func (p *person) SetAge(age int) {
	if age >0 && age < 150 {
		p.age = age
	}else{
		fmt.Println("年龄范围不正确...")
	}
}

func (p *person) GetAge() int {
	return p.age
}
func (p *person) SetSal(sal float64){
    if sal >=3000 && sal <= 30000 {
        p.sal = sal
	}else{
		fmt.Println("年龄范围不正确...")
	}
}

func (p *person) GetSal() float64{
	return p.sal
}
在其他包的main.go进行调用
package main
import (
	"fmt"
	"go_code/object2/encapsulate/model"
)

func main (){
	p :=model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000.0)
	fmt.Println(*p)
	fmt.Println(p.Name,"age =",p.GetAge(),"sal = ",p.GetSal())
}
```

##### &练习：

创建一个程序在model包中定义一个Account结构体：在main函数中体会Golang封装性

1. Account结构体要求具有字段：账号（长度在6-10之间）、余额(必须>20)、密码（必须是六位）
2. 通过SetXxx的方法给Account的字段赋值
3. 在main函数中测试

```go
package model
import (
	"fmt"
)
//定义一个结构体
type account struct{
	accountNo string
	pwd string
	balance float64
}

//工厂模式的函数-构造函数
func NewAccount(accountNo string,pwd string,balance float64) *account{
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号的长度不对...")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码的长度不对...")
		return nil
	}
	if balance < 20 {
		fmt.Println("余额数目不对")
		return nil
	}
	return &account{
		accountNo : accountNo,
		pwd : pwd,
		balance : balance,
	}
}




//方法
//1.存款
func (ac *account) Deposite(money float64,pwd string){
	//看看输入的密码是否正确
	if pwd !=ac.pwd{
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	ac.balance += money
	fmt.Println("存款成功")
}
//取款
func (ac *account) WithDraw(money float64,pwd string){
	//看看输入的密码是否正确
	if pwd !=ac.pwd{
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看存款金额是否正确
	if money <= 0 || money >ac.balance{
		fmt.Println("你输入的金额不正确")
		return
	}

	ac.balance -= money
	fmt.Println("取款成功")
}


//查询余额
func (ac *account)Query(pwd string){
	//看看输入的密码是否正确
	if pwd !=ac.pwd{
		fmt.Println("你输入的密码不正确")
		return
	}
	fmt.Printf("你的账号为=%v 余额=%v \n",ac.accountNo,ac.balance)
}

main。go中进行测试
package main
import (
	"fmt"
	"go_code/object3/ex2/model"
)
func main(){
	account := model.NewAccount("jzh1111","999999",48)
   if account !=nil{
       fmt.Println("创建成功=",*account)
   }else{
	   fmt.Println("创建失败")
   }
}

```

#### 4）面向对象编程思想-继承

##### &1.为什么需要继承

当然是为了代码复用

走代码看看

```go
package main
import (
	"fmt"
)
//编写一个学生考试系统
type Pupil struct {
	Name string
	Age int
	Score int
}

//显示成绩
func (p *Pupil) ShowInfo(){
   fmt.Printf("显示学生名字=%v 年龄=%v 成绩=%v",p.Name,p.Age,p.Score)
}
func (p *Pupil) SetScore(score int) {
	
	p.Score = score
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试")
}
//大学生，研究生。。
type Graduate struct {
	Name string
	Age int
	Score int
}

//显示成绩
func (p *Graduate) ShowInfo(){
   fmt.Printf("显示学生名字=%v 年龄=%v 成绩=%v",p.Name,p.Age,p.Score)
}
func (p *Graduate) SetScore(score int) {
	
	p.Score = score
}

func (p *Graduate) testing() {
	fmt.Println("大学生正在考试")
}

func main(){
  //测试
  var pupil = &Pupil{
	Name : "Tom",
	Age : 10,
  }

  pupil.testing()
  pupil.SetScore(90)
  pupil.ShowInfo()

  //测试
  var graduate = &Graduate{
	Name : "jhon",
	Age : 10,
  }

  graduate.testing()
  graduate.SetScore(90)
  graduate.ShowInfo()
}
//代码冗余 。。。高中生。。。

```

对上面代码做一个小结

1. Pupil和Graduate两个结构体的字段和方法集合一样，但我们欠缺写了相同的代码，代码复用性不强
2. 出现代码冗余，而且代码不利于维护，同时也不利于功能的扩展
3. 解决方法：通过继承的方式解决

##### &2继承的基本介绍和示意图

继承可以**解决代码复用**，让我们的编程更加靠近人类思维

当多个结构体存在相同的属性（字段）和方法时，可以从这些结构体中抽象出结构体（比如刚才的Student），在该结构体中定义这些相同的属性和方法。

其他的结构体不需要重新定义这些属性（字段）和方法，只需嵌套一个Student匿名结构体即可

![](D:\myfile\后端\go语言学习\pic\面向对象\pic13.jpg)

也就是说：在Golang中，如果一个struct嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，**从而实现了继承特性**

##### &3.嵌套匿名结构体的基本语法

```go
type Goods struct{
   Name string
   Price int
}
type Book struct {
     Goods//这里就是嵌套匿名结构体Goods相当于加入了Name和Pirce这两个字段
     Writer string
}
```

&4.应用案例

我们对先前有关student的结构体进行改进操作减少其冗余性,代码实现

```go
package main
import (
	"fmt"
)
//编写一个学生考试系统
type Student struct{
	Name string
	Age int
	Score int
}
//将Pupil和Graduate共有的方法也绑定到Student
func (stu *Student) ShowInfo(){
	fmt.Printf("显示学生名字=%v 年龄=%v 成绩=%v",stu.Name,stu.Age,stu.Score)
}

func (stu *Student)SetScore(score int) {
	
	stu.Score = score
}

//给*Student增加一个方法，那么Pupil 和Graduate都可以使用该方法
func (stu *Student) GetSum(n1 int, n2 int) int{
	return n1 + n2
}

//小学生
type Pupil struct {
	Student  //嵌入Student匿名结构体
}

//这是Pupil结构体特有的方法，保留
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试")
}
//大学生，研究生。。
type Graduate struct {
	Student  //嵌入Student匿名结构体
}
//保留特有的方法
func (p *Graduate) testing() {
	fmt.Println("大学生正在考试")
}

func main(){
  
	//当我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil :=&Pupil{}
	pupil.Student.Name ="tom~"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()
    
	fmt.Println("res=",pupil.Student.GetSum(1,2))
	//大学生进行操作
	graduate :=&Graduate{}
	graduate.Student.Name ="MARY"
	graduate.Student.Age = 28
	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()
	fmt.Println("res=",graduate.Student.GetSum(10,79))

}

```

&4.继承的深入讨论1

1. 结构体可以使用嵌套匿名结构体所有的字段和方法，即：首字母大写或者小写的字段、方法，都可以使用。

   ```go
   package main
   import (
   	"fmt"
   )
   
   type A struct {
   	Name string
   	age int
   }
   
   func (a *A) SayOk(){
      fmt.Println(" A SayOk",a.Name)
   }
   func (a *A) Hello(){
   	fmt.Println(" A Hello",a.Name)
    }
   type B struct {
   	A
   }
   func main(){
      var b B
      b.A.Name = "tom"
      b.A.age = 19
      b.A.SayOk()
      b.A.Hello()
   
   
   }
   ```

   

2. 匿名结构体字段访问可以简化(如果B结构体中没有变量那么就会自动去寻找嵌入的结构体中的字段或方法)。

   ```go
   var b B
      b.A.Name = "tom"       b.Name = "tom"
      b.A.age = 19    ==>    b.age = 19
      b.A.SayOk()            b.SayOk()  
      b.A.Hello()            b.Hello()
   ```

   总结

   1. 当我们直接通过b访问字段或方法时，其执行流程如下比如b.Name

   ​    2.编译器会先看b对应的类型有没有Name，如果有则直接调用B类型的            Name字段，如果没有继续查找，没有找到就报错

   ​    3.如果没有就去看嵌入中的匿名结构体有没有声明Name字段

3. 当结构体和匿名结构体有相同的字段或者方法时，编译器采用**就近访问**原则访问如希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分。

   ```
   package main
   import (
   	"fmt"
   )
   
   type A struct {
   	Name string
   	age int
   }
   
   func (a *A) SayOk(){
      fmt.Println(" A SayOk",a.Name)
   }
   func (a *A) hello(){
   	fmt.Println(" A Hello",a.Name)
    }
   type B struct {
   	A
   }
   func (b *B)SayOk() {
   	fmt.Println("B SayOK",b.Name)
   }
   func main(){
      var b B
   //    b.A.Name = "tom"
   //    b.A.age = 19
   //    b.A.SayOk()
   //    b.A.hello()
      b.Name = "jack"
      b.age  = 100
      b.SayOk()  //B SayOK jack
      b.hello() // A Hello jack
   
   
   }
   ```

   

4. 结构体嵌入两个（或多个）匿名结构体，如两个匿名结构体有相同额字段和方法（**同时结构体本身没有同名的字段和方法**），在访问时，就必须明确指定匿名结构体名字，否则编译报错

   ```go
   package main
   import (
   	"fmt"
   )
   
   type A struct {
   	Name string
   	age int
   }
   type B struct {
   	Name string
   	score float64
   }
   type C struct {
   	A
   	B
   	Name string
   }
   func main(){
      var c C
      //如果 c 无Name字段，而A和B有Name,这时就必须指定匿名结构体名字来区分
      //c.Name = "tom"//会报错，因为不知道选择A里的Name还是B里面的Name
      //c.A.Name = "tom" //具体指定哪一个结构体中的Name
      c.Name = "tom" //成功是因为C有Name
   
   }
   //这个规则对于方法也是有效的
   ```

   

5. 如果一个struct嵌套了一个有名的结构体，这种模式就是组合，如果是组合关系，那么在访问组合结构体的字段或方法时，**必须带上结构体的名字**

   ```go
   type D struct {
   	a A //有名结构体
   }
   func main(){
      var c C
      
      var d D
   //    d.Name= "jack"//报错，必须将结构体名字带上
      d.a.Name = "jack"
   
   
   ```

   

6. 嵌套匿名结构体，也可以在创建结构体变量（实例）时，直接指定各个匿名结构体字段的值

   ```go
   package main
   import (
   	"fmt"
   )
   
   type A struct {
   	Name string
   	age int
   }
   type B struct {
   	Name string
   	score float64
   }
   type C struct {
   	A
   	B
   	Name string
   }
   type D struct {
   	a A //有名结构体
   }
   
   type Goods struct{
   	Name string
   	Price float64
   }
   
   
   type Brand struct{
   	Name string
   	Address string
   }
   type TV struct {
   	Goods
   	Brand
   }
   type TV2 struct {  //也可将指针类型传入进去
   	*Goods
   	*Brand
   }
   func main(){
      var c C
      //如果 c 无Name字段，而A和B有Name,这时就必须指定匿名结构体名字来区分
      //c.Name = "tom"//会报错，因为不知道选择A里的Name还是B里面的Name
      //c.A.Name = "tom" //具体指定哪一个结构体中的Name
      c.Name = "tom" //成功是因为C有Name
      var d D
   //    d.Name= "jack"//报错，必须将结构体名字带上
      d.a.Name = "jack"
   
      tv := TV{ Goods{"电视机001",5000.99},Brand{"海尔","山东"}, }
      
      tv2 := TV{
   	Goods{
   		Price : 6000.90,
   		Name : "电视002",
   	},
   	Brand{
   		Name : "长虹",
   		Address :"北京",
   	},
      }
   //创建之时要把地址传入进去
      tv3 := TV2{ &Goods{"电视机003",7000.99},&Brand{"创维","深圳"}, }
      fmt.Println("TV=",tv)
      fmt.Println("TV2=",tv2)
      fmt.Println("TV3=",*tv3.Goods,*tv3.Brand)
      tv4 := TV2{ 
   	&Goods{
   		Name : "电视机004",
   		Price : 9900.99,
   		},
   		&Brand{
   			Name : "康佳",
   			Address : "上海",
   			}, 
   		}
   		fmt.Println("TV4=",*tv4.Goods,*tv4.Brand)
   }
   ```

   

##### &4.课堂练习

结构体的匿名字段是基本数据类型，如何访问，下面的代码

```go
type A struct {  
    Name string
    Age int
}
type Stu struct {
     A
     int
}
func  main(){
stu :=Stu{}
stu.Name = "tom"
stu.Age = 10
stu.int = 80
fmt.Println(stu)
}
//输出结果为
{{tom 10} 80}
```

说明：

1. 如果一个结构体有int类型的匿名字段，就不能有第二个。
2. 如果需要多个int类型字段，则必须给int指定名字

##### &5.多重继承

多重继承说明

如一个struct嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法。从而实现了多重继承

##### &6.多重继承细节说明

1. 如果嵌入的匿名结构体有相同的**字段名或方法名**，则在访问时，需要通过匿名结构体类型来区分
2. 为了保证代码的简洁性，尽量不要使用多重继承

## 十一、接口

### 1.接口案例

```go
package main
import (
	"fmt"
)

//声明定义一个接口
type Usb interface{
	Start()
	Stop()
}

type Phone struct {

}

type Camera struct {

}
//让phone实现usb接口的方法
func (p Phone) Start(){
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop(){
	fmt.Println("手机停止工作。。。")
}

func (c Camera) Start(){
	fmt.Println("相机开始工作。。。")
}
func (c Camera) Stop(){
	fmt.Println("相机停止工作。。。")
}
//计算机
type Computer struct{

}

//编写一个方法Working方法，接收一个Usb接口类型变量
//只要是实现了Usb接口（所谓实现Usb接口，就是指实现了Usb接口声明所有方法）
func (c Computer) Working (usb Usb){
	//通过usb接口变量来调用Start和Stop方法
     usb.Start()
	 usb.Stop()
}
func main(){

	//测试
	//先创建结构体变量
	computer :=Computer{}
	phone :=Phone{}
	camera := Camera{}

	//关键点
	computer.Working(phone)
	computer.Working(camera)


}
```

### 2.基本介绍

interface类型可以定义一组方法，但是这些不需要实现。并且interface不能包含任何变量，到某个自定义类型（比如结构体Phone）要使用的时候，再根据具体情况把这些方法写出来

### 3.基本语法

```go
type 接口名 interface{
  method1(参数列表)返回值列表
  method2(参数列表)返回值列表
}

实现接口的方法
func (t 自定义类型)method1(参数列表)返回值列表{
 //方法实现
}
func (t 自定义类型)method2(参数列表)返回值列表{
 //方法实现
}
```

小结说明：

1）接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的多态和高内聚低耦合的思想。

2）Golang中的接口，**不需要显式的实现。只要一个变量，含有接口类型中的所有方法，那么这个变量就实现这个接口**。因此,Golang中没有implement这样的关键字

解释：假如有一个不同名字的接口，但是方法和原接口的方法一样，那么实现的方法也是实现了这个接口，golang中接口的实现不是基于名字而是基于方法。可以同时实现两个接口。

```go
type Usb interface{
    Start()
	Stop()
}
type Usb2 interface{
	Start()
	Stop()
}
```

### 4.应用场景介绍

对初学者讲，理解接口不算太难，难的是不知道什么时候使用接口，下面几个例子来解释

1）现在美国要制造轰炸机，武装直升机，专家只需要把飞机需要的功能/规格定下来即可，然后让别人具体实现即可

2）就是做一个项目，在接口中定义规范让其他人去实现所定的规范

### 5.注意事项和细节

1）接口本身不能创建实例，但是**可以指向一个实现了该接口的自定义类型的变量（实例）**

```go
package main
import (
	"fmt"
)

type AInterface interface {
	Say()

}

type Stu struct {
	Name string
}
func (stu Stu)Say(){
	fmt.Println("Stu Say()")
}
func main(){
   var stu Stu  //结构体变量，实现了Say() 实现了AInterface这个接口
   var a AInterface =stu
   a.Say() //Stu Say()
}
```

2）接口中所有的方法都没有方法体，即都是没有实现的方法

3）在Golang中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口

4）一个自定义类型只有实现了某个接口，才能将该自定义类型的实例（变量）赋给接口类型

5）**只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型**(面试会问)

```go
type integer int
type AInterface interface {
	Say()
}

func (i integer)Say() {
	fmt.Println("integer Say i = ",i)
}
func main(){
   var i integer = 10
   var b AInterface = i
   b.Say() //integer Say i =  10
}
```

6）一个自定义类型可以实现多个接口

```go
package main
import (
	"fmt"
)

type AInterface interface {
	Say()

}

type Stu struct {
	Name string
}
func (stu Stu)Say(){
	fmt.Println("Stu Say()")
}

type integer int

func (i integer)Say() {
	fmt.Println("integer Say i = ",i)
}
type BInterface interface{
	Hello()
}
type Monster struct{

}
func (m Monster) Hello(){
	fmt.Println("Monster Hello()~~")
}
func (m Monster) Say(){
	fmt.Println("Monster Say()~~")
}
//此时刻Monster就实现了AInterface和BInterface

func main(){
   var stu Stu  //结构体变量，实现了Say() 实现了AInterface这个接口
   var a AInterface =stu
   a.Say() //Stu Say()
   var i integer = 10
   var b AInterface = i
   b.Say() //integer Say i =  10
//验证monster去实现两个接口
   var monster Monster
   var a2 AInterface = monster //Monster Say()~~
   var b2 BInterface = monster //Monster Hello()~~
   a2.Say()
   b2.Hello()
}
```

7）Golang接口中不能有任何变量

```go
例如，以下写法就是错误的
type AInteger interface {
Name string//这是错误的，不能这样用
test()
}

```

8）**一个接口（比如A接口）可以继承多个别的接口（比如B,C接口），这时如果要实现A接口，也必须将B，C接口的方法也全部实现**

```go
package main
import (
	"fmt"
)
type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}
type AInterface interface {
	BInterface
	CInterface
	test03()
}
//如果需要实现AIerface，就需要把BInterface和CInterface的方法都实现
type Stu struct{

}
//把所要实现的接口所有的方法都实现一下
func (stu Stu) test01(){
  fmt.Println("这是Test01")
}
func (stu Stu) test02(){
	fmt.Println("这是Test02")
}
func (stu Stu) test03(){
	fmt.Println("这是Test03")
}

func main(){
    //实践
	var stu Stu
	var a AInterface = stu
	a.test01()
	a.test02()
	a.test03()
}
```

9）interface类型默认是一个指针（引用类型），如果没有对interface初始化就使用，那么就会输出nil

10）空接口interface{}没有任何方法，**所以所有类型都实现了空接口**

```go
type T interface {

}
func main(){
    //实践
	var stu Stu
	var a AInterface = stu
	a.test01()
	a.test02()
	a.test03()
	var t T =stu
	fmt.Println(t)
	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	fmt.Println(t2,test02)
}
```

查看下列代码看看是否有错

```go
type AInterface interface{
   Test01()
   Test02()
}
type BInterface interface{
   Test01()
   Test03()
}
type CInterface interface{
   AInterface
   BInterface
}
func main(){

}
//这里编译错误，因为CInterface有两个test01().编译器不通过
```

来看第二个代码

```go
type Usb interface{
   Say()
}
type Stu struct {

}
func (this *Stu) Say() {
fmt,Println("Say()")
}

func main() {
var stu Stu = Stu{}
//var u Usb =stu //会报错,因为stu没有实现Say()方法
var u Usb = &stu //改正之后，加个地址就可以了
u.Say()
fmt.Println("here",u)
}
```



### 6.接口编程经典案例

实现对hero结构体切片的排序：sort.Sort(data Interface)

```go
package main
import (
	"fmt"
	"sort"
	"math/rand"
)

//1.声明Hero结构体
type Hero struct {
	Name string
	Age int
}

//2.声明一个Hero的切片类型
type HeroSlice []Hero 

//3.实现Interface 接口
func (hs HeroSlice) Len() int {
	return len(hs )
}
//Less方法就是决定你是用什么标准进行排序
//1.按Hero的年龄从小到大进行排序！！
func (hs HeroSlice) Less(i,j int) bool {
	return hs[i].Age > hs[j].Age
	//修改成对姓名排序
    // return hs[i].Name > hs[j].Name
}

func (hs HeroSlice) Swap(i,j int) {
	// temp := hs[i]
	// hs[i] =hs[j]
	// hs[j] = temp
	//简洁的交换:下面一句话等价于上面三句话
	hs[i],hs[j] = hs[j],hs[i]
}

//声明一个Student结构体
//1.声明Student结构体
type Student struct {
	Name string
	Age int
	Score int
}
//然后将上面那三个方法复制到下面

//将student按成绩从大到小进行排序

//声明一个Stu切片类型
type StuSlice []Student
//3.实现Interface 接口
func (stu StuSlice) Len() int {
	return len(stu)
}
//Less方法就是决定你是用什么标准进行排序
//1.按Hero的年龄从小到大进行排序！！
func (stu StuSlice) Less(i,j int) bool {
	//按成绩进行排序
	return stu[i].Score > stu[j].Score
}

func (stu StuSlice) Swap(i,j int) {
	stu[i],stu[j] = stu[j],stu[i]
}



func main(){

	//先定义一个数组/切片
	var intSlice = []int{0,-1,10,7,90}
    //要求对intSlice切片进行排序
	//1.冒泡排序...
	//2.可以使用系统提供的方法
    sort.Ints(intSlice)
	fmt.Println(intSlice)//[-1 0 7 10 90]
	//请对结构体进行排序
	//1.冒泡排序
	//2.系统提供的方法

	//测试我们是否可以对结构体进行排序
	var heroes HeroSlice
	for i :=0;i < 10; i++{
		hero :=Hero {
			Name : fmt.Sprintf("英雄~%d",rand.Intn(100)),
			Age : rand.Intn(100),
		}
		//将hero append到heros切片
		heroes = append(heroes,hero)
	}
	//看看排序前的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}
	fmt.Println( )
    //调用sort.Sort()
    sort.Sort(heroes)
    //看看排序后的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}

	//这个接口的妙处就是将这个接口的三个方法实现然后只需要将结合挂钩梯放入到
	//sort方法中去就可以
	fmt.Println()

	var studentsl StuSlice
	for i :=0;i < 10; i++{
		stus :=Student {
			Name : fmt.Sprintf("学生~%d",rand.Intn(100)),
			Age : rand.Intn(100),
			Score : rand.Intn(100),
		}
		//将hero append到heros切片
		studentsl = append(studentsl,stus)
	}
	//看看排序前的顺序
	for _, v := range studentsl {
		fmt.Println(v)
	}
	fmt.Println( )
    //调用sort.Sort()
    sort.Sort(studentsl)
    //看看排序后的顺序
	for _, v := range studentsl {
		fmt.Println(v)
	}
}
```

### 7.接口与继承之间的比较

1）大家可能对实现接口和继承比较迷茫了。那么问题来了，那么他们究竟有什么区别呢？

例如有一个猴子会爬树，然而他想学鸟儿飞翔，鱼儿游泳

```go
package main
import (
	"fmt"
)
//,Monkey结构体
type Monkey struct {
	Name string
}

//声明接口
type BirdAble interface {
	Flying()
}

//声明接口
type FishAble interface {
	Swimming()
}
func (this *Monkey) climbing() {
	fmt.Println(this.Name,"生来会爬树")
}

//创建LittleMonkey结构体
type LittleMonkey struct {
	Monkey //继承
}

//让littleMonkey实现BirdAble的Flying()方法
func (this *LittleMonkey) Flying(){
	fmt.Println(this.Name,"通过学习会飞翔")
}
//littleMonkey实现FishAble的Swimming()方法
func (this *LittleMonkey) Swimming(){
	fmt.Println(this.Name,"通过学习会游泳")
}
func main(){
    
	//创建LittleMonkey实例
	monkey := LittleMonkey{
		Monkey {
			Name : "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
  
}
```

对上面代码的小结

1. 当A结构体继承了另外B结构体，那么A结构体就自动继承了B结构体的字段和方法，并且可以直接使用
2. 当A结构体需要扩展功能，同时不希望去破坏继承关系，则可以去实现某个接口即可，因此我们可以认为：实现接口是对继承机制的补充。

**实现接口可以看作是对继承的一种补充**

![](D:\myfile\后端\go语言学习\pic\面向对象\pic14.jpg)

- 接口和继承解决的问题是不同的

**继承的价值**主要在于：解决代码的**复用性**和**可维护性**

**接口的价值**主要在于：设计。设计好各种规范（方法），让其他自定义类型去实现这些方法

- 接口比继承更加灵活 Person Student   BirdAble LittleMonkey

- 接口比继承更加灵活。继承是满足 is --a的关系。则接口只需满足like --a 的关系

- 接口在一定程度上实现**代码解耦**

### 8.面向对象编程--多态

#### 1）基本介绍

变量（实例）具有多种形态。面向对象的第三大特征，在Go语言，多态特征是通过接口实现的。可以按照统一的接口来调用不同的实现。这时接口变量就呈现不同的形态

#### 2）快速入门

在前面的Usb接口案例：Usb usb 既可以接收手机变量，又可以接收相机变量，就体现了Usb接口多态特性。

```go
//编写一个方法Working方法，接收一个Usb接口类型变量
//只要是实现了Usb接口（所谓实现Usb接口，就是指实现了Usb接口声明所有方法）
func (c Computer) Working (usb Usb){//usb变量会根据传入的实参，来判断到底是Phone还是Camera
	//通过usb接口变量来调用Start和Stop方法
     usb.Start()
	 usb.Stop()
}
func main(){

	//测试
	//先创建结构体变量
	computer :=Computer{}
	phone :=Phone{}
	camera := Camera{}

	//关键点
	computer.Working(phone)//传入phon就是phon执行这个方法
	computer.Working(camera)

}
```

#### 3）接口体现多态的两种形式

 多态参数

在前面的Usb接口案例，Usb usb,既可以接收手机变量，又可以接收相机变量，就体现了Usb的接口多态

多态数组

````
演示一个案例给Usb数组，存放phone结构体和Camera结构体变量

```go
package main
import (
	"fmt"
)

//声明定义一个接口
type Usb interface{
	Start()
	Stop()
}

type Phone struct {
    name string
}

type Camera struct {
	name string
}
//让phone实现usb接口的方法
func (p Phone) Start(){
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop(){
	fmt.Println("手机停止工作。。。")
}

func (c Camera) Start(){
	fmt.Println("相机开始工作。。。")
}
func (c Camera) Stop(){
	fmt.Println("相机停止工作。。。")
}
//计算机
type Computer struct{

}

//编写一个方法Working方法，接收一个Usb接口类型变量
//只要是实现了Usb接口（所谓实现Usb接口，就是指实现了Usb接口声明所有方法）
func (c Computer) Working (usb Usb){
	//通过usb接口变量来调用Start和Stop方法
     usb.Start()
	 usb.Stop()
}
func main(){
   //定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
   //这里就体现出多态数组
   var usbArr  [3]Usb
   usbArr[0] = Phone{"iphone"}
   usbArr[1] = Phone{"小米"}
   usbArr[2] = Camera{"佳能"}
   fmt.Println(usbArr) //[{iphone} {小米} {佳能}]
	
}
```




````



### 9.类型断言

#### 1）先看一个需求

代码

```go
type Point struct{
     x int
     y int
}
func main() {
    var a interface{}
    var point Point = Pooint{1,2}
    a = point //OK
    //如何将A赋给一个Point变量？
    var b Point
    b = a /可以吗？ ==》erro
    fmt.Println(b)
}
//改进之后
package main
import (
	"fmt"
)
type Point struct{
	x int
	y int
}
func main() {
   var a interface{}
   var point Point = Point{1,2}
   a = point //OK
   //如何将A赋给一个Point变量？
   var b Point
   //b = a //可以吗？ ==》erro
   b = a.(Point) //解决办法，类型断言
    //b = a.(Point)就是类型断言，表示判断a是否指向Poin类型的变量，如果是就转成了Point类型并赋给b变量，否则报错
   fmt.Println(b) //1 2
}
```

有一个具体的需要，引出了类型断言

#### 2）基本介绍

类型断，由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言，具体的如下

```go
//float32可以是其他类型，比如Point
var t float32
var x interface{}
x = t //OK
y := x.(float32)//转成float32



//float32可以死其他类型，比如Point
var t float32 
var x interface{}
x = t
//转成float,待检查的
y,ok :=a.(float32)
if ok == true {
fmt.Println("convert success")
}else{
fmt.Println("convert fall")
}


```

```go
/类型断言的其他案例
   var x interface{}
   var b2 float32 = 1.1
   x = b2 //空接口可以接收任意类型
   //x = >float32 【使用类型断言】
   y := x.(float32)
   fmt.Printf("y的类型是%T 值是%v",y,y) //y的类型是float32 值是1.1
```

对上面的代码的说明

在进行类型断言时，如果类型不匹配，就会报panic,因此进行类型断言时，要确保原来的空接口指向的就是断言的类型

如何在进行断言时带上检测机制，如果成功就ok否则也不要报panic

```go
//类型断言(带检测)
var x interface{}
var b2 float32 = 1.1
x = b2 //空接口可以接收任意类型
//x = >float32 【使用类型断言】
//y,ok := x.(float64)
//简写

if  y,ok := x.(float64);ok {
	fmt.Println("转换成功了")
    fmt.Printf("y的类型是%T 值是%v",y,y) //y的类型是float32 值是1.1
}else{
	fmt.Println("转换失败")
}
fmt.Println("检测失败了无妨，代码不要停")

}
```

#### 3)类型断言的最佳实践1

在前面的Usb接口案例做改进

给Phone结构体增加一个特有的方法call(),当Usb接口接收的是Phone变量时，还需要调用call方法。

```go
package main
import (
	"fmt"
)

//声明定义一个接口
type Usb interface{
	Start()
	Stop()
}

type Phone struct {
    name string
}

type Camera struct {
	name string
}
//让phone实现usb接口的方法
func (p Phone) Start(){
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop(){
	fmt.Println("手机停止工作。。。")
}

func (p Phone) Call(){
	fmt.Println("手机开始打电话。。。")
}


func (c Camera) Start(){
	fmt.Println("相机开始工作。。。")
}
func (c Camera) Stop(){
	fmt.Println("相机停止工作。。。")
}
//计算机
type Computer struct{

}

//编写一个方法Working方法，接收一个Usb接口类型变量
//只要是实现了Usb接口（所谓实现Usb接口，就是指实现了Usb接口声明所有方法）
func (c Computer) Working (usb Usb){
	//通过usb接口变量来调用Start和Stop方法
     usb.Start()
	 //如果usb是指向Phone结构体变量，则还需要调用Call()方法
	 //类型断言
	 if phone, ok := usb.(Phone); ok {
		phone.Call()
	 }//否则断言失败还是继续执行下列方法
	 usb.Stop()
}

func main(){
   //定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
   //这里就体现出多态数组
   var usbArr  [3]Usb
   usbArr[0] = Phone{"iphone"}
   usbArr[1] = Phone{"小米"}
   usbArr[2] = Camera{"佳能"}
   //fmt.Println(usbArr) //[{iphone} {小米} {佳能}]
   //遍历
   var computer Computer
   for _,v := range usbArr{
          computer.Working(v)
		  fmt.Println()
   }
	
}
执行结果如下：
手机开始工作。。。
手机开始打电话。。。
手机停止工作。。。

手机开始工作。。。
手机开始打电话。。。
手机停止工作。。。

相机开始工作。。。
相机停止工作。。。

```

#### 4）类型断言实践2

写一个函数，循环判断传入参数的类型

```go
package main
import (
	"fmt"
)
//编写一个函数，可以判断输入的参数是什么类型
func TypeJudge(items... interface{}){
	for index, x := range items {
        switch x.(type) {
			case bool : 
				fmt.Printf("第%v个参数是 bool 类型,值是%v\n",index+1,x)
			case float32 : 
				fmt.Printf("第%v个参数是 float32 类型,值是%v\n",index+1,x)
			case float64 : 
				fmt.Printf("第%v个参数是 float64 类型,值是%v\n",index+1,x)	
			case int, int32, int64 : 
				fmt.Printf("第%v个参数是 整数 类型,值是%v\n",index+1,x)		  	  
			case string : 
				fmt.Printf("第%v个参数是 string 类型,值是%v\n",index+1,x)
			default :
			    fmt.Printf("第%v个参数是 类型 不确定,值是%v\n",index+1,x)	
		 }
	}
	
}
func main(){
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name = "tom"
	address := "北京"
	n4 :=300
    
	TypeJudge(n1,n2,n3,name,address,n4)

}
//运行结果如下
第1个参数是 float32 类型,值是1.1
第2个参数是 float64 类型,值是2.3
第3个参数是 整数 类型,值是30
第4个参数是 string 类型,值是tom
第5个参数是 string 类型,值是北京
第6个参数是 整数 类型,值是300
```

5）类型断言的最佳实践3

在前面代码的基础上，增加判断Student类型和*Student类型

```go
package main
import (
	"fmt"
)

//定义Student类型
type Student struct {

}
//编写一个函数，可以判断输入的参数是什么类型
func TypeJudge(items... interface{}){
	for index, x := range items {
        switch x.(type) {
			case bool : 
				fmt.Printf("第%v个参数是 bool 类型,值是%v\n",index+1,x)
			case float32 : 
				fmt.Printf("第%v个参数是 float32 类型,值是%v\n",index+1,x)
			case float64 : 
				fmt.Printf("第%v个参数是 float64 类型,值是%v\n",index+1,x)	
			case int, int32, int64 : 
				fmt.Printf("第%v个参数是 整数 类型,值是%v\n",index+1,x)		  	  
			case Student : 
				fmt.Printf("第%v个参数是 Student 类型,值是%v\n",index+1,x)
			case *Student : 
				fmt.Printf("第%v个参数是 *Student 类型,值是%v\n",index+1,x)
			case string : 
				fmt.Printf("第%v个参数是 string 类型,值是%v\n",index+1,x)
						
			default :
			    fmt.Printf("第%v个参数是 类型 不确定,值是%v\n",index+1,x)	
		 }
	}
	
}


func main(){
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name = "tom"
	address := "北京"
	n4 :=300

	stu1 := Student{}
	stu2 := &Student{}
    
	TypeJudge(n1,n2,n3,name,address,n4,stu1,stu2)

}
```



## 十二、项目

### 1.项目开发流程图

![](D:\myfile\后端\go语言学习\pic\面向对象\pic15.jpg)

### 2.家庭收支记账软件项目

1）需求说明

- 模拟实现基于文本界面的《家庭记账软件》

- 该软件能够记录家庭的收入、支出，并能够打印收支明细表

- 项目采用分级菜单的方式，主菜单如下：

  ```
  --------家庭收支记账软件-------
           1.收支明细
           2.登记收入
           3.登记支出
           4.退出
           
           请选择(1-4):
           
  ```

#### 2）项目代码实现

实现基本功能（先使用面向过程，后面改成面向对象）

编写文件TestMyAccount.go 完成基本功能

1. 功能1：先完成可以显示主菜单，并且可以退出
2. 功能2：完成可以显示明细和登记收入的功能
3. 功能3：完成了登记支出的功能

#### 3）具体功能实现

功能1：先完成可以显示主菜单，并且可以退出

思路分析：给出的界面完成，主菜单的显示，当用户输入4的时候就退出

```go
package main
import (
	"fmt"
)

func main(){

	//声明一个变量，保存接收用户输入的选项
	key := ""
	//声明一个变量，控制是否退出for循环
	loop := true

	//显示这个主菜单
	for {
		fmt.Println("--------家庭收支记账软件---------")
		fmt.Println("         1.收支明细")
		fmt.Println("         2.登记收入")
		fmt.Println("         3.登记支出")
		fmt.Println("         4.退出软件")
		fmt.Print("请选择(1-4)")

		fmt.Scanln(&key)

		switch key {
		case "1" :
			fmt.Println("1.收支明细")
		case "2" :
			fmt.Println("2.登记收入")
	    case "3" :
			fmt.Println("3.登记支出")
		case "4" :
			loop = false	
		default :
		    fmt.Println("请输入正确的选项")			
		}

		if !loop {
			break
		}
	}
	fmt.Println("你退出了家庭记账软件的使用")
}
```

功能2：完成可以显示明细和登记收入的功能

思路分析:

1.因为需要显示明细，我们定义一个变量details string来记录

2.还需要定义变量来记录余额(balance)，每次支出的收支的金额(money)，以及收支说明(note)

走代码

```go
    //声明一个变量统计余额
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
    note := ""
	//收支的详情
	//当有收支发生的时候，就对details进行拼接处理
	details := "收支\t账户余额\t收支金额\t说明"

在case的操作
case "2" :
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money //修改账户余额
			fmt.Println("本次收入的说明：")
			fmt.Scanln(&note)
			//将这个收入情况，拼接到details变量当中
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v",balance,money,note)
```

功能3完成登记支出的功能

思路分析：登记支出的功能和登记收入的功能类似做一些修改即可

```go
case "3" :
			fmt.Println("本次支出的金额：")
			fmt.Scanln(&money)
			//这里需要做出一个必要的判断
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			balance -=money
			fmt.Println("本次的支出说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v",balance,money,note)
```

项目改进

1.用户输入4时，给出提示"你确定要退出吗？y/n",必须输入正确的y/n，否则循环输入指令，直到输入y或者n

```go
case "4" :
			fmt.Println("您确定要退出吗？ y/n")
			choice :=" "
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n"{ //输了y/n就break出去
					break
				}
				fmt.Println("您的输入有误请重新输入 y/n")
			}
			if choice == "y" {
				loop = false	
			}
			
			
```

2.当没有任何收支明细时，提示“当前没有收支明细。。。来一笔把!”

```go
case "1" :
			fmt.Println("------------当前收支明细记录--------")
			if flag {
				fmt.Println(details)
			}else{
				fmt.Println("您当前没有支出记录，来一笔吧！")
			}
```

3.在支出时，判断余额是否够，并给出相应的提示

```go
case "3" :
			fmt.Println("本次支出的金额：")
			fmt.Scanln(&money)
			//这里需要做出一个必要的判断
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			balance -=money
			fmt.Println("本次的支出说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v",balance,money,note)
			flag = true
```

面向过程的家庭记账收支软件全部代码

```go
package main
import (
	"fmt"
)

func main(){

	//声明一个变量，保存接收用户输入的选项
	key := ""
	//声明一个变量，控制是否退出for循环
	loop := true
	//声明一个变量统计余额
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
    note := ""
	//定义一个变量记录是否有收支的行为
	flag := false
	//收支的详情
	//当有收支发生的时候，就对details进行拼接处理
	details := "收支\t账户余额\t收支金额\t说明"
	


	//显示这个主菜单
	for {
		fmt.Println("\n--------家庭收支记账软件---------")
		fmt.Println("         1.收支明细")
		fmt.Println("         2.登记收入")
		fmt.Println("         3.登记支出")
		fmt.Println("         4.退出软件")
		fmt.Print("请选择(1-4)")

		fmt.Scanln(&key)

		switch key {
		case "1" :
			fmt.Println("------------当前收支明细记录--------")
			if flag {
				fmt.Println(details)
			}else{
				fmt.Println("您当前没有支出记录，来一笔吧！")
			}
		case "2" :
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money //修改账户余额
			fmt.Println("本次收入的说明：")
			fmt.Scanln(&note)
			//将这个收入情况，拼接到details变量当中
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v",balance,money,note)
	        flag = true
		case "3" :
			fmt.Println("本次支出的金额：")
			fmt.Scanln(&money)
			//这里需要做出一个必要的判断
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			balance -=money
			fmt.Println("本次的支出说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v",balance,money,note)
			flag = true
		case "4" :
			fmt.Println("您确定要退出吗？ y/n")
			choice :=" "
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n"{ //输了y/n就break出去
					break
				}
				fmt.Println("您的输入有误请重新输入 y/n")
			}
			if choice == "y" {
				loop = false	
			}
			
		default :
		    fmt.Println("请输入正确的选项")	

		}

		if !loop {
			break
		}
	}
	fmt.Println("你退出了家庭记账软件的使用")
}
```



4.将面向过程的代码改为面向对象的方法编写myFamilyAccount.go,并使用testMyFamilyAccount.go去完成测试。

思路分析

把记账软件的功能封装到一个结构体中，然后调用该结构体的方法来实现记账，显示明细就可以了，结构体的名字为FamilyAccount

再通过main方法中创建一个结构体FamilyAccount实例，实现记账即可

代码实现，代码不需要重新写，只需要引用上侧代码

```go
package objectTestAcc
import (
	"fmt"
)

type FamilyAccount struct {
	//声明必须字段
	
	//声明一个字段，保存接收用户输入的选项
	key string
	//声明一个字段，控制是否退出for循环
	loop bool
	//声明一个字段统计余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
    note string
	//定义一个字段记录是否有收支的行为
	flag bool
	//收支的详情
	//当有收支发生的时候，就对details进行拼接处理
	details string
}
//编写一个构造方法返回一个FamilyAccount实例 
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key : "",
		loop : true,
		balance : 10000.0,
		money : 0.0,
		note : "",
		flag : false,
		details :  "收支\t账户余额\t收支金额\t说明",
	}
}

//将显示明细写成一个方法
func (this *FamilyAccount) ShowDetails(){
    fmt.Println("------------当前收支明细记录--------")
	if this.flag {
		fmt.Println(this.details)
	}else{
		fmt.Println("您当前没有支出记录，来一笔吧！")
	}
}

//将登记收入写成一个方法和*FamilyAccount绑定
func (this *FamilyAccount) Income(){
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money //修改账户余额
	fmt.Println("本次收入的说明：")
	fmt.Scanln(&this.note)
	//将这个收入情况，拼接到details变量当中
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v",this.balance,this.money,this.note)
	this.flag = true
}
//将支出也绑定到一个方法当中
func (this *FamilyAccount) Pay(){
	fmt.Println("本次支出的金额：")
	fmt.Scanln(&this.money)
	//这里需要做出一个必要的判断
	if this.money > this.balance {
		fmt.Println("余额不足")
	}
	this.balance -=this.money
	fmt.Println("本次的支出说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v",this.balance,this.money,this.note)
	this.flag = true
}

//将退出系统写成一个方法
func (this *FamilyAccount) exit(){
	fmt.Println("您确定要退出吗？ y/n")
	choice :=" "
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n"{ //输了y/n就break出去
			break
		}
		fmt.Println("您的输入有误请重新输入 y/n")
	}
	if choice == "y" {
		this.loop = false	
	}
}

//为该结构体绑定相应的方法
//显示主菜单
func (this *FamilyAccount) MainMenu(){
	for {
		fmt.Println("\n--------家庭收支记账软件---------")
		fmt.Println("         1.收支明细")
		fmt.Println("         2.登记收入")
		fmt.Println("         3.登记支出")
		fmt.Println("         4.退出软件")
		fmt.Print("请选择(1-4)")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1" :
			this.ShowDetails()
		case "2" :
			this.Income()
		case "3" :
			this.Pay()
		case "4" :
			this.exit()	
		default :
		    fmt.Println("请输入正确的选项")	
		}

		if !this.loop {
			break
		}
	}
}
建立一个main方法
package main
import (
	"fmt"
	"go_code/project/objectTestAcc"
)

func main() {
	fmt.Println("这个是面向对象的方式完成")
	objectTestAcc.NewFamilyAccount().MainMenu()

}



```

### 3.客户信息管理系统

#### 1）项目需求说明

模拟实现基于文本界面的《客户信息管理软件》

该软件能够实现对客户对象的插入、修改和删除（用切片实现），并能够打印客户明细表 多个对象协同工作

#### 2）界面设计

![](D:\myfile\后端\go语言学习\pic\面向对象\pic17.jpg)

添加客户界面

![](D:\myfile\后端\go语言学习\pic\面向对象\pic18.jpg)

修改客户界面

![](D:\myfile\后端\go语言学习\pic\面向对象\pic19.jpg)

删除客户界面

![](D:\myfile\后端\go语言学习\pic\面向对象\pic20.jpg)

客户列表的界面

![](D:\myfile\后端\go语言学习\pic\面向对象\pic21.jpg)

#### 3）项目框架图

![](D:\myfile\后端\go语言学习\pic\面向对象\pic22.jpg)

#### 4）流程

功能说明

当用户运行程序，可以看到主菜单，当输入5时，可以退出该软件

思路分析

编写customerView.go另外可以把customer.go和customerDervice.go协商

代码实现

customerManager/model/customer.go

```go
package model
// import (
// 	"fmt"
// )
//声明一个customer结构体，表示一个客户信息
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//编写一个工厂模式，返回一个Customer的实例

func NewCustomer(id int,name string, gender string,
	age int，phone string,email string) Customer {
		return Customer{
			Id : id,
			Name : name,
			Gender : gender,
			Age : age,
			Phone : phone,
			Email : email,
		}
	}


```

customerManagerservice/customerService.go

```go
package service
import (
	"go_code/project/customerManager/model"
)

//该CustomerService ,完成对Customer的操作，包括增删改查
type CustomerService struct {
	customers []model.Customer
    //声明一个字段，表示当前切片含有多少客户
	//该字段后面，还可以作为新客户的id+1
	customerNum int
}
```

customerManager/view/customerView.go

```go
package main
import (
	"fmt"
)

type customerView struct {
     //定义必要字段
	 key string //接收用户输入
	 loop bool //是否循环显示菜单

}

//显示主菜单
func (this *customerView) mainView() {
	for{
		fmt.Println("--------客户信息管理系统------------")
		fmt.Println("         1.添加客户   ")
		fmt.Println("         2.修改客户   ")
		fmt.Println("         3.删除客户   ")
		fmt.Println("         4.客户列表   ")
		fmt.Println("         5.退出   ")
		fmt.Println("请选择(1-5): ")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			fmt.Println("添加客户")
		case "2":
		    fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			fmt.Println("客户列表")
		case "5":
			this.loop = false
		default :
		    fmt.Println("你的输入有误，请重新输入...")						
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统的使用")
}
func main() {
	//在主函数中，创建一个customerView并运行显示主菜单...
	customerView := customerView{
		key : "",
		loop : true,	
	}
	//显示主菜单
	customerView.mainView()
}

```

#### 5）完成显示客户列表的功能

思路分析

![](D:\myfile\后端\go语言学习\pic\面向对象\批次3.jpg)

代码实现

customerManager/model/customer.go

```go
package model
import (
	"fmt"
)
//声明一个customer结构体，表示一个客户信息
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//编写一个工厂模式，返回一个Customer的实例

func NewCustomer(id int,name string, gender string,
	age int,phone string,email string) Customer {
	return Customer{
			Id : id,
			Name : name,
			Gender : gender,
			Age : age,
			Phone : phone,
			Email : email,
	}
}
//增加了这个方法
//返回用户的信息,格式化的字符串
func (this Customer) GetInfo() string{
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",this.Id,this.Name,
	this.Gender,this.Age,this.Phone,this.Email)
	return info
} 


```

customerManagerservice/customerService.go

```go
package service
import (
	"go_code/project/customerManager/model"
)

//该CustomerService ,完成对Customer的操作，包括增删改查
type CustomerService struct {
	customers []model.Customer
    //声明一个字段，表示当前切片含有多少客户
	//该字段后面，还可以作为新客户的id+1
	customerNum int
}

//编写一个方法，可以返回一个*customerService实例
func NewCustomerService() *CustomerService {
	//为了可以看到客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1,"张三","男",20,"112","zs@sohu.com")
	customerService.customers = append(customerService.customers ,customer)
	return customerService
}

//返回客户切片
func (this *CustomerService) List()[]model.Customer{
    return this.customers
}


```

customerManager/view/customerView.go

```go
package main
import (
	"fmt"
	"go_code/project/customerManager/service"
)

type customerView struct {
     //定义必要字段
	 key string //接收用户输入
	 loop bool //是否循环显示菜单
	 //增加一个字段customerService
     customerService   *service.CustomerService
}

//显示所有的客户信息
func (this *customerView) list(){
     //首先获取到当前所有的客户信息（在切片中）
	 customers := this.customerService.List()
	 //显示
	 fmt.Println("----------客户列表--------------")
	 fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	 for i :=0;i<len(customers);i++ {
		fmt.Println(customers[i].GetInfo())
	 }
	 fmt.Printf("\n--------客户列表完成------------\n\n")
}

//显示主菜单
func (this *customerView) mainView() {
	for{
		fmt.Println("--------客户信息管理系统------------")
		fmt.Println("         1.添加客户   ")
		fmt.Println("         2.修改客户   ")
		fmt.Println("         3.删除客户   ")
		fmt.Println("         4.客户列表   ")
		fmt.Println("         5.退出   ")
		fmt.Println("请选择(1-5): ")


		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			fmt.Println("添加客户")
		case "2":
		    fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			this.list()
		case "5":
			this.loop = false
		default :
		    fmt.Println("你的输入有误，请重新输入...")						
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统的使用")
}
func main() {
	//在主函数中，创建一个customerView并运行显示主菜单...
	customerView := customerView{
		key : "",
		loop : true,	
	}
	//完成对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainView()
}

```

#### 6）添加客户功能

功能说明

![](D:\myfile\后端\go语言学习\pic\面向对象\pic24.jpg)

思路分析

![](D:\myfile\后端\go语言学习\pic\面向对象\pic25.jpg)

代码实现

需要编写CustomerView和customerService，Customer类

规定，新添加的学院的id就是他是第几个加入的

customerManager/model/customer.go

```go
//编写一个工厂模式，返回二种Customer的实例方法，不带id
func NewCustomer2(name string, gender string,
	age int,phone string,email string) Customer {
	return Customer{
			Name : name,
			Gender : gender,
			Age : age,
			Phone : phone,
			Email : email,
	}
}
```

customerManagerservice/customerService.go

```go
增加一个方法
//添加客户到customer切片中
func (this *CustomerService) Add(customer model.Customer) bool{

	//我们确定一个分配id的规则，就是添加的顺序
	this.customerNum ++
	customer.Id = this.customerNum
	this.customers = append(this.customers,customer)
    return true
}


```

customerManager/view/customerView.go

```go
编写一个add方法调用servic蹭的Add()
//得到用户的输入，信息构建新的客户，并完成添加
func (this *customerView) add() {
	fmt.Println("------------添加客户------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话号码:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的Customer实例
	//注意：id号没有让用户输入，id号是唯一的，让系统分配即可
	customer := model.NewCustomer2(name,gender,age,phone,email)
	//调用
	if this.customerService.Add(customer) {
		fmt.Println("------------添加完成------------")
	}else{
		fmt.Println("------------添加失败------------")
	}
}
下面的switch方法也要改一下
case "1":
			this.add()
```

#### 7）删除客户功能

功能说明

![](D:\myfile\后端\go语言学习\pic\面向对象\pic26.jpg)

思路分析

需要编写CustomerView和CustomerService

![](D:\myfile\后端\go语言学习\pic\面向对象\pic27.jpg)

代码实现

customerManager/model/customer.go:无变化

customerManagerservice/customerService.go

```go

增加了这两个方法，一个删除一个查找id
//根据id删除客户（从切片中删除）
func (this *CustomerService) Delete(id int )bool {
	index :=this.FindById(id)
	//如果index ==-1说明没有这个客户
	if index== -1 {
		return false
	}
		//如何从切片中删除一个元素
	this.customers = append(this.customers[:index],this.customers[index+1:]...)
	return true
	
}

//根据Id查找客户在在切片对应中的下标,返回-1
func (this *CustomerService) FindById(id int) int {
	//默认为-1
	index := -1
	//遍历this.customers切片
	for i :=0;i < len(this.customers);i++ {
		if this.customers[i].Id ==id {
			//找到了
			index = i
		}
	}
	return index
}

```

customerManager/view/customerView.go

```go
增加这个方法
//得到用户输入的id删除该id对应的客户
func (this *customerView) delete() {
	fmt.Println("------------删除客户------------")
	fmt.Println("请选择待删除的客户编号(-1退出)：")
	id :=-1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}

	fmt.Println("确认是否删除(Y/N): ")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
			//调用service中的delete方法
		if this.customerService.Delete(id) {
			fmt.Println("------------删除成功------------")
	}else{
		fmt.Println("------------删除失败，输入的id号不存在------------")
	}
  }
}

```

8）完善退出确认功能

功能说明：

要求用户在退出时提示“是否退出(Y/N),用户必须输入y/n否则循环提示

思路分析：需编写CustomerView

代码实现

在customerManager/view/customerView.go增加这个方法

```go
//退出软件
func (this *customerView) exit(){
	fmt.Println("确定是否退出(Y/N): ")
	for {
	fmt.Scanln(&this.key)
	if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n"{
		break
	  }
	  fmt.Println("您的输入有误，请重新输入(Y/N) : ")
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
}
然后在switch中修改一下
case "5":
			this.exit()
```

#### 8）修改客户的功能

功能说明：根据id进行对客户的修改操作

思路：依旧在customerService和customerView中进行编写操作

代码实现

customerManagerservice/customerService.go

```go
//根据id进行修改客户信息的操作
func (this *CustomerService) Update(customer model.Customer) bool {
	index :=this.FindById(customer.Id)
	//如果index ==-1说明没有这个客户
	if index== -1 {
		return false
	}
	//将customer插入到指定的位置并对customers进行更新操作,就将原来位置的customer用一个新的customer进行替换操作
	this.customers = append(append(this.customers[:index],customer),this.customers[index+1:]...)
    return true
}

//根据Id查找客户在在切片对应中的下标,返回-1
func (this *CustomerService) FindById(id int) int {
	//默认为-1
	index := -1
	//遍历this.customers切片
	for i :=0;i < len(this.customers);i++ {
		if this.customers[i].Id ==id {
			//找到了
			index = i
		}
	}
	return index
}
```

customerManager/view/customerView.go

```go
//修改客户的操作
func (this *customerView) update() {
	fmt.Println("------------修改客户------------")
    fmt.Println("请选择修改客户的编号(-1的话就退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话号码:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的Customer实例
	//注意：id号没有让用户输入，id号是唯一的，让系统分配即可
	customer := model.NewCustomer(id,name,gender,age,phone,email)
	//调用
	if this.customerService.Update(customer) {
		fmt.Println("------------修改成功------------")
	}else{
		fmt.Println("------------修改失败------------")
	}
}
```

再外加一个简单的登录操作使得项目更加完善

在customerManager/view/customerView.go中进行编写

```go
//简单登录功能的时间
func (this *customerView) Login (){
	account :=""
	pwd :=""
	for {
    fmt.Println("请输入账号： ")
	fmt.Scanln(&account)
	fmt.Println("请输入密码")
	fmt.Scanln(&pwd)
	if account == "7758258" && pwd =="111"{
		fmt.Println("恭喜你！正在进入系统！")
		break
	   }
	   fmt.Println("您的输入的账号或者密码有误，请重新输入: ")	   	
	}
	this.mainView()
}
func main() {
	//在主函数中，创建一个customerView并运行显示主菜单...
	customerView := customerView{
		key : "",
		loop : true,	
	}
	//完成对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.Login()
}
```



#### 9）完整代码的展示如下

customerManager/model/customer.go

```go
package model
import (
	"fmt"
)
//声明一个customer结构体，表示一个客户信息
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//编写一个工厂模式，返回一个Customer的实例
func NewCustomer(id int,name string, gender string,
	age int,phone string,email string) Customer {
	return Customer{
			Id : id,
			Name : name,
			Gender : gender,
			Age : age,
			Phone : phone,
			Email : email,
	}
}

//编写一个工厂模式，返回二种Customer的实例方法，不带id
func NewCustomer2(name string, gender string,
	age int,phone string,email string) Customer {
	return Customer{
			Name : name,
			Gender : gender,
			Age : age,
			Phone : phone,
			Email : email,
	}
}

//返回用户的信息,格式化的字符串
func (this Customer) GetInfo() string{
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",this.Id,this.Name,
	this.Gender,this.Age,this.Phone,this.Email)
	return info
} 


```

customerManagerservice/customerService.go

```go
package service
import (
	"go_code/project/customerManager/model"
)

//该CustomerService ,完成对Customer的操作，包括增删改查
type CustomerService struct {
	customers []model.Customer
    //声明一个字段，表示当前切片含有多少客户
	//该字段后面，还可以作为新客户的id+1
	customerNum int
}

//编写一个方法，可以返回一个*customerService实例
func NewCustomerService() *CustomerService {
	//为了可以看到客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1,"张三","男",20,"112","zs@sohu.com")
	customerService.customers = append(customerService.customers ,customer)
	return customerService
}

//返回客户切片
//一定要使用指针的方式
func (this *CustomerService) List()[]model.Customer{
    return this.customers
}

//添加客户到customer切片中
//必须要用指针的方式，保证一直用的都是一个CustomerService
func (this *CustomerService) Add(customer model.Customer) bool{

	//我们确定一个分配id的规则，就是添加的顺序
	this.customerNum ++
	customer.Id = this.customerNum
	this.customers = append(this.customers,customer)
    return true
}

//根据id删除客户（从切片中删除）
func (this *CustomerService) Delete(id int )bool {
	index :=this.FindById(id)
	//如果index ==-1说明没有这个客户
	if index== -1 {
		return false
	}
		//如何从切片中删除一个元素
	this.customers = append(this.customers[:index],this.customers[index+1:]...)
	return true
	
}

//根据id进行修改客户信息的操作
func (this *CustomerService) Update(customer model.Customer) bool {
	index :=this.FindById(customer.Id)
	//如果index ==-1说明没有这个客户
	if index== -1 {
		return false
	}
	//将customer插入到指定的位置并对customers进行更新操作,就将原来位置的customer用一个新的customer进行替换操作
	this.customers = append(append(this.customers[:index],customer),this.customers[index+1:]...)
    return true
}

//根据Id查找客户在在切片对应中的下标,返回-1
func (this *CustomerService) FindById(id int) int {
	//默认为-1
	index := -1
	//遍历this.customers切片
	for i :=0;i < len(this.customers);i++ {
		if this.customers[i].Id ==id {
			//找到了
			index = i
		}
	}
	return index
}



```

customerManager/view/customerView.go

```go
package main
import (
	"fmt"
	"go_code/project/customerManager/service"
	"go_code/project/customerManager/model"
)

type customerView struct {
     //定义必要字段
	 key string //接收用户输入
	 loop bool //是否循环显示菜单
	 //增加一个字段customerService
     customerService   *service.CustomerService
}

//显示所有的客户信息
func (this *customerView) list(){
     //首先获取到当前所有的客户信息（在切片中）
	 customers := this.customerService.List()
	 //显示
	 fmt.Println("----------客户列表--------------")
	 fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	 for i :=0;i<len(customers);i++ {
		fmt.Println(customers[i].GetInfo())
	 }
	 fmt.Printf("\n--------客户列表完成------------\n\n")
}

//得到用户的输入，信息构建新的客户，并完成添加
func (this *customerView) add() {
	fmt.Println("------------添加客户------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话号码:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的Customer实例
	//注意：id号没有让用户输入，id号是唯一的，让系统分配即可
	customer := model.NewCustomer2(name,gender,age,phone,email)
	//调用
	if this.customerService.Add(customer) {
		fmt.Println("------------添加完成------------")
	}else{
		fmt.Println("------------添加失败------------")
	}
}

//修改客户的操作
func (this *customerView) update() {
	fmt.Println("------------修改客户------------")
    fmt.Println("请选择修改客户的编号(-1的话就退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话号码:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的Customer实例
	//注意：id号没有让用户输入，id号是唯一的，让系统分配即可
	customer := model.NewCustomer(id,name,gender,age,phone,email)
	//调用
	if this.customerService.Update(customer) {
		fmt.Println("------------修改成功------------")
	}else{
		fmt.Println("------------修改失败------------")
	}
}

//得到用户输入的id删除该id对应的客户
func (this *customerView) delete() {
	fmt.Println("------------删除客户------------")
	fmt.Println("请选择待删除的客户编号(-1退出)：")
	id :=-1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}

	fmt.Println("确认是否删除(Y/N): ")
	choice := ""
	for {
	fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" || choice =="n" || choice =="N"{
			break
		}
		fmt.Println("您的输入有误请重新输入(Y/N): ")
	}
	if choice == "y" || choice == "Y" {
			//调用service中的delete方法
		if this.customerService.Delete(id) {
			fmt.Println("------------删除成功------------")
	}else{
		fmt.Println("------------删除失败，输入的id号不存在------------")
	}
  } else{
	this.mainView()
  }
}

//退出软件
func (this *customerView) exit(){
	fmt.Println("确定是否退出(Y/N): ")
	for {
	fmt.Scanln(&this.key)
	if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n"{
		break
	  }
	  fmt.Println("您的输入有误，请重新输入(Y/N) : ")
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
}

//显示主菜单
func (this *customerView) mainView() {
	for{
		fmt.Println("--------客户信息管理系统------------")
		fmt.Println("         1.添加客户   ")
		fmt.Println("         2.修改客户   ")
		fmt.Println("         3.删除客户   ")
		fmt.Println("         4.客户列表   ")
		fmt.Println("         5.退出   ")
		fmt.Println("请选择(1-5): ")


		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
		    this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default :
		    fmt.Println("你的输入有误，请重新输入...")						
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统的使用")
}

//简单登录功能的时间
func (this *customerView) Login (){
	account :=""
	pwd :=""
	for {
    fmt.Println("请输入账号： ")
	fmt.Scanln(&account)
	fmt.Println("请输入密码")
	fmt.Scanln(&pwd)
	if account == "7758258" && pwd =="111"{
		fmt.Println("恭喜你！正在进入系统！")
		break
	   }
	   fmt.Println("您的输入的账号或者密码有误，请重新输入: ")	   	
	}
	this.mainView()
}
func main() {
	//在主函数中，创建一个customerView并运行显示主菜单...
	customerView := customerView{
		key : "",
		loop : true,	
	}
	//完成对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.Login()
}

```

10）项目展示

1.登录

![](D:\myfile\后端\go语言学习\pic\面向对象\pic28.jpg)

2.客户列表

![](D:\myfile\后端\go语言学习\pic\面向对象\pic29.jpg)

3.添加客户

![](D:\myfile\后端\go语言学习\pic\面向对象\pic30.jpg)

4.修改客户

![](D:\myfile\后端\go语言学习\pic\面向对象\pic31.jpg)

5.删除客户

![](D:\myfile\后端\go语言学习\pic\面向对象\pic32.jpg)

6.退出

![](D:\myfile\后端\go语言学习\pic\面向对象\pic33.jpg)

## 十三、文件操作

### 1.基本介绍

文件对于我们并不陌生，文件是数据源（保存数据的地方）一种，比如大家经常使用的word文件，txt文件，excel文件...都是文件。文件最主要的作用就是保存数据，它既可以保存一张图片，也可以保存视频声音

文件在程序中是以流的形式来操作的

![](D:\myfile\后端\go语言学习\pic\文件\pic1.jpg)

- **流**：数据在数据源（文件）和程序（内存）之间经历的路径

- **输入流**：数据从数据源（文件）到程序（内存）的路径

- **输出流**：数据从程序（内存）到数据源（文件）的路径

os.File结构体封装所有文件相关操作

### 2.常用文件操作函数和方法

1）打开一个文件进行操作

os.Open(name string)(*File,error)

2）关闭一个文件：

File.Close()

3）其他的函数和方法在案例提

案例演示

```go
package main
import (
	"fmt"
	"os"
)
func main(){
	//打开文件
	//概念说明： file的叫法
	//1. file 叫file对象
	//1. file 叫file指针
	//1. file 叫file文件句柄
	file , err := os.Open("D:/test/test01/test.txt")
	if err !=nil {
		fmt.Println("open file err=",err)
	}
	//输出下文件，看看文件是什么,看出file/就是一个指针
	fmt.Printf("file=%v",file) //file=&{0xc0420705a0}

   //关闭文件
  err = file.Close()
  if err != nil {
	fmt.Println("close file err=",err)
  }

}
```

### 3.关于文件操作应用实例

1）读取文件的内容并显示在终端（带缓冲区的方式），使用os.Open,file.Close, bufio.NewReader(),reader.ReadString函数和方法

```go
package main
import (
	"fmt"
	"os"
	"io"
	"bufio"
)
func main(){
	//打开文件
	//概念说明： file的叫法
	//1. file 叫file对象
	//1. file 叫file指针
	//1. file 叫file文件句柄
	file , err := os.Open("D:/test/test01/test.txt")
	if err !=nil {
		fmt.Println("open file err=",err)
	}

	//当函数退出时，要及时的关闭file
	/*
    const (
		defaultBufSize  =4096 //默认缓冲区4096个字节
	)
	*/
	defer file.Close()  //要及时关闭file句柄，否则会有内存泄露

	//创建一个*Read ,带缓冲
	reader :=bufio.NewReader(file)
	//循环的读取文件的内容
	for {
		str,err := reader.ReadString('\n') //读到一个换行符就结束一次
		if err == io.EOF { //io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Printf(str)
	}
	fmt.Println("文件读取结束")
}
   
```

2）读取文件的内容并显示在终端（使用ioutil一次将整个文件读入到内存中），这种方式适用于文件不大的情况。相关方法和函数(ioutil.ReadFile)

```go
package main
import (
	"fmt"
	"io/ioutil"
	
)
func main(){
	//使用ioutil.ReaderFile一次性将文件读取到位
	file := "D:/test/test01/test.txt"
	content, err := ioutil.ReadFile(file)
	if err !=nil {
		fmt.Println("read file err=%v",err)
	}

	//把读取到的内容显示到终端
    fmt.Printf("%v",string(content)) //[]byte
    //因为我们没有显示的open文件，因此也不需要显示的close文件
	//因为。文件Open和Close背封装到ReadFile函数内部
	
}
```

注意：这种只适合文件不太大的情况使用

### 4.写文件操作应用实例（创建文件并写入文件）                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      

#### 1）基本介绍

```go
func OpenFile(name string,flag int,perm FileMode)(file *File,err error)
```

说明：os.OpenFile是一个更一般性的文件打开函数，它会使用指定的选项（如O_RDONLY）、指定的模式(如0666等)打开指定名称的文件。如果操作成功，返回的文件对象可用于i/o如果出错，错误底层类型是*PathError.

第二个参数：文件打开模式（可以组合），第三个参数：权限控制(linux)

![](D:\myfile\后端\go语言学习\pic\文件\pic2.jpg)

#### 2）基本应用实例-方式一

（1）创建一个新文件，写入内容 5句"hello Gardon"

```go
package main
import (
	"fmt"
	"os"
	"bufio"
)
func main(){
    //创建一个新文件，写入内容 5句"hello Gardon"
	//1.打开一个文件 "D:/test/test01/test.txt"
	filePath := "D:/test/test01/abc.txt"
	file, err := os.OpenFile(filePath,os.O_WRONLY | os.O_CREATE,0666)
	if err != nil {
		fmt.Printf("open file err=%v",err)
		return
	}
	//及时关闭file句柄，防止内存泄露
	defer file.Close()

	//准备写入6句话
	str := "hello Gardon\r\n"  // \r\n表示换行
	//写入时，使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存的，因此在调用WriterString方法时，其实内存是先写入缓存的
	//所以需要调用Flush()方法，将缓存的数据
	//真正写入到文件中，否则文件中会没有数据
	writer.Flush()
```

（2）打开一个存在的文件中，将原来的内容覆盖成新的内容10句 “你好,爸爸"

```go
package main
import (
	"fmt"
	"os"
	"bufio"
)
func main(){
   //打开一个存在的文件中，将原来的内容覆盖成新的内容10句 “你好,爸爸"
	//1.打开一个文件 "D:/test/test01/test.txt"
	filePath := "D:/test/test01/abc.txt"
	file, err := os.OpenFile(filePath,os.O_WRONLY | os.O_TRUNC,0666)
	if err != nil {
		fmt.Printf("open file err=%v",err)
		return
	}
	//及时关闭file句柄，防止内存泄露
	defer file.Close()

	//准备写入10句话:你好,爸爸
	str := "你好,爸爸!\r\n"  // \n 表示换行
	//写入时，使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存的，因此在调用WriterString方法时，其实内存是先写入缓存的
	//所以需要调用Flush()方法，将缓存的数据
	//真正写入到文件中，否则文件中会没有数据
	writer.Flush()
}
```

（3）打开一个存在的文件，在原来的内容追加内容“ABCI ENGLISH!”

```go
package main
import (
	"fmt"
	"os"
	"bufio"
)
func main(){
	//打开一个存在的文件，在原来的内容追加内容“ABCI ENGLISH!”
	//1.打开一个文件 "D:/test/test01/test.txt"
	filePath := "D:/test/test01/abc.txt"
	file, err := os.OpenFile(filePath,os.O_WRONLY | os.O_APPEND,0666)
	if err != nil {
		fmt.Printf("open file err=%v",err)
		return
	}
	//及时关闭file句柄，防止内存泄露
	defer file.Close()

	str := "ABCI ENGLISH!\r\n"  // \n 表示换行
	//写入时，使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存的，因此在调用WriterString方法时，其实内存是先写入缓存的
	//所以需要调用Flush()方法，将缓存的数据
	//真正写入到文件中，否则文件中会没有数据
	writer.Flush()
}
```

（4）打开一个存在的文件，将原来的内容读出显示在终端，并且追加"hello 北京"

```go
package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)
func main(){
	//打开一个存在的文件，将原来的内容读出显示在终端，并且追加"hello 北京"
	//1.打开一个文件 "D:/test/test01/abc.txt"
	//这是一个既要读又要写的操作
	filePath := "D:/test/test01/abc.txt"
	file, err := os.OpenFile(filePath,os.O_RDWR | os.O_APPEND,0666)
	if err != nil {
		fmt.Printf("open file err=%v",err)
		return
	}
	//及时关闭file句柄，防止内存泄露
	defer file.Close()

	//先读取原来文件的内容，并显示在终端
	reader := bufio.NewReader(file)
	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF { //如果读到文件末尾
            break
		}
		//显示到终端
		fmt.Print(str)
	}
//写到文件中
	str := "hello 北京\r\n"  // \n 表示换行
	//写入时，使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存的，因此在调用WriterString方法时，其实内存是先写入缓存的
	//所以需要调用Flush()方法，将缓存的数据
	//真正写入到文件中，否则文件中会没有数据
	writer.Flush()
}

```

使用os.OpenFile(),bufio.NewWriter(),“Writer的方法WriteString完成上面的任务

3）基本应用实例-方式二

编写一个程序，将一个文件的内容，写入到另一个文件，注意：这两个文件已经存在了

说明：

1）使用ioutil.ReadFile /ioutil.WriteFile 完成文件的任务

```go
package main
import (
	"fmt"
	"io/ioutil"
)
func main(){
	//将D:/test/test01/abc.txt文件的内容导入到D:/test/test01/kkk.txt中

	//1.首先将D:/test/test01/abc.txt内容读取到内存
	
	filePath :="D:/test/test01/abc.txt"
	filePath2 :="D:/test/test01/kkk.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		//说明读取文件有错误
		fmt.Printf("read file err=%v\n",err)
		return 
	}
//2.将读取到的内容写入：D:/test/test01/kkk.txt中
	err = ioutil.WriteFile(filePath2,data,0666)
	if err != nil {
		fmt.Printf("write file err=%v\n",err)
	}
	//会覆盖掉写入之后的文件中的所有内容

}

```

### 5.判断文件是否存在

golang判断文件或文件夹是否存在的方法是使用os.Stat()函数返回的错误进行判断：

1）如果返回的错误为nil,说明文件或文件夹存在

2）如果返回的错误类型使用so.IsNotExist()判断为true,说明文件或文件夹不存在

3）如果返回的错误为其他类型，则不确定是否存在

```go
//自己写了一个函数
func PathExists(path string)(bool,error){
	_,err :=os.Stat(path)
	if err == nil { //文件或目录存在
		return true,nil
	}
	if os.IsNotExist(err){
		return false,nil
	}
	return false,err
}
              
```

6.文件编程应用实例

拷贝文件

说明：将一张图片拷贝到另外一个目录下io包

```go
func Copy (dst Writer,src Reader)(written int64,err error)
```

```go
package main
import (
	"fmt"
	"io"
	"os"
	"bufio"
)

//自己写一个函数，接收两个文件路径 srcFileName dstFileName
func CopyFile(dstFileName string,srcFileName string)(written int64,err error){
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open file err=%v\n",err)
	}
	//用完了需要关闭
	defer srcFile.Close()
	//通过srcFile，获取到Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName :不能单纯地打开因为你不确定是否存在
	dstFile, err := os.OpenFile(dstFileName,os.O_WRONLY | os.O_CREATE,0666)
    if err != nil {
		fmt.Printf("open file err=%v\n",err)
		return
	}
	
	//通过dstFile,获取到writer
	writer := bufio.NewWriter(dstFile)//用完了需要关闭
	defer dstFile.Close()

	return io.Copy(writer,reader)

}
func main() {
	//将D:/test/dog.jpg拷贝到D:/test/test01/dog1.jpg
	//调用CopyFile完成文件的拷贝
	srcFile := "D:/test/dog.jpg"
	dstFile := "D:/test/test01/dog1.jpg"
	_, err :=CopyFile(dstFile,srcFile)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Printf("拷贝错误err=%v\n",err)
	}
}
```

### 6.统计英文、数字、空格和其他字符数量

说明：统计一个文件中含有的英文、数字-、空格及其它字符数量

```go
package main
import (
	"fmt"
	"os"
	"io"
	"bufio"
)
//定义个结构体，用于保存统计结构1
type CharCount struct {
	CharCount int //记录英文个数
	NumCount int //记录数字的个数
	SpaceCount int //记录空格的个数
	OtherCount int //记录其他字符的个数
}


func main(){
	//思路：打开一个文件。创建一个Reader
	//每读取一行，就去统计该行有多少个 英文、数字、空格和其他字符
	//然后将它们保存到一个结构体当中
fileName := "D:/test/abc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err =%v\n",err)
		return
	}
     defer file.Close()
	//定义一个结构体实例
	var count CharCount
	//创建一个Reder
	reader := bufio.NewReader(file)

	//开始循环读取fileName的内容
	for {
		str,err :=reader.ReadString('\n')
		if err == io.EOF {  //读到文件末尾
			break
		}
		//遍历str进行统计
		for _, v := range str {
			
		   switch {
			case v >= 'a' && v <= 'z' :
					fallthrough //穿透处理
			case v >= 'A' && v <= 'A' :
				count.CharCount ++ 	
			case v == ' ' || v =='\t' :
				count.SpaceCount ++	
			case v >= '0' && v <='9' :
				count.NumCount ++ 	
			default :
				count.OtherCount ++				 
			}
		}

	}
	//输出统计的结构
	fmt.Printf("字符的个数=%v,数字的个数=%v,空格的个数=%v,其他字符的个数=%v",count.CharCount,count.NumCount,count.SpaceCount,count.OtherCount)
}
```

## 十四、命令行参数

有一个需求

我们希望能够获取到命令行输入的各种参数，该如何处理？如下，我们执行一个可执行文件并附带一个参数

```
D:\myfile\GO\project\src\go_code\exec\count\main> test.exe tom c:/aa/bb/config.init 88
```

### 1.使用os.Args对参数进行解析

基本介绍

os.Args是一个string的切片，用来存储所有的命令行参数

应用案例

请编写一段代码，可以获取命令行的各种参数

```go
package main
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行的参数值",len(os.Args))
	//遍历os.Args切片，就可以得到所有命令行输入的参数值
	for i,v :=range os.Args {
		fmt.Printf("Args[%v]=%v\n",i,v)
	}
}
```

使用go build去编译一个可执行文件test.exe进行测试

![](D:\myfile\后端\go语言学习\pic\文件\pic3.jpg)

### 2.flag包解析命令行参数

flag包用来解析命令行参数

说明：前面的方式是比较原生的方式，对解析参数不是特别的方便，特别是带有指定参数形式的命令行

比如：cmd>main.exe -f C:/aaa.txt -p 200 -u root 这样的命令行，go设计者给我们提供了flag包，，可以方便的解析命令行参数，而且参数顺序可以随意

请编写一段代码，可以获取命令行的各个参数

```go
package main
import (
	"fmt"
	"flag"
)
func main() {

	 //定义几个变量，用于接收命令行的参数值
	 var user string
	 var pwd string
	 var host string
	 var port int
	 //&user 就是接收用户命令行中输入的-u后面的参数
	 //"u",就是-u 指定参数
	 //"",默认值
	 //"用户名，默认为空" 说明
	 flag.StringVar(&user,"u","","用户名，默认为空")
	 flag.StringVar(&pwd,"pwd","","密码，默认为空")
	 flag.StringVar(&host,"h","localhost","主机名，默认为localhost")
     flag.IntVar(&port,"port",3306,"端口号默认为3306")

	 //这里有一个非常重要的操作，转换，必须调用该方法
	 flag.Parse()

	 //输出结果
	 fmt.Printf("user=%v pwd=%v host=%v port=%v",
	user,pwd,host,port)

}
```

测试结果

![](D:\myfile\后端\go语言学习\pic\文件\pic4.jpg)

将顺序打乱或者不传参数再次进行测验

![](D:\myfile\后端\go语言学习\pic\文件\pic5.jpg)

## 十五、json以及序列化

### 1.概述

JSON(JavaScript Object Notation)是一种轻量级的数据交换格式，易于人阅读和编写，同时也易于机器解析和生成。key-val

JSON是2001年开始推广使用的数据格式，目前已成为主流的数据格式

JSON易于机器解析和生成，并有效地提升网络传输效率，通常程序在网络传输时会先将数据（结构体、map等）序列化成json字符串时，在反序列化恢复成原来的数据类型（结构体、map等）。这种方式已然成为各个语言的标准

![](D:\myfile\后端\go语言学习\pic\文件\pic6.jpg)

### 2.json应用场景图

![](D:\myfile\后端\go语言学习\pic\文件\pic7.jpg)

### 3.json数据格式说明

**在JS语言中，一切都是对象**。因此，任何支持的类型都可以通过JSON来表示，例如字符串、数字、对象、数组等

JSON键值对是用来保存 数据的一种方式

**键/值对**组合中的键名写在前面并引用双引号“”包裹，使用冒号:分隔，然后紧接着值：

[{"key1":val1,"key2":val2,"key3":val3,"key4":[val4,val5]},

{"key1":val1,"key2":val2,"key3":val3,"key4":[val4,val5]}]

比如：

```json
{"firstName": "Json"}
比如：
{"name":"tom","age":18,"address":["北京","上海"]}
比如：
[{"name":"tom","age":18,"address":["北京","上海"]},
{"name":"tom","age":18,"address":["北京","上海"]}]
```

任何数据类型都可以转换为json格式

json在线验证网站www.json.cn

![](D:\myfile\后端\go语言学习\pic\文件\pic8.jpg)

### 4.json的序列化

#### 1）介绍

json序列化是指，将有key-value结构的数据类型（比如结构体、map、切片）序列化成json字符串的操作

#### 2）应用案例

这里我们介绍一下结构体、map和切片的序列化，其他数据类型的序列化类似

```go
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
```

注意事项，对于结构体的序列化，如果我们希望序列化后的key的名字，由我们自己重新制定，那么可以给struct指定一个tag标签

```go
//定义一个结构体
type Monster struct {
	Name string `json:"monster_name"`//运用反射机制
	Age int `json:"monster_age"`
	Birthday string
	Sal float64
	Skill string
}
//这样做可以指定key值
```

序列化后：monster序列化后={"monster_name":"牛魔王","monster_age":500,"Birthday":"2011-11-11","Sal":8000,"Skill":"牛魔拳"}

### 5.json的反序列化

### 1）介绍

json反序列化是指，将json字符串反序列化成对应的数据类型（比如结构体、map、切片）的操作

### 2）应用案例

这里我们介绍一下将jason字符串反序列化成结构体、map和切片

代码演示

```go
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
```

**对上面代码的注意事项**

- 在反序列化一个json字符串时，要确保反序列化后的数据类型和原来序列化前的数据类型一致

- 如果json字符串是通过程序获取获取到的，则不需要对 “”进行转义处理\"，因为转义处理已经包含在内部了

## 十六、单元测试

### 1.引子

先看一个需求，怎样确定他运行的结果是正确的

```go
func addUpper (n int) int {
	res := 0
	for i :=1;i <=n;i++ {
		res +=i
	}
	return res
}
```

传统的方法解决：

在main函数中，调用addUpper函数，看看实际输出的结果是否与你预期的结果一致，如果一致，则说明函数正确。否则函数有错误，然后修改错误

```go
package main
import (
	"fmt"
)
//一个被测试函数
func addUpper (n int) int {
	res := 0
	for i :=1;i <=n;i++ {
		res +=i
	}
	return res
}
func main() {
	//传统的测试方法，就是在main函数中使用看看结果是否正确
    res :=addUpper(10)
	if res != 55 {
		fmt.Printf("adUpper错误，返回值=%v 期望值=%v\n",res,55)
	} else {
        fmt.Printf("adUpper正确，返回值=%v 期望值=%v\n",res,55)
	}
	
}
```

**传统方法的缺点分析**

- 不方便，我们需要在main函数中去调用，这样就需要去修改main函数，如果现在项目正在运行，就可能去停止项目。
- 不利于管理，因为当我们测试多个函数或者多个模块时，都需要写在main函数中，不利于我们的管理和清晰我们的思路
- 引出单元测试。->testing测试框架，可以很好的解决问题

### 2.单元测试-基本介绍

go语言中自带一个轻量记得测试框架testing和自带的go test命令来完成单元测试和性能测试，testing框架和其他语言中的测试框架类似，可以基于该框架写相应的压力测试用例。通过单元测试，可以解决以下问题

1）确保每个函数是可运行的，并且运行结果是正确的

2）确保写出来的代码性能是好的

3）单元测试及时的发现程序设计或实现的逻辑错误，使问题及早暴露，便于问题的定位解决，而性能测试的重点在于发现程序设计上的一些问题，让程序能够在高并发的情况下还能保持稳定

使用go的单元测试，对addUpper和sub函数进行测试

注意：测试时，可能需要暂时退出360（因为360可能认为生成的测试用例的程序是木马）

### 3.代码实现

```go
package main

//一个被测试函数
func AddUpper (n int) int {
	res := 0
	for i :=1;i <=n;i++ {
		res +=i
	}
	return res
}

//求两个数的差
func getSub(n1 int,n2 int) int {
	return n1 - n2
}

```

cal_test.go

```go
package main
import (
	_"fmt"
	"testing" //引入go的testing框架包
)

//编写测试用例，去测试，去测试addUpper函数是否正确   、
func TestAddUpper(t *testing.T) {
  
	//调用
	res := AddUpper(10)
	if res != 55 {
		//fmt.Println("AddUpper(10)执行错误，期望值=%v实际值=%v\n",55,res)
		t.Fatalf("AddUpper(10)执行错误，期望值=%v实际值=%v\n",55,res)
	}

	//如果正确，输出日志
	t.Logf("AddUpper(10)执行正确...")
}                                                                                                                                                                                                                                                         
```

sub_test.go

```go
package main
import (
	_"fmt"
	"testing" //引入go的testing框架包
)

//编写测试用例，去测试，去测试sub函数是否正确   、
func TestGetSub(t *testing.T) {
  
	//调用
	res := getSub(10,3)
	if res != 7 {
		
		t.Fatalf("getSub(10)执行错误，期望值=%v实际值=%v\n",7,res)
	}

	//如果正确，输出日志
	t.Logf("getSub(10)执行正确...")
}                                                                                                                                                           
```

在cmd中执行go test -v就可以对此函数进行测试操作了

![](D:\myfile\后端\go语言学习\pic\文件\pic9.jpg)

单元测试的运行原理

![](D:\myfile\后端\go语言学习\pic\文件\pic10.jpg)

### 4.单元测试的细节说明

- 测试用例文件名必须以_test.go结尾，比如cal_test.go，cal不是固定的

- 测试用例函数必须以Test开头，一般来说就是Test_被测试的函数名，比如TestAddUpper.

- TestAddUpper(t *testing.T)的形参类型必须是*testing.T

- 一个测试用例文件中，可以有多个测试用例函数，比如TestUpper.TestSub

- 运行测试用例的指令为

  1. cmd > go test [如果运行正确，无日志，错误时，会输出日志]
  2. cmd>go test -v [运行正确或者错误，都输出日志]

- 当出现错误时，可以用t.Fatalf来格式化输出错误信息，并退出程序

- t.Logf（“”）方法可以输出相应的日志

- 测试用例函数，并没有放在main函数中，也执行了，这就是测试用例的方便之处

- PASS表示测试用例运行成功，FAIL表示测试用例运行失败

- 测试单个文件一定要带上被测试的源文件

  go test -v cal.test,go cal.go

- 测试单个方法

  go test -v -test.run TestAddUpper

- sd

  

### 5.单元测试的综合案例

1）编写一个Monter结构体，字段Name,Age,Skill

2）给Monster绑定方法Store，可以将一个Monster变量(对象)，序列化后保存到文件中

3）给Monster绑定方法ReStore，可以将一个序列化的Monster,从文件中读取，并反序列化为Monster对象

4）编程测试用例文件store_go编写测试用例函数TestStore和TestRestore进行测试

monster.go

```go
package monster
import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)
type Monster struct {
	Name string
	Age int
	Skill string
}

//给Monster绑定方法Store，可以将一个Monster变量(对象)，序列化后保存到文件中
func (this *Monster) Store() bool{

	//先序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err = ", err)
		return false
	}
	//保存到文件
		filePath := "D:/test/test02/monster.ser"
		err = ioutil.WriteFile(filePath, data,0666)
		if err != nil {
			fmt.Println("write file  err = ", err)
			return false
	}
	return true
	//保存到文件中
}





//给Monster绑定方法ReStore，可以将一个序列化的Monster,从文件中读取，
// 并反序列化为Monster对象
func (this *Monster) ReStore() bool {

	//1.先从文件中读取序列化字符串
	filePath := "D:/test/test02/monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Read file  err = ", err)
		return false
	}

	//2.使用读取到的data []byte，对反序列化
	err = json.Unmarshal(data,this)
	if err != nil {
        fmt.Println("Unmarshal  err = ", err)
		return false
	}
	return true
}
```

monster_test.go

```go
package monster
import (
	"testing"
)
//测试用例，测试Store方法
func TestStore(t *testing.T) {

	//先创建一个Monster实例
	monster := &Monster {
		Name : "红孩儿",
		Age : 10,
		Skill : "吐火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()错误，希望为=%v 实际为=%v",true,res)
	}
	t.Logf("monster.Store()测试成功")

}

func TestReStore(t *testing.T) {
	//创建一个Monster实例，不需要指定字段的值
	var monster = &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore()错误，希望为=%v 实际为=%v",true,res)
	}

	//进一步判断
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore()错误，希望为=%v 实际为=%v",true,monster.Name)

	}
	t.Logf("monster.ReStore()测试成功")

	}	

```

cmd运行

```
D:\myfile\GO\project\src\go_code\TestUnit\demo2>go test -v -test.run TestReStore
=== RUN   TestReStore
--- PASS: TestReStore (0.00s)
        moster_test.go:35: monster.ReStore()测试成功
PASS
ok      go_code/TestUnit/demo2  0.191s
```

将测试文件中改一下

```
D:\myfile\GO\project\src\go_code\TestUnit\demo2>go test -v -test.run TestReStore
=== RUN   TestReStore
--- FAIL: TestReStore (0.00s)
        moster_test.go:32: monster.ReStore()错误，希望为=true 实际为=红孩儿~
FAIL
exit status 1
FAIL    go_code/TestUnit/demo2  0.181s

```

## 十七、goroutine(协程)和chan

### 1.goroutine入门

需求：要求统计1-20000的数字中，哪些是素数

分析思路

1）传统的方法，就是使用一个循环，循环的判断各个数是不是素数

2）使用并发或者并行的方式，将统计素数的任务分配给多个goroutine去完成。这是就会使用到goroutine去完成，这时就会使用goroutine

### 2.goroutine基本介绍

#### -1.进程和线程说明

1）进程就是程序在操作系统中的一次执行过程，是系统进行资源分配和调度的基本单位

2）线程是进程的一个执行实例吗，是程序执行的最小单位，他是比进程更小的能独立运行的基本单位

3）一个进程可以创建和销毁多个线程，同一个进程中的多个线程可以并发执行

4）一个程序至少有一个进程，一个进程至少有一个线程

#### -2.程序、进程和线程的关系示意图

![](D:\myfile\后端\go语言学习\pic\进程\pic1.jpg)

-3.并发和并行

1）多线程程序在单核上运行，就是并发

2）多线程程序在多核上运行，就是并行



**并发**：因为是在一个CPU上，比如有10个线程，每个线程执行10毫秒（进行轮换操作），从人的角度来看，好像这10个线程都在运行，但是从微观上看，在某一个时间点来看，其实只有一个线程在执行，这就是并发

**并行**：因为是在多个CPU上(比如有10个CPU)，比如有10个线程，每个线程执行10毫秒（各自在不同的CPU上执行），从人的角度上看，这10个线程都在运行，但是从微观上看，在某一个时间点，也是同时有10个线程在执行，这就是并行

#### -3.Go协程和Go主线程

1）Go主线程（有程序员直接称为线程/也可以理解成进程）：一个Go线程上，可以起多个协程，你可以这样理解，**协程是轻量级的线程**

2）Go协程的特点

- 有独立的栈空间
- 共享程序堆空间
- 调度由用户控制
- 协程是轻量级的线程

示意图

![](D:\myfile\后端\go语言学习\pic\进程\pic2.jpg)

### 3.案例说明

请编写一个程序，完成如下功能

1）在主线程（可以理解成进程）中，开启一个goroutine,该协程每隔1秒输出"hello world"

2）在主线程也每隔一秒输出"hello golang",输出10次后，退出程序

3）要求主线程和goroutine同时执行

4）画出主线程和协程执行流程图

代码实现

```go
package main
import (
	"fmt"
	"strconv"
	"time"
)
/*
1）在主线程（可以理解成进程）中，开启一个goroutine,该协程每隔1秒输出"hello world"

2）在主线程也每隔一秒输出"hello golang",输出10次后，退出程序

3）要求主线程和goroutine同时执行
*/
//编写一个函数，每隔1秒输出"hello world
func test () {
	for i := 1; i <= 10; i++ {
		fmt.Println("test()hello world"+strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {
	go test() //开启了一个协程
    for i := 1; i <= 10; i++ {
		fmt.Println("main()hello world"+strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
```

执行结果如下，我们可以发现主线程和go协程是同时执行的

![](D:\myfile\后端\go语言学习\pic\进程\pic3.jpg)

go主线程与go协程的执行示意图

![](D:\myfile\后端\go语言学习\pic\进程\pic4.jpg)

### 4.小结

1）主线程是一个物理线程，直接作用在cpu上的，是重量级的，非常消耗cpu资源，

2）协程从主线程开启的，是轻量级的线程，是逻辑态。对资源消耗相对少

3）golang的协程机制是重要的特点，可以轻松开启上万个协程。其他编程语言的并发机制是一般基于线程的，开启过多的线程，资源耗费大，这里就凸显出golang在并发上的优势了

### 5.MPG模式基本介绍

![](D:\myfile\后端\go语言学习\pic\进程\pic5.jpg)

1）M：操作系统的主线程（是物理线程）

2）P：协程执行需要的是上下文

3）G：协程

![](D:\myfile\后端\go语言学习\pic\进程\pic6.jpg)

![](D:\myfile\后端\go语言学习\pic\进程\pic7.jpg)

### 6.设置Golang运行的CPU数

介绍：为了充分利用多cpu的优势，在Golang程序中，设置运行的cpu数目

```go
package main
import (
	"fmt"
	"runtime"
)

func main() {
	//获取当前系统CPU的数量
	num := runtime.NumCPU()
	//我这里设置num -1的cpu运行go程序
	runtime.GOMAXPROCS(num)
	fmt.Println("num=",num)
}
```

1）go1.8后，默认让程序运行在多个核上，可以不用设置了

2）go1.8前，还是要设置一下，可以更高效的利用CPU了

### 7.协程并发（并行）资源竞争的问题

需求:现在要计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中，最后显示出来。要求使用goroutine完成



分析思路：

1）使用goroutine来完成，效率高，但是会出现并发/并行安全问题

2）这里就提出了不同1goroutine如何通信的问题

代码实现

1）使用goroutine来完成（看看使用goroutine并发完成会出现什么问题？



2）在运行某个程序时，如何知道是否存在资源竞争的问题，方法很简单。在编译该程序时增加一个参数 -race即可

![](D:\myfile\后端\go语言学习\pic\进程\pic9.jpg)

会发现map有些有值有些没有值，各个协程出现了资源竞争的问题

3）示意图

![](D:\myfile\后端\go语言学习\pic\进程\pic8.jpg)

他们之间会出现资源竞争的问题

### 8.全局互斥锁解决资源竞争

不同的goroutine之间如何通信

1）全局变量加锁同步

2）channel

使用全局变量加锁同步改进程序

因为没有针对全局变量m加锁，因此会出现资源竞争的问题，代码会出现报错提示concurrent map writes

解决方案,-1加入互斥锁

```go
package main
import (
	"fmt"
	"time"
	"sync"
)
//需求:现在要计算1-200的各个数的阶乘，
// 并且把各个数的阶乘放入到map中，最后显示出来。要求使用goroutine完成

//思路
//1.编写一个函数，来计算各个数的阶乘，并放入到map中
//2.我们爱动的协程是多个，统计的结果放入到map中
//2.map应该做出一个全局的

var (
	myMap = make(map[int]int,10) 
	//声明一个全局的互斥锁
	//lock是一个全局的互斥锁
	//sync 是包：synchornized 同步
	//Mutex是互斥的意思
	lock sync.Mutex
)

//test函数就是计算n的阶乘，将这个结果放入到map中
func test(n int) {
	res := 1
	for i :=1; i <=n;i++ {
		res *= i
	}

	//这里我们将res放入到myMap中
	//加锁
	lock.Lock()
    myMap[n]= res//concurrent map writes
	//解锁
   lock.Unlock()
}

func main() {
	//我们这里开启多个协程完成这个任务[200个协程]
	for i :=1; i <=15; i++ {
		go test(i)
	}
	//休眠10秒
	time.Sleep(time.Second * 5)
	//输出结果，遍历结果
	lock.Lock()
	for i,v :=range myMap {
		fmt.Printf("map[%d]=%d\n",i,v)
	}
    lock.Unlock()
}
```

我们的数的阶乘很大，结果会越界，我们可以改成sum +=uint64(i)



加锁解释

![](D:\myfile\后端\go语言学习\pic\进程\pic10.jpg)

## 十八、管道

### 1.为什么要使用channel

前面使用全局变量加锁同步来解决goroutine的通讯，但不完美

1）主线程在等待所有goroutine全部完成的时间很难确定。我们这里设置10秒，仅仅只是估算

2）如果主线程休眠时间长了，会加长等待时间，如果等待时间短了，可能还有goroutine处于工作状态，这时也会随着主线程的退出而销毁

3）通过全局变量加锁同步来实现通讯，也并不利用多个协程对全局变量的读写操作

4）上面的种种分析都在呼唤一个新的通讯机制-channel

### 2.channel的介绍

1）channel本质就是一个数据结构-队列

2）数据是先进先出[FIFIO frist in first out]

3）线程安全，多goroutine访问时，不需要加锁，就是说在channel本身就是线程安全的

4）channel是有类型的，一个string的channel只能存放string数据

![](D:\myfile\后端\go语言学习\pic\进程\pic11.jpg)

**channel是线程安全，多个协程作同一个管道时，不会发生资源竞争的问题**

### 3.管道的定义/声明channel

var 变量名 chan 数据类型

举例

```go
var intChan chan int (intChan用于存放int数据)
var mapChan chan map[int]string (mapChan用于存放map[int]string类型)
var perChan chan Person
var perChan2 chan *Person
...
```

说明

1）channel是引用类型

2）channel必须初始化才能写入数据，即make后才能使用

3）管道是有类型的 intChan只能写入整数int

管道的初始化，写入数据到管道，从管道读取数据以及基本的注意事项

```go
package main
import (
	"fmt"
)
func main() {
	//演示一下管道的使用
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int,3)

	//2.看看intChan是什么
	fmt.Printf("intchan的值是=%v intChan本身的地址=%p\n",intChan,&intChan)
    
	//3.像管道写入数据
	intChan<-10
	num := 211
	intChan<- num

	//注意点，当我们在给管道写入数据时，不能超过其容量
	intChan<- 50
	//intChan<- 98 //会报错
	//4.输出看看管道的长度和cap（容量）
	fmt.Printf("channel len =%v cap=%v\n",len(intChan),cap(intChan)) // 3,3
    
	//5.从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Printf("num2=%v\n",num2) //10
	fmt.Printf("channel len =%v cap=%v\n",len(intChan),cap(intChan))//2,3
	 
	//6.在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
    num3 := <-intChan
    num4 := <-intChan
    // num5 := <-intChan
    // fmt.Println("num3=",num3,"num4=",num4,"num5=",num5)//报错


	}   

```

### 4.channel使用的注意事项

1）channel中只能存放指定的数据类型

2）channel的数据放满后，就不能在放入了

3）如果从channel取出数据后，可以继续放入

4）在没有使用协程的情况下，如果channel数据取完了，再取就会报deadlock

### 5.读写channel案例演示

![](D:\myfile\后端\go语言学习\pic\进程\pic12.jpg)

![](D:\myfile\后端\go语言学习\pic\进程\pic13.jpg)

![](D:\myfile\后端\go语言学习\pic\进程\pic14.jpg)

![](D:\myfile\后端\go语言学习\pic\进程\pic15.jpg)

![](D:\myfile\后端\go语言学习\pic\进程\pic16.jpg)

![](D:\myfile\后端\go语言学习\pic\进程\pic17.jpg)

```go
package main
import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
}

func main() {
	//定义一个存放任意数据类型的管道3个数据
	// var  allChan chan interface{}
    allChan := make(chan interface{},3)

	allChan<-10
	allChan<-"tom jack"
	cat :=Cat{"小花猫",4}
	allChan<- cat
 
//我们希望获得管道中的第三个元素，则先将前2个推出
<-allChan
<-allChan

newCat :=<-allChan //从管道中取出来的cat是什么
fmt.Printf("newCat=%T,newCat=%v\n",newCat,newCat)//newCat=main.Cat,newCat={小花猫 4}
//下面的写法是错误的,编译不通过，则使用类型断言就可以通过
// fmt.Printf("newCat.Name=%v",newCat.Name)
a :=newCat.(Cat)
fmt.Printf("newCat.Name=%v",a.Name)//newCat.Name=小花猫
}

```

### 6.channel的遍历和关闭

#### -1.channel的关闭

使用内置函数close可以关闭channel,当channel关闭后，就不能再向channel写数据了，但是仍然可以从谈channel读取数据

```go
package main
import (
	"fmt"
)
func main() {
	intChan :=make(chan int,3)
	intChan<- 100
	intChan<- 200
	close(intChan) //close
	//这时不能够再写入到数channel
	//intChan<- 300 //panic: send on closed channel
	fmt.Println("okok~")
	//当管道关闭后，读取数据是可以的
	n1 := <-intChan
	fmt.Println("n1=",n1)
//输出如下
   // okok~
     //n1= 100
}
```



#### -2.channel的遍历

channel支持for-range的方式进行遍历，请注意两个细节

1）在遍历时，如果channel没有关闭，则会出现deadlock的错误

2）在遍历时，如果cahnnel已经关闭，则会正常遍历数据，遍历完成后，就会退出遍历

代码演示

```go
package main
import (
	"fmt"
)
func main() {
	intChan :=make(chan int,3)
	intChan<- 100
	intChan<- 200
	close(intChan) //close
	//这时不能够再写入到数channel
	//intChan<- 300 //panic: send on closed channel
	fmt.Println("okok~")
	//当管道关闭后，读取数据是可以的
	n1 := <-intChan
	fmt.Println("n1=",n1)

	//遍历管道
    intChan2 :=make(chan int,100)
	for i :=0; i < 100; i++ {
        intChan2 <- i *2 //放入100个数据进去管道之中
	}
	//遍历:这种遍历是错误的，因为遍历过程中管道的长度会变化
	// for i :=0; i < len(intChan2);++ {

	// }
	//在遍历时，如果channel没有关闭，则回出现deadlock的错误

	//在遍历时，如果cahnnel已经关闭，则会正常遍历数据，遍历完成后，就会退出遍历
	close(intChan2)
	for v := range intChan2 {
        fmt.Println("v=",v)
	}
}
```

7.应用案例

-1.应用案例1

请完成goroutine和channel协同工作案例，具体要求

1）开启一个writeData协程，向管道intChan中写入50个整数

2）开启一个readData协程，从管道inChan中读取writeData写入的数据

3）注意：writeData和readData操作的是同一个管道

4）主线程需要等到writeData协程都完成工作才能退出

思路分析

![](D:\myfile\后端\go语言学习\pic\进程\pic18.jpg)

看代码演示：

```go
package main
import (
	"fmt"
	_"time"
)
//writeDtata
func writeData(intChan chan int) {
	for i :=0;i<=50;i++ {
		//放入数据
		intChan<- i
		fmt.Println("writeData",i)
		// time.Sleep(time.Second )
	}
	close(intChan)//关闭管道，不影响读
}

////readDtata
func readData(intChan chan int,exitChan chan bool) {
	for {
		v,ok := <-intChan
		if !ok {
			break
		}
		//time.Sleep(time.Second )
		fmt.Printf("readData 读到的数据=%v\n",v)
	}
	//readData 读取完数据后，即任务完成
	exitChan<- true //数据读取完之后就网退出管道加入一个1
	close(exitChan)
}
func main() {
	//创建两个管道
	intChan := make(chan int,50)
	exitChan :=make(chan bool,1 )

	go writeData(intChan)
	go readData(intChan,exitChan)
	//time.Sleep(time.Second * 10)
	for {
		_, ok :=<-exitChan
		if !ok {
			break
		}
	}
}
```

### 7.管道阻塞的机制

#### -1.应用实例2 --阻塞

```go
func main() {
 intChan :=make(chan int, 10) //10->50的话数据一下就放入了
 exitChan :=make(chan bool,1)
 //go readData(intChan,exitChan)
 
 //就是为了等待readData的协程完成
 for ——=range exitChan{
 fmt.Println("ok...")
 }
}
```

问题：如果注销掉go readData(intChan, exitChan)程序会怎么样

答：如果只是向管道写入数据，而没有读取，就会出现阻塞而deadLock，原因是intChan容量是10，而writeData会写入50个数据，因此会阻塞在writeData的ch <-i

#### 2-应用实例3

1）需求：要求统计1 200000的数字中，哪些是素数？这个问题在本章开篇就提出了，
现在我们有goroutine和channel的知识后，就可以完成了[测试数据：80000]

2）分析思路：

- 传统的方法，就是使用一个循环，循环的判断各个数是不是素数。

- 使用并发/并行的方式，将统计素数的任务分配给多个（4个)goroutine去完成,

  完成任务时间短。

1.画出分析思路

![](D:\myfile\后端\go语言学习\pic\进程\pic19.jpg)

2.代码实现

```go
package main
import (
	"fmt"
	"time"
)
//向intChan放入 1-8000个数
func putNum(intChan chan int){
	for i := 0 ;i<8000; i++{
		intChan<- i
	}
	//关闭intChan
	close(intChan)
}

//从intchan中取出数据，并判断是否为素数,如果是就放入到primeChan
func primeNum(intChan chan int,primeChan chan int,exitChan chan bool){

	//使用for循环
	
	var flag bool
	for {
		time.Sleep(time.Millisecond)
		num,ok := <-intChan
		if !ok { //intChan取不到的时候，就退出这个主for循环
			break
		}
		flag = true //假设是素数
		//判断num是不是素数
		for i :=2;i<num;i++{
			if num %i ==0 { //说明i不是素数
				flag = false
				break
			}
		}
		if flag {
			//将这个数就放入到primeChan中
			primeChan<- num
		}

	}
	fmt.Println("有一个协程因为取不到数据没退出！")
	//这里我们还不能关闭primeChan
	//向exitChan 写入true
	exitChan<- true
}
func main() {
	intChan :=make(chan int,1000)
	primeChan :=make(chan int,2000) //放入结果
	//标识退出的管道
	exitChan :=make(chan bool ,4) //4个
	
	//开启一个协程，向intChan放入 1-8000个数
	go putNum(intChan)
	//开启4个协程，从intchan中取出数据，并判断是否为素数,如果是就放入到primeChan
	for i :=0;i<4; i++{
		go primeNum(intChan,primeChan,exitChan)
	}
	//这里我们主线程，进行处理
	go func() {
		for i :=0;i<4; i++{
			<-exitChan
		}
		//当我们从exitChan祛除了4个结果，就可以放心关闭primeChan
		close(primeChan)
	}()

	//遍历primeChan
	for {
		res,ok := <-primeChan
		if !ok {
			break
		}
		//将结果输出
		fmt.Printf("素数为=%d\n",res)
	}
	fmt.Println("main主线程退出")


}
```

说明：使用goroutine完成后，可以在使用传统的方法来统计一下，看看完成这个
任务，各自耗费的时间是多少?[用map保存primeNum]

使用go协程后，执行的速度，比普通方法提高至少4倍

### 8.channel使用细节和注意事项



#### 1）channel可以声明为只读，或者只写性质

```go
package main
import (
	"fmt"
)
func main() {
	//管道可以生命为只读或只写

	//1.在默认的情况下，管道是双向的。
	// var chan1 chan int //可读可写


	//2.声明为只写
	var chan2 chan<- int
	chan2 = make(chan int,3)
	chan2<- 20
	// num := <-chan2 err在这个管道中不可以读
	fmt.Println("chan2=",chan2)

	//3.声明为只读
	var chan3 <-chan int
	num2 := <-chan3
	// chan3<- 30 err 会报错，因为该管道为只读
	fmt.Println("num2=",num2)
}
```

#### 2）channel只读和只写的最佳实践案例

```go
package main
import (
	"fmt"
)

//ch chan<- int,这样ch就只能写操作
func send (ch chan<- int,exitChan chan struct{}){
	for i :=0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}
//ch <- chan int,这样ch就只能读操作了
func recv(ch <-chan int,exitChan chan struct{}){
	for {
		v,ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan <- a
}
func main() {
	var ch chan int
	ch = make(chan int , 10)
	exitChan :=make(chan struct{},2)
	go send(ch,exitChan)
	go recv(ch,exitChan)
	var total = 0
	for _= range exitChan {
		total ++
		if total == 2 {
			break
		}
	}
	fmt.Println("结束...")
}
```

#### 3）使用select可以解决从管道取数据的阻塞问题

```go
package  main
import (
	"fmt"
	"time"
)
func main() {

	//使用select可以解决从管道读取数据阻塞问题

	//1.先定义一个管道 10个数据 int
	intChan :=make(chan int, 10)
	for i := 0 ; i < 10 ;i ++{
		intChan<- i
	}
	//2.定义一个管道5个数据string
	stringChan :=make (chan string , 5)
	for i := 0; i < 5 ; i++ {
		stringChan <- "hello" +fmt.Sprintf("%d",i)
	}

	//传统方法遍历管道时，如果不关闭会阻塞而导致 deadlock

	//问题，在实际开发中，可能我们不好确定什么时候关闭该管道
	//可以使用select 方式解决
	label:
	for {
		select  {
		case v := <-intChan :  //注意：这里如果 intChan一直没有关闭，不会导致deadlocks,会自动到下一个case
			fmt.Printf("从intChan读取的数据%d\n",v)
			time.Sleep(time.Second)
		case v := <-stringChan :
			fmt.Printf("从stringChan读取的数据%s\n",v)	
			time.Sleep(time.Second)
		default :
			fmt.Println("都取不到，不玩了，你可以加入逻辑")	
			time.Sleep(time.Second)
			return
			break label
		}

	}

}
```

#### 4）goroutine中使用recover。解决协程中出现panic,导致程序崩溃问题

```go
package main
import (
	"fmt"
	"time"
)
//函数1
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}
//函数2
func test(){
	//这里试用贴defer + recover
	defer func() {
		//捕获test爬出的panic
		if err := recover(); err !=nil {
			fmt.Println("test()发生错误",err)
		}
	}()
	//定义了一个map
	var myMap map[int]string
	myMap[0] = "golang" //erro
}
func main(){
	go sayHello()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=",i)
		time.Sleep(time.Second)
	}
}
```

输出结果如下

![](D:\myfile\后端\go语言学习\pic\进程\pic20.jpg)

**说明：如果我们起了一个协程，但...是这个协程出现了panic，如果我们没有捕获这个panic。就会造成整个程序崩溃，这时我们可以在goroutine中使用recover来捕获panic,进行处理，这些即使这个协程发生的问题，但是主线程任然不受影响，可以继续运行**

## 十九、反射

### 1、反射的使用场景

#### 1）结构体标签的应用

![](D:\myfile\后端\go语言学习\pic\进程\pic21.jpg)

#### 2）使用反射机制编写函数的适配器（桥连接）

![](D:\myfile\后端\go语言学习\pic\进程\pic22.jpg)

### 2、反射的基本介绍

#### -1.基本介绍

1）反射可以在运行时动态获取变量的各种信息，比如变量的类型（type），类别(kind)

2）如果是结构体变量，还可以获取到结构体本身的信息（包括结构体的字段、方法）

3）通过反射，可以修改变量的值，可以调用关联的方法

4）使用反射，需要import ("reflect")

#### -2.反射的图解

![](D:\myfile\后端\go语言学习\pic\进程\pic23.jpg)

#### -3.反射重要的函数和概念

1）reflect.TypeOf(变量名)。获取变量的类型，返回reflect.Type类型

2）reflect.ValueOf(变量名)。获取变量的值，返回reflect.Value类型reflect.Value是一个结构体类型。通过reflect.Value.可以获取到关于该变量的很多信息

3）变量、interface{}和reflect.Value是可以相互转换的，这点在实际开发中，会经常使用到。

![](D:\myfile\后端\go语言学习\pic\进程\pic24.jpg)

#### 

### 3.反射快速入门

#### -1.请编写一个函数，演示对（基本数据类型、interface{}、reflect.Value）进行反射的基本操作。

代码演示

```go
package main
 import (
	"fmt"
	"reflect"
 )

 //专门反射演示
 func reflectTest01(b interface{}) {
	//通过反射获取的传入的变量 type,kind,值
	//1.先获取到reflect.Type
	rTyp :=reflect.TypeOf(b)
	fmt.Println("rTyp=",rTyp) //int
	//2.获取到reflectValue
	rVal :=reflect.ValueOf(b) //100
	// n1 := 10
	// n2 := 2 + rVal
	// fmt.Println("n2=",n2)//erro 

	// n1 := 10
	n2 := 2 + rVal.Int()
	fmt.Println("n2=",n2) //102

	fmt.Printf("rVal=%v type=%T\n",rVal,rVal) //rVal=100 type=reflect.Value

	//下面我们将rVal转成interface{}
	iv := rVal.Interface()
	//将interface{}通过断言转成需要的类型
	num2 := iv.(int)
	fmt.Println("num2=",num2) //num2= 100
 }
 func main() {
	//-1.请编写一个函数，演示对（基本数据类型、interface{}、reflect.Value）进行反射的基本操作。
	
	//1、先定义一个int
	var num int = 100
	reflectTest01(num)
 }
```

#### -2.请编写一个案例，演示对（结构体类型、interface{}、reflect.Value）进行反射的基本操作

```go
 //专门演示对结构体的反射
 func reflectTest02(b interface{}) {
	//通过反射获取的传入的变量 type,kind,值
	//1.先获取到reflect.Type
	rTyp :=reflect.TypeOf(b)
	fmt.Println("rTyp=",rTyp) //rTyp= main.Student

	//2.获取到reflectValue
	rVal :=reflect.ValueOf(b)
	
	//下面我们将rVal转成interface{}
	iv := rVal.Interface()
	fmt.Printf("iv=%v ,iv type=%T\n",iv,iv) //iv={Tom 20} ,iv type=main.Student
	//将interface{}通过断言转成需要的类型
	// fmt.Printf("iv=%v ,iv type=%T name\n",iv,iv,iv.Name) //erro无法取出name的值
	//所以先要进行断言
	stu,ok :=iv.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n",stu.Name)//stu.Name=Tom
	}

	
 }

 type Student struct {
	Name string
	Age int
 }
 func main() {


	//2.定义一个Student的实例
	stu := Student{
		Name : "Tom",
		Age : 20,
	}

	reflectTest02(stu)

 }
```

### 4.反射的注意事项和细节说明

#### 1）reflect.Value.Kind，获取变量的类别，返回的是一个常量

补充常量是知识

常量介绍

- 常量使用const修改

- 常量在定义的时候，必须初始化

- 常量不能修改

- 常量只能修饰bool、数值类型(int,float系列)、string类型

  语法： const identifier [type] = value

  ```go
  package main
  import (
  	"fmt"
  )
  func main() {
  	var num int
  	num = 90 
  	//常量声明的时候必须赋值
  	const tax int = 90
  	// tax = 10 //常量是不能修改的
  	fmt.Println(num,tax)
      //常量只能修饰bool、数值类型(int,float系列)、string类型
      //const b =num /3 erro
  }
  ```

  常量使用注意事项

  1》比较简洁的写法

  ```
  func main () {
  const(
  	a = 1
  	b = 2
  )
  }
  ```

  2》还有一种专业的写法

  ```
  func main(){
  	const(
  	a = iota
      b
      c
  	)
  	fmt.Println(a,b,c)//0,1,2
  }
  ```

  3》Golang中没有常量名必须大写的规范

  4》仍然通过首字母的大小来控制常量的访问范围

#### 2）Type是类型，kind是类别，Type和Kind可能是相同的，也可能是不同的

比如: var num int = 10 num 的Type是int, Kind也是int

比如：var stu Student stu 的Type是**包名.Student**,Kind是**struct**（Kind的等级 >Type的等级）

3）通过反射可以让变量在interface{}和Reflect.Value之间相互转换

4）使用反射的方式来获取变量的值（并返回对应的类型），要求数据类型匹配，比如x是int,那么就应该使用reflect.Value(x)Int(),而不能使用其他的，否则报panic

5）通过反射的值来修改变量。注意当使用SetXxx方法来设置需要通过对应的指针类型来完成，这样才能改变传入的值，同时需要使用到reflect.Value.Elem()方法

```go
package main
import (
	"fmt"
	"reflect"
)

//通过反射，修改
//num int 的值
//修改student的值

func reflect01(b interface{}){
	//获取到reflect.Value
	rVal := reflect.ValueOf(b)
	rVal.Elem().SetInt(20)
}
func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println("num=",num) //20


	//你可以这样理解这句话：rVal.Elem()
	// num := 9
	//  ptr *int = &num
	//  num2 :=*ptr
}
```

6）如何理解rVal.Elem()

```go
//你可以这样理解这句话：rVal.Elem()
	// num := 9
	//  ptr *int = &num
	//  num2 :=*ptr
}
```

### 5.反射练习题

1）给你一个变量 var v float64 = 1.2.请使用反射来得到它的reflect.Value,然后获取对应的Type,Kind和值，并将reflect.Value转换成interface{}.,再将interface{}转换成float64

```go
package main
 import (
	"fmt"
	"reflect"
 )
 func reflect01(b interface{})  {
	num := reflect.ValueOf(b)
    kind1 :=num.Kind()
	iv :=num.Interface()
	fmt.Printf("b的reflect.Value是=%v,kind值为=%v,num转换为interface的值为=%v",num,kind1,iv)
 }
 func main() {
	var n float64 =65.9
	reflect01(n)
 }
```

2）给字符串改名题

```go
        var str string = "tom"
        fs :=reflect.ValueOf(&str) //这里要改成地址
        fs.Elem().SetString("jackma")
        fmt.Println(str)
```

### 6.反射的最佳实践

1）使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值
2）使用反射的方式来获取结构体的tag标签,遍历字段的值，修改字段值，调用结构体方法
val.Method(0).Call() 调用第一个方法
方法的排序默认是按照函数名的排序(ASCII码)
定义了两个函数testl 和test2， 定义一个适配器函数用作统一 处理接口
使用反射操作任意结构体类型
使用反射创建并操作结构体

```go
package main
import (
	"fmt"
	"reflect"
)

//定义了一个Monster结构体
type Monster struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Score float32 `json:"成绩"`
	Sex string
}

// 方法: 返回两数之和
func(s Monster) GetSum(n1, n2 int) int{
	return n1 + n2
}

// 方法: 接收四个值, 给S赋值
func(s Monster) Set(name string, age int, score float32, sex string){
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

//方法，显示s的值
func (s Monster) Print() {
	fmt.Println("--start--")
	fmt.Println(s)
	fmt.Println("--end--")
}

func TestStruct(a interface{}){
	//获取reflect.Type类型
	typ := reflect.TypeOf(a)
	//获取reflect.Value类型
	val := reflect.ValueOf(a)
	//获取到a对应的类别
	kd := val.Kind()
	//如果传入的不是struct,就退出
	if kd != reflect.Struct {
		fmt.Println("Expect struct")
		return
	}
	//获取到该结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num) //4
	//变量结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d:值为=%v\n", i, val.Field(i))
		//获取到struct 标签,注意需要通过reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		//如果该字段于tag标签就显示，否则就不显示
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}
	//获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	//var params []reflect.Value
	//方法的排序默认是按照函数名的排序(ASCII码)
	val.Method(1).Call(nil)//获取到第二个方法。调用它
	
	//调用结构体的第1个方法Method(0)
	var params []reflect.Value //声明了[]reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //传入的参数是[]reflect.Value,返回[]reflect.Value
	fmt.Println("res=", res[0].Int()) //返回结果,返回的结果是[]reflect.Value*/
}

func main(){
	var a Monster = Monster{
		Name: "黄鼠狼精",
		Age: 400,
		Score: 30.8,
	} 
	//将Monster实例传递给TestStruct函数
	TestStruct(a)
}

```

2）使用反射操作任意结构体类型

```go
type user struct {
	UserId string
	Name string
 }

 func TestReflectStruct(t *testing.T){
	var {
		model *user
		sv reflect.Value
	}
	model = &user{}
	sv = reflect.ValueOf(model)
	t.Log("reflect.ValueOf",sv.kind().String())
	sv=sv.Elem()
	t.Log("reflect.ValueOf",,sv.kind().String())
	sv.FieldByName("userId").SetString("12345678")
	sv.FieldByName("Name").SetString("nickname")
	t.Log("model",model)
 }
```

## 二十、网络编程

### 1、网络编程的基本介绍

Golang的主要设计目标之一就是面向大规模后端服务程序，网络通信这块是服务端程序必不可少也是至关重要的一部分

网络编程有两种

1）TCP socket编程，是网络编程的主流。之所以交TCP socket编程，是因为底层是基于Tcp/ip协议的，比如QQ聊天

2）b/s结构的http编程。我们使用浏览器去访问服务器时，使用的就是http协议，而http底层依旧是tcp socke实现的，比如京东商城(这属于go web开发范畴)

### 2.网络编程的基础知识

#### 1）协议(tcp/ip)

TCP/IP(Transmission Control Protocal)的简写，中文译名为传输控制协议/因特网互联协议，又叫网络通信协议，这个协议是Internet最基本的协议 internet国际互联网的基础，简单的说，就是由网络层的IP协议和传输层的TCP协议组成的

协议的抽象理解图

![](D:\myfile\后端\go语言学习\pic\网络编程\pic1.jpg)

#### 2）OSI与TCP/ip参考模型

![](D:\myfile\后端\go语言学习\pic\网络编程\pic2.jpg)

详细结构（模拟qq好友发送数据经过的网络协议层）

![](D:\myfile\后端\go语言学习\pic\网络编程\pic3.jpg)

#### 3）ip地址

概述：每个internet上的主机和路由器都有一个ip地址，他包括网络和主机号，ip地址有ipv4(32位)和ipv6(128位)，可以通过ipconfig来查看

#### 4）端口(port)介绍

我们这里所指的端口不是指物理意义上的端口，而是特指TCP/IP协议上的端口，是逻辑意义上的端口

如果把ip地址比作一个房子，端口就是出入这间房子的门。真正的房子只有几个门，但是一个ip地址的端口可以有65536(256*256)个之多！端口是通过端口号来标记的，端口号只有整数，范围是0到65535（256*256-1）

**端口--分类**

- 0是保留端口
- 1-1024是固定端口 又叫有名端口，即被某些程序固定使用，一般程序员不使用， 22：SSH远程登录协议 23：telnet使用 21：Ftp使用 25：smtp服务使用 80：iis使用 7：echo服务
- 1025-65535是动态端口这些端口，程序员是可以使用的（尽量使用40000以上的端口，这样不易冲突，更好地进行监听操作）

![](D:\myfile\后端\go语言学习\pic\网络编程\pic4.jpg)

端口(port)-使用注意

- 在计算机（尤其是做服务器）要尽可能的少开端口
- 一个端口只能被一个程序监听
- 如果使用netstat -an 可以查看本机有哪些端口在监听
- 可以使用netstat -anb来查看监听端口的pid,在结合任务管理器关闭不安全的端口

#### 5）tcp socket编程的客户端和服务器端

下图为Golang socket编程中客户端和服务器的网络分布

![](D:\myfile\后端\go语言学习\pic\网络编程\pic5.jpg)

### 3.socket编程快速入门

项目示意图

![](D:\myfile\后端\go语言学习\pic\网络编程\pic6.jpg)

1）服务端的处理流程

-1.监听端口

-2.创建客户端的tcp连接，建立客户端和服务端的连接

-3.创建goroutine，处理该连接的请求（通常客户端会通过连接来发送请求包）

server.go

```go
package main
import (
	"fmt"
	"net" //做网络socket开发时。net包含有我们需要所有的方法和函数
	// "io"
)

func process(conn net.Conn) {
	//这里我们循环的接收客户端发送的数据
	defer conn.Close() //关闭conn

	for {
		//创建一个新的切片
		buf := make([]byte,1024)
		//1.等待客户端通过conn发送信息
		//2.如果客户端没有write[发送]，那么协程就阻塞在这里
		fmt.Printf("服务器在等待客户端%s 发送信息"+ conn.RemoteAddr().String())
		n,err :=conn.Read(buf) //从conn读取
		if err != nil {
			fmt.Printf("客户端退出 err=%v\n",err)
			return // !!!
		}
		//3.显示客户端发送的内容到服务器的终端
		fmt.Println(string(buf[:n]))
		
	}
}
func main() {
	fmt.Println("服务器开始监听")
	//net.Listen("tcp","0.0.0.0:8888")
	//1.tcp表示使用网络协议是tcp
	//2.0.0.0.0:8888 表示在本地监听8888端口
	listen, err :=net.Listen("tcp","0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err")
		return
	}
	defer listen.Close()  //延时关闭listen

	//循环等待客户端连接我
	for {
		//等待客户端连接诶
	    fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=",err)
		}else {
			fmt.Println("Accept() suc conn=%v 客户端ip为=%v\n",conn,conn.RemoteAddr().String())
		}
		//这里准备启动一个协程为客户端服务

		go process(conn)
	}
	// fmt.Printf("Listen successfully=%v\n",listen)
}
```

2）客户端的处理流程

-1.建立与服务端的链接

-2.发送请求数据，接收服务器端返回的结果数据

-3.关闭连接

3）客户端功能

-1.编写一个客户端程序，能连接到服务端的8888端口

-2.客户端可以发送单行数据，然后就退出

-3.能通过终端输入数据（输入一行发送一行），并发送给服务器端

-4.在终端输入exit表示退出程序

client.go

```go
package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main() {
	conn, err :=net.Dial("tcp","192.168.31.102:8888")
	if err != nil {
		fmt.Println("client dial err=",err)
		return
	}
	//功能1.客户端可以发送单行数据，然后就退出
	reader:= bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入【终端】
    for {
	//从终端读取一行用户输入，并准备发送给服务器
	line, err :=reader.ReadString('\n')
	if err != nil {
		fmt.Println("readerString err=",err)
	}
	//如果用户输入的是exit就退出
	line = strings.Trim(line, "\r\n") 
	if line == "exit"{
		fmt.Println("客户端退出..")
		break
	}

	//再将line发送给服务器
	_, err =conn.Write([]byte(line+ "\n"))
	if err != nil {
		fmt.Println("conn.Write err=",err)
	}
	//fmt.Printf("客户端发送了%d字节的数据，并退出",n)
   }


}
```

运行效果图

![](D:\myfile\后端\go语言学习\pic\网络编程\pic8.jpg)

### 4.经典项目-海量用户即时通讯系统

#### 1）项目开发流程

需求分析 -->设计阶段-->编码实现-->测试阶段--->实施

#### 2）需求分析

-1 用户注册

-2 用户登录

-3 显示在线用户列表

-4 群聊（广播）

-5 点对点聊天（私聊）

-6 离线留言

#### 3）界面设计

![](D:\myfile\后端\go语言学习\pic\网络编程\pic9.jpg)

项目开发前技术准备

项目要保存用户信息和消息数据，因此我们需要学习数据库(Redis和mysql),这里我们选择redis,先学习如何在golang中使用redis

## 二十一、redis的学习与使用

### 1.Redis的基本介绍

1）.Redis是NoSQL数据库，不是传统的关系型数据库

官网: http://redis.io/ 和http://www.redis.cn/

2）.Redis：REmote DIctionary Server(远程字典服务器)，Redis的性能非常高，单机可以达到15W gps.通常适合做缓存，也可以持久化

3）是完全开源免费的，高性能的(key/value)分布式内存数据库，基于内存运行并支持持久化Nosql数据库，是最热门的NoSql数据库之一，也称为数据结构服务器

### 2.Redis的安装下载安装包即可

网址：https://github.com/tporadowski/redis/releases

下载好后的目录结构，其中有两个是redis的关键文件

redis客户端文件：**redis-cli.exe**

redis服务器文件：**redis-server.exe**

![](D:\myfile\后端\go语言学习\pic\网络编程\pic11.jpg)



**redis结构示意图**

![](D:\myfile\后端\go语言学习\pic\网络编程\pic10.jpg)

### 3.Redis的基本使用

#### 1）Redis的启动：

启动Redis的服务器端程序(redis-server.exe),直接双击即可运行

![](D:\myfile\后端\go语言学习\pic\网络编程\pic12.jpg)

#### 2）Redis的操作的三种方式

![](D:\myfile\后端\go语言学习\pic\网络编程\pic13.jpg)

```
说明：
1）使用telnet操作Redis
2）使用redis-cli.exe操作Redis,双击即可进入客户端
3）Golang操作redis
```

redis基本指令

http://redisdoc.com(已失效)

![](D:\myfile\后端\go语言学习\pic\网络编程\pic14.jpg)

![](D:\myfile\后端\go语言学习\pic\网络编程\pic15.jpg)

#### 3）说明：Redis安装好后，默认有16个数据库，初始默认使用0号库，编号0...15

-1.添加key-val [set]

-2.查看当前redis的所有key [key * 获取key对应的值[get key]

-3.切换redis数据库[select index]

-4.如何查看当前数据库的key-val数量[dbsize]

![](D:\myfile\后端\go语言学习\pic\网络编程\pic16.jpg)

 -5.清空当前数据库的key-val和清空所有数据库的key -val[flush db flushall]

### 4.Redis的五大数据类型和CRUD操作

#### 1）Redis的五大数据类型是：**String(字符串)**、**Hash(哈希)**、**List(列表)**、**Set(集合)**和**zset(sorted set :有序集合)**

#### 2）String

(字符串)-介绍

-1.string是redis的基本类型，一个key对应一个value

-2.string是类型是二进制安全的。除普通的字符串外，也可以存放图片数据。

-3.redis中字符串value最大值是512M

举例，存放一个地址信息:

address北京天安门

说明：

key :address

value : 北京天安门

```redis
127.0.0.1:6379> set address beijing
OK
127.0.0.1:6379> get address
"beijing"
127.0.0.1:6379>

```

-4.String（字符串）-CRUD

举例说明Redis的String字符串的CRUD操作

set如果存在就相当于修改，不存在就是添加 /get/del

```go
127.0.0.1:6379> del address
(integer) 1
127.0.0.1:6379> get address
(nil)
127.0.0.1:6379>
```

-5.String(字符串)-使用细节和注意事项

- **setex(set with expire)键秒值**

```
127.0.0.1:6379> setex mess01 10 hello,you
OK
127.0.0.1:6379> get mess01
"hello,you"
127.0.0.1:6379> get mess01
"hello,you"
127.0.0.1:6379> get mess01
"hello,you"
127.0.0.1:6379> get mess01
"hello,you"
127.0.0.1:6379> get mess01
(nil)  //mess01有效期就10秒，10秒后就销毁
127.0.0.1:6379>
```

- **mset[同时设置一个或多个key-value对]**

​           MSET key vlaue [key value ...]

​          同时设置一个或多个key-value对

​          如果某个给定key值已经存在，那么MSET会用新值覆盖原来的旧值

```
127.0.0.1:6379> mset worker01 tom worker02 scott
OK
127.0.0.1:6379> get worker01
"tom"
127.0.0.1:6379> get worker02
"scott"
127.0.0.1:6379> mget worker01 worker02
1) "tom"
2) "scott"
```

#### 3）Hash

(哈希，类似于golang里的Map)-介绍

Redis hash是一个键值对象集合。var user1 map[string]string

Redis hash是一个string类型的field和value的映射表，hash特别适合用于存储对象

举例。存放一个User信息

```
user1 name  张三 age 10
说明
key : user1
name 张三 和 age 30 就是两对 field-value
```

实操：演示添加user的信息案例(name age job)

```
127.0.0.1:6379> hset user1 name "smith"
(integer) 1
127.0.0.1:6379> hset user1 age 30
(integer) 1
127.0.0.1:6379> hset user1 job "golang coder"
(integer) 1
127.0.0.1:6379> hget user1 name
"smith"
127.0.0.1:6379> hget user1 age
"30"
127.0.0.1:6379> hget user1 job
"golang coder"
127.0.0.1:6379>
```

 Hash使用细节和注意事项

在给user设置name和age时，前面我们是一步一步设置，使用hmset和hmget可以一次性来设置多个field的值和返回多个field的值

```
127.0.0.1:6379> hmset user2 name jerry age 110 job "java coder"
OK
127.0.0.1:6379> hmget user2 name age job
1) "jerry"
2) "110"
3) "java coder"
127.0.0.1:6379>
```

**hlen统计一个hash有几个元素**

```
127.0.0.1:6379> hlen user2
(integer) 3
127.0.0.1:6379>
```

hexists key field

查看哈希表key中，给定域field是否存在

```
127.0.0.1:6379> hexists user2 name
(integer) 1
127.0.0.1:6379>
```

#### 4）List

-1.(列表)-介绍

列表是简单的字符串列表，按照插入顺序排序，你可以添加一个元素到列表头部（左边）或者尾部（右边）

List的本质是一个链表，List的元素的有序的，元素的值可以重复

举例，存放多个地址信息

```
city 北京 天津 上海
说明
北京 天津 上海 就是三个元素
```

```
127.0.0.1:6379> lpush city beijing shanghai tianjing
(integer) 3
127.0.0.1:6379> lrange city 0 -1（0表示开头 -1表示尾部）
1) "tianjing"
2) "shanghai"
3) "beijing"
127.0.0.1:6379>
```

-2.关于list的基本命令介绍

lpush:从左边插入数据

rpush:从右边插入数据

lrange:  用法 lrange key start stop :返回列表key中指定区间内的元素，区间以偏移量start和stop指定,下标(index)参数start和stop都是以0为底，也就是说，以0表示列表的第一个元素，以1表示列表的第二个元素，以此类推。你也可以使用函数负数下标，以-1表示列表的最后一个元素，-2表示列表的倒数第二个元素，以此类推

lpop :从我们列表的左边弹出一个数据

rpop :从我们列表的右边弹出一个数据

del : 删掉这个列表

实战演示

```
127.0.0.1:6379> lpush herolist aa bb cc
(integer) 3
127.0.0.1:6379> lrange herolist 0 -1
1) "cc"
2) "bb"
3) "aa"
127.0.0.1:6379> rpush herolist dd eee
(integer) 5
127.0.0.1:6379> lrange herolist 0 -1
1) "cc"
2) "bb"
3) "aa"
4) "dd"
5) "eee"
127.0.0.1:6379> lpop herolist
"cc"
127.0.0.1:6379> lrange herolist 0 -1
1) "bb"
2) "aa"
3) "dd"
4) "eee"
127.0.0.1:6379> rpop herolist
"eee"
127.0.0.1:6379> lrange herolist 0 -1
1) "bb"
2) "aa"
3) "dd"
127.0.0.1:6379> del herolist
(integer) 1
127.0.0.1:6379> lrange herolist 0 -1
(empty list or set)
127.0.0.1:6379> 
```

-3.List-使用细节和注意事项

（1）index，按照索引下标获得元素（从左到右，编号从0开始）

（2）LLEN key ，返回列表key的长度，如果key不存在，则key被解释为一个空列表，返回0

（3）List的其他说明

- List数据，可以从左或者右，插入添加
- 如果值全移除，对应的键也就消失了

#### 5）Set

(集合)-介绍

Redis的Set是string类型的无序集合。

底层是HashTble数据结构，Set也是存放很多字符串元素，字符串元素是无序的，而且元素的值不能重复

举例，存放多个邮件列表信息

```
email sggg@sohu.com tom@sohu.com
说明
key :email
sggg@sohu.com tom@sohu.com 这两个就是元素

```

关于set的基本命令

sadd :添加元素

smembers ：取出所有值

sismember ：判断是否存在该成员

srem ：删除指定值

```
127.0.0.1:6379> sadd emails tom@sohu.com jack@qq.com
(integer) 2
127.0.0.1:6379> smembers emails
1) "tom@sohu.com"
2) "jack@qq.com"
127.0.0.1:6379> sadd emails kk@yy.com uu@qq.com
(integer) 2
127.0.0.1:6379> smembers emails
1) "uu@qq.com"
2) "tom@sohu.com"
3) "jack@qq.com"
4) "kk@yy.com"
127.0.0.1:6379> sadd emails jack@qq.com
(integer) 0  //0表示添加失败，因为set的元素是不可以重复的
127.0.0.1:6379> sismember emails tom@sohu.com
(integer) 1 //查看set里面是否含有tom@souhu.com这个元素 1代表存在 0代表不存在
127.0.0.1:6379> sismember emails tom~@sohu.com
(integer) 0
127.0.0.1:6379> srem emails tom@sohu.com
(integer) 1 //演示删除元素
127.0.0.1:6379> smembers emails
1) "kk@yy.com"
2) "uu@qq.com"
3) "jack@qq.com"
127.0.0.1:6379>

```

## 二十二、go操作redis

### 1、安装第三方开源的redis库

1）使用第三方开源的redis库：github.com/garyburd/redigo/redis

2）在使用redis前，先安装第三方Redis库，在GOPATH路径下执行安装指令：

```
D:\myfile\go\project>go get github.com/garyburd/redigo/redis
```

3）安装之后，可以看到如下包

![](D:\myfile\后端\go语言学习\pic\网络编程\pic17.jpg)

特别说明：在安装Redis 库之前，确保已经安装并配置了git,因为是从github上下载安装库的，需要使用到git

### 2、使用go语言开始进行操作

Set/Get接口

#### 1）操作string

说明：通过golang添加和获取key-value【比如name-tom~】

```go
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
```

#### 2）操作hash

说明：通过Golang对redis操作Hash数据类型

```go
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
```

#### 3）批量操作Set/Get数据

说明：通过Golang对Redis操作，一次操作可以Set/Get多个key-val数

```go
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
```

#### 4）给数据设置有效时间

说明：通过golang对redis操作，给key-value设置有效时间

core code

```
//给name数据设置有效时间10s
_,err =c.Do("expire","name",10)
```

5）操作List

说明：通过golang对redis操作list数据类型

corecode

```
_,err=c.Do("lpush","herolist","no1.宋江",30,"no2.吴用",28)
```

#### 5）课堂练习

![](D:\myfile\后端\go语言学习\pic\网络编程\pic19.jpg)

```go
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
```

#### 6）Redis连接池

说明：通过Golang对redis的操作，还可以通过redis连接池，流程如下

-1.事先初始化一定数量的连接，放入到连接池

-2.当go需要操作redis时，直接从redis连接池取出连接即可

-3.这样可以节省临时获取redis连接的时间，从而提高效率

core code

```go
var pool *redis.Pool
pool = &redis.Pool{
   MaxIdle:8
   MaxActive:0
   IdleTimeout:100
   Dial:func()(redis.Conn,error){
   	return redis.Dial("tcp","localhost:6379")
   },
}
c :=pool.Get()
```

-4.示意图

![](D:\myfile\后端\go语言学习\pic\网络编程\pic21.jpg)

```go
package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
//定义一个全局的pool
var pool *redis.Pool

//当启动程序时就初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:8, //最大空闲连接数
		MaxActive:0,//表示和数据库的最大连接数，0表示没有限制
		IdleTimeout:100,//最大空闲时间
		Dial:func()(redis.Conn,error){//初始化连接的代码。连接哪个ip
			return redis.Dial("tcp","localhost:6379")
		},
	 }
}
func main() {
	//先从pool池取出一个连接
	conn := pool.Get()
	defer conn.Close()

	_, err :=conn.Do("Set","name","汤姆猫~~")
	if err != nil {
		fmt.Println("conn.Do err=",err)
		return
	}

	//取出
	r,err :=redis.String(conn.Do("Get","name"))
	if err != nil {
		fmt.Println("conn.Do err=",err)
		return
	}

	fmt.Println("r=",r)


}
```

## 二十三、海量用户即时通讯系统

### 1、项目开发前技术准备

项目要保存用户信息和信息数据，因此我们需要学习数据(redis或者mysql),这里我们选择redis

### 2.实现功能-显示客户端登录菜单

![](D:\myfile\后端\go语言学习\pic\网络编程\pic22.jpg)

代码编写

clien包下的main.go

```go
package main
import (
	"fmt"
	"os"
)

//定义两个变量，一个表示用户的id,一个表示用户的密码
var userId int
var userPwd string

func main() {
	//接收用户的选择
    var key int
	//判断是否还继续显示菜单
	var loop = true
    
	for loop{
		fmt.Println("-----------欢迎登录多人聊天系统------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择 1-3：")

		fmt.Scanf("%d\n",&key)
		switch key {
		case 1 :
			fmt.Println("登录聊天室")
			loop=false
		case 2 :
			fmt.Println("注册用户")	
			loop=false
		case 3 :
			fmt.Println("退出系统")	
			//loop=false
			os.Exit(0)
		default:
			fmt.Println("输入有误，请输入1-3")	
		}
	}

	//根据用户的输入，显示新的提示信息
	if key ==1 {
		//说明用户要登录了
		fmt.Println("请输入用户的id")
		fmt.Scanf("%d\n",&userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%s\n",&userPwd)
		//先把登录函数，写到另外一个文件,先写login.go
		err := login(userId,userPwd)
		if err != nil {
			fmt.Println("登录失败")
		}else {
			fmt.Println("登录成功")
		}
	}else if key ==2 {
		fmt.Println("进行用户注册的逻辑....")

	}

}
```

clien包下的login.go

```go
package main
import (
	"fmt"
)
//写一个函数，完成登录操作
func login(userId int,userPwd string) (err error) {

	//下一个就要开始定协议
	fmt.Printf("userId = %d userPwd = %s\n",userId,userPwd)
	return nil
}
```

### 3.实现功能-完成用户登录

要求：完成指定用户的验证，用户id=100，密码pwd=123456可以登录，其他用户不能登录

理解从client到server中的程序执行流程，如图所示【Message组成的示意图。并发送一个message的流程介绍】

![](D:\myfile\后端\go语言学习\pic\网络编程\pic23.jpg)

#### -1.完成客户端可以该长度值发送消息长度，服务器端可以正常接收到

分析思路

1）先确定消息Message的格式

2）发送消息示意图

![](D:\myfile\后端\go语言学习\pic\网络编程\pic24.jpg)

代码展示

sever

main.go

```go
package main
import (
	"fmt"
	"net"
)

//处理和客户端的通讯
func process(conn net.Conn){
    //这里需要延时关闭
	defer conn.Close()
	//循环地读客户端发送的信息
	for {
		buf := make([]byte,8096)
		fmt.Println("读取客户端发送的数据...")
		n, err :=conn.Read(buf[:4])
		if n != 4 || err !=nil {
			fmt.Println("conn.Read err=",err)
			return
		}
		fmt.Println("独到的buf=",buf[:4])
	}

}
func main() {
	//提示信息
	fmt.Println("服务器在8889端口监听....")
	listen, err := net.Listen("tcp","0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=",err)
		return
	} 
	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=",err)
		} 

		//一旦连接成功，则则启动一个协程和客户端保持通讯。。
		go process(conn)
	}
}
```

common层的message

message.go

```go
package message

const (
	LoginMesType   = "LoginMes"
	LoginResMesType  = "LoginResMes"
)

type Message struct {
	Type string `json:"type"`//消息的类型
	Data string `json:"data"`//消息的数据
}

//定义两个消息。。后面需要再添加
type LoginMes struct {
	UserId int `json:"userId"`//用户Id
	UserPwd string `json:"userPwd"`//用户密码
	UserName string `json:"userName"`//用户名
}

type LoginResMes struct {
	Code int `json:"code"`//返回状态码 500表示该用户未注册 200表示登录成功
	Error string `json:"error"`//返回错误信息
}
```

client层

login.go

```go

package main
import (
	"fmt"
	"net"
	"encoding/json"
	"encoding/binary"
	"go_code/chatroom/common/message"
)
//写一个函数，完成登录操作
func login(userId int,userPwd string) (err error) {

	//下一个就要开始定协议
	// fmt.Printf("userId = %d userPwd = %s\n",userId,userPwd)
	// return nil

	//1.连接到服务器端
	conn, err :=net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=",err)
		return
	}

	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType 

	//3.创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd 

	//4.将loginMes序列化
	data, err :=json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Mashal err=",err)
		return
	}
	//5.将data赋给了mes.Data字段
	mes.Data = string(data)
	//6.将mes进行序列化
	data, err =json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Mashal err=",err)
		return
	}
	//7.到这个时候，data就是我们要发送的消息
    //7.1先把data的长度发送给服务器
	//先获取data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen) //将该、长度转成了byte类型是数据
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err !=nil {
		fmt.Println("connWrite(buf) fail ",err)
		return
	}

	fmt.Printf("客户端发送数据的消息长度=%d 内容是=%s",len(data),string(data))
	return


}
```

main.go

```go
package main
import (
	"fmt"
	"os"
)

//定义两个变量，一个表示用户的id,一个表示用户的密码
var userId int
var userPwd string

func main() {
	//接收用户的选择
    var key int
	//判断是否还继续显示菜单
	var loop = true
    
	for loop{
		fmt.Println("-----------欢迎登录多人聊天系统------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择 1-3：")

		fmt.Scanf("%d\n",&key)
		switch key {
		case 1 :
			fmt.Println("登录聊天室")
			loop=false
		case 2 :
			fmt.Println("注册用户")	
			loop=false
		case 3 :
			fmt.Println("退出系统")	
			//loop=false
			os.Exit(0)
		default:
			fmt.Println("输入有误，请输入1-3")	
		}
	}

	//根据用户的输入，显示新的提示信息
	if key ==1 {
		//说明用户要登录了
		fmt.Println("请输入用户的id")
		fmt.Scanf("%d\n",&userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%s\n",&userPwd)
		//先把登录函数，写到另外一个文件,先写login.go
		err := login(userId,userPwd)
		if err != nil {
			fmt.Println("登录失败")
		}else {
			fmt.Println("登录成功")
		}
	}else if key ==2 {
		fmt.Println("进行用户注册的逻辑....")

	}

}
```

#### -2.完成客户端可以发送消息，服务器端可以接收到消息并根据客户端发送的消息判断用户的合法性并返回相应的消息

思路分析

1）让客户端发送消息本身

2）服务器端接收到消息，然后反序列化成对应的消息结构体

3）服务器端根据反序列化的消息，判断是否登录用户是合法，返回LoginReMes

4）客户端解析返回的LoginReMes,显示对应界面

5）这里我们需要做一些函数的封装

cient/login.go在结尾添加这些coding

```go
//发送消息本身
	_, err = conn.Write(data)
	if err !=nil {
		fmt.Println("connWrite(data) fail ",err)
		return
	}
	//休眠20秒
	time.Sleep(10 * time.Second)
	fmt.Println("休眠了20秒..")
	//这里还需要处理服务器端返回的消息
	return
```

在server/main.go中我们做了以下改动

将读数据的过程封装了一个函数

```go
package main
import (
	"fmt"
	"net"
	"encoding/json"
	"encoding/binary"
	"go_code/chatroom/common/message"
	//"errors"
	"io"
)
func readPkg(conn net.Conn)(mes message.Message,err error){
	buf := make([]byte,8096)
		fmt.Println("读取客户端发送的数据...")
		//conn.Read()只有在conn没有被关闭的情况下，才会阻塞
		//如果客户端关闭conn则，就不会阻塞
		_, err =conn.Read(buf[:4]) //read出buf中的数据
		if err !=nil {
			//fmt.Println("conn.Read err=",err)
			//err = errors.New("read pkg header error")
			return
		}

		//根据buf[:4]转成uint32类型
		var pkgLen uint32
		pkgLen=binary.BigEndian.Uint32(buf[0:4])

		//根据pkgLen读取消息内容
		n, err :=conn.Read(buf[:pkgLen])
		if n != int(pkgLen) || err !=nil {
			//err = errors.New("read pkg body error")
			return
		}

		//把pkgLen 反序列化成 -->message.Message
		//技术就是一层窗户纸
		json.Unmarshal(buf[:pkgLen],&mes)
		if err != nil {
			fmt.Println("json.Unmarshal err=",err)
			return
		}

		return
}

//处理和客户端的通讯
func process(conn net.Conn) {
    //这里需要延时关闭
	defer conn.Close()
	
	//循环地读客户端发送的信息
	for {
		//这里我们将读取数据包，直接封装成一个函数readPkg(),返回Message,Err
		mes, err :=readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出...")
				return
			}else {
				fmt.Println("readpkg err=",err)
			}
			return
		}
		fmt.Println("mes=",mes)
	}
}
//main函数下的则没有改变
func main() {
	//提示信息
	fmt.Println("服务器在8889端口监听....")
	listen, err := net.Listen("tcp","0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=",err)
		return
	} 
	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=",err)
		} 

		//一旦连接成功，则则启动一个协程和客户端保持通讯。。
		go process(conn)
	}
}
```

#### -3.能够完成登录，并提示信息

server/main.go

```go

添加了发送信息给客户端的代码
func writePkg(conn net.Conn,data []byte)(err error) {

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen) //将该、长度转成了byte类型是数据
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err !=nil {
		fmt.Println("connWrite(buf) fail ",err)
		return
	}

	//发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err !=nil {
		fmt.Println("connWrite(data) fail ",err)
		return
	}
	return

}

//编写一个函数serverProcessLogin函数，专门处理登录请求
func serverProcessLogin(conn net.Conn,mes *message.Message)(err error){
	//核心代码
	//1.先从mes中取出mes.Data,并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err =json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=",err)
		return
	}

	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type=message.LoginResMesType

	//2.再声明一个LoginResMes
	var loginResMes message.LoginResMes

	//如果用户的id=100,密码=123456认为合法，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
	} else {
		//不合法
		loginResMes.Code = 500  //500状态码表示用户不存在
		loginResMes.Error = "该用户不存在，请注册再使用。。。"
	}

	//3.将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("Marshal fail err=",err)
	}

	//4.将data赋值给resMes
	resMes.Data = string(data) 

	//5.对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("Marshal fail err=",err)
		return
	}
	//6.发送data 我们将其封装为writePkg
    err = writePkg(conn,data)
	return

}


//编写一个ServerProcessMes函数
//功能 ：根据客户端发送的消息种类不同，决定调用哪个函数处理
func serverProcessMes(conn net.Conn,mes *message.Message)(err error) {
	switch mes.Type {
	case message.LoginMesType :
		//处理登录的逻辑
		err = serverProcessLogin(conn,mes)
	case message.RegisterMesType :
		//处理注册
	default :
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}


在process进行了改定
//处理和客户端的通讯
func process(conn net.Conn) {
    //这里需要延时关闭
	defer conn.Close()
	
	//循环地读客户端发送的信息
	for {
		//这里我们将读取数据包，直接封装成一个函数readPkg(),返回Message,Err
		mes, err :=readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出...")
				return
			}else {
				fmt.Println("readpkg err=",err)
			}
			return
		}
        //增加了这段代码进行调用这个函数
		err = serverProcessMes(conn,&mes)
		if err != nil {
			return
		}
	}
}
```

client/utils(增加了一个utils.go用于read的write的操作)

```go
package main
import (
	"fmt"
	"net"
	"encoding/json"
	"encoding/binary"
	"go_code/chatroom/common/message"
)

func readPkg(conn net.Conn)(mes message.Message,err error){
	buf := make([]byte,8096)
		fmt.Println("读取客户端发送的数据...")
		//conn.Read()只有在conn没有被关闭的情况下，才会阻塞
		//如果客户端关闭conn则，就不会阻塞
		_, err =conn.Read(buf[:4]) //先读取之前发送的数据长度
		if err !=nil {
			//fmt.Println("conn.Read err=",err)
			//err = errors.New("read pkg header error")
			return
		}

		//根据buf[:4]转成uint32类型
		var pkgLen uint32
		pkgLen=binary.BigEndian.Uint32(buf[0:4])

		//根据pkgLen（data数据的长度）读取消息内容
		n, err :=conn.Read(buf[:pkgLen])
		if n != int(pkgLen) || err !=nil {
			//err = errors.New("read pkg body error")
			return
		}

		//把pkgLen 反序列化成 -->message.Message
		//技术就是一层窗户纸
		json.Unmarshal(buf[:pkgLen],&mes)
		if err != nil {
			fmt.Println("json.Unmarshal err=",err)  //json的反序列化失败!
			return
		}

		return
}

func writePkg(conn net.Conn,data []byte)(err error) {

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen) //将该、长度转成了byte类型是数据
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err !=nil {
		fmt.Println("connWrite(buf) fail ",err)
		return
	}

	//发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err !=nil {
		fmt.Println("connWrite(data) fail ",err)
		return
	}
	return

}
```

client/login.go

```go
//在末尾加入了如下的代码
//这里还需要处理服务器端返回的消息
	mes, err = readPkg(conn) //mes 就是

	if err != nil {
		fmt.Println("readPkg(conn) err=",err)
		return
	}
	
	//将mes的Data部分反序列化为LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	}else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return


}
```

#### -4.程序结构的改进

说明：前面的程序虽然完成了功能，但是没有结构，系统的可读性、拓展性和维护性都不好，因此需要对程序的结构进行改进

##### 1）画出程序框架图

![](D:\myfile\后端\go语言学习\pic\网络编程\pic25.jpg)

##### 2）步骤

（1）先把分析出来的文件，创建好，然后放到相应的文件夹中

###### server层后端项目结构图

![](D:\myfile\后端\go语言学习\pic\网络编程\pic26.jpg)

（2）现在根据各个文件完成的任务和作用不同，将main.go的代码剥离到对应的文件即可

（3）先修改了utils.go

```go
package utils
 import (
	"fmt"
	"net"
	"encoding/json"
	"encoding/binary"
	"go_code/chatroom/common/message"
 )

 //将这些方法关联到结构体当中
 type Transfer struct {
	//分析应该有哪些字段
	Conn net.Conn
	Buf [8096]byte  //这是传输时使用缓冲
 }

 func (this *Transfer) ReadPkg()(mes message.Message,err error){
		fmt.Println("读取客户端发送的数据...")
		//conn.Read()只有在conn没有被关闭的情况下，才会阻塞
		//如果客户端关闭conn则，就不会阻塞
		_, err =this.Conn.Read(this.Buf[:4]) //先读取之前发送的数据长度
		if err !=nil {
			//fmt.Println("conn.Read err=",err)
			//err = errors.New("read pkg header error")
			return
		}

		//根据buf[:4]转成uint32类型
		var pkgLen uint32
		pkgLen=binary.BigEndian.Uint32(this.Buf[0:4])

		//根据pkgLen（data数据的长度）读取消息内容
		n, err :=this.Conn.Read(this.Buf[:pkgLen])
		if n != int(pkgLen) || err !=nil {
			//err = errors.New("read pkg body error")
			return
		}

		//把pkgLen 反序列化成 -->message.Message
		//技术就是一层窗户纸
		json.Unmarshal(this.Buf[:pkgLen],&mes)
		if err != nil {
			fmt.Println("json.Unmarshal err=",err)  //json的反序列化失败!
			return
		}

		return
}

func (this *Transfer) WritePkg(data []byte)(err error) {

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4],pkgLen) //将该、长度转成了byte类型是数据
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err !=nil {
		fmt.Println("connWrite(this.Buf) fail ",err)
		return
	}

	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err !=nil {
		fmt.Println("connWrite(data) fail ",err)
		return
	}
	return

}

```

（4）修改了process2/userProcess.go

```go
package process2
import (
	"fmt"
	"net"
	"encoding/json"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
)

type UserProcess struct {
	//字段
	Conn net.Conn
}

//编写一个函数serverProcessLogin函数，专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message)(err error){
	//核心代码
	//1.先从mes中取出mes.Data,并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err =json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=",err)
		return
	}

	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type=message.LoginResMesType

	//2.再声明一个LoginResMes
	var loginResMes message.LoginResMes

	//如果用户的id=100,密码=123456认为合法，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
	} else {
		//不合法
		loginResMes.Code = 500  //500状态码表示用户不存在
		loginResMes.Error = "该用户不存在，请注册再使用。。。"
	}

	//3.将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("Marshal fail err=",err)
	}

	//4.将data赋值给resMes
	resMes.Data = string(data) 

	//5.对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("Marshal fail err=",err)
		return
	}
	//6.发送data 我们将其封装为writePkg
	//因为使用了分层模式(mvc),我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
    err = tf.WritePkg(data)
	return
}
```

（5）修改了main/processor.go

```go
package main
import (
	"fmt"
	"net"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"go_code/chatroom/server/process"
	"io"
)

//先创建一个Processor的结构体
type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes函数
//功能 ：根据客户端发送的消息种类不同，决定调用哪个函数处理
func (this *Processor) serverProcessMes(mes *message.Message)(err error) {
	switch mes.Type {
	case message.LoginMesType :
		//处理登录的逻辑
		//创建一个UserProcess实例
		up := &process2.UserProcess{
			Conn : this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType :
		//处理注册
	default :
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) process2()(err error){
		//循环地读客户端发送的信息
		for {
			//这里我们将读取数据包，直接封装成一个函数readPkg(),返回Message,Err
			//创建一个Transfer实例完成读包任务
			tf := &utils.Transfer{
				Conn : this.Conn,
			}
			mes, err :=tf.ReadPkg()
			if err != nil {
				if err == io.EOF {
					fmt.Println("客户端退出，服务器端也退出...")
					return err
				}else {
					fmt.Println("readpkg err=",err)
				}
				return err
			}
			err = this.serverProcessMes(&mes)
			if err != nil {
				return err
			}
		}
}
```

修改了main/main.go

```go
package main
import (
	"fmt"
	"net"
)

//处理和客户端的通讯
func process(conn net.Conn) {
    //这里需要延时关闭
	defer conn.Close()
	
	//这里调用总控，创建一个processor实例
	processor := &Processor{
		Conn : conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯的协程错误=err",err)
		return 
	}
}
func main() {
	//提示信息
	fmt.Println("服务器[新的结构]在8889端口监听....")
	listen, err := net.Listen("tcp","0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=",err)
		return
	} 
	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=",err)
		} 

		//一旦连接成功，则则启动一个协程和客户端保持通讯。。
		go process(conn)
	}
}
```

修改客户端。先画出程序的框架图，再写代码

###### client层后端项目结构图

![](D:\myfile\后端\go语言学习\pic\网络编程\pic27.jpg)

（2）先把各个文件放到对应的文件夹[包]

![](D:\myfile\后端\go语言学习\pic\网络编程\pic28.jpg)

（3）将server/utils.go拷贝到client/utils/utils.go

（4）创建了client/process/userProcess.go

```go
package process
import (
	"fmt"
	"net"
	"encoding/json"
	"encoding/binary"
	"go_code/chatroom/common/message"
	"go_code/chatroom/client/utils"
)
type UserProcess struct {
	//暂时不需要字段
}
//给关联一个用户登录的方法
//写一个函数，完成登录操作
func (this *UserProcess) Login(userId int,userPwd string) (err error) {

	//下一个就要开始定协议
	// fmt.Printf("userId = %d userPwd = %s\n",userId,userPwd)
	// return nil

	//1.连接到服务器端
	conn, err :=net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=",err)
		return
	}

	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType 

	//3.创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd 

	//4.将loginMes序列化
	data, err :=json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Mashal err=",err)
		return
	}
	//5.将data赋给了mes.Data字段
	mes.Data = string(data)
	//6.将mes进行序列化
	data, err =json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Mashal err=",err)
		return
	}
	//7.到这个时候，data就是我们要发送的消息
    //7.1先把data的长度发送给服务器
	//先获取data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen) //将该、长度转成了byte类型是数据
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err !=nil {
		fmt.Println("connWrite(buf) fail ",err)
		return
	}

	//fmt.Printf("客户端发送数据的消息长度=%d 内容是=%s",len(data),string(data))
	//发送消息本身
	_, err = conn.Write(data)
	if err !=nil {
		fmt.Println("connWrite(data) fail ",err)
		return
	}
	//休眠20秒
	// time.Sleep(10 * time.Second)
	// fmt.Println("休眠了20秒..")
	//这里还需要处理服务器端返回的消息
	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn : conn,
	}
	mes, err = tf.ReadPkg() //mes 就是

	if err != nil {
		fmt.Println("readPkg(conn) err=",err)
		return
	}
	
	//将mes的Data部分反序列化为LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200 {
		//fmt.Println("登录成功")

		//这里我们还需要再客户端启动一个协程
		//该协程保持和服务器端的通讯，如果服务器有数据推送给客户端
		//则可以接受并显示在客户端的终端
		go serverProcessMes(conn)
		//1.显示登录成功后的菜单[循环显示]
		for {
			ShowMenu()
		}
	}else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
```

说明：该文件就是在原来login.go做了一个改进，封装到userProcess结构体

（5）创建了server/process/server.go

```go
package process
import (
	"fmt"
	"os"
	"go_code/chatroom/client/utils"
	"net"
)

//显示登录后的界面..
func ShowMenu(){
	fmt.Println("----------恭喜xxx登录成功--------")
	fmt.Println("          1.显示用户在线列表     ")
	fmt.Println("          2.发送消息            ")
	fmt.Println("          3.信息列表            ")
	fmt.Println("          4.退出系统            ")
	fmt.Println("请选择(1-4): ")
	var key int
	fmt.Scanf("%d\n",&key)
	switch key {
		case 1:
			fmt.Println("显示用户在线列表")
		case 2:
			fmt.Println("发送消息")
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("你选择退出系统 ")	
			os.Exit(0)	
		default:
			fmt.Println("你输入的选项不正确")		
	}
}
//和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	//创建一个transfer实例，不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn : conn,
	}
	for {
		fmt.Printf("客户端正在等待读取服务器发送的消息")
		mes, err:=tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=",err)
			return
		}
		//如果读取到消息，又是下一步处理逻辑
		fmt.Printf("mes=%v",mes)


	}

}
```

（6）client/main/main.go

```go
package main
import (
	"fmt"
	"os"
	"go_code/chatroom/client/process"
)

//定义两个变量，一个表示用户的id,一个表示用户的密码
var userId int
var userPwd string

func main() {
	//接收用户的选择
    var key int
	//判断是否还继续显示菜单
	// loop = true
    
	for true{
		fmt.Println("-----------欢迎登录多人聊天系统------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择 1-3：")

		fmt.Scanf("%d\n",&key)
		switch key {
		case 1 :
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n",&userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n",&userPwd)
			//完成登录
			//1.创建一个UserProcess的实例
			up :=&process.UserProcess{}
			up.Login(userId,userPwd)
			//loop=false
		case 2 :
			fmt.Println("注册用户")	
			//loop=false
		case 3 :
			fmt.Println("退出系统")	
			//loop=false
			os.Exit(0)
		default:
			fmt.Println("输入有误，请输入1-3")	
		}
	}
}
```

#### -5.应用redis

##### 1）在Redis手动添加测试用户，并画图+说明注意（后面通过程序注册用户）

![](D:\myfile\后端\go语言学习\pic\网络编程\pic29.jpg)

手动直接在redis增加一个用户信息

![](D:\myfile\后端\go语言学习\pic\网络编程\pic30.jpg)

##### 2）如输入的用户名密码正确在Redis中存在则登录，否则退出系统，并给出相应的提示信息

- 1.用户不存在，你也可以重新注册，再登录
- 2.你的密码不正确

##### 3）代码实现

###### (1)先编写了server/model/user.go

```go
package model

//定义一个用户的结构体
type User struct {
	//确定字段信息
	//为了序列化和反序列化成功
	//用户信息的json字符串与结构体字段对应的Tag名字一致
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}
```

###### （2）先编写了server/model/error.go

```go
package model
import (
	"errors"
)

//根据业务逻辑的需要，自定义一些错误

var (
	ERROR_USER_NOTEXIST = errors.New("用户不存在。。")
	ERROR_USER_EXIST = errors.New("用户已存在。。")
	ERROR_USER_PWD = errors.New("密码错误")

)
```

###### （3）编写了server/model/userDao.go

```go
package model
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
)

//我们在服务器启动后，就初始化一个UserDao实例
//把它做成全局的变量，在需要和redis操作时，就直接使用即可
var (
	MyUserDao *UserDao
)
//定义一个UserDao结构体
//完成对User 结构体的各种操作

type UserDao struct {
	pool *redis.Pool 
}

//使用工厂模式创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao){
	userDao = &UserDao{
	pool:pool,
  }
  return
}

//写方法,应该提供哪个方法呢
//1,根据用户id返回一个User实例+err
func (this *UserDao) getUserById(conn redis.Conn,id int) (user *User,err error) {

	//通过给定的id去redis去查询用户
	res,err := redis.String(conn.Do("HGet","users",id))
	if err != nil {
		//错误
		if err == redis.ErrNil {//表示在users中没有找到对应的id
			err= ERROR_USER_NOTEXIST
		}
		return
	}
	user = &User{}

	//这里我们需要反序列化成一个User实例
	err = json.Unmarshal([]byte(res),user)
	if err != nil {
		fmt.Println("json.Unmarshal Err=",err)
		return
	}
	return

}

//完成登录的校验 Login
//1.Login 完成对用户的验证
//2.如果用户的id和pwd都正确，则返回一个User实例
//3.如果用户的id和pwd有错误，则返回对应的错误信息

func (this *UserDao)Login(userId int,userPwd string)(user *User,err error){

	//先从UserDao链接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	user,err = this.getUserById(conn,userId)
	if err != nil {
		return
	}
	//这时证明用户是获取到了
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

```

###### （4）编写了server/main.redis.go

```go
package main
import (
	"github.com/garyburd/redigo/redis"
	"time"

)

//定义一个全局的pool
var pool *redis.Pool

func initPool(address string,maxIdle,maxActive int,idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle: maxIdle, //最大空闲连接数
		MaxActive: maxActive,//表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: idleTimeout,//最大空闲时间
		Dial:func()(redis.Conn,error){//初始化连接的代码。连接哪个ip
			return redis.Dial("tcp",address)
		},
	 }
}
```

###### （5）编写了server/process/userProcess.go改进登录方式以及错误类型

```go
//我们需要到redis数据库去完成验证
	//1.使用model.MyUserDao到redis去验证
	user, err := model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)
	
	if err != nil {
		if err ==model.ERROR_USER_NOTEXIST {
			loginResMes.Code = 500
		    loginResMes.Error = err.Error()
		}else if err ==model.ERROR_USER_PWD {
			loginResMes.Code = 403
		    loginResMes.Error = err.Error()
		}else {
			loginResMes.Code = 505
		    loginResMes.Error = "服务器内部错误..."
		}
		
		//这里我们先测试成功，然后再返回具体的错误信息
	}else{
		loginResMes.Code = 200
		fmt.Println(user,"登录成功")
	}
```

###### （6）改进server/main/main.go(加了一个初始化redis连接池的函数)

```go
func init(){
	//当服务器启动时，我们就去初始化我们的redis的连接池
	initPool("localhost:6379",16,0,300 * time.Second)
    initUserDao()
}

//这里我们编写一个函数完成对UserDao的初始化任务
func initUserDao() {
	//这里的pool本身就是一个全局的变量
	//这里需要注意一个初始化的顺序问题
	//initPool,在initUserDao
	model.MyUserDao =model.NewUserDao(pool)
}
```

### 4.完成用户注册操作

#### 1）要求

完成注册功能，将用户信息录入到Redis中

思路分析，并完成代码

思路分析的示意图

#### 2）具体代码

##### （1）common/message/user.go(从server/model下复制过来的。记住要复制而不是剪切还要改包名)

```go
package message

// User 定义一个用户的结构体
type User struct {
	//确定字段信息
	//为了序列化和反序列化成功
	//用户信息的json字符串与结构体字段对应的Tag名字一致
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
```

##### （2）common/message/message.go增加了关于注册消息的代码

```
type RegisterMes struct {
	User User `json:"user"` //类型就是User结构体

 }
type RegisterResMes struct {
	Code int `json:"code"` //返回状态码400表示该用户已经占用 200表示登录注册成功
	Error string `json` //返回错误信息
}
```

##### （3）server/process/userProcess(增加了一个方法)

```go
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error){
	//1.先从mes中取出mes.Data,并直接反序列化成RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	//2.再声明一个RegisterMes
	var registerResMes message.RegisterResMes

	//我们需要到redis数据库去完成注册
	//1.使用model.MyUserDao到redis去注册
	err= model.MyUserDao.Register(&registerMes.User)
	if err !=nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册时发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}

	//3.将loginResMes 序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("Marshal fail err=", err)
	}

	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("Marshal fail err=", err)
		return
	}
	//6.发送data 我们将其封装为writePkg
	//因为使用了分层模式(mvc),我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return

}
```

##### （4）server/model/userDao(增加了一个Register方法对数据库进行添加的操作)

```go
func (this *UserDao)Register(user *message.User)(err error){

	//先从UserDao链接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	_,err = this.getUserById(conn,user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	//这时说明id在redis还没有，则可以完成注册
	data, err :=json.Marshal(user) //序列化
	if err != nil {
		return
	}
	//入库
	_,err = conn.Do("HSet","users",user.UserId,string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err=",err)
		return
	}
	return

}
```

##### （5）在client/main/main.go进行了调用操作

```go
case 2 :
			fmt.Println("注册用户")	
			fmt.Println("请输入用户id")	
			fmt.Scanf("%d\n",&userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n",&userPwd)
			fmt.Println("请输入用户的名字(昵称)")
			fmt.Scanf("%s\n",&userName)
			//2.调用UserProcess，完成注册的请求
			up :=&process.UserProcess{}
			up.Register(userId,userPwd,userName)
```

##### (6)client/process/userProcess.go(添加一个Register的方法)

```go
func (this *UserProcess) Register(userId int,userPwd string,userName string)(err error){
	//1.连接到服务器端
	conn, err :=net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=",err)
		return
	}

	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3.创建一个RegisterMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd 
	registerMes.User.UserName = userName

	//4.将registerMes序列化
	data, err :=json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Mashal err=",err)
		return
	}

	//5.将data赋给了mes.Data字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data, err =json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Mashal err=",err)
		return
	}
	
	//7.到这个时候，data就是我们要发送的消息
    //7.1先把data的长度发送给服务器
	//先获取data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen) //将该、长度转成了byte类型是数据
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err !=nil {
		fmt.Println("connWrite(buf) fail ",err)
		return
	}

	fmt.Printf("客户端发送数据的消息长度=%d 内容是=%s",len(data),string(data))
	//发送消息本身
	_, err = conn.Write(data)
	if err !=nil {
		fmt.Println("connWrite(data) fail ",err)
		return
	}

	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn : conn,
	}

	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=",err)
	}

	mes, err = tf.ReadPkg() //mes 就是RegisterResMes

	if err != nil {
		fmt.Println("readPkg(conn) err=",err)
		return
	}

	//将mes的Data部分反序列化为RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data),&registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，你重新登录一把")
		os.Exit(0)
	}else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return

}
```

### 5.实现功能-完成登录时能返回当前在线用户

1）用户登陆后，可以得到当前在线用户列表 思路分析、示意图代码实现

用户登陆后，可以得到当前在线用户列表

（1）在服务器端维护一个onlineUsers map[int] *UserProcess

（2）创建一个新的文件userMgr.go,完成功能，对onlineUsers这个map进行增删改查

（3）在loginResMes增加一个字段 User []int 将在线的用户ID返回

（4）当用户登陆后，可以显示当前在线用户列表

2）示意图

![](D:\myfile\后端\go语言学习\pic\网络编程\pic31.jpg)

#### 3）代码实现

##### （1）编写了server/process/userMgr.go

```go
package process
import (
	"fmt"
)
//因为UserMge实例在服务其中有且只有一个
//因为在很多的地方，都会使用，因此，我们
//将其定义为全局变量
var (
	userMgr *UserMgr
)
type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对userMge的初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers : make(map[int]*UserProcess,1024),
	}
}

//完成对onlineUsers的添加
func (this *UserMgr) AddOnlinesUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除
func (this *UserMgr) DeleteOnlinesUser(userId int ) {
	delete(this.onlineUsers,userId)
}

//返回当前所有在线的用户
func (this *UserMgr)GetAllUsers() map[int]*UserProcess {
	return this.onlineUsers
}

//根据id返回对应的值
func(this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess,err error){
	//如何从map中取出一个值，待检测的方式
	up, ok := this.onlineUsers[userId]
	if !ok { //说明你要查找的用户，当前不在线
		err = fmt.Errorf("用户id不存在",userId)
		return
	} 
	return
}
```

##### （2）server/process/userProcess.go(在login成功的地方加入代码)

```go
} else {
		loginResMes.Code = 200
		//这里，因为用户登录成功，我们就把登录成功的用户放入到userMgr中
		//将登录成功的用户的userId赋给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlinesUser(this)
		//将当前在线用户的id放入到loginResMes.UsersId
		//遍历userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers{
			loginResMes.UsersId = append(loginResMes.UsersId,id)
		}
		fmt.Println(user, "登录成功")
	}

```

（3）client

/process/userProcess.go(在login成功的地方加入代码)

```go
//现在可以显示当前在线的列表 遍历loginResMes.UsersId
		fmt.Println("当前在线用户列表如下")
		for _, v := range loginResMes.UsersId {
			//如果我们要求不显示自己在线，下面我们增加一个代码
			if v == userId {
				continue
			}
			
			fmt.Println("用户id:\t",v)
		}
		fmt.Println("\n\n")
```

#### 4）当一个新的用户上线后，其他已经登录的用户也能获取最新在新用户列表

思路1：

当有一个用户上线后，服务其就马上把维护的onlineUser map整体推送

思路2：

服务其有自己的策略，每隔一段时间，把维护的onlineUsers map整体推送

思路3：

（1）当一个用户上线后，服务器就把A用户的上线信息推送给所有在线用户即可

（2）客户端也要维护一个map,map中记录了他的好友（目前就是所有人）map[int]User

（3）客户端和服务器的通讯通道要依赖于serverProcess协程

代码实现

（1）在server/process/userMgr.go

```go
package process
import (
	"fmt"
	"go_code/chatroom/common/message"
)

//客户端要维护的Map
var onlineUsers map[int]*message.User = make(map[int]*message.User,10)

//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历一把onlineUsers
	fmt.Println("当前在线用户列表：")
	for id,_ := range onlineUsers{
		//如果不显示自己
		fmt.Println("用户id:\t\t",id)
	}
}

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	
	//适当的优化
	user,ok :=onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &message.User{
			UserId : notifyUserStatusMes.UserId,
		}	
	}

	user.UserStatus = string(notifyUserStatusMes.Status)
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}
```

（2）server/process/userProcess.go

```go
//这里我们编写通知所有在线用户的方法
//这个id要通知其他的在线用户，我上线
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {

	//遍历 onlineUsers ,然后一个一个的发送 NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		//过滤掉自己
		if id == userId {
			continue
		}
		//开始通知【单独的写一个方法】
		up.NotifyMeOnline(userId)
	}

}

func (this *UserProcess) NotifyMeOnline(userId int){

	//组装我们的NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err",err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值给mes.Data
	mes.Data = string(data)

	//对message再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err",err)
		return
	}
	//发送,创建一个transfer实例发送
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOline err=",err)
		return
	}

}

下面调用
//通知其他的用户我上线了
		this.NotifyOthersOnlineUser(loginMes.UserId)
```

（3）common/message/message.go

```go
//为了配合服务器端推送用户状态变化类型
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用户id
	Status int `json:"status"` //用户的状态
}
```

(4)客户端client/process/userMgr.go

```go
package process
import (
	"fmt"
	"go_code/chatroom/common/message"
)

//客户端要维护的Map
var onlineUsers map[int]*message.User = make(map[int]*message.User,10)

//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历一把onlineUsers
	fmt.Println("当前在线用户列表：")
	for id,_ := range onlineUsers{
		//如果不显示自己
		fmt.Println("用户id:\t\t",id)
	}
}

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	
	//适当的优化
	user,ok :=onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &message.User{
			UserId : notifyUserStatusMes.UserId,
		}	
	}

	user.UserStatus = string(notifyUserStatusMes.Status)
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}
```

（5）client/main/server.go

```go
case 1:
			//fmt.Println("显示用户在线列表")
			outputOnlineUser()
		case 2:


//如果读取到消息，又是下一步处理逻辑
		switch mes.Type {
			case message.NotifyUserStatusMesType : //有人上线了
			//1.取出 NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data),&notifyUserStatusMes)
			//2.把这个用户的信息，状态保存在客户端的map[int]User中
			updateUserStatus(&notifyUserStatusMes)
			//处理

			default :
			fmt.Println("服务其端返回了未知的消息类型")	

		}
```

### 6.完成登录可以完成群聊操作

#### -1.步骤1 ：

当一个用户上线后，可以将群聊消息发给服务器。服务器可以接收到

##### 1）思路分析

![](D:\myfile\后端\go语言学习\pic\网络编程\pic32.jpg)

（1）新增一个消息结构体

（2）新增一个model CurUser

（3）在smsProcess增加相应的方法 SendGroupMes,

##### 2）代码实现

（1）common/message/message.go

```go
//增加一个SmsMes //发送的
type SmsMes struct {
	Content string   `json:"content"` //内容
	User //匿名结构体，继承
}

```

（2）client/model/curUser.go

```go
package model
import (
	"net"
	"go_code/chatroom/common/message"
)

//因为在客户端，我们很多地方会使用到curUser,我们将其作为一个全局的
type CurUser struct {
	Conn net.Conn
	message.User
}
```

（3）client/process/smsProcess.go

```go
package process
import (
	"fmt"
	"encoding/json"
	"go_code/chatroom/common/message"
	"go_code/chatroom/client/utils"
)

type SmsProcess struct {

}

//发送群聊的消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {

	//1.创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.创建一个SmsMes 实例
	var smsMes message.SmsMes
	smsMes.Content = content //内容
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//3.序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal err=",err.Error())
		return
	}
	mes.Data = string(data)

	//4.对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println(" json.Marshal err=",err.Error())
		return
	}

	//5.将mes发送给服务器
	tf := &utils.Transfer{
		Conn : CurUser.Conn,
	}

	//6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupsMes err=",err.Error())
		return
	}

	return
}
```

（4）测试

![](D:\myfile\后端\go语言学习\pic\网络编程\pic33.jpg)

#### -2.步骤2.

服务器可以将接收到的消息，群发给所有在线用户(发送者除外)

##### 1）思路分析

![]()

（1）在服务器端接收到SmsMes消息![pic34](D:\myfile\后端\go语言学习\pic\网络编程\pic34.jpg)

（2）在server/process/SmsProcess.go文件增加群发消息的方法

（3）在客户端还要增加去处理服务器端转发的群发消息

##### 2）代码实现

（1）server/main/processor.go[在server中调用转发消息的方法]

```go
//处理注册
		up := &process2.UserProcess{
			Conn : this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType :
		//创建一个SmsProcess实例完成转发群聊消息。	
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)

```

（2）client/process/smsMes.go

```go
package process
import (
	"fmt"
	"encoding/json"
	"go_code/chatroom/common/message"
	"go_code/chatroom/client/utils"
)

type SmsProcess struct {

}

//发送群聊的消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {

	//1.创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.创建一个SmsMes 实例
	var smsMes message.SmsMes
	smsMes.Content = content //内容
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//3.序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal err=",err.Error())
		return
	}
	mes.Data = string(data)

	//4.对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println(" json.Marshal err=",err.Error())
		return
	}

	//5.将mes发送给服务器
	tf := &utils.Transfer{
		Conn : CurUser.Conn,
	}

	//6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupsMes err=",err.Error())
		return
	}

	return
}
```

（3）client/process/smsMgr.go

```go
package process
import (
	"fmt"
	"encoding/json"
	"go_code/chatroom/common/message"

)

func outputGroupMes(mes *message.Message) {//这个地方一定是SmsMes
	//显示即可
	//1.反序列化mes.Data
	var smsMes message.SmsMes 
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=",err.Error())
		return
	}

	//显示信息
	info := fmt.Sprintf("用户id:\t%d 对大家说:\t%s",smsMes.UserId,smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
```

（4）client/process/server.go

```go
case message.SmsMesType : //有人群发消息了
				outputGroupMes(&mes)
```

##### 3）拓展功能要求

1.可以实现私聊（点对点聊天）

2.如果一个登录用户离线，就把这个人从在线列表中去掉

3.实现离线留言，在群聊时，如果某个用户没有在线，当登录后，可以接受到离线的消息

## 二十四、数据结构

### 1.数据结构的介绍

- 数据结构是一门研究算法的学科，自从有了数据结构，学好数据结构可以编写除更加漂亮，更加有效率的代码
- 要学习好数据结构就要多多考虑如何将生活中遇到的问题，用程序去实现解决
- 程序=数据结构+算法

### 2.稀疏数组

#### -1.基本介绍

当一个数组中大部分元素为0，或者为同一个值的数值时，可以使用稀疏数组来保存该数组

稀疏数组的处理方法是：

1）**记录数组一共有几行几列。有多少个不同的值**



2）**把具有不同值的元素的行列及值记录在一个小规模的数组中，从而缩小程序的规模**

举例：

![](D:\myfile\后端\go语言学习\pic\数据结构\pic1.jpg)

#### -2.稀疏sparsearry数组

先看一个实际的需求

编写的五子棋程序中，有存盘退出和续上盘的功能。

![](D:\myfile\后端\go语言学习\pic\数据结构\pic2.jpg)

-3.应用实例

![](D:\myfile\后端\go语言学习\pic\数据结构\pic3.jpg)

1）使用稀疏数组，来保存类似前面的二维数组(棋盘、地图等)

2）把稀疏数组存盘，并且可以重新恢复为原来的二维数组

3）整体思路分析

4）代码实现

```go
package main
import (
	"fmt"
)
type ValNode struct {
	row int
	col int
	val int
}
func main() {
	//1.先创建一个原始数组
	var chessmap [11][11]int
	chessmap[1][2] = 1 //黑子
	chessmap[2][3] = 2 //黑子

	//2.输出看看原始的数组
	for _,v :=range chessmap {
		for _,v2 := range v {
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}

	//3.转成一个稀疏数组
	//思路
	//（1）遍历chessmap，如果我们发现有一个元素的值不为0，创建一个node结构体
	//（2）将其放入到对应的切片即可

	var sparseArr []ValNode
    //标准的稀疏数组应该还有一个记录元素的二维数组的规模(行和列，默认值)
	valNode := ValNode{
		row : 11,
		col : 11,
		val : 0,
	}
	sparseArr = append(sparseArr,valNode)

	for i,v :=range chessmap {
		for j,v2 := range v {
			if v2 !=0 {
				//创建一个ValNode值结点
				valNode := ValNode{
					row : i,
					col : j,
					val : v2,
				}
				//将该结点存入切片中
				sparseArr = append(sparseArr,valNode)
			}
		}
	} 
	//输出稀疏数组
	fmt.Println("当前的稀疏数组是::::")
	for i,valNode := range sparseArr {
		fmt.Printf("%d:  %d %d %d\n",i,valNode.row,valNode.col,valNode.val)
	}

	//将这个稀疏数组，存盘 d://chessmap.data

	//如何恢复原始的数组

	//1，打开这个d://chessmap.data =>恢复原始数组
	//2、这里使用稀疏数组恢复

	//先创建一个原始数组
	var chessmap2 [11][11]int

	//遍历 sparse [遍历文件每一行]
	for i,valNode :=range sparseArr {
		if i != 0 { //跳过第一行记录的值
			// println()
			chessmap2[valNode.row][valNode.col] = valNode.val
		}
		
	}

	//看看chessmap2是不是恢复了
	fmt.Println("恢复后的原始数据")
	for _,v :=range chessmap2 {
		for _,v2 := range v {
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}

}
```

实现效果

![](D:\myfile\后端\go语言学习\pic\数据结构\pic4.jpg)



### 3.队列

使用场景：银行排队的案例，先进先出。留言

#### -1.队列介绍

队列是一个有序列表，可以用数组或是链表来实现

遵循**先入先出**的原则，即：先存入队列的数据，要先取出，后存入的要后取出

#### -2.数组模拟队列

理论：

- 队列本身是有序列表，若使用数组的结构来存储队列的数据，组队列数组的声明如下，其中maxsize是该队列的最大容量

- 因为队列的输出、输入是分别从前后端来处理，因此需要两个变量front及rear分别记录队列前后端的下标；front会随着数据输出而改变，而rear则是随着数据输入而改变如图所示：

  ![](D:\myfile\GO\project\src\go_code\pic\数据结构\pic5.jpg)

当我们将数据存入队列时称为"addqueue",addqueue的处理需要有两个步骤

1）将尾指针往后移：res+1,front==real

2）若尾指针rear小于等于队列的最大下标MaxSize-1，则将数据存入rear所指的数组元素种，否则无法存入数据

思路分析

1.创建一个数组array,是作为队列的一个字段

2.front初始化为-1

3.rear表示队列尾部，初始化为-1

4.完成队列的基本查找

addQueue//加入数据到队列

GetQueue//从队列中取出数据

ShowQueue //显示队列

代码

```go
package main
import (
	"fmt"
	"os"
	"errors"
)
//使用一个结构体管理队列
type Queue struct {
	maxSize int
	array [5]int//数组 ==》模拟队列
	front int //表示指向队列列首
	rear int //表示指向队列的尾部
}

//添加数据到队列
func (this *Queue) AddQueue(val int)(err error){

	//先判断队列是否已满
	if this.rear == this.maxSize -1 {//重要的提示；rear是队列尾部（含最后这个元素）
		return errors.New("queue full!")
	}

	this.rear++ //rear后挪一下
	this.array[this.rear] = val
	return 
}

//从队列中取出数据
func (this *Queue) GetQueue() (val int,err error) {
	//先判断队列是否为空
	if this.rear == this.front {//队空
		return -1,errors.New("queue empty!")
	}
	this.front ++
	val =this.array[this.front]
	fmt.Println("front的值为：",this.front)
	return val ,err

}

//显示队列,找到队首遍历到队尾
func (this *Queue)ShowQueue(){
	fmt.Println("队列当前的情况是:")
	//this.front不包含队首元素
	for i :=this.front +1; i <=this.rear; i++ {
		fmt.Printf("array[%d]=%d\t",i,this.array[i])
	}
	fmt.Println(" ")
}


//编写一个函数测试一下
func main(){
	//先创建一个队列
	queue := &Queue{
		maxSize : 5,
		front : -1,
		rear : -1,		
	}
	
	var key string
	var val int

	for {
		fmt.Println("1.输入add 表示添加数据到队列")
		fmt.Println("2.输入get 表示从该队列获取到元素")
		fmt.Println("3.输入show 表示显示队列")
		fmt.Println("4.输入exit 退出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err :=queue.AddQueue(val)
			if err != nil{
				fmt.Println(err.Error())
			}else{
				fmt.Println("加入队列ok")
			}
		case "get":
			val,err :=queue.GetQueue() 
			if err != nil {
				fmt.Println(err.Error())
			}else{
				fmt.Println("从队列中取出了一个数=",val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
    }
}
	
```

​    现在可以简单的实现一个队列但是存在一个问题，就是当存放一个数入队列后再取出该数后这个数组位置就不能用了，如果存满后都取走的话就无法添加数入队列

1)上面代码实现了基本队列结构，但是没有有效的利用数组空间

2)请思考如何使用数组实现一个环形队列

改进在取值的func进行改进操作

```go
//从队列中取出数据
func (this *Queue) GetQueue() (val int,err error) {
	//先判断队列是否为空
	if this.rear == this.front {//队空
		return -1,errors.New("queue empty!")
	}
	this.front ++
	val =this.array[this.front]
	fmt.Println("front的值为：",this.front)
	//当front==rear=4的时候对这两个变量的值进行重制-1就可以了
	if this.front==this.maxSize-1{
		this.front= -1
		this.rear = -1
	}
	return val ,err

}
```

但是还是有缺陷，因为这样做的话只有全部元素取完才可以进行添加元素操作，必须是rear=front=maxSize-1才可以。

继续改进吧！

-3.数组实现环形列表

思路

   对前面的数组模拟队列的优化，充分利用数组因此将数组看做是一个环形的（通过取模的方式来实现即可）

提醒：

1）尾索引的下一个为头索引时表示队列满，即将队列容量空出一个作为约定，这个在做判断队列满的时候需要(tail +1)%maxSize ==head满

2）tail ==head[空]

分析思路

1）什么时侯表示    队列满？

(tail +1)%maxSize=head

2）tail == head表示空

3）初始化时：tail=0  head =0

4)怎么统计该队列有多少元素

(tail + maxSize -head) % maxSize

代码

```go
package main
import (
	"fmt"
	"os"
	"errors"
)

//使用一个结构题管理环形队列
type CircleQueue struct{
	maxSize int //4
	array [5]int //数组
	head int //指向对头
	tail int//指向队尾
}

//如队列AddQueue(push) GetQueue(pop)
//往队列中推入一个元素(入队列)
func (this *CircleQueue) Push(val int) (err error){
	if this.IsFull(){
		return errors.New("queue full")
	}
//分析出this.tail在队列尾部但是包含最后的元素
	this.array[this.tail]=val  //把值给尾部
	this.tail =(this.tail + 1) % this.maxSize
	return
}
//从队列中弹出一个元素(出队列)
func (this *CircleQueue) Pop()(val int,err error){

	if this.IsEmpty(){
		return 0,errors.New("queue empty")
	}
	//取出,head是指向队首的，含队首元素
	val =this.array[this.head]
	this.head=(this.head + 1) % this.maxSize
	return 
}

//显示队列
func (this *CircleQueue) ListQueue(){
	fmt.Println("环形队列情况如下：")
	//取出当前队列有多少个元素
	size := this.Size()
	if size == 0{
		fmt.Println("队列为空")
	}

	//设计一个辅助变量，指向head
	tempHead :=this.head
	for i := 0;i < size; i++{
		fmt.Printf("arr[%d]=%d\t",tempHead,this.array[tempHead])
		tempHead = (tempHead  +1) %this.maxSize
	}
	fmt.Println()
}

//判断环形队列为满
func (this *CircleQueue) IsFull() bool {
	return (this.tail +1) % this.maxSize == this.head
}
//判断环形队列是否空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}
//取出环形队列有多少个元素
func (this *CircleQueue) Size() int {
	//这是一个关键算法
	return (this.tail + this.maxSize - this.head) % this.maxSize
}
func main(){

	//初始化一个环形队列
	queue := &CircleQueue{
		maxSize : 5,
		head : 0,
		tail : 0,
	}
	var key string
	var val int

	for {
		fmt.Println("1.输入add 表示添加数据到队列")
		fmt.Println("2.输入get 表示从该队列获取到元素")
		fmt.Println("3.输入show 表示显示队列")
		fmt.Println("4.输入exit 退出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err :=queue.Push(val)
			if err != nil{
				fmt.Println(err.Error())
			}else{
				fmt.Println("加入队列ok")
			}
		case "get":
			val,err :=queue.Pop() 
			if err != nil {
				fmt.Println(err.Error())
			}else{
				fmt.Println("从队列中取出了一个数=",val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
    }
}

```

运行结果如下：将元素push到队列，满了之后对其进行pop再push的效果图

![](D:\myfile\GO\project\src\go_code\pic\数据结构\pic6.jpg)

**这个要默写下来！！！！**

