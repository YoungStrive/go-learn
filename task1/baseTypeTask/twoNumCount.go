package main

import "fmt"

func main() {
	numArr := []int{1, 2, 3, 4, 5}
	targetNum := 7
	result := []int{}
	var lenFlag = len(numArr)
	for i := 0; i < len(numArr); i++ {
		var j = numArr[i]
		if j > targetNum {
			continue
		}
		if i+1 < lenFlag {
			var k = numArr[i+1]
			if j+k == targetNum {
				result = append(result, i, i+1)
			}
		}

	}
	fmt.Println("The result is: ", result)

}
