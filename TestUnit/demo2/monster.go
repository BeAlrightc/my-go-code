package monster
import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)
type Monster struct {
	Name string
	Age int
	Skill string
}

//给Monster绑定方法Store，可以将一个Monster变量(对象)，序列化后保存到文件中
func (this *Monster) Store() bool{

	//先序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err = ", err)
		return false
	}
	//保存到文件
		filePath := "D:/test/test02/monster.ser"
		err = ioutil.WriteFile(filePath, data,0666)
		if err != nil {
			fmt.Println("write file  err = ", err)
			return false
	}
	return true
	//保存到文件中
}





//给Monster绑定方法ReStore，可以将一个序列化的Monster,从文件中读取，
// 并反序列化为Monster对象
func (this *Monster) ReStore() bool {

	//1.先从文件中读取序列化字符串
	filePath := "D:/test/test02/monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Read file  err = ", err)
		return false
	}

	//2.使用读取到的data []byte，对反序列化
	err = json.Unmarshal(data,this)
	if err != nil {
        fmt.Println("Unmarshal  err = ", err)
		return false
	}
	return true
}