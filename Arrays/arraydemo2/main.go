package main
import (
	"fmt"
)

func main(){
	//演示for -range遍历数组
	heroes :=[...]string{"刘备","张飞","关羽"}
	fmt.Println(heroes)
	//for -range遍历
	for i,v :=range heroes{
		fmt.Printf("i=%v,v=%v\n",i,v)//i=0,v=刘备i=1,v=张飞i=2,v=关羽s
	    fmt.Printf("heroes[%d]=%v\n",i,heroes[i])
	}

	for _,v :=range heroes{
		fmt.Printf("元素的值=%v\n",v)
	}
}