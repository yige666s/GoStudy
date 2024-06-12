package main

// Go语言中strconv包实现了基本数据类型和其字符串表示的相互转换。
// func Transfer() {
// 这是C语言遗留下的典故。这里的 a 指的是 array 字符数组，
// C语言中没有string类型而是用字符数组(array)表示字符串，所以Itoa对很多C系的程序员很好理解。
// s1 := "100"
// i1, _ := strconv.Atoi(s1) // string转int
// fmt.Println(i1)

// i2 := 200
// s2 := strconv.Itoa(i2) //int转string
// fmt.Println(s2)

// Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。
// i3, _ := strconv.ParseInt("-127", 10, 8)
// fmt.Println(i3)
// i4, _ := strconv.ParseUint("127", 16, 8)
// fmt.Println(i4)
// Format系列函数实现了将给定类型数据格式化为string类型数据的功能。
// func FormatBool(b bool) string
// func FormatInt(i int64, base int) string
// func FormatUint(i uint64, base int) string
// s2 := strconv.FormatFloat(3.1415926, 'E', -1, 32)
// fmt.Println(s2)
// strconv.IsPrint() //返回一个字符是否可打印
// }

// func HttpClient() {
// 	resp, err := http.Get("https://www.baidu.com/") // 这里的url要写全了，不能只写baidu.com
// 	if err != nil {
// 		fmt.Println("get failed,%v\n", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body) // 注意go中的函数返回值如果没有被接收会产生错误
// 	if err != nil {
// 		fmt.Println("read from body failed: %v", err)
// 		return
// 	}
// 	fmt.Println(string(body))
// }

// func GetwithParam() {
// 	infoApiUrl := "http://127.0.0.1:8080/getinfo"
// 	// infoApi需要传递的参数
// 	data := url.Values{}
// 	data.Set("name", "jack")
// 	u, err := url.ParseRequestURI(infoApiUrl) //解析请求的infoapiurl
// 	if err != nil {
// 		fmt.Println("parser url failed,err:%v", err)
// 	}
// 	u.RawQuery = data.Encode() // URL编码
// 	fmt.Println(u.String())

// 	resp, err := http.Get(u.String())
// 	if err != nil {
// 		fmt.Println("get failed, err :", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	content, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("read from body failed:%v", err)
// 		return
// 	}
// 	fmt.Println(string(content))
// }

// func getHandler(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	data := r.URL.Query()
// 	fmt.Println(data.Get("name"))
// 	answer := `{"status": "ok"}` // 这里使用 `` 定义原生字符串，当用到 " \ 等需要转义的字符时不需要加 \ 即可保留原语义
// 	w.Write([]byte(answer))
// }

// func PostTest() {
// 	Registerurl := "http://127.0.0.1/register"
// 	// 表单数据，json格式
// 	contentType := "application/json"
// 	data := `{
// 		"name":"jack",
// 		"passwd":"111111"
// 	}`
// 	resp, err := http.Post(Registerurl, contentType, strings.NewReader(data))
// 	if err != nil {
// 		fmt.Println("post failed:", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	content, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("read from resp failed", err)
// 		return
// 	}
// 	fmt.Println(string(content))
// }

// func PostHandler(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
// 	r.ParseForm()
// 	fmt.Println(r.PostForm)
// 	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("passwd"))
// 	// 2. 请求类型是application/json时从r.Body读取数据
// 	content, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Println("read from resquest failed", err)
// 		return
// 	}
// 	fmt.Println(string(content))
// 	answer := `{"status" : "ok" }`
// 	w.Write([]byte(answer))
// }

//	func NewClient() {
//		// 要管理代理、TLS配置、keep-alive、压缩和其他设置，创建一个Transport：
//		tr := &http.Transport{
//			TLSClientConfig:    &tls.Config{RootCAs: pool},
//			DisableCompression: true,
//		}
//		// 要管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client：
//		client := &http.Client{Transport: tr}
//		resp, _ := client.Get("https://example.com")
//		fmt.Println(resp.Body)
//		// Client和Transport类型都可以安全的被多个goroutine同时使用。出于效率考虑，应该一次建立、尽量重用
//	}
// func HelloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "hello world", time.Now())
// }

// func HelloWebTest() {
// 	http.HandleFunc("/hello", HelloHandler)  // 创建路由绑定handler函数
// 	err := http.ListenAndServe(":9090", nil) // 启动服务器监听
// 	if err != nil {
// 		fmt.Println("server listen failed", err)
// 		return
// 	}
// }

// func sayHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello 沙河！")
// }

//	func main2() {
//		http.HandleFunc("/", sayHello)
//		err := http.ListenAndServe(":9090", nil)
//		if err != nil {
//			fmt.Printf("http server failed, err:%v\n", err)
//			return
//		}
//	}
// func myhandler(w http.ResponseWriter, r *http.Request) {}
// func Mysrv() {
// 	s := &http.Server{
// 		Addr:           ":8090",
// 		Handler:        myhandler,
// 		ReadTimeout:    10 * time.Second,
// 		WriteTimeout:   10 * time.Second,
// 		MaxHeaderBytes: 1 << 20,
// 	}
// }
