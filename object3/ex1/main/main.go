package main
import (
	"fmt"
	"go_code/object3/ex1/model"
)

func main() {
	a:=model.NewAcount("11000")
	var m int=a.SetBalance(2000.0)
    var b int = a.SetPwd("00000000")
    if m ==1 && b==1{
		fmt.Printf("恁创建的账户名是:%v 余额为:%v 密码设置为:%v",a.Achao,a.GetBalance(),a.GetPwd())
	}else{
		fmt.Println("账户创建失败")
	}
		
	}
