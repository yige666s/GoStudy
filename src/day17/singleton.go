package main

import "sync"

// 在标准库sync中找到了Once类型。它能保证某个操作仅且只执行一次。
// once.Do(func()) 内部实现使用的 check-lock-check + atomic 方法,是安全地实现此目标的首选方式
// 所有刚转到Go语言的新开发人员都必须真正了解并发安全性如何工作以更好地改进其代码

type Singleton struct{}

var instance *Singleton
var once sync.Once

func GetInstaance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
