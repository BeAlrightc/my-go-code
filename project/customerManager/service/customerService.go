package service
import (
	"go_code/project/customerManager/model"
)

//该CustomerService ,完成对Customer的操作，包括增删改查
type CustomerService struct {
	customers []model.Customer
    //声明一个字段，表示当前切片含有多少客户
	//该字段后面，还可以作为新客户的id+1
	customerNum int
}

//编写一个方法，可以返回一个*customerService实例
func NewCustomerService() *CustomerService {
	//为了可以看到客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1,"张三","男",20,"112","zs@sohu.com")
	customerService.customers = append(customerService.customers ,customer)
	return customerService
}

//返回客户切片
//一定要使用指针的方式
func (this *CustomerService) List()[]model.Customer{
    return this.customers
}

//添加客户到customer切片中
//必须要用指针的方式，保证一直用的都是一个CustomerService
func (this *CustomerService) Add(customer model.Customer) bool{

	//我们确定一个分配id的规则，就是添加的顺序
	this.customerNum ++
	customer.Id = this.customerNum
	this.customers = append(this.customers,customer)
    return true
}

//根据id删除客户（从切片中删除）
func (this *CustomerService) Delete(id int )bool {
	index :=this.FindById(id)
	//如果index ==-1说明没有这个客户
	if index== -1 {
		return false
	}
		//如何从切片中删除一个元素
	this.customers = append(this.customers[:index],this.customers[index+1:]...)
	return true
	
}

//根据id进行修改客户信息的操作
func (this *CustomerService) Update(customer model.Customer) bool {
	index :=this.FindById(customer.Id)
	//如果index ==-1说明没有这个客户
	if index== -1 {
		return false
	}
	//将customer插入到指定的位置并对customers进行更新操作,就将原来位置的customer用一个新的customer进行替换操作
	this.customers = append(append(this.customers[:index],customer),this.customers[index+1:]...)
    return true
}

//根据Id查找客户在在切片对应中的下标,返回-1
func (this *CustomerService) FindById(id int) int {
	//默认为-1
	index := -1
	//遍历this.customers切片
	for i :=0;i < len(this.customers);i++ {
		if this.customers[i].Id ==id {
			//找到了
			index = i
		}
	}
	return index
}


