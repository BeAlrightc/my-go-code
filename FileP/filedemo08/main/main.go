package main
import (
	"fmt"
	"io/ioutil"
)
func main(){
	//将D:/test/test01/abc.txt文件的内容导入到D:/test/test01/kkk.txt中

	//1.首先将D:/test/test01/abc.txt内容读取到内存
	
	filePath :="D:/test/test01/abc.txt"
	filePath2 :="D:/test/test01/kkk.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		//说明读取文件有错误
		fmt.Printf("read file err=%v\n",err)
		return 
	}
//2.将读取到的内容写入：D:/test/test01/kkk.txt中
	err = ioutil.WriteFile(filePath2,data,0666)
	if err != nil {
		fmt.Printf("write file err=%v\n",err)
	}
	//会覆盖掉写入之后的文件中的所有内容

}

func PathExists(path string)(bool,error){
	_,err :=os.Stat(path)
	if err == nil {
		return true,nil
	}
	if os.IsNotExist(err){
		return false,nil
	}
	return false,err
}
