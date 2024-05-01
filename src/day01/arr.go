package main

func Arr() {
	// var a [3]int                                // 大小不可变
	// var b = [3]string{"liming", "jack", "lucy"} // 初始化时赋值使用列表必须使用 var T =
	// var c = [...]int{1, 2, 3}                   // 自动推导长度

	// var d = [3][2]int{
	// 	{1, 2},
	// 	{4, 5},
	// 	{6, 7},
	// }

	// //支持的写法
	// a1 := [...][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }
	// 不支持多维数组的内层使用... ，无法自动推导数量
	// b1 := [3][...]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }

	// var t1 [2]*int // 指针数组
	// var t2 *[3]int // 数组指针

	// 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
}

func SumofIndex(arr []int, target int) [][]int {
	indexs := [][]int{}

	indexMap := make(map[int]int) // 存储元素和下标

	for i, v := range arr {
		diff := target - v

		if index, found := indexMap[diff]; found { // 查找是否存在差值
			indexs = append(indexs, []int{index, i}) // 添加两个对应位置的索引
		}

		indexMap[v] = i // 存储元素值和下标
	}
	return indexs
}
