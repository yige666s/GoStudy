package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 建立与服务端的链接
// 进行数据收发
// 关闭链接

func EchoCli() {
	conn, err := net.Dial("tcp", "127.0.0.1:8099")
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	defer conn.Close() // 关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputinfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputinfo) == "Q" { // 输入q客户端退出
			return
		}
		_, err = conn.Write([]byte(inputinfo)) // 发送输入的数据
		if err != nil {
			return
		}

		buf := [512]byte{}          // 接收缓冲区
		n, err := conn.Read(buf[:]) // 接收数据到缓冲区
		if err != nil {
			fmt.Println("recv failed : ", err)
			return
		}
		fmt.Println(string(buf[:n])) // 打印接收到的数据
	}
}
