package main
import (
	"fmt"
	"time"
)
func main(){
   //看看日期和时间相关的函数和方法使用
   //1.获取当前时间
   now := time.Now()
   fmt.Printf("now=%v now type=%T",now,now)
   //输出结果如下
//    now=2023-08-23 18:37:45.9279802 +0800 CST m=+0.008001001 now type=time.Time
//2.通过now可以获取到年，月，日，时分秒
fmt.Printf("年=%v\n",now.Year())
fmt.Printf("月=%v\n",int(now.Month()))//转成数字
fmt.Printf("日=%v\n",now.Day())
fmt.Printf("时=%v\n",now.Hour())
fmt.Printf("分=%v\n",now.Minute())
fmt.Printf("秒=%v\n",now.Second())


//格式化日期和时间
//way1
fmt.Printf("当前年月日 %02d-%02d-%02d %02d:%02d:%02d \n",
now.Year(),now.Month(),now.Day(),
now.Hour(),now.Minute(),now.Second())//当前年月日 2023-08-23 19:17:28

dateStr := fmt.Sprintf("当前年月日 %d-%d-%d-%d-%d-%d \n",now.Year(),
now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())

fmt.Printf("dateStr=%v",dateStr)//dateStr=当前年月日 2023-8-23-19-21-36
//way2
fmt.Printf(now.Format("2006/01/02 15:04:05"))
fmt.Println()
fmt.Printf(now.Format("2006-01-02"))
fmt.Println()
fmt.Printf(now.Format("15:04:05"))
fmt.Println()
//输出结果如下所示：
// 2023/08/23 19:26:21
// 2023-08-23
// 19:26:21


//每隔0.1秒就打出一个数字，打印到100时就退出
// i := 0
// for {
// 	i++
// 	fmt.Println(i)
// 	//休眠
// 	time.Sleep(time.Millisecond * 100)
// 	if i== 100{
// 		break
// 	}
// }
//
//unix和unixnano的使用
fmt.Printf("unix时间戳=%v unixnano时间戳=%v",now.Unix(),now.UnixNano())
//输出有以下结果：
//unix时间戳=1692792690 unixnano时间戳=1692792690199200800



}