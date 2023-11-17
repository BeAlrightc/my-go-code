package model
import (
	"fmt"
)
//定义一个结构体
type account struct{
	accountNo string
	pwd string
	balance float64
}

//工厂模式的函数-构造函数
func NewAccount(accountNo string,pwd string,balance float64) *account{
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号的长度不对...")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码的长度不对...")
		return nil
	}
	if balance < 20 {
		fmt.Println("余额数目不对")
		return nil
	}
	return &account{
		accountNo : accountNo,
		pwd : pwd,
		balance : balance,
	}
}




//方法
//1.存款
func (ac *account) Deposite(money float64,pwd string){
	//看看输入的密码是否正确
	if pwd !=ac.pwd{
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	ac.balance += money
	fmt.Println("存款成功")
}
//取款
func (ac *account) WithDraw(money float64,pwd string){
	//看看输入的密码是否正确
	if pwd !=ac.pwd{
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看存款金额是否正确
	if money <= 0 || money >ac.balance{
		fmt.Println("你输入的金额不正确")
		return
	}

	ac.balance -= money
	fmt.Println("取款成功")
}


//查询余额
func (ac *account)Query(pwd string){
	//看看输入的密码是否正确
	if pwd !=ac.pwd{
		fmt.Println("你输入的密码不正确")
		return
	}
	fmt.Printf("你的账号为=%v 余额=%v \n",ac.accountNo,ac.balance)
}