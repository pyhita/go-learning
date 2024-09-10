package main

import (
	"net"
	"strings"
)

type User struct {
	Addr string
	Name string
	Channel chan string
	Conn net.Conn

	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	addr := conn.RemoteAddr().String()

	user := &User {
		Addr: addr,
		Name: addr,
		Channel: make(chan string),
		Conn: conn,
		server: server,
	}

	// 开启协程监听user channel
	go user.ListenMessage()

	return user;
}

// 用户上线
func (this *User) OnLine() {
	// 新建立一个user 加入到map
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	// 向channel里面写入广播消息
	this.server.BroadCast(this, "已上线")
}

// 用户下线
func (this *User) OffLine() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	this.server.BroadCast(this, "下线")
}

func (this *User) sendMessage(msg string) {
	this.Conn.Write([]byte(msg))
}

// 用户处理消息
func (this *User) DoMessage(msg string) {
	// 处理当前用户查询指令
	if msg == "who" {
		for addr, _ := range this.server.OnlineMap {
			sendMsg := "用户 [" + addr + "]" + "在线" + "\n"
			this.sendMessage(sendMsg)
		}
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 用户修改用户名
		newName := strings.Split(msg, "|")[1]
		// 判断用户名是否已经使用
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.sendMessage("当前用户名已经被使用\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.Name = newName
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.sendMessage("用户名修改成功了\n")
		}
	} else if len(msg) > 3 && msg[:3] == "to|" {
		// 用户私聊功能，消息格式：to|name|msg
		msgs := strings.Split(msg, "|")
		remoteName := msgs[1]
		this.server.mapLock.Lock()
		remoteUser, ok := this.server.OnlineMap[remoteName]
		this.server.mapLock.Unlock()

		if !ok {
			this.sendMessage("用户名输入存在，私聊用户不存在 \n")
			return
		}

		// 发送私聊消息
		sendMsg := remoteName + "对您说：" + "[" + msgs[2] + "] \n"
		remoteUser.sendMessage(sendMsg)
	} else {
		this.server.BroadCast(this, msg)
	}
}


// 监听用户channel的消息
func (this *User) ListenMessage() {
	for {
		msg := <-this.Channel
		this.Conn.Write([]byte(msg + "\n"))
	}
}