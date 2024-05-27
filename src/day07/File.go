package main

import (
	"fmt"
	"io"
	"os"
)

func FileTest() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("Open failed")
		return
	}
	defer file.Close()
	var buf = make([]byte, 128)
	n, err := file.Read(buf)
	if err == io.EOF {
		fmt.Println("Read Over")
		return
	}
	if err != nil {
		fmt.Println("read file failed,err :", err)
		return
	}

	fmt.Printf("已经读取了%d个字节数据\n", n)
	fmt.Println(string(buf[:n]))
}
