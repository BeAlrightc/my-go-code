package main
import (
	"fmt"
)
type A struct {  
    Name string
    Age int
}
type Stu struct {
     A
     int
}
func  main(){
stu :=Stu{}
stu.Name = "tom"
stu.Age = 10
stu.int = 80
fmt.Println(stu)
}