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
 //在举一个案例
 type myFunt func(int,int)int //这时myFun就是 func(int,int)int类型

 func myFunc2(funvar myFunt,num1 int,num2 int) int{
	return funvar(num1,num2)
}
//支持对函数返回值命名
func cal(n1 int,n2 int) (sum int,sub int){
	sum = n1 +n2
	sub = n1 - n2
	return  //这样就不用这样return sum sub之类额
	}

func main(){
	/*
   a :=getSum
   fmt.Printf("a的类型%T,getSum类型是%T\n",a,getSum)

   res := a(10,40) //等价 res =:= getSum(10,40)
   fmt.Println("res=",res)

   //看案例
   res2 :=myFunc(getSum,50,60)
   fmt.Println("res2=",res2) //110

   type myInt int //给int取了别名 在go中myInt和int虽然都是
   //int类型，但是在go中还是认为myInt 和int是不同的数据类型

   var num1 myInt
   var num2 int
   num1 = 40
  // num2= num1 //报错不是同一个类型
  num2 = int(num1) //这样进行转换就可以
   fmt.Println("num1=",num1) //40

  */
   
   //案例结果
   res3 :=myFunc2(getSum,500,600)
   fmt.Println("res2=",res3) //1100

   a,b := cal(1,2)
   fmt.Printf("a=%v,b=%v",a,b) //3,-1,a是sum b是sub
}