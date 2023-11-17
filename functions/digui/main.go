package main

import(
	"fmt"
)

func test(n int){
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=",n)
 }
 func test2 (n int) {
	 if n > 2{
	   n--
	   test2(n)
	 }else{
		fmt.Println("n=",n)
	 }
 }

//斐波那契数
 func feibo(n int) int{
	if n==1 || n==2{
		return 1
	}else{
	return feibo(n-1)+feibo(n-2)
	}
 }
 //求函数的值 当f(1)=3,f(n)=2*f(n-1)+1
func han(n int) int {
	if n==1{
		return 3
	}else{
		return 2*han(n-1)+1
	}
}
//猴子吃桃问题
func MonkeyPeach(n int) int {
	if n>10 || n<0{
		fmt.Println("输入的天数不对")
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

func main(){
  // test(4) //输出 n=2 n=2 n=3
   //fmt.Println(" ")
  // test2(4) //输出n=2 
  //测试
//   fmt.Println("res=",feibo(1))//1
//   fmt.Println("res=",feibo(2))//1
//   fmt.Println("res=",feibo(3))//2
//   fmt.Println("res=",feibo(4))//3
 //  fmt.Println("res=",han(1)) //3
  fmt.Println("res=",MonkeyPeach(9)) 
  //fmt.Println("res=",maonkey2(10)) 
}
