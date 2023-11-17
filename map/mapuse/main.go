package main
import(
	"fmt"
)
func main(){
	studentsMap := make(map[string]map[string]string) 
	studentsMap["no1"] = make(map[string]string,3)
	studentsMap["no1"]["name"] = "tom"
	studentsMap["no1"]["sex"] = "男"
	studentsMap["no1"]["adrress"] = "北京"
 //第二个学生
    studentsMap["no2"] = make(map[string]string,3)
	studentsMap["no2"]["name"] = "jhon"
	studentsMap["no2"]["sex"] = "男"
	studentsMap["no2"]["adrress"] = "上海"

	fmt.Println(studentsMap["no1"])
	fmt.Println(studentsMap["no2"])
	//打印结果如下
	/*
    map[sex:男 adrress:北京 name:tom]
    map[name:jhon sex:男 adrress:上海]
	*/
}