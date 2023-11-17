package main
import (
	_"fmt"
	"testing" //引入go的testing框架包
)

//编写测试用例，去测试，去测试addUpper函数是否正确   、
func TestAddUpper(t *testing.T) {
  
	//调用
	res := AddUpper(10)
	if res != 55 {
		//fmt.Println("AddUpper(10)执行错误，期望值=%v实际值=%v\n",55,res)
		t.Fatalf("AddUpper(10)执行错误，期望值=%v实际值=%v\n",55,res)
	}

	//如果正确，输出日志
	t.Logf("AddUpper(10)执行正确...")
}                                                                                                                                                                                                                                                                                                                                                                                                
