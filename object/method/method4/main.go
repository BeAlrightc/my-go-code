package main
import (
	"fmt"
)

type MethodUtils struct{
    //字段...
}
//给MethodUtils编写方法
func (m MethodUtils) print(){
	for i :=1; i<=10;i++{
		for j :=1;j<=8;j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
//编写一个方法，提供m和n两个参数，方法中打印m*n的矩形
func (mu MethodUtils) print2(m int, n int){
	for i :=1; i<=m;i++{
		for j :=1;j<=n;j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
} 

//编写一个方法算该矩形的面积(可以接收长len和宽width),
//将其作为方法返回值。在main方法中调用该方法，接收返回的
//面积值并打印
func  (mu MethodUtils) area(length float32,width float32)float32{
      return length * width
}
func (mu MethodUtils) qi(m int) {
	if m %2==0{
		fmt.Printf("%v是偶数",m)
	}else{
        fmt.Printf("%v是奇数",m)
	}
}
/*
根据行、列、字符打印对应行数和列数的字符，
比如：行：3，列：2，字符 * ，则打印相应的效果
*/
func (mu *MethodUtils) printe(n int, m int,key string) {
	for i :=1;i<= n;i++{
		for j :=1;j<=m;j++{
			fmt.Print(key)
		}
		fmt.Println()
	}
}

/*
定义小小计算器结构体(Calcuator),实现加减乘除四个功能

实现形式1：分4个方法完成

实现形式2：用一个方法搞定
*/
//实现形式1：
type Calcuator struct{
	Num1 float64
	Num2 float64
}
//求和
func (calcuator *Calcuator) getSum() float64{
   
	return calcuator.Num1 +calcuator.Num2
}
//求差
func (calcuator *Calcuator) getJian() float64{
   
	return calcuator.Num1 -calcuator.Num2
}
//求积
func (calcuator *Calcuator) getJi() float64{
   
	return calcuator.Num1 *calcuator.Num2
}
//求商
func (calcuator *Calcuator) getShang() float64{
   
	return calcuator.Num1 /calcuator.Num2
}


//实现形式2
func (calcuator *Calcuator) getRes(operator byte) float64 {
   res :=0.0
   switch operator {
   case '+':
	   res = calcuator.Num1 +  calcuator.Num2
   case '-':
	   res = calcuator.Num1 -  calcuator.Num2
   case '*':
	   res = calcuator.Num1 *  calcuator.Num2
   case '/':
	   res = calcuator.Num1 / calcuator.Num2
   default:
	   fmt.Println("运算符输入有误")
	}
	return res
}


func main(){
	/*
   编写结构体(MethodUtils),编程一个方法，
   方法不需要参数，在方法中打印一个10 *8 的矩形，
   在main方法中调用该方法
	*/
	var m MethodUtils
	// m.print()
	// m.print2(3,4)
	// res :=m.area(4.0,5.6)
	// fmt.Println("该矩形的面积为：",res)
	// m.qi(8)
	m.printe(7,6,"#")

	 n := Calcuator{24.0,8.0}
	he :=n.getSum()
	cha :=n.getJian()
	ji :=n.getJi()
	shangha :=n.getShang()
	fmt.Printf("和为%v 差为%v 积为%v 商为%v\n",he,cha,ji,shangha)

	s := Calcuator{32.0,8.0}
	res :=s.getRes('*')
	fmt.Println("运算的结果为：",res)
}