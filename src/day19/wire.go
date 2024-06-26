package main

import (
	"context"
	"errors"
	"io"
	"log"
	"os"
	"sync"

	"github.com/google/wire"
)

// 控制反转（Inversion of Control，缩写为IoC），是面向对象编程中的一种设计原则，可以用来减低计算机代码之间的耦合度
// 最常见的方式叫做依赖注入（Dependency Injection，简称DI）。依赖注入是生成灵活和松散耦合代码的标准技术，通过明确地向组件提供它们所需要的所有依赖关系。

// NewBookRepo 创建BookRepo的构造函数
// func NewBookRepo(db *gorm.DB) *BookRepo {
// 	// 构造函数NewBookRepo在创建BookRepo时需要从外部将依赖项db作为参数传入，
// 	// 我们在NewBookRepo中无需关注db的创建逻辑，实现了代码解耦
// 	return &BookRepo{db: db}
// }

// Wire 是一个的 Google 开源的依赖注入工具，通过自动生成代码的方式在编译期完成依赖注入。
// wire中有两个核心概念：提供者（provider）和注入器（injector）。
// Wire中的提供者就是一个可以产生值的普通函数。
type Z struct { // 必须大写
	Value int
}
type Y struct { // 必须大写
	Value int
}

func NewY() (Y, error) {
	return Y{Value: 1}, nil
}

// NewZ 返回一个Z对象，当传入依赖的value为0时会返回错误。

func NewZ(ctx context.Context, y Y) (Z, error) {
	if y.Value == 0 {
		return Z{}, errors.New("cannot provide z when value is zero")
	}
	return Z{Value: y.Value + 1}, nil
}

// Provider函数可以分组为提供者函数集（provider set）。
// 使用wire.NewSet 函数可以将多个提供者函数添加到一个集合中。如果经常同时使用多个提供者函数，这非常有用。
var Providerset = wire.NewSet(NewZ, NewY)

// 应用程序中是用一个injector来连接Provider，injector就是一个按照依赖顺序调用Provider。
// 要声明一个注入器函数只需要在函数体中调用wire.Build。这个函数的返回值也无关紧要，只要它们的类型正确即可
// go:build wireinject,确保wire.go不会参与最终的项目编译
func initZ(ctx context.Context) (Z, error) {
	wire.Build(Providerset)
	return Z{}, nil
}

// 结构体Provider
type Foo int
type Bar int

func FooProvider() Foo { return 1 }
func BarProvider() Bar { return 2 }

// 有时防止结构体的某些字段被注入器填充很有必要，尤其是在将*传递给wire.Struct的时候
type FooBar struct {
	mu    sync.Mutex `wire:"-"` // 你可以用wire:"-"标记字段，使wire忽略这些字段
	MyFoo Foo
	MyBar Bar
}

var Set = wire.NewSet(
	FooProvider,
	BarProvider,
	wire.Struct(new(FooBar), "MyFoo", "MyBar"),
)

type Foo1 struct {
	x int
}

func injector() Foo1 {
	wire.Build(wire.Value(Foo1{x: 0}))
	return Foo1{}
}

// 对于接口值，使用 InterfaceValue：
func injectReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
}

// 结构字段Provider
type Foo2 struct {
	S string
	N int
	F float64
}

func providerFoo2() Foo2 {
	return Foo2{S: "Hello, World!", N: 1, F: 3.14}
}
func injectedMessage() string {
	// 可以使用wire.FieldsOf直接使用结构体的字段，而无需编写一个类似proivderS的函数
	wire.Build(providerFoo2(), wire.FieldsOf(new(Foo2), "S", "N"))
}

// clean函数，用于清理providerSource对象
type Logger log.Logger
type Path string

// cleanup函数的签名必须是func()，并且保证在提供者的任何输入的cleanup函数之前调用。
func providerFile(log Logger, path Path) (*os.File, func(), error) {
	f, err := os.Open(string(path))
	if err != nil {
		return nil, nil, err
	}
	clean := func() {
		if err := f.Close(); err != nil {
			log.Log(err)
		}
	}
	return f, clean, err
}

// 备用注入器语法
// 如果你厌倦了在注入器函数声明的末尾编写类似return Foo{}, nil的语句，那么你可以简单粗暴地使用panic
func injectFoo2() Foo2 {
	panic(wire.Build( /*...*/ ))
}
