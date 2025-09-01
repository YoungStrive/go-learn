package main

import "fmt"

func main() {

	var a int = 1
	/* 声明指针变量 */
	var a1 *int = &a

	c := add(a1)
	fmt.Println("c:", c)

}

func add(a *int) int {
	return 10 + *a
}
