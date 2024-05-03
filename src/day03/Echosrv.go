package main

import (
	"bufio"
	"fmt"
	"net"
)

// 因为Go语言中创建多个goroutine实现并发非常方便和高效，可以每建立一次链接就创建一个goroutine去处理。
// 监听端口
// 接收客户端请求建立链接
// 创建goroutine处理链接

func Worker(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn) // 创建协程处理
		var buf [128]byte               // 缓冲区
		n, err := reader.Read(buf[:])   // 读取数据到buf
		if err != nil {
			fmt.Println("read from clecnt failed, err:", err)
			break
		}
		recvstr := string(buf[:n]) // 读取缓冲区中的数据
		fmt.Println("收到数据: ", recvstr)
		conn.Write([]byte(recvstr)) // 发送数据
	}
}

func Echosrv() {
	listen, err := net.Listen("tcp", "127.0.0.1:8099")
	if err != nil {
		fmt.Println("listen failed: ", err)
		return
	}

	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed: ", err)
			continue
		}
		Worker(conn) // 启动一个goroutine处理连接
	}
}
