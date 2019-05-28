package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("[]"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid(""))
	fmt.Println(isValid("(("))
}

func isValid(s string) bool {
	strLen := len(s)

	if strLen % 2 != 0 {
		return false
	}

	var collect = make(map[string][]int)
	collect["("] = []int{}
	collect["["] = []int{}
	collect["{"] = []int{}

	var symbolMapping = make(map[string]string)
	symbolMapping[")"] = "("
	symbolMapping["]"] = "["
	symbolMapping["}"] = "{"

	var i int

	for i = 0; i < strLen; i++ {
		c := string(s[i])

		if c == "[" || c == "(" || c == "{" {
			collect[c] = append(collect[c], i)
		} else {
			indexPosCollects := collect[symbolMapping[c]]

			if len(indexPosCollects) == 0 {
				return false
			}

			lastPos := indexPosCollects[len(indexPosCollects) - 1]
			if (i - lastPos) % 2 != 1 {
				return false
			}

			collect[symbolMapping[c]] = indexPosCollects[:len(indexPosCollects) - 1]
		}
	}

	if len(collect["("]) != 0 || len(collect["["]) != 0 || len(collect["{"]) != 0 {
		return false
	}

	return true
}
