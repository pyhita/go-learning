package main

import "net"

type User struct {
	Addr string
	Name string
	Channel chan string
	conn net.Conn
}

func NewUser(conn net.Conn) *User {
	addr := conn.RemoteAddr().String()

	user := &User {
		Addr: addr,
		Name: addr,
		Channel: make(chan string),
		conn: conn,
	}

	// 开启协程监听user channel
	go user.ListenMessage()

	return user;
}


func (this *User) ListenMessage() {
	for {
		msg := <-this.Channel
		this.conn.Write([]byte(msg + "\n"))
	}
}