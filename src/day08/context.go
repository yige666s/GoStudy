package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// Context类型，专门用来简化 对于处理单个请求的多个 goroutine 之间
// 与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。

// func gen(ctx context.Context) <-chan int {
// 	dst := make(chan int)
// 	n := 1
// 	go func() {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			case dst <- n:
// 				n++
// 			}
// 		}
// 	}()
// 	return dst
// }

// func ContextTest() {
// 	// Background(),返回顶级Context
// 	// WithCacel,WithDeadline,WithTimeout,Withvalue
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	for n := range gen(ctx) {
// 		fmt.Println(n)
// 		if n == 5 {
// 			break
// 		}
// 	}
// }

// 调用服务端API时在客户端实现超时控制

// // Server
// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	num := rand.Intn(2) // 返回[0,2)随机数
// 	if num == 0 {
// 		time.Sleep(time.Second * 5)
// 		fmt.Fprintf(w, "resp slow")
// 		return
// 	}
// 	fmt.Fprintf(w, "resp quick")
// }

// func server() {
// 	http.HandleFunc("/", indexHandler)
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// Client
type respdata struct {
	resp *http.Response
	err  error
}

func docall(ctx context.Context) {
	trans := http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: &trans,
	}
	respChan := make(chan *respdata, 1)
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000", nil)
	if err != nil {
		fmt.Println("new request failed,err", err)
		return
	}

	//使用带有超时上下文创建新的request
	req = req.WithContext(ctx)

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		resp, err := client.Do(req) // 发送请求返回相应
		if err != nil {
			fmt.Println("request failed", err)
		}
		fmt.Println("resp is ", resp)
		ret := &respdata{
			resp: resp,
			err:  err,
		}
		respChan <- ret
		wg.Done()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("call api timeout")
	case result := <-respChan:
		fmt.Println("success")
		if result.err != nil {
			fmt.Println("call server api failed", result.err)
			return
		}
		defer result.resp.Body.Close()
		data, _ := io.ReadAll(result.resp.Body)
		fmt.Printf("resp:%v\n", string(data))
	}
}

func Client() {
	// 定义超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	// 调用cancel释放goroutine资源
	defer cancel()
	docall(ctx)
}
