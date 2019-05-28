package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World c"))
	fmt.Println(lengthOfLastWord(""))
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord("        "))
}

// 用内置方法实现
func lengthOfLastWord(s string) int {
	strs := strings.Split(strings.Trim(s, " "), " ")
	return len(strs[len(strs) - 1])
}

