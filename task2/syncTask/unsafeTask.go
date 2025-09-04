package main

import (
	"fmt"
	"time"
)

type UnsafeCounter struct {
	count int
}

// 增加计数
func (c *UnsafeCounter) Increment() {
	c.count += 1
}

func main() {
	countNum := UnsafeCounter{}

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
