package main

import (
	"fmt"
	"sync"
	"time"
)

// 线程安全的计数器
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// 增加计数
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	c.count++
	defer c.mu.Unlock()
}

func main() {
	countNum := SafeCounter{}

	// 启动10个goroutine同时增加计数
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				countNum.Increment()
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("countNum:", countNum.count)
}
