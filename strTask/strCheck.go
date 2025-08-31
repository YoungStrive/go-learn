package main

import "fmt"

func main() {
	str := "{()}[]"
	var checkResult bool = isValid(str)
	fmt.Println("字符串是否有效：", checkResult)
}

func isValid(str string) bool {
	n := len(str)
	if n%2 != 0 {
		return false
	}
	st := []rune{}
	for _, c := range str {
		//fmt.Println("字符：", string(c))
		switch c {
		case '{':
			//入栈 对应的右括号
			st = append(st, '}')
		case '(':
			//入栈 对应的右括号
			st = append(st, ')')
		case '[':
			//入栈 对应的右括号
			st = append(st, ']')
			//那么就是右括号的 就出栈
		default:
			if len(st) == 0 || st[len(st)-1] != c {
				return false // 没有左括号，或者左括号类型不对
			}
			st = st[:len(st)-1] // 出栈

		}
	}

	return len(st) == 0

}
