package main
import (
	"fmt"
	"go_code/object2/encapsulate/model"
)

func main (){
	p :=model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000.0)
	fmt.Println(*p)
	fmt.Println(p.Name,"age =",p.GetAge(),"sal = ",p.GetSal())
}