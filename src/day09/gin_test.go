package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// httptest 用于Mockhttp请求
func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name   string // 测试用例名
		param  string // 参数
		expect string //期望值
	}{
		{"base case", `{"name": "jack"}`, "hello jack"},
		{"bad case", "", "we need a name"},
	}

	r := SetupRouter() //创建一个路由

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock http请求
			req := httptest.NewRequest(
				"POST",                      // 请求方法
				"/hello",                    // 请求路径
				strings.NewReader(tt.param), // 请求参数
			)

			// mock 相应记录器
			w := httptest.NewRecorder()
			// mock Server处理http请求并响应
			r.ServeHTTP(w, req)
			// 校验状态码是否正确
			assert.Equal(t, http.StatusOK, w.Code)
			// 校验响应内容
			var resp map[string]string
			err := json.Unmarshal([]byte(w.Body.Bytes()), &resp) // 解析返回的数据
			assert.Nil(t, err)
			assert.Equal(t, tt.expect, resp["msg"])
		})
	}
}
