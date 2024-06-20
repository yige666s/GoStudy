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

// 扫描或遍历所有key
// 在Redis中可以使用KEYS prefix* 命令按前缀查询所有符合条件的 key，
// go-redis库中提供了Keys方法实现类似查询key的功能。
// 针对这种需要遍历大量key的场景，go-redis中提供了一个简化方法——Iterator
// delKeysByMatch 按match格式扫描所有key并删除
func delKeysByMatch(match string, timeout time.Duration) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	iter := rdb.Scan(ctx, 0, match, 0).Iterator()
	// iter := rdb.SScan(ctx, "set-key", 0, "perfix:*", 0).Iterator()		// set
	// iter := rdb.HScan(ctx, "hash-key", 0, "perfix:*", 0).Iterator()        // hash
	// iter := rdb.ZScan(ctx, "sorted-hash-key", 0, "prefix:*", 0).Iterator() //zset
	for iter.Next(ctx) {
		err := rdb.Del(ctx, iter.Val()).Err()
		if err != nil {
			panic(err)
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

// Redis Pipeline 允许通过使用单个 client-server-client 往返执行多个命令来提高性能
// 节省了执行命令的网络往返时间（RTT）。
// 在那些我们需要一次性执行多个命令的场景下，就可以考虑使用 pipeline 来优化。
func PipelinedTest() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	cancel()

	// Pipelined 方法会在函数退出时调用 Exec。
	var incr *redis.IntCmd
	var boolres *redis.BoolCmd
	_, err := rdb.Pipelined(ctx, func(p redis.Pipeliner) error {
		incr = p.Incr(ctx, "pipelined_counter")
		boolres = p.Expire(ctx, "pipelined_counter", time.Hour)
		return nil
	})
	if err != nil {
		panic(err)
	}
	// 在pipeline执行后获取到结果
	fmt.Println(incr.Val(), boolres.Val())

	// 可以遍历 pipeline 命令的返回值依次获取每个命令的结果
	cmds, _ := rdb.Pipelined(ctx, func(p redis.Pipeliner) error {
		for i := 0; i < 100; i++ {
			p.Get(ctx, fmt.Sprintf("key%d", i))
		}
		return nil
	})

	for _, cmdres := range cmds {
		// cmdres被类型断言为*redis.StringCmd类型。类型断言是Go语言中的一种操作，用于将接口类型的变量转换为具体的类型。
		// 语法是：x.(T)，其中x是接口类型的变量，T是你要断言的具体类型。如果断言成功，结果就是具体类型的值；如果失败，会引发panic。
		fmt.Println(cmdres.(*redis.StringCmd).Val())
	}
}

// Multi/exec能够确保在multi/exec两个语句之间的命令之间没有其他客户端正在执行命令。
// 在这种场景我们需要使用 TxPipeline 或 TxPipelined 方法将 pipeline 命令使用 MULTI 和EXEC包裹起来。
// 通常搭配 WATCH命令来执行事务操作。从使用WATCH命令监视某个 key 开始，
// 直到执行EXEC命令的这段时间里，如果有其他用户抢先对被监视的 key 进行了替换、更新、删除等操作，
// 那么当用户尝试执行EXEC的时候，事务将失败并返回一个错误，用户可以根据这个错误选择重试事务或者放弃事务。
func watchDemo(ctx context.Context, key string, rdb *redis.Client) error {
	return rdb.Watch(ctx, func(tx *redis.Tx) error {
		n, err := tx.Get(ctx, key).Int() // 获取key的值
		if err != nil && err != redis.Nil {
			return err
		}

		// 假设操作耗时5秒,5秒内我们通过其他的客户端修改key，当前事务就会失败
		time.Sleep(5 * time.Second)

		_, err = tx.Pipelined(ctx, func(p redis.Pipeliner) error {
			p.Set(ctx, key, n+1, time.Hour) // 修改key的值
			return nil
		})
		return err
	}, key)
}
