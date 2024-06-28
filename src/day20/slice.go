package main

import "fmt"

func SliceOperate() {
	// 赋值
	a := []int{1, 2, 3}
	b := make([]int, len(a))  // 申请内存
	copy(b, a)                // 拷贝内容
	b = append([]int{}, a...) // 将a拷贝到空slice中赋给b
	// a[:0:0] 这部分代码创建了一个新的切片，它的长度和容量都为0。这确保了append操作不会复用原始切片a的底层数组，而是创建一个新的底层数组。
	// append函数会将 a 中的所有元素追加到新创建的长度和容量都为0的切片中，这样最终得到一个与 a 内容相同但底层数组不同的切片 b
	b = append(a[:0:0], a...)
	fmt.Println(b)

	// 剪切
	i := 1
	j := 2
	a = append(a[:i], a[j:]...) // 将i-j位置的元素剪切掉

	// 删除
	a = append(a[:i], a[i+1:]...) // 删除索引为i的位置元素
	a[i] = a[len(a)-1]            // 使用最后一个元素覆盖索引为i的位置的元素，但会破坏元素顺序
	a = a[:len(a)-1]              // 截断最后一个元素

	// 内存泄漏处理

}
