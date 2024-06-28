package main

import "unsafe"

// 因为 CPU 访问内存时，并不是逐个字节访问，而是以字（word）为单位访问。比如 64位CPU的字长（word size）为8bytes，
// 那么CPU访问内存的单位也是8字节，每次加载的内存数据也是固定的若干字长，如8words（64bytes）、16words(128bytes）等。

type Foo struct {
	A int8 // 1
	B int8 // 1
	C int8 // 1
}

type Bar struct { // struct的内存对齐情况与字段的声明顺序有关
	x int32 // 4
	y *Foo  // 8
	a Foo   // 3
	b Foo   // 4
}

type demo0 struct { // C++中空类大小为1;每个空结构体实例在内存中都有唯一的地址，尽管它们不占用实际的内存空间。
	a struct{}
}

type demo1 struct {
	m struct{} // 0
	n int8     // 1
}

type demo2 struct { // 空结构体类型作为结构体的最后一个字段时，
	// 如果有指向该字段的指针，那么就会返回该结构体之外的地址。为了避免内存泄露会额外进行一次内存对齐。
	n int8
	m struct{}
}

var set map[int]struct{} // 使用空结构体作为map的值来实现一个类似 Set 的数据结构

// 空结构体通道通常用于传递信号，而不是实际的数据。这种信号传递可以用来通知事件的发生，例如超时或取消操作
// 空结构体（struct{}）不占用任何内存空间，因此在通道上传递空结构体的开销非常小。
// 使用空结构体可以明确表达仅需信号传递，而不需要传递实际数据的意图
var Interrupted = make(chan struct{})

// 在 x86 平台上原子操作需要强制内存对齐是因为在 32bit 平台下进行 64bit 原子操作要求必须 8 字节对齐，否则程序会 panic

func MemModel() {
	// var f Foo
	// fmt.Println(unsafe.Sizeof(f))
	// var b Bar
	// fmt.Println(unsafe.Sizeof(b))
	// fmt.Println(unsafe.Alignof(b))
	// fmt.Println(unsafe.Alignof(b.y))
	// var a1 demo0
	// var a2 demo1
	// var a3 demo2
	// fmt.Println(unsafe.Sizeof(a1)) // 0
	// fmt.Printf("%p\n", &a1)
	// fmt.Println(unsafe.Sizeof(a2)) // 1
	// fmt.Println(unsafe.Sizeof(a3)) // 2
}

// 缓存系统中是以缓存行（cache line）为单位存储的。缓存行是2的整数幂个连续字节，一般为32-256个字节。最常见的缓存行大小是64个字节。
// 当多线程修改互相独立的变量时，如果这些变量共享同一个缓存行，就会无意中影响彼此的性能，这就是伪共享(false sharing)
// 当不同的线程同时读写同一个cache line上不同数据时就可能发生false sharing。false sharing会导致多核处理器上严重的系统性能下降。
// 在一些需要防止CacheLine伪共享的时候，也需要进行特殊的字段对齐
type poolLocal struct {
	poolLocalInternal

	// Prevents false sharing on widespread platforms with
	// 128 mod (cache line size) = 0 .
	pad [128 - unsafe.Sizeof(poolLocalInternal{})%128]byte
}

// hot path 是指执行非常频繁的指令序列。
// 如果要访问结构体的其他字段，除了结构体指针外，还需要计算与第一个值的偏移(calculate offset)。
// 在机器码中，偏移量是随指令传递的附加值，CPU 需要做一次偏移值与指针的加法运算，才能获取要访问的值的地址。因为访问第一个字段的机器代码更紧凑，速度更快。
// 通过将常用字段放置在结构体的第一个位置上减少CPU要执行的指令数量，从而达到更快的访问效果。
// src/sync/once.go
type Once struct {
	done uint32
	m    Mutex
}
