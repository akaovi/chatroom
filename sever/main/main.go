package main

import (
	"chatroom_sever/sever/process"
	"fmt"
	"net"
)

func processall(conn net.Conn) {
	fmt.Println("客户端: ", conn.LocalAddr().String(), "接入成功...")
	defer conn.Close()
	pm := &process.ProcessMain{
		Conn: conn,
	}
	pm.Processmain()

}

func main() {
	fmt.Println("服务器在9977端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:9977")
	if err != nil {
		fmt.Println("Dial err: ", err)
		return
	}
	defer listen.Close()

	// 监听成功, 接收客户端的链接
	for {
		fmt.Println("等待客户端的连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(conn.LocalAddr().String(), "连接失败...")
		}

		go processall(conn)
	}
}
