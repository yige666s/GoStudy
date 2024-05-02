package main

import (
	"encoding/json"
	"fmt"
)

// Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。
// Go语言中通过 [结构体的内嵌] 再配合 [接口] 比面向对象具有更高的扩展性和灵活性。

type size_t uint // 类型定义
type byte = int8 // 类型别名
type rune = int32

func type1() {
	var a size_t
	var b byte
	fmt.Printf("%T,%T", a, b) //main.size_t, int8,其中size_t只存在于代码中，编译后会没有size_t类型
}

// Go语言中通过struct来实现面向对象。
// 只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
type person struct {
	name  string
	age   int8
	score int8
}

func structTest() {
	// 列表初始化
	p1 := person{"lili", 19, 80} // p1是一个对象
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.age, p1.name, p1.score)

	p2 := new(person)
	p2.name = "jack"
	fmt.Println(p2.name, p2.age, p2.score)

	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次[new]实例化操作。
	// 列表初始化
	p3 := &person{"lucy", 10, 66} // p3是一个指针
	fmt.Printf("%T\n", p3)
	fmt.Printf("%#v\n", p3)

	p4 := person{}
	fmt.Printf("%#v\n", p4)

	// 结构体占用一块连续的内存,空结构体是不占用空间的。
}

type student struct {
	name string
	age  int
}

func interview1() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

// struct没有构造函数，自定义构造函数
func NewPerson(name string, age int8, score int8) *person {
	return &person{name, age, score}
}

// Go语言中的方法（Method）是一种作用于特定类型变量的函数。
// 这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。
// 方法与函数的区别是，[函数不属于任何类型]，[方法属于特定的类型]
// 值接收
func (p person) Dream() {
	fmt.Println("相当于person的成员函数")
}

// 什么时候应该使用指针类型接收者
// 1. 需要修改接收者中的值
// 2. 接收者是拷贝代价比较大的大对象
// 3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

// 指针接收,可改变接收者内部的变量值
func (p *person) SetAge(Newage int) {
	p.age = int8(Newage)
}

func ConstructTest() {
	p1 := NewPerson("lili", 19, 80)
	fmt.Printf("%T\n", p1)
	fmt.Printf("%#v\n", p1)
	p1.Dream()
}

// Person 结构体Person类型
// 默认会采用类型名作为字段名，结构体要求字段名称必须唯一
// 因此一个结构体中同种类型的匿名字段只能有一个
type Person struct {
	bool
	string
	int
}

// 嵌套结构体
// Address 地址结构体
type Address struct {
	Province string
	City     string
}

// User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

func CombineStruct() {
	user1 := User{
		"lili",
		"woman",
		Address{
			"hebei",
			"qinhuangdao",
		},
	}
	fmt.Printf("%#v\n", user1)
}

// 结构体集成其实是通过组合的方式实现的
// Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

// Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func heritage() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}

// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
func jsonTest() { // json 序列化与反序列化
	json.Marshal()
	json.Unmarshal()
}

// 因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意
type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) SetDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams) // slice是引用传递，需要重新开辟空间进行拷贝，否则两个指针指向同一地址
}

func Addtest() {
	p1 := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams) // ?
}

type stu struct {
	id    int8
	name  string
	age   int8
	score int8
}

var stuInfo = []stu{}

func (p1 *stu) infoPrint() {
	for _, v := range stuInfo {
		fmt.Printf("%#v\n", v)
	}
}

func (p1 *stu) addstu(id int8, name string, age int8, score int8) {
	stutmp := stu{id, name, age, score}
	stuInfo = append(stuInfo, stutmp)
}

func (p *stu) setName(newname string) {
	p.name = newname
}

func (p *stu) delestu(id_del int8) {
	for i, v := range stuInfo {
		if v.id == id_del {
			stuInfo = append(stuInfo[:i], stuInfo[i+1:]...)
		}
	}
}
