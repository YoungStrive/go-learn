package main

import "fmt"

func main() {
	str := "Hello World"
	var checkResult bool = isValid(str)
	fmt.Println("字符串是否有效：", checkResult)
}

func isValid(str string) bool {
	n := len(str)
	fmt.Println("字符串长度为：", n)
	if n%2 != 0 {
		return false
	}
	return true

}
