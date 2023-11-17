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
