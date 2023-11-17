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