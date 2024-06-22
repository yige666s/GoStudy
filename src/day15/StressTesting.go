package main

// 1. 响应时间(RT) ：指系统对请求作出响应的时间.
// 2. 吞吐量(Throughput) ：指系统在单位时间内处理请求的数量
// 3. QPS每秒查询率(Query Per Second) ：“每秒查询率”，是一台服务器每秒能够响应的查询次数，
//			是对一个特定的查询服务器在规定时间内所处理流量多少的衡量标准
// 4. TPS(TransactionPerSecond)：每秒钟系统能够处理的交易或事务的数量
// 5. 并发连接数：某个时刻服务器所接受的请求总数

// ab全称Apache Bench，是Apache自带的性能测试工具,指定同时连接数、请求数以及URL
// ab -n 10000 -c 100 -t 10 "http://127.0.0.1:8080/api/v1/posts?size=10"
// ab -n 10000 -c 100 -t 10 -p post.json -T "application/json" "http://127.0.0.1:8080/api/v1/post"

// wrk是一款开源的HTTP性能测试工具，它和上面提到的ab同属于HTTP性能测试工具
// 它比ab功能更加强大，可以通过编写lua脚本来支持更加复杂的测试场景。
// wrk -t8 -c100 -d30s --latency http://127.0.0.1:8080/api/v1/posts?size=10

// go-wrk是Go语言版本的wrk,使用方法同wrk类似，基本格式如下
// go-wrk [flags] url
// go-wrk -t=8 -c=100 -n=10000 "http://127.0.0.1:8080/api/v1/posts?size=10"
