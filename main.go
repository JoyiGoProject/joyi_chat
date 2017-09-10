package main

import (
	_ "joyiyi_chat/routers"
	"net"

	"github.com/astaxie/beego"
)

func main() {
	//首先设置监听
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	CheckErr(err)
	defer listen_socket.Close()

	//循环接收数据
	for {
		conn, err := listen_socket.Accept()
		CheckErr(err)
		go ProcessInfo(conn)
	}
	beego.Run()
}

func ProcessInfo(conn net.Conn) {
	buf := make([]byte, 1024)
	defer conn.Close()

	//不断进行读取
	for {
		_, err := conn.Read(buf)
		CheckErr(err)
		beego.Info("Has this message: %s", string(buf))
	}
}

func CheckErr(err error) {
	if err != nil {
		beego.Error("Error：%s", err)
		return
	}
}
