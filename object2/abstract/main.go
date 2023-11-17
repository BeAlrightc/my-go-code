package main
import (
	"fmt"
)
//定义一个结构体
type Account struct{
	AccountNo string
	Pwd string
	Balance float64
}

//方法
//1.存款
func (account *Account) Deposite(money float64,pwd string){
	//看看输入的密码是否正确
	if pwd !=account.Pwd{
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance += money
	fmt.Println("存款成功")
}
//取款
func (account *Account) WithDraw(money float64,pwd string){
	//看看输入的密码是否正确
	if pwd !=account.Pwd{
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看存款金额是否正确
	if money <= 0 || money >account.Balance{
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance -= money
	fmt.Println("取款成功")
}

//查询余额

func (account *Account)Query(pwd string){
	//看看输入的密码是否正确
	if pwd !=account.Pwd{
		fmt.Println("你输入的密码不正确")
		return
	}
	fmt.Printf("你的账号为=%v 余额=%v \n",account.AccountNo,account.Balance)
}


func main(){
	//测试一下
	account := Account{
		AccountNo : "gs111111",
		Pwd : "666666",
		Balance  : 1000.0,
	}
	for{
		//这里可以做的更加灵活，就是让用户通过控制台来输入命令
		//菜单
	account.Query("666666")
	account.Deposite(200,"666666")
	account.Query("666666")
	account.WithDraw(150.0,"666666")
	account.Query("666666")
	}

}