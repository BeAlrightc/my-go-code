package main
import (
	"fmt"
)

type Person struct{
	Name string
	Age int
}

func main() {
	//方式1

	//方式2
	p2 := Person{"mary",20}
	// p2.Name = "Tom"
	// p2.Age = 18
	fmt.Println(p2)

	//方式3-8
	//案例：var person *Person = new (Person)

	var p3 *Person= new(Person)
	//因为p3是一个指针，因此标准的给字段赋值方式
	//(*p3).Name = "smith"也可以这样写 p3.Name = "smith"
	//原因是go的设计者为了程序员使用方便，在底层会对p3.Name = "smith"进行优化
	//底层会给p3加上取值运算 (*p3).Name = "smith"
	(*p3).Name = "smith"
	p3.Name = "jhon"
	(*p3).Age = 30
	fmt.Println(*p3)//{jhon 30}

	//方式4-{}案例：var person *Person = &Person{}
	var person *Person = &Person{"jerry",60}
	//因为person是一个指针，因此标准的访问其字段的方法是
    //(*peron).Name = "scott"
	//go的设计者为了方便,也可以person.Name = "scott"
	//底层会进行优化操作
	// (*person).Name = "scott"
	// (*person).Age= 30
	fmt.Println(*person)

	var p4 Person
	p4.Name= "小明"
	p4.Age = 23
	var p5 *Person = &p4  //定义一个指针将p4的地址赋给p5
    //fmt.Println(*p5.Age)这种用法是错误的

	fmt.Println((*p5).Age)//10
	fmt.Println(p5.Age) //10

	p5.Name = "tom~"
	fmt.Printf("p5.Name=%v p4.Name=%v \n",p5.Name,p4.Name)
	fmt.Printf("p5.Name=%v p4.Name=%v \n",(*p5).Name,p4.Name)
	/*
	输出结果如下所示
	p5.Name=tom~ p4.Name=tom~
    p5.Name=tom~ p4.Name=tom~

	*/
	fmt.Printf("p4的地址是%p\n",&p4)
	//p4的地址是0xc042056420
	
	fmt.Printf("p5的地址是%p p5的值是%p\n",&p5,p5)
	//p5的地址是0xc04207a020 p5的值是0xc042056420

}