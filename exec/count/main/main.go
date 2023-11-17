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
		//str = []rune(str)
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