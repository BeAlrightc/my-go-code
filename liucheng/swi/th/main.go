package main

import(
	"fmt"
)
func main() {
	//3,根据用户指定月份，打印该月份所属的季节。
	// 3,4,5.春季 6,7,8夏季 9,10,11秋季 12,1,
	// 冬季
	var month byte
	fmt.Println("请输入月份")
	fmt.Scanln(&month)
	switch month {
	case 3,4,5:
		fmt.Println("春季")
	case 6,7,8:
		fmt.Println("夏季")	
	case 9,10,11:
		fmt.Println("秋季")	
	case 12,1,2:
		fmt.Println("冬季")	
	default:
		fmt.Println("您输错了月份")	
	}




	//2.对学生成绩大于60分的，输出”合格"。
	// 低于60的输出“不合格”。（注：输入的成绩不能大于100）
   /* var score float64
    fmt.Println("请输入成绩")
	fmt.Scanln(&score)
	switch int(score / 60){
	    case 1:
			fmt.Println("及格")	
		case 0:
			fmt.Println("不及格")	
	    default:
			fmt.Println("输入有误...")			
	}*/

	//1，使用switch将小写字符改为大写字母
	/*var char byte
	fmt.Println("请输入相应的字符")
	fmt.Scanf("%c",&char)
	switch char{
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
    case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default :
	fmt.Println("other")						
	}*/
}