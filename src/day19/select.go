package main

import "fmt"

// select语句中所有case中的表达式都必须是channel的发送或接收操作
func SelectTest() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	select {
	// select 会一直阻塞下去，直到其中的一个 channel 转为就绪状态时执行对应case分支的代码。
	// 如果多个channel同时就绪的话则随机选择一个case执行。
	case <-ch1:
		fmt.Println("ch1")
	case ch2 <- 1:
		fmt.Println("ch2")
	}

	//空的 select 语句会直接阻塞当前的goroutine，使得该goroutine进入无法被唤醒的永久休眠状态。
	// select {}

	// 如果select中只包含一个case，那么该select就变成了一个阻塞的channel读/写操作。
	select {
	case <-ch1:
		fmt.Println("wait for ch1 channel data")
	default:
		// default语句，用于当其他case都不满足时执行一些默认操作,当于做了一个非阻塞的channel读取操作。
		fmt.Println("default")
	}
}

func SelectWithPriorityDemo(ch1, ch2 <-chan int, stopch chan struct{}) {
	for {
		select {
		case <-stopch:
			return
		case job1 := <-ch1:
			fmt.Println(job1)
		case job2 := <-ch2:
		priority:
			for {
				select { //内部循环使用 select 语句，优先处理 ch1 中的消息。
				case job1 := <-ch1:
					fmt.Println(job1) // 如果 ch1 中有消息，将其打印出来
				default:
					break priority //如果 ch1 中没有消息，default 分支将触发，跳出内部循环
				}
			}
			fmt.Println(job2) // 跳出内部循环后，打印 job2，即 ch2 中接收到的消息。
		}
		// 这种设计保证了 ch1 的消息优先于 ch2 的消息被处理，
		// 即便 ch2 先接收到消息，处理 ch2 消息之前也会处理所有 ch1 中的消息。
	}
}

// kubernetes/pkg/controller/nodelifecycle/scheduler/taint_manager.go
// func (tc *NoExecuteTaintManager) worker(worker int, done func(), stopCh <-chan struct{}) {
// 	defer done()

// 	// 当处理具体事件的时候，我们会希望 Node 的更新操作优先于 Pod 的更新
// 	// 因为 NodeUpdates 与 NoExecuteTaintManager无关应该尽快处理
// 	// -- 我们不希望用户(或系统)等到PodUpdate队列被耗尽后，才开始从受污染的Node中清除pod。
// 	for {
// 		select {
// 		case <-stopCh:
// 			return
// 		case nodeUpdate := <-tc.nodeUpdateChannels[worker]:
// 			tc.handleNodeUpdate(nodeUpdate)
// 			tc.nodeUpdateQueue.Done(nodeUpdate)
// 		case podUpdate := <-tc.podUpdateChannels[worker]:
// 			// 如果我们发现了一个 Pod 需要更新，我么你需要先清空 Node 队列.
// 		priority:
// 			for {
// 				select {
// 				case nodeUpdate := <-tc.nodeUpdateChannels[worker]:
// 					tc.handleNodeUpdate(nodeUpdate)
// 					tc.nodeUpdateQueue.Done(nodeUpdate)
// 				default:
// 					break priority
// 				}
// 			}
// 			// 在 Node 队列清空后我们再处理 podUpdate.
// 			tc.handlePodUpdate(podUpdate)
// 			tc.podUpdateQueue.Done(podUpdate)
// 		}
// 	}
// }
