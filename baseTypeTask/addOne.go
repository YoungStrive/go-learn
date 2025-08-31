package main

import "fmt"

func main() {
	numArray := []uint32{2, 3, 2, 1}
	count := len(numArray)
	//	var num1 = 1
	var numFlag uint32 = 10
	var flagCount int = count - 1
	var resultNum uint32 = 0
	for i := 0; i < count; i++ {
		var num2 = numArray[i]
		var result uint32 = customPower(numFlag, flagCount)
		resultNum += result * num2
		flagCount--
	}
	fmt.Println(resultNum + 1)

}

func customPower(base uint32, exponent int) uint32 {
	var result uint32 = 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
