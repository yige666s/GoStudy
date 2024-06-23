package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// 使用匿名结构体添加字段
type Userinfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func anonymousStructDemo() {
	u1 := Userinfo{
		ID:   123456,
		Name: "jack",
	}
	b, _ := json.Marshal(struct {
		*Userinfo
		Token string `json:"token"`
	}{
		&u1,
		"aksndlandkdsa",
	})
	fmt.Printf("str:%s\n", b)
}

type Comment struct {
	Content string
}

type Image struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func anonymousStructDemo2() {
	c1 := Comment{
		Content: "永远不要高估自己",
	}
	i1 := Image{
		Title: "赞赏码",
		URL:   "https://www.liwenzhou.com/images/zanshang_qr.jpg",
	}
	// struct -> json string
	b, err := json.Marshal(struct {
		*Comment
		*Image
	}{&c1, &i1})
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	// json string -> struct
	jsonStr := `{"Content":"永远不要高估自己","title":"赞赏码","url":"https://www.liwenzhou.com/images/zanshang_qr.jpg"}`
	var (
		c2 Comment
		i2 Image
	)
	if err := json.Unmarshal([]byte(jsonStr), &struct {
		*Comment
		*Image
	}{&c2, &i2}); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("c2:%#v i2:%#v\n", c2, i2)
}

// 自定义解析时间字段,如2020-04-05 12:25:42

type CustomTime struct {
	time.Time
}

const ctlayout = "2006-01-02 15:04:05"

var nilTime = (time.Time{}).Unix()

func (c *CustomTime) Unmarsha1JSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		c.Time = time.Time{}
		return
	}
	c.Time, err = time.Parse(ctlayout, s)
	return
}

func (c *CustomTime) Marsha1JSON() ([]byte, error) {
	if c.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", c.Time.Format(ctlayout))), nil
}

func (c *CustomTime) IsSet() bool {
	return c.Time.UnixNano() != nilTime
}

type Post struct {
	CreateTime CustomTime `json:"create_time"`
}

func timeFieldDemo() {
	p1 := Post{CreateTime: CustomTime{time.Now()}}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal p1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	jsonStr := `{"create_time":"2020-04-05 12:25:42"}`
	var p2 Post
	if err := json.Unmarshal([]byte(jsonStr), &p2); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}

type Stu struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Score  int     `json:"score"`
	Weight float64 `json:"-"`
}

func JsonTagDemo() {
	s1 := Stu{
		Name:   "jack",
		Age:    18,
		Score:  99,
		Weight: 66.6,
	}
	b1, _ := json.Marshal(s1)
	fmt.Printf("str:%s\n", b1)

	b2 := `{"name":"jack","age":18,"score":99}`
	var s2 Stu
	json.Unmarshal([]byte(b2), &s2)
	fmt.Println("s2 :", s2)
}
