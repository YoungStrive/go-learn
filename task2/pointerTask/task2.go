package main

import "fmt"

func main() {
	var num = []int{2, 3, 4, 5, 6}
	var num1 *[]int = &num
	double(num1)

}

func double(a *[]int) []int {
	num := *a
	result := []int{}
	for _, v := range num {
		result = append(result, (v * 2))
	}
	fmt.Println("Result: ", result)
	return result

}
