package main

import "fmt"

// 接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。
// 类似于Java中的Interface或者C++中的虚函数
// 一个接口类型就是一组方法的集合，它规定了需要实现的所有方法。相较于使用结构体类型，当我们使用接口类型说明相比于它是什么更关心它能做什么。

type Writer interface {
	Write([]byte) error
}

type Stu struct {
}

func (s *Stu) StuWrite(w Writer) {
	w.Write([]byte{})
}

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

func interfaceTest() {
	var c WashingMachine = haier{}
	c.dry()
	c.wash()
}

// 通常我们在使用空接口类型时不必使用type关键字声明，可以像下面的代码一样直接使用interface{}
var x interface{} // 定义一个空接口

// 空接口作为函数参数,接收任意类型的函数参数。
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

// 使用空接口实现可以保存任意值的字典。
func NoneInferface() {
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}

// 由于接口类型变量能够动态存储不同类型值的特点，所以很多初学者会滥用接口类型（特别是空接口）来实现编码过程中的便捷。
// 只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。
// 切记不要为了使用接口类型而增加不必要的抽象，导致不必要的运行时损耗
// 接口是一种类型，一种抽象的类型。区别于我们在之前章节提到的那些具体类型（整型、数组、结构体类型等），它是一个只要求实现特定方法的抽象类型

type WriteLoger interface {
	WriteLog()
}

type Terminal struct {
}

func (t *Terminal) WriteLog() {
	fmt.Println("Wrtiring log on terminal")
}

type File struct {
}

func (f *File) WriteLog() {
	fmt.Println("writing log to File")
}

func WriterlogTest() {
	var w WriteLoger = &Terminal{}
	w.WriteLog()
	w = &File{}
	w.WriteLog()
}
