package main

import "fmt"

func controller() {
	if isture := true; isture { // 先赋值再判断
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	for { // 无限循环
		count := 0
		if count == 100 {
			break
		}
		fmt.Println(count)
		count++
	}

	for range 2 {
		fmt.Println("hello world") // 打印两遍hello world
	}

	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中值")
	default:
		fmt.Println("无效输入")
	}

	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	}
}

func PrintMutilList() {
	for i := 0; i < 9; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%d*%d=%2d\n", j, i, j*i)
		}
		fmt.Println() // 打印换行
	}
}
