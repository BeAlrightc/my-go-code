package objectTestAcc
import (
	"fmt"
)

type FamilyAccount struct {
	//声明必须字段
	
	//声明一个字段，保存接收用户输入的选项
	key string
	//声明一个字段，控制是否退出for循环
	loop bool
	//声明一个字段统计余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
    note string
	//定义一个字段记录是否有收支的行为
	flag bool
	//收支的详情
	//当有收支发生的时候，就对details进行拼接处理
	details string
}
//编写一个构造方法返回一个FamilyAccount实例 
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key : "",
		loop : true,
		balance : 10000.0,
		money : 0.0,
		note : "",
		flag : false,
		details :  "收支\t账户余额\t收支金额\t说明",
	}
}

//将显示明细写成一个方法
func (this *FamilyAccount) ShowDetails(){
    fmt.Println("------------当前收支明细记录--------")
	if this.flag {
		fmt.Println(this.details)
	}else{
		fmt.Println("您当前没有支出记录，来一笔吧！")
	}
}

//将登记收入写成一个方法和*FamilyAccount绑定
func (this *FamilyAccount) Income(){
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money //修改账户余额
	fmt.Println("本次收入的说明：")
	fmt.Scanln(&this.note)
	//将这个收入情况，拼接到details变量当中
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v",this.balance,this.money,this.note)
	this.flag = true
}
//将支出也绑定到一个方法当中
func (this *FamilyAccount) Pay(){
	fmt.Println("本次支出的金额：")
	fmt.Scanln(&this.money)
	//这里需要做出一个必要的判断
	if this.money > this.balance {
		fmt.Println("余额不足")
	}
	this.balance -=this.money
	fmt.Println("本次的支出说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v",this.balance,this.money,this.note)
	this.flag = true
}

//将退出系统写成一个方法
func (this *FamilyAccount) exit(){
	fmt.Println("您确定要退出吗？ y/n")
	choice :=" "
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n"{ //输了y/n就break出去
			break
		}
		fmt.Println("您的输入有误请重新输入 y/n")
	}
	if choice == "y" {
		this.loop = false	
	}
}

//为该结构体绑定相应的方法
//显示主菜单
func (this *FamilyAccount) MainMenu(){
	for {
		fmt.Println("\n--------家庭收支记账软件---------")
		fmt.Println("         1.收支明细")
		fmt.Println("         2.登记收入")
		fmt.Println("         3.登记支出")
		fmt.Println("         4.退出软件")
		fmt.Print("请选择(1-4)")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1" :
			this.ShowDetails()
		case "2" :
			this.Income()
		case "3" :
			this.Pay()
		case "4" :
			this.exit()	
		default :
		    fmt.Println("请输入正确的选项")	
		}

		if !this.loop {
			break
		}
	}
}
