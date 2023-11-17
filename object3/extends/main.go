package main
import (
	"fmt"
)
//编写一个学生考试系统
type Student struct{
	Name string
	Age int
	Score int
}
//将Pupil和Graduate共有的方法也绑定到Student
func (stu *Student) ShowInfo(){
	fmt.Printf("显示学生名字=%v 年龄=%v 成绩=%v",stu.Name,stu.Age,stu.Score)
}

func (stu *Student)SetScore(score int) {
	
	stu.Score = score
}

//给*Student增加一个方法，那么Pupil 和Graduate都可以使用该方法
func (stu *Student) GetSum(n1 int, n2 int) int{
	return n1 + n2
}

//小学生
type Pupil struct {
	Student  //嵌入Student匿名结构体
}

//这是Pupil结构体特有的方法，保留
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试")
}
//大学生，研究生。。
type Graduate struct {
	Student  //嵌入Student匿名结构体
}
//保留特有的方法
func (p *Graduate) testing() {
	fmt.Println("大学生正在考试")
}

func main(){
  
	//当我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil :=&Pupil{}
	pupil.Student.Name ="tom~"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()
    
	fmt.Println("res=",pupil.Student.GetSum(1,2))
	//大学生进行操作
	graduate :=&Graduate{}
	graduate.Student.Name ="MARY"
	graduate.Student.Age = 28
	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()
	fmt.Println("res=",graduate.Student.GetSum(10,79))

}
