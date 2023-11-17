package utils

import (
	"fmt"
)

//将计算的功能放入一个函数中，然后在需要的时候进行使用
//为了让马哥其他的包文件使用Cal函数，需要将c大写，类似java中public
func Cal (n1 float64,n2 float64,operator byte) float64{
	
	var res float64
	switch operator {
	case '+':
          res = n1 + n2
	case '-':
		  res = n1 - n2
	case '*':
		  res = n1 * n2
	case '/':
		  res = n1 / n2	  
	default:
		 fmt.Println("操作符错误")  
	}
	return res
}