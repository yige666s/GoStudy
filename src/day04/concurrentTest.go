package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// 协程（coroutine）：非操作系统提供而是由[用户自行创建和控制的用户态‘线程’]，比线程更轻量级

// Go语言中使用 goroutine 非常简单，只需要在函数或方法调用前加上go关键字就可以创建一个 goroutine ，从而让该函数或方法在新创建的 goroutine 中执行。
// 匿名函数也支持使用go关键字创建 goroutine 去执行。

// 全局等待组变量
var wg sync.WaitGroup

func Hello() {
	fmt.Println("hello world")
	wg.Done()
}

func TheLast() {
	fmt.Println("the last ?")
	wg.Done()
}

func GoroutineTest() {
	wg.Add(2) // 登记2个goroutine
	go Hello()
	go TheLast()
	fmt.Println("go hello")
	// go func() {
	// 	fmt.Println("the last ?")
	// 	wg.Done()
	// }()
	wg.Wait() // 阻塞等待登记的goroutine完成
}

func hello1(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("hello", i)
}
func main2() {
	// goroutine 的调度是随机的
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello1(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}

// goroutine 的调度是Go语言运行时（runtime）层面的实现，是完全由 Go 语言本身实现的一套调度系统——go scheduler。
// 它的作用是按照一定的规则将所有的 goroutine 调度到操作系统线程上执行。
//在经历数个版本的迭代之后，目前 Go 语言的调度器采用的是 GPM 调度模型

func GMPTest() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

// Go语言采用的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。

func chanTest() {
	c1 := make(chan int, 1) // 声明一个缓冲区大小为1的通道，数据结构是一个FIFO队列
	c1 <- 10
	x := <-c1
	fmt.Println(x)
	close(c1)
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main3() {
	ch := make(chan int)
	go recv(ch) // 创建一个 goroutine 从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}

func f3(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main4() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	f3(ch)
}

func Producer() <-chan int {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

func Consumer(ch <-chan int) {
	sum := 0
	for v := range ch {
		sum += v
	}
	fmt.Println(sum)
}

func main5() {
	ch := Producer()
	Consumer(ch)
}

func selectTest() {
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

// 使用once实现单例
type sigleton struct{}

var instance *sigleton
var once sync.Once

func Getinstance() *sigleton {
	once.Do(func() {
		instance = &sigleton{}
	})
	return instance
}

// Go 语言中内置的 map 不是并发安全的，Go语言的sync包中提供了一个开箱即用的并发安全版 map——sync.Map
var m = sync.Map{}

func SyncMap() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Println(value)
		}(i)
	}
	wg.Wait()
}

// 针对整数数据类型（int32、uint32、int64、uint64）我们还可以使用原子操作来保证并发安全，
// 通常直接使用原子操作比使用锁操作效率更高。Go语言中原子操作由内置的标准库sync/atomic提供
// atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用。这些函数必须谨慎地保证正确使用。
// 除了某些特殊的底层应用，使用通道或者 sync 包的函数/类型实现同步更好。

type Counter interface {
	Inc()
	Load() int64
}

type AtomicCounter struct {
	counter int64
}

func (p *AtomicCounter) Inc() {
	atomic.AddInt64(&p.counter, 1)
}
func (p *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&p.counter)
}

func AtomicTest(c Counter) {
	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}

func main6() {
	c := AtomicCounter{}
	AtomicTest(&c)
}

func RandomInt64() <-chan int {
	jobchan := make(chan int, 2)
	for i := 0; i < 1; i++ {
		rnd := rand.Int()
		jobchan <- rnd
		fmt.Println(rnd)
	}
	return jobchan
}

func Sumofin64(c <-chan int) <-chan int {
	resultchan := make(chan int, 2)
	sum := 0
	for i := 0; i < 1; i++ { // 注意这个地方开的协程数量和上面发送的随机数数量有关，随机数数量如果少于协程数量会死锁
		wg.Add(1)
		go func() {
			var x int = <-c
			for {
				a := x % 10
				sum += int(a)
				if x /= 10; x == 0 {
					break
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	resultchan <- sum
	return resultchan
}

func Calc(ch <-chan int) {
	x := <-ch
	fmt.Println(x)
}

func main7() {
	c1 := RandomInt64()
	c2 := Sumofin64(c1)
	Calc(c2)
}
