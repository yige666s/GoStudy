package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func JudgeByTime(now time.Time) int {
	switch hour := now.Hour(); {
	case hour > 8 && hour < 20:
		return 10
	case hour > 20 && hour < 23:
		return 1
	}
	return -1
}

func TestJudgeBuTime(t *testing.T) {
	tests := []struct {
		name string
		arg  time.Time
		want int
	}{
		{name: "work time", arg: time.Date(2024, 6, 15, 12, 25, 15, 0, time.UTC), want: 10},
		{name: "night", arg: time.Date(2024, 6, 15, 21, 14, 54, 0, time.UTC), want: 1},
		{name: "MidNight", arg: time.Date(2024, 6, 15, 23, 14, 15, 0, time.UTC), want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, JudgeByTime(tt.arg))
		})
	}
}

type App struct {
	logger
}

func (a *App) start() {
	a.logger.info("app start....")
}

// NewApp 构造函数，将依赖项注入
// 依赖注入就是指在创建组件（Go 中的 struct）的时候接收它的依赖项，而不是它的初始化代码中引用外部或自行创建依赖项。
// 通过在构造函数中隐式创建依赖项，这样的代码强耦合、不易扩展，也不容易编写单元测试。可以通过使用依赖注入的方式，将构造函数中的依赖作为参数传递给构造函数。
func NewApp(lg logger) *App {
	return &App{logger: lg}
}

// SOLID原则
// S(Single) 单一职责原则
// O(OpenClose)	一个软件实体，如类、模块和函数应该对扩展开放，对修改关闭。
// L(LiReplace)	认为“程序中的对象应该是可以在不改变程序正确性的前提下被它的子类所替换的”的概念。
// I(Isolation) 许多特定于客户端的接口优于一个通用接口。
// D(Dependency) 应该依赖抽象，而不是某个具体示例。
