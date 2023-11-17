package main
import (
	"fmt"
	"reflect"
)

//通过反射，修改
//num int 的值
//修改student的值

func reflect01(b interface{}){
	//获取到reflect.Value
	rVal := reflect.ValueOf(b)
	rVal.Elem().SetInt(20)
}
func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println("num=",num) //20


	//你可以这样理解这句话：rVal.Elem()
	// num := 9
	//  ptr *int = &num
	//  num2 :=*ptr
}