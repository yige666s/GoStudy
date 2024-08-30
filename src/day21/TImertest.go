package main

import (
	"fmt"
	"sync"
	"time"
)

type Mytimer struct {
	pp       uintptr            // goroutinue调度处理器
	when     int64              // 执行时间
	period   int64              // 执行周期
	f        func(any, uintptr) // 回调函数
	arg      any                // 回调函数参数
	seq      uintptr            // 定时器序列号
	nextwhen int64              // 下次执行时间
	status   uint32             // 定时器状态
}

// 一次性定时器
func timerTest() {
	timer := time.NewTimer(5 * time.Second)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		<-timer.C
		fmt.Println("hello world")
		wg.Done()
	}()

	wg.Wait()
}

// 重复定时任务
func timerTicker() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println("hello world")
		}
	}
}

// goroutinue + sleep
func sleepTest() {
	go func() {
		for {
			time.Sleep(3 * time.Second)
			fmt.Println("hello world")
		}
	}()
	select {}
}
