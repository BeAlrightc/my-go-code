package main
import(
	"fmt"
	"sort"
)
func main(){
	//map排序
	map1 :=make(map[int]int,10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	fmt.Println(map1)//map[10:100 1:13 4:56 8:90]

	//如何按照map的key的顺序进行排序输出
	//1.先将map的key放入到切片中
	//2.对切片排序、
	//3.遍历切片，然后按照key来输出map的值
	var keys []int
	for k,_:=range map1{
		keys = append(keys,k)
	}
	//排序
	sort.Ints(keys)
	fmt.Println(keys)

	//遍历切片，然后按照key来输出map的值
	for _, k := range keys{
		fmt.Printf("map[%v]=%v\n",k,map1[k])
	}
}