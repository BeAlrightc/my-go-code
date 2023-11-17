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
   