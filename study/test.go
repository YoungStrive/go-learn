package main

import "fmt"

func TestFunc() {
	println("pkg1.TestFunc() called")
}
func main() {
	str1 := []string{"(", ")", "{", "}", ")", "(", ")"}
	stack := Stack{}
	for _, v := range str1 {
		stack.Push(v)
	}
	popped, _ := stack.Pop()
	fmt.Println(popped)
	fmt.Println(stack.elements)

}

type Stack struct {
	elements []string
}

func (s *Stack) Push(value string) {
	s.elements = append(s.elements, value)
}

// Pop 方法移除并返回栈顶元素
func (s *Stack) Pop() (string, bool) {
	if len(s.elements) == 0 {
		return "", false // 栈为空时返回0和false
	}
	index := len(s.elements) - 1    // 获取栈顶元素的索引
	element := s.elements[index]    // 获取栈顶元素的值
	s.elements = s.elements[:index] // 移除栈顶元素，通过重新切片实现
	return element, true            // 返回被移除的元素和true表示成功
}
