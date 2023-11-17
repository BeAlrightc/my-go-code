package main
import(
	"fmt"
)
func modifyUsers(users map[string]map[string]string,name string){
  //判断users中是否有name
//   v , ok := users[name]
   if users[name] != nil {
	    //有这个用户
		users[name]["pws"] = "8888"
   }else {
	//没有这个用户
	users[name] = make(map[string]string,2)
	users[name]["pws"] = "8888"
	users[name]["nicname"] = "昵称" + name //示意
   }
}
func main(){
	users :=make(map[string]map[string]string)
	users["smith"] =make(map[string]string,2)
	users["smith"]["pwd"] = "999999"
	users["smith"]["nickname"] = "小花猫"
    modifyUsers(users,"tom")
    modifyUsers(users,"mary")
    modifyUsers(users,"smith")


	fmt.Println(users)
	//输出结果为：map[mary:map[pws:8888 nicname:
	//昵称mary] smith:map[nickname:小花猫 pws:8888 pwd:999999] tom:map[pws:8888 nicname:昵称tom]]
}