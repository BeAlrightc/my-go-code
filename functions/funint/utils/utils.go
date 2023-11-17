package utils
import (
	"fmt"
)
var Age int
var Name string
//Aage 和name全局变量我们需要在main.go中使用
//但是我们需要初始化Age和Name

//init函数完成初始化
func init(){
	fmt.Println("init函数执行")
  Age = 100
  Name= "初始化"
}