package main

import (
	"fmt"
	"time"
)

// 发送数据
func sendNum(ch chan<- int) {
	for i := 1; i < 30; i++ {
		ch <- i
	}
	//关闭通道
	close(ch)

}

// 只接收channel的函数
func receiveNum(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
	}
}

func main() {
	ch := make(chan int, 10)
	go sendNum(ch)
	go receiveNum(ch)
	// 使用select进行多路复用
	timeout := time.After(10 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(500 * time.Millisecond)
		}
	}

}
