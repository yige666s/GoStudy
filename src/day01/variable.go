package main

// func variable() {

// 	// Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。每个变量会被初始化成其类型的默认值，
// 	// 例如： 整型和浮点型变量的默认值为  0  。字符串变量的默认值为  空字符串  。
// 	// 布尔型变量默认为  false  。 切片、函数、指针变量的默认为  nil  。
// 	var name string
// 	var age int
// 	var isOK bool
// 	var (
// 		a string
// 		b int
// 		c bool
// 	)
// 	fmt.Println(name, age, isOK, a, b, c)

// 	var d string = "jack"
// 	var e int = 18
// 	fmt.Println(d, e)

// 	f := 10
// 	fmt.Println(f)
// }

// // 函数外的每个语句都必须以关键字开始（var、const、func等）
// // :=不能使用在函数外。
// // _多用于占位，表示忽略值。

// var g int = 100

// const h int = 1.000
// const (
// 	a = 100.00
// 	b // 与a的值相同
// )

// // iota是go语言的常量计数器，只能在常量的表达式中使用。
// // iota在const关键字出现时将被重置为0。
// // const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)
// // 使用iota能简化定义，在定义枚举时很有用。

// const (
// 	n1 = iota // 0
// 	n2 = 100  // 100
// 	n3 = iota // 2
// 	n4        // 3
// )
// const (
// 	a1, b1 = iota + 1, iota + 2 // 1, 2
// 	c1, d1
// 	e1, f1
// )
