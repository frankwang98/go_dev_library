package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建一个 UDP 地址结构
	addr, err := net.ResolveUDPAddr("udp", ":8888")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// 创建一个 UDP 连接监听器
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP Server Running...")

	for {
		// 接收来自客户端的消息
		buf := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		// 将接收到的消息打印到控制台
		fmt.Printf("Received message from %s: %s\n", remoteAddr.String(), string(buf[:n]))

		// 向客户端发送响应
		reply := []byte("Hi, client!")
		_, err = conn.WriteToUDP(reply, remoteAddr)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
	}
}
