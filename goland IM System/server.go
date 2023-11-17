package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	//在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//消息广播的channel
	Message chan string
}

// NewServer 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// BroadCast 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

// ListenMessager 监听Message广播channel的goroutine,一旦有消息就发送给全部在线User
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		this.mapLock.Lock()
		//将msg发送给全部在线的User
		for _, cli := range this.OnlineMap {
			cli.c <- msg
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) Handler(conn net.Conn) {
	//当前链接的业务
	//fmt.Println("链接建立成功")

	user := NewUser(conn, this)

	user.Online()

	//监听用户是否活跃的channel
	isLive := make(chan bool)

	//接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Online()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}
			//提取用户的消息（去除\n)
			msg := string(buf[:n-1])

			////将得到的消息进行广播
			//this.BroadCast(user, msg)
			//用户针对msg进行消息处理
			user.DoMessage(msg)

			//用户的任意操作代表当前用户是一个活跃的
			isLive <- true
		}
	}()

	//当前handler阻塞
	for {
		select {
		case <-isLive:
			//当前用户是活跃的，应该重置定时器
			//不做任何事，为了激活select,更新下面定时器

		case <-time.After(time.Second * 300):
			//已超时
			//将当前的User强制的关闭
			user.SendMsg("你被踢了")
			//销毁资源
			close(user.c)

			//关闭连接
			conn.Close()

			//退出当前handle
			return //runtime.Goexit()
		}
	}
}

// Start 启动服务器的接口
func (this *Server) Start() {
	//TODO 1. socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//close listen socket
	defer listener.Close()

	//启动监听Message的goroutine
	go this.ListenMessager()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Listener accept err:", err)
			continue
		}
		//do handler
		go this.Handler(conn)
	}
}

//func main() {
//	server := NewServer("127.0.0.1", 8888)
//	server.Start()
//}
