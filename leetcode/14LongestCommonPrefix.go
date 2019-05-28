package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
	fmt.Println(longestCommonPrefix([]string{}))
	fmt.Println(longestCommonPrefix([]string{"aa", "a"}))
}

func longestCommonPrefix(strs []string) string {
	var prefix string = ""
	var length int = 0

	for {
		if len(strs) == 0 || length > len(strs[0]) {
			return prefix
		}

		var tmpPrefix = strs[0][:length]

		for _, str := range strs {
			str = string(str)

			if length > len(str) || str[:length] != tmpPrefix {
				return prefix
			}
		}

		length++
		prefix = tmpPrefix
	}
}