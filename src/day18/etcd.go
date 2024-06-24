package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
)

// etcd是近几年比较火热的一个开源的、分布式的键值对数据存储系统，
// 提供共享配置、服务的注册和发现，本文主要介绍etcd的安装和使用。
// etcd是近几年比较火热的一个开源的、分布式的键值对数据存储系统，提供共享配置、服务的注册和发现，本文主要介绍etcd的安装和使用。

func Put_Get() {
	// 连接etcd
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = cli.Put(ctx, "jack", "abc")
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "q1mi")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}

	// watch
	rch := cli.Watch(context.Background(), "jack") // <-chan WatchResponse

	// 使用 for 循环遍历 rch 通道，逐一处理接收到的 WatchResponse（监视响应）。
	// 每当键 "jack" 发生变化（例如，值被设置、删除等），rch 通道会收到一个新的 WatchResponse。
	for wresp := range rch {
		for _, ev := range wresp.Events { // 遍历 WatchResponse 中的所有事件,
			// 每个 WatchResponse 中包含一个 Events 字段，它是一个事件列表，表示键的变化。
			// 对于每个事件 ev：
			// ev.Type：事件类型，表示是键值对的创建、更新还是删除。
			// ev.Kv.Key：发生变化的键。
			// ev.Kv.Value：键的新值（如果是删除事件，这个值可能为空）。
			fmt.Printf("Type :%s key: %s value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}

	// lease租约
	// 自动删除：由于键值对绑定了一个5秒的租约，当租约到期时（5秒后），etcd 服务器会自动删除这个键值对。
	resp1, err1 := cli.Grant(context.TODO(), 5)
	if err1 != nil {
		log.Fatal(err)
	}

	// 5秒钟之后, /nazha/ 这个key就会被移除
	_, err = cli.Put(context.TODO(), "/nazha/", "dsb", clientv3.WithLease(resp1.ID))
	if err != nil {
		log.Fatal(err)
	}

	// keepalive
	// 通过创建一个租约并持续续租，确保与该租约绑定的键值对不会过期
	ch, err2 := cli.KeepAlive(context.TODO(), resp1.ID)
	if err2 != nil {
		log.Fatal(err2)
	}
	for {
		// the key 'foo' will be kept forever
		ka := <-ch
		fmt.Println("ttl:", ka.TTL)
	}
}

// 分布式锁
func DisturbuteLock() {
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 创建两个单独的会话用来演示锁竞争
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	m1 := concurrency.NewMutex(s1, "/my-lock/")

	s2, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s2.Close()
	m2 := concurrency.NewMutex(s2, "/my-lock/")

	// 会话s1获取锁
	if err := m1.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("acquired lock for s1")

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		// 等待直到会话s1释放了/my-lock/的锁
		if err := m2.Lock(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	if err := m1.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("released lock for s1")

	<-m2Locked
	fmt.Println("acquired lock for s2")

}
