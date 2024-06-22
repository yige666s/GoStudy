package main

//编译可以通过以下命令或编写 makefile 来操作。
// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/bluebell
// ├── bin
// │   └── bluebell
// ├── conf
// │   └── config.yaml
// ├── static
// │   ├── css
// │   │   └── app.0afe9dae.css
// │   ├── favicon.ico
// │   ├── img
// │   │   ├── avatar.7b0a9835.png
// │   │   ├── iconfont.cdbe38a0.svg
// │   │   ├── logo.da56125f.png
// │   │   └── search.8e85063d.png
// │   └── js
// │       ├── app.9f3efa6d.js
// │       ├── app.9f3efa6d.js.map
// │       ├── chunk-vendors.57f9e9d6.js
// │       └── chunk-vendors.57f9e9d6.js.map
// └── templates
//     └── index.html
// nohup 用于在系统后台不挂断地运行命令，不挂断指的是退出执行命令的终端也不会影响程序的运行。
// sudo nohup ./bin/bluebell conf/config.yaml > nohup_bluebell.log 2>&1 &
// 当然我们也可以通过以下命令查看 bluebell 相关活动进程：
// ps -ef | grep bluebell
// root      6338  4048  0 08:43 pts/0    00:00:00 ./bin/bluebell conf/config.yaml
// root      6376  4048  0 08:43 pts/0    00:00:00 grep --color=auto bluebell

// 在需要静态文件分离、需要配置多个域名及证书、需要自建负载均衡层等稍复杂的场景下，
// 我们一般需要搭配第三方的web服务器（Nginx、Apache）来部署我们的程序。
