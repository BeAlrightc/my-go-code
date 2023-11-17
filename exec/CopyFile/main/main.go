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