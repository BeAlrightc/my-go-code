package main
import(
	"fmt"
)
func main(){
	//map的声明和注意事项，map是无序的数据结构
	var a map[string]string
	//使用、map前，需要先make,make的作用就是给map分配数据空间
	a  = make(map[string]string,10) //最大可以放10对
	a["ao1"]="宋江"
	a["ao2"]="吴用"
	a["ao3"]="李逵"
	a["ao4"]="林冲"
	a["ao5"]="吴用"
	a["ao5"]="无名" //会覆盖同key的值
    fmt.Println(a)

	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)//map[no1:北京 no2:天津 no3:上海]

	//第三种方式
	heroes := map[string]string{
		"heroe1" : "宋江",
		"heroe2" : "林冲",
	}
    fmt.Println(heroes)//map[heroe2:林冲 heroe1:宋江]
}