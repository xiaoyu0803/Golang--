package server

import (
	"chat/user"
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Server struct {
	Ip            string
	Port          int
	OnlineUserMap map[string]*user.User
	MapLock       sync.RWMutex
	msg           chan string
}

//创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:            ip,
		Port:          port,
		OnlineUserMap: make(map[string]*user.User),
		msg:           make(chan string),
	}

	return server
}

func (this *Server) ListenMsg() {
	for {
		msg := <-this.msg
		this.MapLock.Lock()
		for _, cli := range this.OnlineUserMap {
			cli.C <- msg
		}
		this.MapLock.Unlock()
	}
}

func (this *Server) BroadcastMsg(user *user.User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.msg <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	//...当前链接的业务
	log.Println("链接建立成功")
	user := user.NewUser(conn)
	this.MapLock.Lock()
	this.OnlineUserMap[user.Name] = user
	this.MapLock.Unlock()
	this.BroadcastMsg(user, "已上线")
	var islive = make(chan bool)
	//接受客户端消息
	go func() {
		buf := make([]byte, 10240)

		for {
			n, err := conn.Read(buf)
			if n == 0 {
				this.BroadcastMsg(user, "下线了")
				this.MapLock.Lock()
				delete(this.OnlineUserMap, user.Name)
				this.MapLock.Unlock()
				return
			}
			if err != nil && err != io.EOF {
				log.Println("err", err)
				return
			}
			msg := string(buf[:n-1])
			if msg == "who" {
				this.MapLock.Lock()
				msgOnline := ""
				for k := range this.OnlineUserMap {
					msgOnline = msgOnline + k + "\n"
				}

				//查询在线
				user.Sendmsg(msgOnline)
				this.MapLock.Unlock()
			} else if len(msg) > 7 && msg[:7] == "rename|" {
				this.MapLock.Lock()
				newname := strings.Split(msg, "|")[1]
				if _, ok := this.OnlineUserMap[newname]; ok {
					user.Sendmsg("名字已存在")
				}
				delete(this.OnlineUserMap, user.Name)

				this.OnlineUserMap[newname] = user
				user.Name = newname
				user.Sendmsg(fmt.Sprintf("已经改名为%s", user.Name))
				this.MapLock.Unlock()
			} else if len(msg) > 4 && msg[:3] == "to|" {
				this.MapLock.Lock()
				toname := strings.Split(msg, "|")[1]
				tomsg := strings.Split(msg, "|")[2]
				if toname == "" {
					user.Sendmsg("私聊名为空")
				}
				touser, ok := this.OnlineUserMap[toname]
				if !ok {
					user.Sendmsg("用户名不存在")
					this.MapLock.Unlock()
					continue
				}
				touser.Sendmsg(fmt.Sprint(user.Name, ":", tomsg))
				this.MapLock.Unlock()
			} else {
				this.BroadcastMsg(user, msg)
				islive <- true
			}
		}

	}()
	for {

		select {
		case <-islive:
		case <-time.After(time.Second * 100):
			user.Sendmsg("你被踢了")
			close(user.C)
			conn.Close()
			runtime.Goexit()
		}
	}
}

//启动服务器的接口
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//close listen socket
	defer listener.Close()
	go this.ListenMsg()
	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		//do handler
		go this.Handler(conn)
	}
}
