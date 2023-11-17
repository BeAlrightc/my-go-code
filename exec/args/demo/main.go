package main
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行的参数值",len(os.Args))
	//遍历os.Args切片，就可以得到所有命令行输入的参数值
	for i,v :=range os.Args {
		fmt.Printf("Args[%v]=%v\n",i,v)
	}
}