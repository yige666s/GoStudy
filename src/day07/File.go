package main

import (
	"bufio"
	"flag"
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

func FileTest2() {
	content, err := os.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed, ", err)
	}
	fmt.Println(content)
}

func FileWrite() {
	str := "file write test"
	err := os.WriteFile("text.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("file write failed, err", err)
	}
}

func FileWrite2() {
	file, _ := os.OpenFile("text2.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	defer file.Close()
	writer := bufio.NewWriter(file)           // 参数是一个文件句柄
	writer.WriteString("bufioWrite string\n") // 将数据写入缓存
	writer.Flush()                            //将缓存中的内容写入文件
}

func FileCopy() {
	src, err := os.OpenFile("text.txt", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("src open faield")
		return
	}
	defer src.Close()

	dst, err := os.OpenFile("text2.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("dst open failed")
		return
	}
	defer dst.Close()

	HasCopy, _ := io.Copy(dst, src)
	fmt.Println(HasCopy)
}

func Cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n') // 遇到'\n'停止读取
		if err != nil {
			fmt.Fprintf(os.Stdout, "%s", buf)
			break
		}
		fmt.Fprintf(os.Stdout, "%s\n", buf)
	}
}

func CatTest() {
	flag.Parse()          // 从命令行解析参数
	if flag.NArg() == 0 { // 没有命令行参数
		Cat(bufio.NewReader(os.Stdin)) // 从标准输入接收
	}
	for i := 0; i < flag.NArg(); i++ { // NArg返回命令函参数数量
		f, err := os.Open(flag.Arg(i)) // Arg返回第i个命令行参数, f为命令行接受的文件名
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed\n", flag.Arg(i))
			continue
		}
		Cat(bufio.NewReader(f)) //从文件中读取内容并打印
	}
}
