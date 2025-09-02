package main

import (
	"fmt"
	"sync"
)

// 奇数
func oddNum(wg *sync.WaitGroup) {
	// 当goroutine完成时，通知WaitGroup
	defer wg.Done()
	numArr := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	for _, num := range numArr {
		if num%2 != 0 {
			fmt.Println(num)
		}
	}
}

func main() {
	// 创建一个wait group等待两个goroutine完成
	var wg sync.WaitGroup
	wg.Add(2)
	go oddNum(&wg)
	go evenNum(&wg)
	// 确保所有goroutine都已完成
	wg.Wait()
	fmt.Println("都执行完成了")

}

// 偶数
func evenNum(wg *sync.WaitGroup) {
	defer wg.Done()
	numArr := []int{2, 3, 4, 5, 6, 7, 9, 10}
	for _, num := range numArr {
		if num%2 == 0 {
			fmt.Println(num)
		}
	}
}
