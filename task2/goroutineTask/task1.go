package main

import (
	"fmt"
)

// 奇数
func oddNum(numArr []int) {

	fmt.Println(numArr)

	for _, num := range numArr {
		fmt.Println(num)
		if num%2 != 0 {
			fmt.Println(num)
		}
	}
}

func main() {
	fmt.Println("Starting the program")
	numArr := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	go oddNum(numArr)
	go evenNum()

}

// 偶数
func evenNum() {
	fmt.Println("为何不执行")
	numArr := []int{2, 3, 4, 5, 6, 7, 9, 10}
	for _, num := range numArr {
		if num%2 == 0 {
			fmt.Println(num)
		}
	}
}
