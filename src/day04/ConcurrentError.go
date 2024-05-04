package main

import (
	"fmt"
	"time"
)

// 在启用 goroutine 去执行任务的场景下，如果想要 recover goroutine中可能出现的 panic 就需要在 goroutine 中使用 recover

func f2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover outer panic:%v\n", r)
		}
	}()
	// 开启一个goroutine执行任务
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recover inner panic:%v\n", r)
			}
		}()
		fmt.Println("in goroutine....")
		// 只能触发当前goroutine中的defer
		panic("panic in goroutine")
	}()

	time.Sleep(time.Second)
	fmt.Println("exit")
}

// errgroup包能为处理公共任务的子任务而开启一组 goroutine 提供同步、error 传播和基于context 的取消功能。
// fetchUrlDemo2 使用errgroup并发获取url内容
// func fetchUrlDemo2() error {
// 	g := new(errgroup.Group) // 创建等待组（类似sync.WaitGroup）
// 	var urls = []string{
// 		"http://pkg.go.dev",
// 		"http://www.liwenzhou.com",
// 		"http://www.yixieqitawangzhi.com",
// 	}
// 	for _, url := range urls {
// 		url := url // 注意此处声明新的变量
// 		// 启动一个goroutine去获取url内容
// 		g.Go(func() error {
// 			resp, err := http.Get(url)
// 			if err == nil {
// 				fmt.Printf("获取%s成功\n", url)
// 				resp.Body.Close()
// 			}
// 			return err // 返回错误
// 		})
// 	}
// 	if err := g.Wait(); err != nil {
// 		// 处理可能出现的错误
// 		fmt.Println(err)
// 		return err
// 	}
// 	fmt.Println("所有goroutine均成功")
// 	return nil
// }
