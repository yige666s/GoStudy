package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiURL   = "https://at-open-api.tal.com/mathai/ocr/search"
	authCode = "YOUR_AUTH_CODE" // TODO 请替换为实际的鉴权码
)

type RequestPayload struct {
	Question string `json:"question"`
	TopNum   int    `json:"top_num"`
	UserID   string `json:"user_id,omitempty"`
}

type Tag struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type Result struct {
	Score      float64 `json:"score"`
	QuestionID string  `json:"question_id"`
	Question   string  `json:"question"`
	Answer     string  `json:"answer"`
	Tags       []Tag   `json:"tags"`
}

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		Results []Result `json:"results"`
	} `json:"result"`
}

func main() {
	// 构造请求数据
	payload := RequestPayload{
		Question: "https://s.tiku.100tal.com/xes_souti/paisou2/images/202402/23/6c8fe47c02754e8cc6f9f30a3e4d0998.jpeg",
		TopNum:   5,
		UserID:   "optional_user_id",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error marshalling payload: %v\n", err) // TODO 处理错误
		return
	}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-mathai", authCode)

	// 发送HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	// 打印响应
	fmt.Printf("Response status: %s\n", resp.Status)
	fmt.Printf("Response body: %s\n", body)

	// 解析响应
	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return
	}

	// 打印解析后的结果
	fmt.Printf("Response code: %d\n", apiResp.Code)
	fmt.Printf("Response message: %s\n", apiResp.Message)
	fmt.Printf("Results:\n")
	for _, result := range apiResp.Result.Results {
		fmt.Printf("Question ID: %s, Score: %.2f, Question: %s, Answer: %s\n",
			result.QuestionID, result.Score, result.Question, result.Answer)
	}
}
