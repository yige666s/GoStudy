package main

// fmt包实现了类似C语言 printf 和 scanf 的格式化I/O。主要分为向外输出内容和获取输入内容两大部分
// func fmtTest() {
// // Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
// fmt.Fprint(os.Stdout, "hello world\n")
// fmt.Fprintln(os.Stdout, "hello world")
// name := "liming"
// fmt.Fprintf(os.Stdout, "my name is %s", name)
// fmt.Println("")
// // Sprint系列函数会把传入的数据生成并返回一个字符串。
// s1 := fmt.Sprint("jack")
// name := "lucy"
// age := 19
// s2 := fmt.Sprintf("my name is %s,age is %d", name, age)
// s3 := fmt.Sprintln("jack")
// fmt.Println(s1, s2, s3)
// Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
// err := errors.New("first errror")
// w := fmt.Errorf("wrap a errror : %w", err)
// fmt.Println(w)
// }

// func Locidentifier() {
// fmt.Printf("%v\n", 100)
// s := struct{ name string }{"jack"}
// fmt.Printf("%v\n", s)
// fmt.Printf("%#v\n", s)
// fmt.Printf("%T\n", s)
// fmt.Printf("100%%\n")
// n := 65
// fmt.Printf("%b\n", n)
// fmt.Printf("%c\n", n)
// fmt.Printf("%d\n", n)
// fmt.Printf("%o\n", n)
// fmt.Printf("%x\n", n)
// fmt.Printf("%X\n", n)
// f := 123.45
// fmt.Printf("%b\n", f)
// fmt.Printf("%E\n", f)
// fmt.Printf("%e\n", f)
// fmt.Printf("%f\n", f)
// fmt.Printf("%g\n", f)
// fmt.Printf("%G\n", f)
// s := "hello world"
// fmt.Printf("%s\n", s)
// fmt.Printf("%q\n", s)
// fmt.Printf("%x\n", s)
// fmt.Printf("%X\n", s)
// a := 10
// fmt.Printf("%p\n", &a)
// fmt.Printf("%#p\n", &a)
// a := 12.34
// fmt.Printf("%f\n", a)
// fmt.Printf("%10.4f\n", a)
// }

func IOTest() {
	// Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
	// var (
	// 	name    string
	// 	age     int
	// 	married bool
	// )
	// fmt.Scan(&name, &age, &married)
	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
	// var (
	// 	name    string
	// 	age     int
	// 	married bool
	// )
	// fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married) // 输入格式为 1:小王子 2:28 3:false
	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
	// Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
	// var (
	// 	name    string
	// 	age     int
	// 	married bool
	// )
	// fmt.Scanln(&name, &age, &married)
	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	// 有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio包来实现。
	// reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	// text, _ := reader.ReadString('\n')
	// text = strings.TrimSpace(text)
	// fmt.Printf("%#v", text)

	// Fscan, Fscanln, Fscanf这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，
	// 只不过它们不是从标准输入中读取数据而是从 io.Reader 中读取数据。

	// Scan, Scanln, Scanf这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，
	// 只不过它们不是从标准输入中读取数据而是从  指定字符串  中读取数据。
}

// func Timedemo() {
// 	now := time.Now()
// 	fmt.Printf("current time:%v\n", now)
// 	year := now.Year()
// 	month := now.Month()
// 	day := now.Day()
// 	hour := now.Hour()
// 	minute := now.Minute()
// 	second := now.Second()
// 	fmt.Println(year, month, day, hour, minute, second)
// }

// timezoneDemo 时区示例
// func timezoneDemo() {
// 	// 中国没有夏令时，使用一个固定的8小时的UTC时差。
// 	// 对于很多其他国家需要考虑夏令时。
// 	secondsEastOfUTC := int((8 * time.Hour).Seconds())
// 	// FixedZone 返回始终使用给定区域名称和偏移量(UTC 以东秒)的 Location。
// 	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

// 	// 如果当前系统有时区数据库，则可以加载一个位置得到对应的时区
// 	// 例如，加载纽约所在的时区
// 	newYork, err := time.LoadLocation("America/New_York") // UTC-05:00
// 	if err != nil {
// 		fmt.Println("load America/New_York location failed", err)
// 		return
// 	}
// 	fmt.Println()
// 	// 加载上海所在的时区
// 	//shanghai, err := time.LoadLocation("Asia/Shanghai") // UTC+08:00
// 	// 加载东京所在的时区
// 	//tokyo, err := time.LoadLocation("Asia/Tokyo") // UTC+09:00

// 	// 创建时间对象需要指定位置。常用的位置是 time.Local（当地时间） 和 time.UTC（UTC时间）。
// 	//timeInLocal := time.Date(2009, 1, 1, 20, 0, 0, 0, time.Local)  // 系统本地时间
// 	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
// 	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)
// 	sameTimeInNewYork := time.Date(2009, 1, 1, 7, 0, 0, 0, newYork)

// 	// 北京时间（东八区）比UTC早8小时，所以上面两个时间看似差了8小时，但表示的是同一个时间
// 	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
// 	fmt.Println(timesAreEqual)

// 	// 纽约（西五区）比UTC晚5小时，所以上面两个时间看似差了5小时，但表示的是同一个时间
// 	timesAreEqual = timeInUTC.Equal(sameTimeInNewYork)
// 	fmt.Println(timesAreEqual)
// }

// timestampDemo 时间戳
// func timestampDemo() {
// 	now := time.Now()        // 获取当前时间
// 	timestamp := now.Unix()  // 秒级时间戳
// 	milli := now.UnixMilli() // 毫秒时间戳 Go1.17+
// 	micro := now.UnixMicro() // 微秒时间戳 Go1.17+
// 	nano := now.UnixNano()   // 纳秒时间戳
// 	fmt.Println(timestamp, milli, micro, nano)
// }

// func tickDemo() {
// 	// 使用time.Tick(时间间隔)来设置定时器，定时器的本质上是一个通道（channel）
// 	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
// 	for i := range ticker {
// 		fmt.Println(i) //每秒都会执行的任务
// 	}
// }

// // formatDemo 时间格式化
// func formatDemo() {
// 	now := time.Now()
// 	// 格式化的模板为 2006-01-02 15:04:05

// 	// 24小时制
// 	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
// 	// 12小时制
// 	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))

// 	// 小数点后写0，因为有3个0所以格式化输出的结果也保留3位小数
// 	fmt.Println(now.Format("2006/01/02 15:04:05.000")) // 2022/02/27 00:10:42.960
// 	// 小数点后写9，会省略末尾可能出现的0
// 	fmt.Println(now.Format("2006/01/02 15:04:05.999")) // 2022/02/27 00:10:42.96

// 	// 只格式化时分秒部分
// 	fmt.Println(now.Format("15:04:05"))
// 	// 只格式化日期部分
// 	fmt.Println(now.Format("2006.01.02"))
// }

// func parseDemo() {
// 	// 其中time.Parse在解析时不需要额外指定时区信息。
// 	timeobj, _ := time.Parse("2006/01/02 15:04:05", "2024/01/01 12:00:00")
// 	fmt.Println(timeobj)
// 	timeobj2, _ := time.Parse(time.RFC3339, "2024-11-13T12:14:24+08:00")
// 	fmt.Println(timeobj2)
// }

// parseDemo 解析时间
// func parseDemo() {
// 	now := time.Now()
// 	fmt.Println(now)
// 	// 加载时区
// 	loc, _ := time.LoadLocation("Asia/Shanghai")
// 	// 按照指定时区和指定格式解析字符串时间
// 	timeObj, _ := time.ParseInLocation("2006/01/02 15:04:05", "2022/10/05 11:25:20", loc)
// 	fmt.Println(timeObj)
// 	fmt.Println(timeObj.Sub(now))
// }

// func parctest() {
// 	// now := time.Now()
// 	// fmt.Println(now.Format("2017/06/19 20:30:05"))
// 	t1 := time.Now()
// 	time.Sleep(time.Second * 2)
// 	t2 := time.Now()
// 	timdura := t2.Sub(t1)
// 	fmt.Println(timdura)
// }

// func ArgsTest() {
// 	// 如果你只是简单的想要获取命令行参数，可以像下面的代码示例一样使用os.Args来获取命令行参数。
// 	//os.Args是一个[]string,一个存储命令行参数的字符串切片，它的第一个元素是执行文件的名称。
// 	if len(os.Args) > 0 {
// 		for i, v := range os.Args {
// 			fmt.Printf("args[%d]=%v\n", i, v)
// 		}
// 	}
// }

// func ArgsParse() {
// 	name := flag.String("name", "jack", "姓名")
// 	age := flag.Int("age", 18, "年龄")
// 	married := flag.Bool("married", true, "是否结婚")
// 	delay := flag.Duration("d", 0, "延迟时间间隔")

// 	// 解析命令行参数
// 	flag.Parse()

// 	fmt.Println(name, age, married, delay)
// 	// 返回命令行参数后的其他参数
// 	fmt.Println(flag.Args())
// 	// 返回命令行参数后的其他参数个数
// 	fmt.Println(flag.NArg())
// 	// 返回使用的命令行参数个数
// 	fmt.Println(flag.NFlag())
// }
