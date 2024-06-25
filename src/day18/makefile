.PHONY: all build run gotool clean help	 #.PHONY用来定义伪目标。不创建目标文件，而是去执行这个目标下面的命令。

BINARY="myapp"	# BINARY：定义一个变量，表示生成的二进制文件的名称。


all:gotool build # 默认目标，如果不指定任何目标，Make 会执行这个目标

build:
	CGO_ENABLED=0 GOOS=linux GORACH=amd64 go build -o ${BINARY}

run:	# 使用 go run 命令执行当前目录下的 Go 程序。@ 符号表示不打印命令本身。
	@go run ./	

gotool:	# 运行 Go 工具进行代码格式化和静态检查。
	go fmt ./ 
	go vet ./

clean: # 清理生成的二进制文件
	@if [ -f ${BINARY}] ; then rm ${BINARY} ; fi

help: # 使用 echo 打印帮助信息。@ 符号表示不打印命令本身。
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"