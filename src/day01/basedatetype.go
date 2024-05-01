package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

func datatype() {
	// 十进制
	a := 10
	fmt.Printf("%d %b\n", a, a)

	// 八进制
	b := 0666
	fmt.Printf("%o\n", b)

	// 十六进制
	c := 0xffe
	fmt.Printf("%x %X\n", c, c)

	print(math.MaxFloat32)
	print(math.MaxFloat64)

	fmt.Printf("%f\n", math.E)
	fmt.Printf("%.2f\n", math.Pi)

	c1 := 1 + 2i // 1为实部，2为虚部
	c2 := 2 + 3i
	fmt.Println(c1, c2)

	// 布尔类型变量的默认值为false。
	// Go 语言中不允许将整型强制转换为布尔型.
	// 布尔型无法参与数值运算，也无法与其他类型进行转换。
	var b2 bool
	fmt.Println(b2)

	s1 := "你好世界"
	fmt.Println(s1)

	s2 := `你好
		请问，
		今天是几号
	`
	fmt.Println(s2)
	fmt.Println("s1" + "s2")
	fmt.Println(len((s2)))
	fmt.Println(strings.Split("hello world", " "))
	// strings.HasPrefix() // 判断前缀
	// strings.HasSuffix() // 判断后缀
	// strings.Index() // 判断字串位置
	// strings.LastIndex() // 判断最后一个字串位置
	// strings.Join()

	// ut := 'i' // uint8 表示一个ASCII字符
	// re := '我' // rune 表示一个UTF8字符, 实际为int32

	s1 = "big"
	bytes2 := []byte(s1) // string 转为[]byte数组进行修改
	fmt.Println(string(bytes2))

	s1 = "hello沙河小王子"
	count := 0
	for _, v := range s1 {
		if utf8.RuneLen(v) > 1 { // 中文字符在UTF8中长度大于1个字节
			count++
		}
	}
	fmt.Println(count)
}
