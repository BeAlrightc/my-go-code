package main
import (
	"fmt"
)

type A struct {
	Name string
	age int
}
type B struct {
	Name string
	score float64
}
type C struct {
	A
	B
	Name string
}
type D struct {
	a A //有名结构体
}

type Goods struct{
	Name string
	Price float64
}


type Brand struct{
	Name string
	Address string
}
type TV struct {
	Goods
	Brand
}
type TV2 struct {
	*Goods
	*Brand
}
func main(){
   var c C
   //如果 c 无Name字段，而A和B有Name,这时就必须指定匿名结构体名字来区分
   //c.Name = "tom"//会报错，因为不知道选择A里的Name还是B里面的Name
   //c.A.Name = "tom" //具体指定哪一个结构体中的Name
   c.Name = "tom" //成功是因为C有Name
   var d D
//    d.Name= "jack"//报错，必须将结构体名字带上
   d.a.Name = "jack"

   tv := TV{ Goods{"电视机001",5000.99},Brand{"海尔","山东"}, }
   
   tv2 := TV{
	Goods{
		Price : 6000.90,
		Name : "电视002",
	},
	Brand{
		Name : "长虹",
		Address :"北京",
	},
   }

   tv3 := TV2{ &Goods{"电视机003",7000.99},&Brand{"创维","深圳"}, }
   fmt.Println("TV=",tv)
   fmt.Println("TV2=",tv2)
   fmt.Println("TV3=",*tv3.Goods,*tv3.Brand)
   tv4 := TV2{ 
	&Goods{
		Name : "电视机004",
		Price : 9900.99,
		},
		&Brand{
			Name : "康佳",
			Address : "上海",
			}, 
		}
		fmt.Println("TV4=",*tv4.Goods,*tv4.Brand)



}