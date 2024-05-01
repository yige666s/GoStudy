package main

import "fmt"

func operatorstudy(nums []int) int { // 找出只出现一次的数字
	result := 0
	for _, v := range nums {
		v ^= result
	}
	return result
}

func operatorTest() {
	fmt.Println(1 >> 7)
	v := 2
	v >>= 4
	fmt.Println(v)
}
