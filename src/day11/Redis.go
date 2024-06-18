package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// func RedisClient() {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "",
// 		DB:       0,  // 数据库
// 		PoolSize: 20, // 连接池大小
// 		TLSConfig: &tls.Config{
// 			MinVersion: tls.VersionTLS12,
// 			// Certificates: []tls.Certificate{cert},
// 			// ServerName: "your.domain.name",
// 		},
// 	})

// 	// 从URL中解析Redis配置
// 	// opt, err := redis.ParseURL("redis://<user>:<passwd>@localhost:6379/<db>")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// rdb := redis.NewClient(opt)

// 	// redis Sentinel模式
// 	rdb1 := redis.NewFailoverClient(&redis.FailoverOptions{
// 		MasterName:    "master-name",
// 		SentinelAddrs: []string{":9126", ":9127", ":9128"},
// 	})

// 	// redis cluster模式
// 	// go-redis 支持按延迟或随机路由命令
// 	rdb2 := redis.NewClusterClient(&redis.ClusterOptions{
// 		Addrs: []string{":7000", ":7001", ":7002", ":7003", "7004", "7005"},
// 	})
// }

// func RedisCommand() {

// 	rdb := redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 		DB:   0,
// 	})

// 	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
// 	defer cancel()

// 	val, err := rdb.Get(ctx, "key").Result()         // 直接执行命令获取val,err
// 	cmder := rdb.Get(ctx, "key")                     // 先获取到命令对象
// 	val1 := rdb.Get(ctx, "key").Val()                // 直接执行命令获取值
// 	err1 := rdb.Set(ctx, "key", 10, time.Hour).Err() // 直接执行命令获取err
// 	fmt.Println(val, err)
// 	fmt.Println(cmder.Val(), cmder.Err())

// 	// do方法执行任意命令
// 	val2, err2 := rdb.Do(ctx, "set", "key", 10, "EX", 3600).Result()
// 	val3, err3 := rdb.Do(ctx, "get", "key").Result()

// 	redis_nil_test := func(key, defaultValue string) (string, error) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
// 		defer cancel()

// 		val, err := rdb.Get(ctx, key).Result()
// 		if err != nil {
// 			if errors.Is(err, redis.Nil) {
// 				return defaultValue, nil // 查询的key为空
// 			}
// 			return "other error", err //其他错误
// 		}
// 		return val, nil
// 	}
// }

func ZsetTest() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	// key
	zsetkey := "language_rank"
	// value
	// languages := []redis.Z{
	// 	{Score: 90.0, Member: "Golang"},
	// 	{Score: 98.0, Member: "Java"},
	// 	{Score: 95.0, Member: "Python"},
	// 	{Score: 97.0, Member: "JS"},
	// 	{Score: 99.0, Member: "C/C++"},
	// }
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// ZADD
	// err := rdb.ZAdd(ctx, zsetkey, languages...).Err()
	// if err != nil {
	// 	fmt.Printf("zadd failed ,err :%v\n", err)
	// 	return
	// }
	// fmt.Println("zadd success")

	// // INCRESE
	// newScore, err := rdb.ZIncrBy(ctx, zsetkey, 10, "golang").Result()
	// if err != nil {
	// 	fmt.Println("")
	// }
	// fmt.Printf("newscore is %f\n", newScore)

	// 取分数最高的三个
	ret := rdb.ZRevRangeWithScores(ctx, zsetkey, 0, 2).Val()
	for _, v := range ret {
		fmt.Println(v.Member, v.Score)
	}

	// 取95-100分之间的数据
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err := rdb.ZRangeByScoreWithScores(ctx, zsetkey, op).Result()
	if err != nil {
		fmt.Printf("ZrangeByScoreWithScores failed, err :%v\n", err)
		return
	}
	for _, v := range ret {
		fmt.Println(v.Member, v.Score)
	}
}

// TODO 扫描或遍历所有key
