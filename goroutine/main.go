package main
import (
	"fmt"
	"runtime"
)

func main() {
	//获取当前系统CPU的数量
	num := runtime.NumCPU()
	//我这里设置num -1的cpu运行go程序num
	fmt.Println("cpuNum=",num)
	runtime.GOMAXPROCS(num-1)
	fmt.Println("num=",num)
}