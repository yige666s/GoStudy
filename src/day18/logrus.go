package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	// the package is named "airbrake"
)

// var logger = log.New()

// func logrus() {
// 	// log.WithFields(log.Fields{"animal": "dog"}).Info("there be a dog")

// 	// // 设置日志输出
// 	// logger.Out = os.Stdout
// 	// logger.WithFields(log.Fields{"animal": "dog", "size": "10"}).Info("there be a lot dog")

// 	// // 日志级别
// 	// log.Trace("Something very low level.")
// 	// log.Debug("Useful debugging information.")
// 	// log.Info("Something noteworthy happened!")
// 	// log.Warn("You should probably take a look at this.")
// 	// log.Error("Something failed but I'm not quitting.")
// 	// // 记完日志后会调用os.Exit(1)
// 	// // log.Fatal("Bye.")
// 	// // 记完日志后会调用 panic()
// 	// // log.Panic("I'm bailing.")
// 	// // logger.SetLevel(log.DebugLevel)	// 适合debug程序
// 	// logger.SetLevel(log.InfoLevel) // 会记录info及以上级别 (warn, error, fatal, panic)

// 	// Logrus鼓励通过日志字段进行谨慎的结构化日志记录，而不是冗长的、不可解析的错误消息。
// 	// logger.WithFields(log.Fields{ // WithFields调用可选
// 	// 	"event": event,
// 	// 	"topic": topic,
// 	// 	"key":   key,
// 	// }).Fatal("failed to send event")

// 	// 默认字段，在请求的上下文中记录request_id和user_ip。
// 	request_id := "19823y21kd"
// 	user_ip := "127.0.0.1"
// 	requestLogger := log.WithFields(log.Fields{
// 		"request_id": request_id,
// 		"user_ip":    user_ip,
// 	})
// 	requestLogger.Info("soemthing happened on the request")
// 	requestLogger.Warn("something not great happened")

// }

// func init() {
// 	log.AddHook(airbrake.NewHook(123, "xyz", "production"))
// 	hook, err := logrus_syslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, "")
// 	if err != nil {
// 		log.Error("Unable to connect to local syslog daemon")
// 	} else {
// 		log.AddHook(hook)
// 	}

// 	// 两种格式化方法
// 	logger2 := log.TextFormatter{
// 		ForceColors:      true,
// 		DisableTimestamp: true,
// 	}
// 	logger3 := log.JSONFormatter{}

// 	// 日志中添加调用函数名,会增加性能开销。
// 	log.SetReportCaller(true)

// 	//默认的logger在并发写的时候是被mutex保护的，比如当同时调用hook和写log时mutex就会被请求
// 	log.New().SetNoLock() // 取消锁保护
// }

var logger0 = log.New()

func Init() {
	logger0.Formatter = &log.JSONFormatter{} // Log as JSON instead of the default ASCII formatter.
	f, _ := os.Create("./gin.log")
	logger0.Out = f // 设置日志输出地址
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = logger0.Out // Gin中的中间件和debug输出地址
	logger0.Level = log.InfoLevel   // 设置日志输出级别
}

func main() {
	Init()
	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		logger0.WithFields(log.Fields{
			"animals": "walrus",
			"size":    10,
		}).Warn("A group of walrus meerges from the ocean")
		ctx.String(http.StatusOK, "hello world!")
	})
	r.Run(":8080")
}
