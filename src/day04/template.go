package main

import (
	"fmt"
)

func min[T int | float64](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

func templateTest() {
	// fmin := min[int]
	y := min[float64](1.1, 2.2)
	x := min(1, 2)
	fmt.Println(x, y)
}

type slice[T int | string] []T //自定义int/string类型的切片类型
type Map[K int | string, V int | string] map[K]V
type Tree[T interface{}] struct { // T类型是一个空接口，可以接收任何类型数据
	left, right *Tree[T] // 左右子树用数组指针表示
	value       T
}

type T interface { // 定义类型集合
	int | string | bool | float32
}
