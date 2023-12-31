package main
import (
	"fmt"
	"go_code/project/customerManager/service"
	"go_code/project/customerManager/model"
)

type customerView struct {
     //定义必要字段
	 key string //接收用户输入
	 loop bool //是否循环显示菜单
	 //增加一个字段customerService
     customerService   *service.CustomerService
}

//显示所有的客户信息
func (this *customerView) list(){
     //首先获取到当前所有的客户信息（在切片中）
	 customers := this.customerService.List()
	 //显示
	 fmt.Println("----------客户列表--------------")
	 fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	 for i :=0;i<len(customers);i++ {
		fmt.Println(customers[i].GetInfo())
	 }
	 fmt.Printf("\n--------客户列表完成------------\n\n")
}

//得到用户的输入，信息构建新的客户，并完成添加
func (this *customerView) add() {
	fmt.Println("------------添加客户------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话号码:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的Customer实例
	//注意：id号没有让用户输入，id号是唯一的，让系统分配即可
	customer := model.NewCustomer2(name,gender,age,phone,email)
	//调用
	if this.customerService.Add(customer) {
		fmt.Println("------------添加完成------------")
	}else{
		fmt.Println("------------添加失败------------")
	}
}

//修改客户的操作
func (this *customerView) update() {
	fmt.Println("------------修改客户------------")
    fmt.Println("请选择修改客户的编号(-1的话就退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话号码:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的Customer实例
	//注意：id号没有让用户输入，id号是唯一的，让系统分配即可
	customer := model.NewCustomer(id,name,gender,age,phone,email)
	//调用
	if this.customerService.Update(customer) {
		fmt.Println("------------修改成功------------")
	}else{
		fmt.Println("------------修改失败------------")
	}
}

//得到用户输入的id删除该id对应的客户
func (this *customerView) delete() {
	fmt.Println("------------删除客户------------")
	fmt.Println("请选择待删除的客户编号(-1退出)：")
	id :=-1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}

	fmt.Println("确认是否删除(Y/N): ")
	choice := ""
	for {
	fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" || choice =="n" || choice =="N"{
			break
		}
		fmt.Println("您的输入有误请重新输入(Y/N): ")
	}
	if choice == "y" || choice == "Y" {
			//调用service中的delete方法
		if this.customerService.Delete(id) {
			fmt.Println("------------删除成功------------")
	}else{
		fmt.Println("------------删除失败，输入的id号不存在------------")
	}
  } else{
	this.mainView()
  }
}

//退出软件
func (this *customerView) exit(){
	fmt.Println("确定是否退出(Y/N): ")
	for {
	fmt.Scanln(&this.key)
	if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n"{
		break
	  }
	  fmt.Println("您的输入有误，请重新输入(Y/N) : ")
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
}

//显示主菜单
func (this *customerView) mainView() {
	for{
		fmt.Println("--------客户信息管理系统------------")
		fmt.Println("         1.添加客户   ")
		fmt.Println("         2.修改客户   ")
		fmt.Println("         3.删除客户   ")
		fmt.Println("         4.客户列表   ")
		fmt.Println("         5.退出   ")
		fmt.Println("请选择(1-5): ")


		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
		    this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default :
		    fmt.Println("你的输入有误，请重新输入...")						
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统的使用")
}

//简单登录功能的时间
func (this *customerView) Login (){
	account :=""
	pwd :=""
	for {
    fmt.Println("请输入账号： ")
	fmt.Scanln(&account)
	fmt.Println("请输入密码")
	fmt.Scanln(&pwd)
	if account == "7758258" && pwd =="111"{
		fmt.Println("恭喜你！正在进入系统！")
		break
	   }
	   fmt.Println("您的输入的账号或者密码有误，请重新输入: ")	   	
	}
	this.mainView()
}
func main() {
	//在主函数中，创建一个customerView并运行显示主菜单...
	customerView := customerView{
		key : "",
		loop : true,	
	}
	//完成对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.Login()
}
