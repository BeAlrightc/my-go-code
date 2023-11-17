package main
import (
	"fmt"
	"sort"
	"math/rand"
)

//1.声明Hero结构体
type Hero struct {
	Name string
	Age int
}

//2.声明一个Hero的切片类型
type HeroSlice []Hero 

//3.实现Interface 接口
func (hs HeroSlice) Len() int {
	return len(hs )
}
//Less方法就是决定你是用什么标准进行排序
//1.按Hero的年龄从小到大进行排序！！
func (hs HeroSlice) Less(i,j int) bool {
	return hs[i].Age > hs[j].Age
	//修改成对姓名排序
    // return hs[i].Name > hs[j].Name
}

func (hs HeroSlice) Swap(i,j int) {
	// temp := hs[i]
	// hs[i] =hs[j]
	// hs[j] = temp
	//简洁的交换:下面一句话等价于上面三句话
	hs[i],hs[j] = hs[j],hs[i]
}

//声明一个Student结构体
//1.声明Student结构体
type Student struct {
	Name string
	Age int
	Score int
}
//然后将上面那三个方法复制到下面

//将student按成绩从大到小进行排序

//声明一个Stu切片类型
type StuSlice []Student
//3.实现Interface 接口
func (stu StuSlice) Len() int {
	return len(stu)
}
//Less方法就是决定你是用什么标准进行排序
//1.按Hero的年龄从小到大进行排序！！
func (stu StuSlice) Less(i,j int) bool {
	//按成绩进行排序
	return stu[i].Score > stu[j].Score
}

func (stu StuSlice) Swap(i,j int) {
	stu[i],stu[j] = stu[j],stu[i]
}



func main(){

	//先定义一个数组/切片
	var intSlice = []int{0,-1,10,7,90}
    //要求对intSlice切片进行排序
	//1.冒泡排序...
	//2.可以使用系统提供的方法
    sort.Ints(intSlice)
	fmt.Println(intSlice)//[-1 0 7 10 90]
	//请对结构体进行排序
	//1.冒泡排序
	//2.系统提供的方法

	//测试我们是否可以对结构体进行排序
	var heroes HeroSlice
	for i :=0;i < 10; i++{
		hero :=Hero {
			Name : fmt.Sprintf("英雄~%d",rand.Intn(100)),
			Age : rand.Intn(100),
		}
		//将hero append到heros切片
		heroes = append(heroes,hero)
	}
	//看看排序前的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}
	fmt.Println( )
    //调用sort.Sort()
    sort.Sort(heroes)
    //看看排序后的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}

	//这个接口的妙处就是将这个接口的三个方法实现然后只需要将结合挂钩梯放入到
	//sort方法中去就可以
	fmt.Println()

	var studentsl StuSlice
	for i :=0;i < 10; i++{
		stus :=Student {
			Name : fmt.Sprintf("学生~%d",rand.Intn(100)),
			Age : rand.Intn(100),
			Score : rand.Intn(100),
		}
		//将hero append到heros切片
		studentsl = append(studentsl,stus)
	}
	//看看排序前的顺序
	for _, v := range studentsl {
		fmt.Println(v)
	}
	fmt.Println( )
    //调用sort.Sort()
    sort.Sort(studentsl)
    //看看排序后的顺序
	for _, v := range studentsl {
		fmt.Println(v)
	}
}