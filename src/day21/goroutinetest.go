package main

import (
	"fmt"
	"sync"
)

func goroutineTest() {
	var wg sync.WaitGroup
	wg.Add(3)

	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	ch3 := make(chan struct{}, 1)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch1
			fmt.Println("123")
			ch2 <- struct{}{}
		}

	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch2
			fmt.Println("456")
			ch3 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch3
			fmt.Println("789")
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}

	wg.Wait()
}
