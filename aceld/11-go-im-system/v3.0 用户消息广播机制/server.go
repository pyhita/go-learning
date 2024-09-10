package main

import (
	"fmt"
	"net"
	"sync"
	"io"
)

type Server struct {
	Ip string
	Port int
	// 在线用户列表
	OnlineMap map[string]*User
	mapLock sync.RWMutex

	// 群发消息的channel
	Message chan string
}

// 创建服务器
func NewServer(Ip string, Port int) *Server {
	server := &Server {
		Ip: Ip,
		Port: Port,
		OnlineMap: make(map[string]*User),
		Message: make(chan string),
	}

	return server
}

func (this *Server) Handler(con net.Conn) {
	fmt.Println("connection succeed!")

	// 新建立一个user 加入到map
	user := NewUser(con)
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()
	// 向channel里面写入广播消息
	this.BroadCast(user, "已上线")

	go func() {
		// 从Socket中阻塞获取数据
		buf := make([]byte, 2048)
		len, err := con.Read(buf)

		if len == 0 {
			this.BroadCast(user, "下线")
			return
		} else if err != nil && err != io.EOF {
			fmt.Println("Read Failed.")
			return
		}

		msg := string(buf[:len-1])
		this.BroadCast(user, msg)
	} ()

	// 阻塞当前handler
	select {}
}

func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg;

	this.Message <- sendMsg
}

func (this *Server) Listener() {
	for {
		// 阻塞获取用户消息
		msg := <-this.Message

		for _, cli := range this.OnlineMap {
			// 广播给已上线用户的channel
			cli.Channel <- msg
		}
	}
}

func (this *Server) Start() {
	listener, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("server start failed !")
		return
	}
	defer listener.Close()

	go this.Listener()

	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println("server start failed !")
			continue
		}
		// 新建一个协程进行处理
		go this.Handler(con)
	}	
}

