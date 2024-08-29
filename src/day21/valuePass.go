package main

import "fmt"

func valuePass() {
	a := 1
	fmt.Printf("%v %v", Add(a), a)
	fmt.Println("")
	b := []int{1}
	Add1(b)
	fmt.Println(b[0])
}

func Add(a int) int {
	a += 1
	return a
}

func Add1(a []int) []int {
	a[0] += 1
	return a
}
