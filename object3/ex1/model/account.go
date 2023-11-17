package model
import (
	"fmt"
)
type account struct{
	Achao string
	balance float64
	pwd string
}


func NewAcount (achao string) *account{
	return &account{
		Achao : achao,
	}
}
func (a *account)SetBalance(balance float64)int{
   if balance <20{
	fmt.Println("余额输入错误")
	return 0
   }else{
	a.balance = balance
	return 1
   }
}
func (a *account)SetPwd(pwd string)int{
	if len(pwd)<6{
		fmt.Println("您输入的密码过短，请重新输入")
		return 0
	}else{
		a.pwd = pwd
		return 1
	}
}


func (a *account)GetBalance()float64{
   return a.balance
}

func (a *account)GetPwd()string{
	return a.pwd
 }




