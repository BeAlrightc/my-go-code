package main
import (
	"fmt"
	"go_code/object3/ex2/model"
)
func main(){
	account := model.NewAccount("jzh1111","999999",48)
   if account !=nil{
       fmt.Println("创建成功=",*account)
   }else{
	   fmt.Println("创建失败")
   }
}