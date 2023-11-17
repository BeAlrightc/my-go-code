 package maintype user struct {
// 	UserId string
// 	Name string
//  }

//  func TestReflectStruct(t *testing.T){
// 	var {
// 		model *user
// 		sv reflect.Value
// 	}
// 	model = &user{}
// 	sv = reflect.ValueOf(model)
// 	t.Log("reflect.ValueOf",sv.kind().String())
// 	sv=sv.Elem()
// 	t.Log("reflect.ValueOf",,sv.kind().String())
// 	sv.FieldByName("userId").SetString("12345678")
// 	sv.FieldByName("Name").SetString("nickname")
// 	t.Log("model",model)
//  }
 import (
	"fmt"
	"reflect"
 )
 func reflect01(b interface{})  {
	num := reflect.ValueOf(b)
    kind1 :=num.Kind()
	iv :=num.Interface()
	fmt.Printf("b的reflect.Value是=%v,kind值为=%v,num转换为interface的值为=%v\n",num,kind1,iv)
 }
//  
 func main() {
	var n float64 =65.9
	reflect01(n)
	var str string = "tom"
	fs :=reflect.ValueOf(&str)
	fs.Elem().SetString("jackma")
	fmt.Println(str)
 }