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
