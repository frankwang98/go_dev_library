package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建一个 UDP 地址结构
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// 创建一个 UDP 连接
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer conn.Close()

	// 向服务器发送消息
	msg := []byte("Hello, Server!")
	_, err = conn.Write(msg)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// 接收服务器的响应消息
	buf := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// 将服务器的响应消息打印到控制台
	fmt.Println("Received Message:", string(buf[:n]))
}
