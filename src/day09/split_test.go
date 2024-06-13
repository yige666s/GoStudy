package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) { // 'Test'开头，'S'大写
	get := Split("a:b:c", ":")         // 程序实际输出
	want := []string{"a", "b", "c"}    // 程序期望输出
	if !reflect.DeepEqual(want, get) { // 比较期望值与实际结果值
		t.Errorf("expected:%v, get:%v", want, get) // 失败给出提示
	}
}

func TestSplitWithComplexSep(t *testing.T) { // 每一个'Test'函数构成一个测试样例
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestSplitWithComplexSepByGroup(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		got := Split("abcd", "bc")
		want := []string{"a", "d"}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected:%v, got:%v", want, got)
		}
	})

	t.Run("case2", func(t *testing.T) {
		got := Split("1,2,3", ",")
		want := []string{"1", "2", "3"}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected:%v, got:%v", want, got)
		}
	})
}

func TestSplitAll(t *testing.T) {
	t.Parallel() // 将 TLog 标记为能够与其他测试并行运行
	// 定义测试表格
	// 这里使用匿名结构体定义了若干个测试用例
	// 并且为每个测试用例设置了一个名称
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
		{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
		{"more sep", "abcd", "bc", []string{"a", "d"}},
		{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}
	// 遍历测试用例
	for _, tt := range tests {
		tt := tt                            // 注意这里重新声明tt变量（避免多个goroutine中使用了相同的变量）
		t.Run(tt.name, func(t *testing.T) { // 使用t.Run()执行子测试
			t.Parallel() // 将每个测试用例标记为能够彼此并行运行
			got := Split(tt.input, tt.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected:%#v, got:%#v", tt.want, got)
			}
		})
	}
}

// 测试覆盖率是指代码被测试套件覆盖的百分比。通常我们使用的都是语句的覆盖率，
// 也就是在测试中至少被运行一次的代码占总代码的比例。在公司内部一般会要求测试覆盖率达到80%左右

func TestWithAssert(t *testing.T) {
	assert := assert.New(t)

	get1 := Split("1,2,3,4", ",")
	res1 := []string{"1", "2", "3", "4"}
	assert.Equal(get1, res1, "they should equal")

	res2 := []string{"1", "2", "3"}
	assert.NotEqual(get1, res2, "they should not equal")

	assert.NotNil(get1)
}
