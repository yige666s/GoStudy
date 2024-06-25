package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var logger = log.New()

func logrus() {
	log.WithFields(log.Fields{"animal": "dog"}).Info("there be a dog")

	// 设置日志输出
	logger.Out = os.Stdout
	logger.WithFields(log.Fields{"animal": "dog", "size": "10"}).Info("there be a lot dog")

	// 日志级别
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// 记完日志后会调用os.Exit(1)
	// log.Fatal("Bye.")
	// 记完日志后会调用 panic()
	// log.Panic("I'm bailing.")
	// logger.SetLevel(log.DebugLevel)	// 适合debug程序
	logger.SetLevel(log.InfoLevel) // 会记录info及以上级别 (warn, error, fatal, panic)

	// Logrus鼓励通过日志字段进行谨慎的结构化日志记录，而不是冗长的、不可解析的错误消息。
	// logger.WithFields(log.Fields{ // WithFields调用可选
	// 	"event": event,
	// 	"topic": topic,
	// 	"key":   key,
	// }).Fatal("failed to send event")

	// 默认字段

}
