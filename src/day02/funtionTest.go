package main

import (
	"fmt"
	"strings"
)

func funtionTest() {
	// Go语言中支持函数、匿名函数和闭包，并且函数在Go语言中属于“一等公民”。
	c := twoSum(1, 2) // 可以不接收返回值
	fmt.Println(c)

}

func twoSum(a int, b int) int {
	return a + b
}

// 可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识。
func intSum(x int, y ...int) {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum += v
	}
	fmt.Println(sum)
}

// 函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回
func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func add(a int, b int) int {
	return a + b
}

func funcVariable() {
	// 使用type关键字来定义一个函数类型
	type calculation func(int, int) int
	c := add // 实例化c变量
	fmt.Println(c(1, 2))
}

// 函数可以作为参数
func add1(x, y int) int {
	return x + y
}
func calc1(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func main1() {
	ret2 := calc1(10, 20, add)
	fmt.Println(ret2) //30
}

// 函数也可以作为返回值：
// func do(s string) (func(int, int) int, error) {
// 	switch s {
// 	case "+":
// 		return add, nil
// 	case "-":
// 		return sub, nil
// 	default:
// 		err := errors.New("无法识别的操作符")
// 		return nil, err
// 	}
// }

// 匿名函数，匿名函数多用于实现回调函数和闭包。
func NoNameFunc() {
	addfunc := func(a int, b int) {
		fmt.Println(a + b)
	}
	addfunc(12, 2) // 变量保存函数

	func(a, b int) { // 匿名函数直接执行
		fmt.Print(a - b)
	}(2, 3)
}

// 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		fmt.Println(x)
		return x
	}
}

func Closure() {
	f := adder(2)
	f(5)
	f(6)
	f(10)
}

// 先被defer的语句最后被执行，最后被defer的语句，最先被执行
// defer语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。
func defertest() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

// 代码的输出结果是? 提示：defer注册要延迟执行的函数时该函数所有的参数都需要确定其值

// A 1 2 3
// B 10 2 12
// BB 10 12 22
// AA 1 3 4
// func calc2(index string, a, b int) int {
// 	ret := a + b
// 	fmt.Println(index, a, b, ret)
// 	return ret
// }

// func main2() {
// 	x := 1
// 	y := 2
// 	defer calc("AA", x, calc2("A", x, y))
// 	x = 10
// 	defer calc("BB", x, calc2("B", x, y))
// 	y = 20
// }

// panic可以在任何地方引发，但recover只有在defer调用的函数中有效。
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		// recover()必须搭配defer使用
		// defer一定要在可能引发panic的语句之前定义,否则无法起作用
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
func panicTest() {
	funcA()
	funcB()
	funcC()
}

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main3() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

func dispatchCoin() int {
	remainscoins := coins
	for _, name := range users { // 遍历每一个user获得对应分配的coins
		conisGiven := calculateCoins(name)
		distribution[name] = conisGiven
		remainscoins -= conisGiven
	}
	return remainscoins
}

func calculateCoins(name string) int {
	coins := 0
	for _, char := range strings.ToLower(name) {
		switch char {
		case 'e', 'E':
			coins += 1
		case 'i', 'I':
			coins += 2
		case 'o', 'O':
			coins += 3
		case 'u', 'U':
			coins += 4
		}
	}
	return coins
}
