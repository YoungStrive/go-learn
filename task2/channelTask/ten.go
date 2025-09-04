package main

import (
	"fmt"
	"sync"
)

// 发送数据
func sendNum(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 11; i++ {
		ch <- i
	}
	//关闭通道
	close(ch)

}

// 只接收channel的函数
func receiveNum(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
	}
}
func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	// 启动发送goroutine
	go sendNum(ch, &wg)

	// 启动接收goroutine
	go receiveNum(ch, &wg)

	wg.Wait()
	fmt.Println("都执行完成了")
}
