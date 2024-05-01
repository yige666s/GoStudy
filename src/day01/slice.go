package main

import (
	"fmt"
	"sort"
)

// 切片（Slice）是一个拥有 相同类型元素 的 可变长度 的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
// 切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
func slicetest() {
	// var a = []string{"jack", "lili", "lucy"} // 不声明容量就是切片
	// fmt.Println(len(a))
	// fmt.Println(cap(a))

	// fmt.Println(a[:2])

	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))

	// 完整切片表达式需要满足的条件是0 <= low <= high <= max <= cap(a)，其他条件和简单切片表达式相同。
	//它会将得到的结果切片的容量设置为max-low。在完整切片表达式中只有第一个索引值（low）可以省略；它默认为0。
	a = [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))

	a1 := make([]int, 2, 10)
	fmt.Println(a1, len(a1), cap(a1))

	// 要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断。
	s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	fmt.Println(len(s3), cap(s3))

	// 下面的代码中演示了拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，即引用计数，这点需要特别注意。
	s1 := make([]int, 5, 10)
	s21 := s1
	s21[0] = 100
	fmt.Println(s21[0])

	// 切片的遍历方式和数组是一致的，支持索引遍历和for range遍历。
	for i, v := range s1 {
		fmt.Println(i, v)
	}

	// Go语言的内建函数append()可以为切片动态添加元素。 可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加…）。
	var s11 []int
	s11 = append(s11, 1)
	s11 = append(s11, 2, 3, 4)
	s12 := []int{5, 6, 7}
	s11 = append(s11, s12...)
	fmt.Println(s11)

	// 每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。
	// “扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值。
	// 切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。
	// 下方为$GOROOT/src/runtime/slice.go源码，可查看go的slice扩容机制

	// 切片的扩容策略
	newcap := old.cap
	doublecap := newcap + newcap
	if cap > doublecap {
		newcap = cap
	} else {
		if old.len < 1024 {
			newcap = doublecap
		} else {
			// Check 0 < newcap to detect overflow
			// and prevent an infinite loop.
			for 0 < newcap && newcap < cap {
				newcap += newcap / 4
			}
			// Set newcap to the requested cap when
			// the newcap calculation overflowed.
			if newcap <= 0 {
				newcap = cap
			}
		}
	}

	c1 := []int{1, 2, 3, 4, 5}
	c11 := make([]int, 5)
	copy(c11, c1) // copy是值传递，切片默认是引用传递
	c11[0] = 100
	fmt.Println(c1[0], c11[0])

	// 要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
	a12 := []int{1, 2, 3, 4, 5}
	a12 = append(a12[:2], a12[2:]...)
	fmt.Println(a12[2])

	//使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序（附加题，自行查资料解答）
	var aa = [...]int{3, 7, 8, 9, 1}
	sort.Ints(aa[:]) // 专门用于int类型的排序函数
	fmt.Println(aa)

}
