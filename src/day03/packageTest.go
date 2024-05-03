package main

import (
	"fmt"

	//Go语言中禁止循环导入包
	d "day03/demo" // 给引入的demo包起一个新的名字 d

	// 如果引入一个包的时候为其设置了一个特殊_作为包名，那么这个包的引入方式就称为匿名引入。
	// 一个包被匿名引入的目的主要是为了加载这个包，从而使得这个包中的资源得以初始化。
	// 被匿名引入的包中的 [init函数将被执行并且仅执行一遍]
	_ "github.com/go-sql-driver/mysql"
)

// 本文介绍了Go语言中如何定义包、如何导出包的内容及如何导入其他包

func packageTest() {
	fmt.Println(d.Mode)
	ret := d.Add(1, 2)
	fmt.Println(ret)
	d.SayHi()
}
