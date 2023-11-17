package main
import (
	"fmt"
)
/*
（1）有一个数列：白眉鹰王、
金毛狮王、紫衫龙王、青翼斧王

猜数游戏：从键盘任意输入一个名称，
判断数列中是否包含此名称
思路：
1.定义一个数组：白眉鹰王、金毛狮王、紫衫龙王、青翼斧王
2.从控制台接收一个名词，依次比较如果发现有就提示
*/
func find1(){
	names := [4]string{"白眉鹰王","金毛狮王","紫衫龙王","青翼斧王"}
	var heroName = " "
	fmt.Println("请输入要查找的人名...")
	fmt.Scanln(&heroName)
	//顺序查找第一种方式
	for i := 0; i < len(names); i++{
		if heroName == names[i]{
			fmt.Printf("找到了%v，下标%v",heroName,i)
			break
		} else if i ==(len(names)-1){
			fmt.Printf("没有找到%v",heroName)
		}
	}
}
//顺序查找第二种方式
func find2(){
	names := [4]string{"白眉鹰王","金毛狮王","紫衫龙王","青翼斧王"}
	var heroName = " "
	fmt.Println("请输入要查找的人名...")
	fmt.Scanln(&heroName)

	index := -1
	for i :=0; i < len(names); i++{
		if heroName == names[i]{
			index =i
			break
		}
	}
	if index != -1{
		fmt.Printf("找到了%v，下标%v",heroName,index)
	}else{
		fmt.Printf("没有找到%v",heroName)
	}

}

func main(){
	// find1()
	find2()
  
}