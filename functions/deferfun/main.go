package main
import (
	"fmt"
)
func sum(n1 int,n2 int)int{
	defer fmt.Println("ok1 n1=",n1)
	defer fmt.Println("ok2 n2=",n2)
	
	//增加一句话
	n1++ //n1=11
	n2++//n2=21
	res := n1 + n2
	fmt.Println("ok3 res=",res)
	return res

}
func main(){
 sum(10,20)
 //输出：
//  ok3 res= 30
// ok2 n2= 20
// ok1 n1= 10
}