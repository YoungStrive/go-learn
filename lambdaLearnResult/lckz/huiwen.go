package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := 12321
	//数字转字符串
	var numStr = strconv.Itoa(num1)
	var reverseNumStr = reverseString(numStr)
	if numStr == reverseNumStr {
		fmt.Println(num1, "是回文数")

	} else {
		fmt.Println(num1, "不是回文数")
	}

}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	//这个是对应的ASCII码，所以输出的结果是121
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)

}
