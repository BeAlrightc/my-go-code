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

}