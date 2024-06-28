package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// 基于偏移量分页
// GET liwenzhou.com/api/v1/books?page=1&size=10
// 第一页：Offset:0, Limit:10
// 第二页：Offset:10, Limit:10
// 第三页：Offset:20, Limit:10
// SELECT id, title FROM books ORDER BY id ASC LIMIT 10 OFFSET 10;
// 缺点：1. 基于偏移量的分页在数据量很大的场景下，查询效率会比较低；2.在并发场景下会出现元素重复或者跳过；3.显式的page参数在支持跳页的同时也会被爬虫并发请求

// 基于游标分页
// cursor——通常是一个不透明字符串。它表示的是这一页数据的最后那个元素，通过这个cursor API 就能准确的返回下一页的数据。
// 缺点：1. 不支持跳页（但现在流行无限滑动翻页）；2. 不太适合多检索条件的场景。
// 在使用基于游标的分页时，通常并不会把具体的cursor数据显式拼接到API URL中，而是使用通常会被命名为next、next_cursor、after或page_token的不透明字符串
type Page struct {
	NextID        string `json:"next_id"`          // cursor
	NextTimeAtUTC int64  `json:"next_time_at_utc"` // 记录分页时间点,为了防止token泄漏后被无限期使用。我们可以限制token在一个合理时间后失效。
	PageSize      int64  `json:"page_size"`        // 记录每页元素数量
}
type Token string // 分页令牌，本质是字符串

// Encide返回分页token
func (p Page) Encode() Token {
	b, err := json.Marshal(p)
	if err != nil {
		return Token("")
	}
	return Token(base64.StdEncoding.EncodeToString(b))
}

// Decode 解析分页信息
func (t Token) Decode() Page {
	var result Page
	if len(t) == 0 {
		return result
	}

	bytes, err := base64.StdEncoding.DecodeString(string(t))
	if err != nil {
		return result
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return result
	}
	return result
}

var data = []string{
	"item1", "item2", "item3", "item4", "item5",
}

// 获取需要的分页数据和下一页的分页token
func GetPage(token Token) ([]string, Token) {
	page := token.Decode()
	start := 0
	// 查找开始位置
	for i, item := range data {
		if item == page.NextID {
			start = i + 1
			break
		}
	}
	// 获取分页数据
	end := start + int(page.PageSize)
	if end > len(data) {
		end = len(data)
	}
	nextId := ""
	if end < len(data) {
		nextId = data[end-1]
	}

	nextPage := Page{
		NextID:        nextId,
		NextTimeAtUTC: page.NextTimeAtUTC, // 可根据需要设置
		PageSize:      page.PageSize,
	}
	return data[start:end], nextPage.Encode()
}

func cursorTest() {
	// 初始化分页
	inititalPage := Page{
		NextID:        "", // 从头开始
		NextTimeAtUTC: 0,  // 根据需要设置
		PageSize:      3,  // 每页3个元素
	}
	token := inititalPage.Encode()

	// 模拟分页获取数据
	// BUG 无法跳出循环
	for {
		items, nextToken := GetPage(token)
		fmt.Println("Fetched items", items)
		if string(nextToken) == "" {
			break // 没有更多数据
		}
		token = nextToken
	}
}
