package main

import "fmt"

func main() {

	arr1 := []int{1, 2, 2, 3, 3, 4, 4, 5}
	map1 := make(map[int]int)
	for _, v := range arr1 {
		// 如果键不存在，ok 的值为 false，count1 的值为该类型的零值
		count1, ok := map1[v]
		//表示有值
		if ok {
			map1[v] = count1 + 1
		} else {
			map1[v] = 1
		}
	}
	fmt.Println(map1)

}
