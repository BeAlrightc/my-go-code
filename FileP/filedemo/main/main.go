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