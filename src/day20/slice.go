package main

import (
	"fmt"
	"sort"
)

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

	//TODO 内存泄漏处理
	var c []*int32
	copy(c[i:], c[j:])
	for k, n := len(c)-(j-i), len(c); k < n; k++ {
		c[k] = nil
	}
	a = a[:len(a)-(j-i)]

	// copy(a[i:], a[i+1:])
	// a[len(a)-1] = nil // 或类型T的零值
	// a = a[:len(a)-1]

	// a[i] = a[len(a)-1]
	// a[len(a)-1] = nil
	// a = a[:len(a)-1]

	// 内部扩张
	a = append(a[:i], append(make([]int, j), a[i:]...)...) //先在ai之前添加j个元素，再将整体添加到前i个元素之后
	// 尾部扩张
	a = append(a, make([]int, j)...) // 尾部添加j个元素
	// 过滤 这里假设过滤的条件已封装为keep函数，使用for range遍历切片a的所有元素逐一调用keep函数进行过滤。
	// n := 0
	// for _, x := range a {
	// 	if keep(x) {
	// 		a[n] = x  // 保留该元素
	// 		n++
	// 	}
	// }
	// a = a[:n]  // 截取切片中需保留的元素
	// }

	// 插入，将元素x插入切片a的索引i处
	// a = append(a[:i],append([]int{x}, a[i:]...)...)	// 底层创建新切片

	// // 弹出
	// x, a := a[len(a)-1], a[:len(a)-1] // 弹出最后一个元素
	// x, a = a[0], a[1:]                // 弹出第一个元素

	// // 头插
	// a = append([]int{x}, a...)

	// 过滤但不分配内存
	// b = a[:0]
	// for _, v := range a {
	// 	if f(v) {
	// 		b = append(b, v)
	// 	}
	// }
	// for i = len(b); i < len(a); i++ {	// 对于需要释放的资源类型元素，将a中不符合的元素置空
	// 	a[i] = nil
	// }

	// 翻转 Reverse
	for i := 0; i < len(a)/2-1; i++ {
		j = len(a) - i - 1
		a[i], a[j] = a[j], a[i]
	}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j+1 {
		a[i], a[j] = a[j], a[i]
	}

	// 洗牌 打乱切片a中元素的顺序
	// for i = len(a) - 1; i > 0; i-- {
	// 	j := rand.Intn(i + 1)
	// 	a[i], a[j] = a[j], a[i]
	// }

	// rand.Shuffle(len(a), func(i, j int) {
	// 	a[i], a[j] = a[j], a[i]
	// })

	// 分批次
	actions := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	batchSize := 3
	batches := make([][]int, 0, (len(actions)+batchSize-1)/batchSize) // 确保即使元素总数不是batchsize的整数倍，容量也足够大以容纳所有批次

	for batchSize < len(actions) {
		actions = actions[batchSize:]                             // 更新actions以删除第一个batchSize元素
		batches = append(batches, actions[0:batchSize:batchSize]) // 从actions中取出第一个batchSize元素并附加到actions
	}
	batches = append(batches, actions) // 循环结束后，任何剩余的元素actions（将少于batchSize元素）都将作为最后一批附加。

	// 原地去重
	in := []int{3, 2, 1, 4, 3, 2, 1, 4, 1} // 切片元素可以是任意可排序的类型
	sort.Ints(in)
	j := 0 // pre指针
	for i := 1; i < len(in); i++ {
		if in[i] == in[j] {
			continue
		}
		j++
		// 需要保存原始数据时
		// in[i], in[j] = in[j], in[i]
		// 只需要保存需要的数据时
		in[j] = in[i]
	}
	result := in[:j+1]
	fmt.Println(result)

}

// TODO LRU : 存在就移到前面，不存在就插入到前
// moveToFront 把needle移动或添加到haystack的前面
func moveToFront(needle string, haystack []string) []string {
	if len(haystack) != 0 && haystack[0] == needle {
		return haystack
	}
	prev := needle
	for i, elem := range haystack {
		switch {
		case i == 0:
			haystack[0] = needle
			prev = elem
		case elem == needle:
			haystack[i] = prev
			return haystack
		default:
			haystack[i] = prev
			prev = elem
		}
	}
	return append(haystack, prev)
}

// var haystack = []string{"a", "b", "c", "d", "e"} // [a b c d e]
// haystack = moveToFront("c", haystack)         // [c a b d e]
// haystack = moveToFront("f", haystack)         // [f c a b d e]

// 滑动窗口 将切片input生成size大小的滑动窗口
func SlidingWindow(size int, input []int) [][]int {
	if len(input) < size {
		return [][]int{input}
	}
	// 以所需的精确大小分配切片
	r := make([][]int, 0, len(input)-size+1)
	for i, j := 0, size; j < len(input); i, j = i+1, j+1 {
		r = append(r, input[i:j])
	}
	return r
}

// a := []int{1, 2, 3, 4, 5}
// res := slidingWindow(2, a)
// fmt.Println(res)
