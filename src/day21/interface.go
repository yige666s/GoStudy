package main

import "fmt"

type StuType int32

const (
	Type1 StuType = iota
	Type2
	Type3
	Type4
)

func EnumsTest() {
	fmt.Println(Type1, Type2, Type3, Type4)
}

type set map[string]struct{}

func NullStructTest() {
	ch1 := make(chan struct{}, 1)
	s := make(set)
}
