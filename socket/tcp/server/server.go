package main
import (
	"fmt"
	"net" //做网络socket开发时。net包含有我们需要所有的方法和函数
	// "io"
)

func process(conn net.Conn) {
	//这里我们循环的接收客户端发送的数据
	defer conn.Close() //关闭conn

	for {
		//创建一个新的切片
		buf := make([]byte,1024)
		//1.等待客户端通过conn发送信息
		//2.如果客户端没有write[发送]，那么协程就阻塞在这里
		fmt.Printf("服务器在等待客户端%s 发送信息"+ conn.RemoteAddr().String()+ " ")
		n,err :=conn.Read(buf) //从conn读取
		if err != nil {
			fmt.Printf("客户端退出 err=%v\n",err)
			return // !!!
		}
		//3.显示客户端发送的内容到服务器的终端
		fmt.Println(string(buf[:n]))
		
	}
}
func main() {
	fmt.Println("服务器开始监听")
	//net.Listen("tcp","0.0.0.0:8888")
	//1.tcp表示使用网络协议是tcp
	//2.0.0.0.0:8888 表示在本地监听8888端口
	listen, err :=net.Listen("tcp","0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err")
		return
	}
	defer listen.Close()  //延时关闭listen

	//循环等待客户端连接我
	for {
		//等待客户端连接诶
	    fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=",err)
		}else {
			fmt.Println("Accept() suc conn=%v 客户端ip为=%v\n",conn,conn.RemoteAddr().String())
		}
		//这里准备启动一个协程为客户端服务

		go process(conn)
	}
	// fmt.Printf("Listen successfully=%v\n",listen)
}