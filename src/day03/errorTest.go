package main

import (
	"errors"
	"fmt"
)

// Go 语言中的错误处理与其他语言不太一样，它把错误当成一种值来处理，更强调判断错误、处理错误，而不是一股脑的 catch 捕获异常
var err = errors.New("EOF")

func PrintError() {
	fmt.Errorf("到达文件尾部,err:%w", err)
}

// OpError 自定义结构体类型
type OpError struct {
	Op string
}

// Error OpError 类型实现error接口
func (e *OpError) Error() string {
	return fmt.Sprintf("无权执行%s操作", e.Op)
}
