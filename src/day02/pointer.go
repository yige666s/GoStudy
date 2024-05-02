package main

import "fmt"

// 区别于C/C++中的指针，Go语言中的指针不能进行偏移和运算，是安全指针。
// &（取地址）和*（根据地址取值）

func Addr1() {
	a := 10
	b := &a
	fmt.Printf("%p %p\n", b, &a)
	c := *b
	fmt.Println(c)
}

// 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
// 而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
// Go语言中new和make是内建的两个函数，主要用来分配内存。
func Mem() {
	// var a *int只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值
	// var a *int
	a := new(int)
	*a = 100
	fmt.Println(*a)
}

// make也是用于内存分配的，区别于new，它只用于slice、map以及channel的内存创建，而且它返回的类型就是这三个类型本身
// func make(t Type, size ...IntegerType) Type
// 在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作

func MakeTest() {
	b := make(map[string]int, 10)
	b["lili"] = 100
	b["jack"] = 90
	fmt.Println(b)
}
