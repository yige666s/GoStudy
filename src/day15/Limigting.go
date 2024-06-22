package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jujuratelimit "github.com/juju/ratelimit"
	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
)

// 限流又称为流量控制（流控），通常是指限制到达系统的并发请求数

// 漏桶算法
// ratelimit 库的基本原理是：
// 固定时间间隔：限制请求的处理速度为固定的时间间隔，即在每个固定的时间间隔内只能处理一个请求。
// 计时器：通过一个计时器来确保请求的处理时间间隔符合预设的速率。

// Take 会阻塞确保两次请求之间的时间走完
// Take 调用平均数为 time.Second/rate.

// func (t *limit) Take() time.Time {
// 	t.Lock()
// 	defer t.Unlock()

// 	now := t.clock.Now()

// 	// 如果是第一次请求就直接放行
// 	if t.last.IsZero() {
// 		t.last = now
// 		return t.last
// 	}

// 	// sleepFor 根据 perRequest 和上一次请求的时刻计算应该sleep的时间
// 	// 由于每次请求间隔的时间可能会超过perRequest, 所以这个数字可能为负数，并在多个请求之间累加
// 	//t.sleeepfor =  t.perrequest + t.last - now，这个值是理论下次请求和当前时间的差值，有正有负
// 	t.sleepFor += t.perRequest - now.Sub(t.last)

// 	// 我们不应该让sleepFor负的太多，因为这意味着一个服务在短时间内慢了很多随后会得到更高的RPS。
// 	if t.sleepFor < t.maxSlack {
// 		t.sleepFor = t.maxSlack
// 	}

// 	// 如果 sleepFor 是正值那么就 sleep
// 	if t.sleepFor > 0 {
// 		t.clock.Sleep(t.sleepFor)
// 		t.last = now.Add(t.sleepFor)
// 		t.sleepFor = 0
// 	} else {
// 		t.last = now
// 	}

// 	return t.last
// }

func LeakyBukketLimit() {
	r1 := ratelimit.New(100)
	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := r1.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}

// 令牌桶其实和漏桶的原理类似，令牌桶按固定的速率往桶里放入令牌，
// 并且只要能从桶里取出令牌就能通过，令牌桶支持突发流量的快速处理。
func TokenLimit() { //Go 官方提供的速率限制库
	limiter := rate.NewLimiter(1, 5) // 每秒1个请求，最多累积5个请求

	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// 我们只需要每次来取令牌的时候计算一下，当前是否有足够的令牌就可以了
// 当前令牌数 = 上一次剩余的令牌数 + (本次取令牌的时刻-上一次取令牌的时刻)/放置令牌的时间间隔 * 每次放置的令牌数
// func (tb *Bucket) currentTick(now time.Time) int64 {
// 	return int64(now.Sub(tb.startTime) / tb.fillInterval)
// }
// func (tb *Bucket) adjustavailableTokens(tick int64) {
// 	if tb.availableTokens >= tb.capacity {
// 		return
// 	}
// 	tb.availableTokens += (tick - tb.latestTick) * tb.quantum
// 	if tb.availableTokens > tb.capacity {
// 		tb.availableTokens = tb.capacity
// 	}
// 	tb.latestTick = tick
// 	return
// }
// func (tb *Bucket) takeAvailable(now time.Time, count int64) int64 {
// 	if count <= 0 {
// 		return 0
// 	}
// 	tb.adjustavailableTokens(tb.currentTick(now))
// 	if tb.availableTokens <= 0 {
// 		return 0
// 	}
// 	if count > tb.availableTokens {
// 		count = tb.availableTokens
// 	}
// 	tb.availableTokens -= count
// 	return count
// }

func TokenLimit2() {
	bucket := jujuratelimit.NewBucket(200*time.Millisecond, 1)
	for i := 0; i < 10; i++ {
		if bucket.TakeAvailable(1) > 0 {
			fmt.Println("allowed")
		} else {
			fmt.Println("denied")
		}
		time.Sleep(200 * time.Millisecond)
	}
}

// 如果要对全站限流就可以注册成全局的中间件，如果是某一组路由需要限流，那么就只需将该限流中间件注册到对应的路由组即可
func ReteLimitMiddleware(fillInterval time.Duration, cap int64) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bucket := jujuratelimit.NewBucket(fillInterval, cap)
		if bucket.TakeAvailable(1) < 1 { // 如果取不到令牌就中断本次请求返回限流提示
			ctx.String(http.StatusOK, "pleaese wait a minute")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
