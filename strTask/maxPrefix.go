package main

import (
	"fmt"
	"strings"
)

func main() {
	resultStr := []string{}
	var arr1 = []string{"book", "cooks", "dookstore", "fookcase", "gookshelf"}
	//找出最小长度的字符串
	var minLenStr = arr1[0]
	for i := 1; i < len(arr1); i++ {
		if len(minLenStr) > len(arr1[i]) {
			minLenStr = arr1[i]
		}
	}
	var runes = []rune(minLenStr)
	for i := 1; i <= len(runes); i++ {
		var prefix = string(runes[0:i])
		boolFlag := true
		for j := 0; j < len(arr1); j++ {
			//如果不包含就为false
			if !strings.Contains(arr1[j], prefix) {
				boolFlag = false
				break
			}
		}
		if boolFlag {
			resultStr = append(resultStr, prefix)
		}
	}

	if len(resultStr) == 0 {
		fmt.Println("no prefix found")
	} else {
		fmt.Println("max prefix:", resultStr[len(resultStr)-1])
	}

}
