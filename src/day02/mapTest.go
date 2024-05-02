package main

import (
	"fmt"
	"strings"
)

// func mapTest() {
// 	// Go语言中提供的映射关系容器为map，其内部使用散列表（hash）实现。
// 	// map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
// 	// map类型的变量默认初始值为nil，需要使用make()函数来分配内存
// 	scoremap := make(map[string]int, 10)
// 	scoremap["lili"] = 80
// 	scoremap["jack"] = 90
// 	scoremap["tony"] = 100

// 	fmt.Println(scoremap)

// 	userinfo := map[string]string{
// 		"username":  "jack",
// 		"userage":   "20",
// 		"useremail": "1234@gmail.com",
// 	}
// 	fmt.Println(userinfo)

// 	// 如果key存在ok为true,value为对应的值；不存在ok为false,value为值类型的零值
// 	value, ok := userinfo["lili"]
// 	fmt.Println(value, ok)

// 	// 基于范围的for循环 ，遍历map时的元素顺序与添加键值对的顺序无关
// 	for k, v := range userinfo {
// 		fmt.Println(k, v)
// 	}

// 	// 使用delete()内建函数从map中删除一组键值对，delete()函数的格式如下：
// 	delete(userinfo, "userage")

// }

// func ReadMapByseq() {
// 	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

// 	var scoreMap = make(map[string]int, 200)

// 	for i := 0; i < 100; i++ {
// 		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
// 		value := rand.Intn(100)          //生成0~99的随机整数
// 		scoreMap[key] = value
// 	}

// 	//取出map中的所有key存入切片keys
// 	var keys = make([]string, 0, 200)
// 	for key := range scoreMap {
// 		keys = append(keys, key)
// 	}
// 	//对切片进行排序
// 	sort.Strings(keys)
// 	//按照排序后的key遍历map
// 	for _, key := range keys {
// 		fmt.Println(key, scoreMap[key])
// 	}
// }

// 元素为map类型的切片，切片每一个元都是一个slice
// func MapSlice() {
// 	mapslice := make([]map[string]string, 3)
// 	for i, v := range mapslice {
// 		fmt.Println(i, v)
// 	}
// 	mapslice[0] = make(map[string]string, 10)
// 	mapslice[0]["name"] = "lili"
// 	mapslice[0]["passwd"] = "123456"
// 	mapslice[0]["email"] = "lili@gmail.com"
// 	for i, v := range mapslice {
// 		fmt.Println(i, v)
// 	}
// }

// 值为切片类型的map，map的value是一个slice
// func sliceMap() {
// 	sliceMap := make(map[string][]string, 3)
// 	fmt.Println(sliceMap)
// 	key := "china"
// 	value := make([]string, 0, 2)
// 	value = append(value, "beijing", "shanghai")
// 	sliceMap[key] = value
// 	fmt.Println(sliceMap)
// }

// 词频统计
func wordcount(s string) {
	wordslice := make([]string, 10, 20)
	wordslice = strings.Split(s, " ")
	countmap := make(map[string]int, 10)
	for _, v := range wordslice {
		countmap[v]++
	}
	for i, v := range countmap {
		fmt.Printf("%s = %d ", i, v)
	}
	fmt.Println("")
}

func test() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...) // slice去掉s[1]后,s[1]变为s[2] = 3，对应的map[q1mi]的value[1] = 3
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
