package main
import (
	_"fmt"
	"testing" //引入go的testing框架包
)

//编写测试用例，去测试，去测试sub函数是否正确   、
func TestGetSub(t *testing.T) {
  
	//调用
	res := getSub(10,3)
	if res != 7 {
		
		t.Fatalf("getSub(10)执行错误，期望值=%v实际值=%v\n",7,res)
	}

	//如果正确，输出日志
	t.Logf("getSub(10)执行正确...")
}                                                                                                                                                                                                                                                                                                                                                                                                

