package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
)

// singleflight处理缓存击穿
// 它提供了重复函数调用抑制机制，使用它可以避免同时进行相同的函数调用。
// 第一个调用未完成时后续的重复调用会等待，当第一个调用完成时则会与它们分享结果，
// 这样以来虽然只执行了一次函数调用但是所有调用都拿到了最终的调用结果。

func getData(id int64) string {
	fmt.Println("query...")
	time.Sleep(10 * time.Second) // 模拟一个比较耗时的操作
	return "hello world!"
}
func DoChanGetData(ctx context.Context, g *singleflight.Group, id int64) (string, error) {
	// DoChan第二个参数是对 key函数"getData"的逻辑封装
	ch := g.DoChan("getData", func() (interface{}, error) { // 在没有其他正在进行的相同 key 的调用时应执行的实际操作
		go func() {
			time.Sleep(100 * time.Millisecond) // 100ms后忘记key,后续重复调用可进行
			g.Forget("getData")
		}()
		ret := getData(id)
		return ret, nil
	})
	// 为了避免第一次调用阻塞所有调用的情况，我们可以结合使用select和DoChan为函数调用设置超时时间。
	select {
	case <-ctx.Done(): // 如果 context 被取消（例如由于超时或手动取消），返回一个错误，表示操作被取消。
		return "", ctx.Err()
	case ret := <-ch: // 如果 DoChan 调用完成，返回结果和错误（如果有的话）。
		return ret.Val.(string), ret.Err
	}
}

func SingleflightTest() {
	g := new(singleflight.Group)
	// 第一次调用
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		v1, err := DoChanGetData(ctx, g, 1)
		fmt.Printf("v1:%v err:%v\n", v1, err)
	}()

	time.Sleep(2 * time.Second)

	// 第二次调用无法进行，等待第一次调用结束赋值
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	v2, err := DoChanGetData(ctx, g, 1)
	fmt.Printf("v2:%v err:%v\n", v2, err)
}

// singleflight将并发调用合并成一个调用的特点决定了它非常适合用来防止缓存击穿。
func getDataFromCache(key string) (string, error) {
	time.Sleep(5 * time.Second)
	if key == "cache" {
		return "cache", nil
	}
	if key == "DB" {
		return "NOT found", errors.New("NOTOUND")
	}

}
func getDataSingleFlight(key string) (interface{}, error) {
	g := new(singleflight.Group)
	value, err, _ := g.Do(key, func() (interface{}, error) {
		// 查缓存
		data, err := getDataFromCache(key)
		if err == nil { // 查到缓存，返回缓存
			return data, nil
		}
		if err == NOTFOUND { // 缓存不命中，查DB
			data, err := getDataFromDB(key)
			if err == nil { // 查询数据库成功，设置缓存，返回数据
				setCache(data)
				return data, nil
			}
			return nil, err // 查询数据库失败，返回错误
		}
		return nil, err // 查缓存出现其他错误直接返回，防止错误传递至DB
	})
	if err != nil {
		return nil, err
	}
	return value, err
}
