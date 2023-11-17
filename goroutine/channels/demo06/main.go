// package main
// import (
// 	"fmt"
// )
// func main() {
// 	//管道可以生命为只读或只写

// 	//1.在默认的情况下，管道是双向的。
// 	// var chan1 chan int //可读可写


// 	//2.声明为只写
// 	var chan2 chan<- int
// 	chan2 = make(chan int,3)
// 	chan2<- 20
// 	// num := <-chan2 err在这个管道中不可以读
// 	fmt.Println("chan2=",chan2)

// 	//3.声明为只读
// 	var chan3 <-chan int
// 	num2 := <-chan3
// 	// chan3<- 30 err 会报错，因为该管道为只读
// 	fmt.Println("num2=",num2)
// }