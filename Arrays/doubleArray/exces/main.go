package main
import (
	"fmt"
)
/*
定义一个二维数组，用于保存三个班，每个班五名同学成绩，
并求出每个班级的平均分、以及所有班级的平均分
*/
func classf(){
	//定义一个二维数组
	var scores [3][5]float64
	//2循环的输入数据
	for i :=0; i < len(scores); i++{
		for j :=0; j < len(scores[i]);j++{
			fmt.Printf("请输入第%d班的第%d个学生的成绩\n",i+1,j+1)
		    fmt.Scanln(&scores[i][j])
		}
	}
	fmt.Println(scores)

	//遍历输出成绩后的二维数组，统计平均分
	totalSum  := 0.0 //定义一个变量用于统计所有班级的分数
	for i :=0; i < len(scores); i++{
		sum := 0.0 //定义一个变量，用于累计各个班级的成绩
		for j :=0; j < len(scores[i]);j++{
			sum += scores[i][j]
		}
		totalSum += sum
		fmt.Printf("第%d班级的总分%v,平均分为%v\n",i+1,sum,
		sum/float64(len(scores[i])))
	}
	fmt.Printf("所有班级额度总分是%v,所有班级的平均分是%v",
	totalSum,totalSum/15)


}

func main(){
classf()
}