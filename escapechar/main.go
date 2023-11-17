package main

import "fmt" //fmt包主要提供格式化和输入的函数
func main() {
	//演示转义字符的使用

	// fmt.Println("tom\tjack")
	// fmt.Println("D:\\myfile\\GO\\project\\src\\go_code\\escapechar")
	// fmt.Println("tom说\"i love you\"")
	//\r表示回车，从当前行的最前面开始输出，覆盖掉以前的内容
	//fmt.Println("天龙八部雪上飞狼\r张飞")
	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
}