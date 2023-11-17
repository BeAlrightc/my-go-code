package main

import(
	"fmt"
)
func main(){
	//案例1
	/*
	 1)统计3个班的成绩情况，每个班5名同学，
	 求出各个班的平均分和所有班级的平均分（学生的成绩从键盘输入
	 ）
	 分析实现思路
	 1.统计一个班成绩情况，每个班有5名同学，求出这个班的平均分
	 2.学生数就5个 
     3.声明一个变量sum统计班级的总分
	 4.定义一个变量统计总成绩得出所有学生的平均分

	 分析思路2
	 要统计三个班，走三次
	 分析思路3
	 1,将代码激活
	 2.定义两个变量表示班级的个数和班级的人数

	 //分析思路3统计班级及格人数
	 1,声明一个变量用于保存及格人数


	 代码实现
	
	var passc int = 0
	var classNum int =2
	var stuNum int = 5
	var totalSum float64 = 0.0
	for j :=1;j <=classNum;j++{
	sum := 0.0
	for i := 1;i<=stuNum;i++{
		var score float64
		fmt.Printf("请输入第%d个班的第%d个学生成绩\n",j,i)
		fmt.Scanln(&score)
        //累计总分
		sum += score
		//判断分数是否及格
		if score >=60{
          passc ++
		}
	}
	fmt.Printf("第%d个班级的平均分是%v\n",j,sum / float64(stuNum))
	//将总分进行累加
	totalSum += sum
	fmt.Println("  ")
 }
 fmt.Printf("各个班级的总成绩是%v各个班级的平均分是%v及格人数是%v",totalSum,totalSum/ float64(classNum * stuNum),passc)
 */
 /*
案例2打印金字塔经典案例
案例分析
代码思路
1.打印一个金字塔
 */
//  var toleve int =9
//  for i := 1;i<=toleve;i++{
// 	//在打印星号前打印空格
// 	for k :=1;k<=toleve-i;k++{
// 		fmt.Print(" ")
// 	}
// 	//在空格后面打印星星
// 	for j :=1;j<=2*i-1;j++{
//       fmt.Printf("*")
// 	}
// 	fmt.Println(" ")
 //}
 //打印空心金字塔
 /*
   *
  * *
  ****
  分析：在我们给每行打印*号时需要考虑一个问题是打印*还是打印空格
  每层的第一个和最后一个就是打印星星，其他的都打印空格
  分析到例外情况，最底层是全部打印星星
 *
 var toleve int =100
 for i := 1;i<=toleve;i++{
	//在打印星号前打印空格
	for k :=1;k<=toleve-i;k++{
		fmt.Print(" ")
	}
	//在空格后面打印星星
	for j :=1;j<=2*i-1;j++{
      if j ==1 || j==2 * i -1 || i==toleve{
		fmt.Printf("*")
	  }else{
		fmt.Print(" ")
	  }
	}
	fmt.Println(" ")
 }
 */

 /*
 打印九九乘法表

 
 */
 for i:=1;i<=9;i++{
	for j:=1;j<=i;j++{
		fmt.Printf("%v * %v= %v  ",i,j,i*j)
	}
	fmt.Println(" ")
 }
}