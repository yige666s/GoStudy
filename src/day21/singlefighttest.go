package main

import (
	"errors"

	"golang.org/x/sync/singleflight"
)

// 提供了重复函数调用抑制机制，使用它可以避免同时进行相同的函数调用。第一个调用未完成时后续的重复调用会等待，
// 当第一个调用完成时则会与它们分享结果，这样以来虽然只执行了一次函数调用但是所有调用都拿到了最终的调用结果。

func getDatafromCache(string) (string, error) {
	err := errors.New("test")
	return "hello,world", err
}

func singlefightTest(g *singleflight.Group, key string) (interface{}, error) {
	v, err, _ := g.Do(key, func() (interface{}, error) {
		// 查缓存
		data, err := getDatafromCache(key)
		if err == nil {
			return data, nil
		}

		if err == errNotFound {
			// 查数据库
			data, err := getDatafromDB(key)
			if err == nil {
				setCache(data)
				return data, nil
			}
			return nil, err
		}
		return nil, err
	})
	if err != nil {
		return nil, err
	}
	return v, nil
}
