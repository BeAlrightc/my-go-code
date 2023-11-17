package main
import (
	"fmt"
)
func modify(arr [3]int){ 
	arr[0] = 100
	fmt.Println("modify的值arr",arr) //100 2 3
	}
func main(){
	var arr = [...]int{1,2,3}
	modify(arr)
	fmt.Println(arr) //1 2 3
	}