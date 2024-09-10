package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip string
	Port int
}

// 创建服务器
func NewServer(Ip string, Port int) *Server {
	server := &Server {
		Ip: Ip,
		Port: Port,
	}

	return server
}

func (this *Server) Handler(con net.Conn) {
	fmt.Println("connection succeed!")
}

func (this *Server) Start() {
	listener, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("server start failed !")
		return
	}
	defer listener.Close()

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

