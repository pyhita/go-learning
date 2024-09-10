package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp string
	ServerPort int
	Name string
	Conn net.Conn
}

func NewClinet(ip string, port int) *Client {
	client := &Client{
		ServerIp: ip,
		ServerPort: port,
	}

	con, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return nil
	}

	client.Conn = con
	return client
}

var ip string
var port int

// ./client -ip 127.0.0.1 -port 8888
func init() {
	// 绑定变量
	flag.StringVar(&ip, "ip地址", "127.0.0.1", "服务器IP地址，默认是(127.0.0.1)")
	flag.IntVar(&port, "端口号", 8888, "服务器端口号，默认是(8888)")
}


func main() {
	flag.Parse()
	client := NewClinet(ip, port)

	if client == nil {
		fmt.Println("连接服务器失败！")
		return
	}

	fmt.Println("连接服务器成功！")

	select {}
}