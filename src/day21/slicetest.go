package main

import (
	"fmt"
)

func printSliceInfo(s []int) {
	fmt.Printf("长度: %d, 容量: %d, 值: %v\n", len(s), cap(s), s)
}
func slicetest() {
	// 初始化一个容量为2的slice
	slice := make([]int, 0, 2)
	s1 := make(map[string]struct{})    // 如果只关心键的存在性，使用 map[string]struct{} 是更好的选择，因为它更加高效
	s2 := make(map[string]interface{}) //表示可以存储任意类型的值。因此，这种写法用于需要存储不同类型的值的场景。通过 interface{}，map 可以存储任何类型的值，比如整数、字符串、结构体等。
	printSliceInfo(slice)              // 长度: 0, 容量: 2

	// 向slice添加元素
	slice = append(slice, 1)
	printSliceInfo(slice) // 长度: 1, 容量: 2

	slice = append(slice, 2)
	printSliceInfo(slice) // 长度: 2, 容量: 2

	// 添加第三个元素，触发扩容
	slice = append(slice, 3)
	printSliceInfo(slice) // 长度: 3, 容量: 4

	// 再次添加元素
	slice = append(slice, 4)
	printSliceInfo(slice) // 长度: 4, 容量: 4

	// 添加第五个元素，再次触发扩容
	slice = append(slice, 5)
	printSliceInfo(slice) // 长度: 5, 容量: 8
}

// 解释：
// 初始 Slice：我们使用 make([]int, 0, 2) 初始化了一个长度为 0、容量为 2 的 slice。
// 添加元素：每次使用 append 向 slice 添加元素时，如果 slice 的容量足够大，Go 语言不会调整容量，只会增加长度。
// 触发扩容：当我们向 slice 中添加第三个元素时，发现容量不够（当前容量为 2），Go 语言会自动扩容。扩容后，容量翻倍，从 2 增加到 4。
// 再次扩容：当我们添加第五个元素时，容量再次不够（当前容量为 4），Go 语言再次扩容，容量从 4 增加到 8。
// 扩容机制：
// Go 语言的 slice 扩容机制通常是按容量的 2 倍增长。当容量超过 1024 时，增长因子会逐渐减小。
// 在扩容时，Go 会分配一个更大的底层数组，并将原 slice 的元素复制到新的数组中。
