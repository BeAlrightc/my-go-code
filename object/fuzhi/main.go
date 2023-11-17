package main
import (
	"fmt"
)
type Stu struct {
	Name string
	Age int
}
func main(){
	//在创建结构体变量时，就直接指定字段的值
	var stu1 = Stu{"tom",23}
	stu2 :=Stu{"晓明",20}
	//创建结构体变量时。把字段名和字段值写在一起
	var stu3 = Stu{ //这种写 法更加稳健不依赖于字段的顺序
		Name : "jack",
		Age : 20,
	}
	//可以类型颠倒
	var stu4= Stu{ //这种写法更加稳健不依赖于字段的顺序
		Age : 20,
		Name : "jack",
	}

	fmt.Println(stu1,stu2,stu3,stu4)

	//方式2.返回结构体的指针类型(!!!)
	var stu5 *Stu= &Stu{"小王",29} //stu2--->地址 ---》结构体数据[xxxx,xxxx]
	stu6 :=&Stu{"小王",39}

	//在创建结构体指针变量是，把字段名和字段值写在一起，这种写法，就不依赖于字段的定义顺序
	var stu7 = &Stu{
		Name : "小李",
		Age : 49,
	}

	var stu8 = &Stu{
		Age :59,
		Name : "小李~",

	}
	fmt.Println(*stu5,*stu6,*stu7,*stu8)//
}