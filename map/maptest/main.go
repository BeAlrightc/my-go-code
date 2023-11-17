package main
import(
	"fmt"
)
func main(){
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)
	//因为no3这个key已经存在，因此下面的这句话就是修改
	cities["no3"] = "深圳"
    fmt.Println(cities)
    //演示删除
	//delete(cities,"no1")
	//fmt.Println(cities)//map[no2:天津 no3:深圳]

    //当delete指定的key不存在时，删除不会操作，也不会报错
	delete(cities,"no5")
	fmt.Println(cities)////map[no2:天津 no3:深圳]
	



	//演示map的查找
	val, ok := cities["no1"]
	if ok{
		fmt.Printf("有no1的key,值为：%v\n",val)
	}else{
		fmt.Println("没有no1的key，不存在这个值")
	}

	//如果希望一次性删除所有的key
	//1.遍历啊所有的key,如何逐一删除[遍历]
	//2.直接make一个新空间
	cities = make (map[string]string)//效率较高
	fmt.Println(cities)
}