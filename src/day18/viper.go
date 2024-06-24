package main

import (
	"fmt"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Viper是适用于Go应用程序（包括Twelve-Factor App）的完整配置解决方案。
// 它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式
// Viper会按照下面的优先级。每个项目的优先级都高于它下面的项目: 显示调用Set设置值，命令行参数（flag），环境变量，配置文件，key/value存储，默认值
// Viper 会依次检查以下支持的配置文件格式（按照固定的顺序）：JSON、TOML、YAML、HCL、TOML、envfile。

func ViperDemo() {
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file :%s \n", err))
	}
	viper.WatchConfig()

	r := gin.Default()
	r.GET("/version", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, viper.GetString("version"))
	})

	if err := r.Run(
		fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
		panic(err)
	}
}

type Config struct {
	host    string `mapstructure:"host"`
	port    int    `mapstructure:"port"`
	Version string `mapstructure:"version"`
}

var conf = new(Config)

func ViperWithStruct() {
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file:%s \n", err))
	}
	viper.WatchConfig()
	// 注意！！！配置文件发生变化后要同步到全局变量Conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config has benn modified")
		if err := viper.Unmarshal(conf); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})

	r := gin.Default()
	// 访问/version的返回值会随配置文件的变化而变化
	r.GET("/version", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, conf.Version)
	})

	if err := r.Run(fmt.Sprintf(":%d", conf.port)); err != nil {
		panic(err)
	}
}
