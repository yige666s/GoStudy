package main

import (
	"encoding/base64"
	"encoding/json"
)

type Page struct {
	NextID         string `json:"next_id"`          // cursor
	NextTimieAtUTC int64  `json:"next_tmie_at_utc"` // 记录分页时间
	PageSize       int64  `json:"page_size"`        // 记录每页的元素
}
type Token string

func (p Page) Encode() Token {
	b, err := json.Marshal(p)
	if err != nil {
		return Token("")
	}
	return Token(base64.StdEncoding.EncodeToString(b))
}

func (p Token) Decode() Page {
	var res Page
	if len(p) == 0 {
		return res
	}

	bytes, err := base64.StdEncoding.DecodeString(string(p))
	if err != nil {
		return res
	}

	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return res
	}

	return res
}

// 游标分页：
// 基于游标的分页通过游标（cursor）来跟踪每一页的数据位置，而不是直接使用偏移量。每次分页请求时，都会带上一个游标，这个游标通常是上一次查询结果中的最后一个记录的唯一标识符（如ID）。
// 理解游标分页的SQL示例
// 假设我们有一本包含100本书的书籍表，id 是唯一且自增的（连续的）。每一页显示10本书。

// 第一页查询（无游标）：
// SQL:
// SELECT id, title FROM books ORDER BY id ASC LIMIT 10;
// 结果：将返回 id 为 1 到 10 的书。
// 第二页查询（带游标）：

// 假设第一页最后一本书的 id 是 10，那么下一页查询的 SQL 为：
// SELECT id, title FROM books WHERE id > 10 ORDER BY id ASC LIMIT 10;
// 结果：返回 id 为 11 到 20 的书。
// 后续查询：
// 每次查询都会根据前一页的最后一个 id 来确定下一页的起点游标。例如，第三页查询将使用 id > 20，第四页使用 id > 30，依此类推。
