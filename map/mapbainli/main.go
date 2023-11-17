package main
import(
	"fmt"
)
func main(){
	//使用for-range遍历map
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
    
	for k,v := range cities {
		fmt.Printf("k=%v v=%v\n",k,v)

	}
	fmt.Printf("cities有%v 对key-value \n",len(cities)) //3对
	//输出结果如下
	//k=no1 v=北京k=no2 v=天津k=no3 v=上海
	
	//使用for-range遍历比较复杂的map
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
 
	for k1,v1 :=range studentsMap{
        fmt.Println("k1=",k1)
		for k2,v2 := range v1 {
			fmt.Printf("\t k2=%v v2 = %v\n",k2,v2)
		}
	}
	/*
    输出结果如下
	k1= no1
         k2=name v2 = tom
         k2=sex v2 = 男
         k2=adrress v2 = 北京
k1= no2
         k2=sex v2 = 男
         k2=adrress v2 = 上海
         k2=name v2 = jhon

	*/
}