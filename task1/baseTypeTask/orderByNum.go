package main

import "fmt"

func main() {

	nums := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	result := []int{}
	map1 := make(map[int]int)
	for i, v := range nums {
		var count = map1[v]

		if count == 0 {
			map1[v] = i + 1
			result = append(result, v)
		} else {
			continue
		}

	}
	fmt.Println(result)
}
