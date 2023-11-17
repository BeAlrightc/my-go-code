package main

import (
	"fmt"
)

func main(){
	//题1
 var day int  =97
 var week int = day/7
 var mo int = day%7
fmt.Printf("还有%d天放假,是%d个星期%d天",day,week,mo)
 
var huashi float32 = 134.2
var sheshi float32 = 5.0 / 9 * (huashi - 100)
fmt.Printf("%v对应的摄氏温度=%v",huashi,sheshi) 19 
}