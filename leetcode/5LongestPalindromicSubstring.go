package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	maxLp := ""

	for index := range s {
		oddLp := extendOddAround(index, s)
		evenLp := extendEvenAround(index, s)

		if len(oddLp) > len(maxLp) {
			maxLp = oddLp
		}

		if len(evenLp) > len(maxLp) && len(evenLp) > 1 {
			maxLp = evenLp
		}
	}

	return maxLp
}

func extendOddAround(point int, context string) string {
	length := 0
	result := ""

	for {
		start := point - length
		end := point + length + 1

		if (point - length < 0 || point + length >= len(context)) || (context[point - length] != context[point + length]) {
			return result
		}

		result = context[start:end]
		length++
	}
}

func extendEvenAround(point int, context string) string {
	length := 0
	result := ""

	for {
		start := point - length
		end := point + length + 1

		if start < 0 || end >= len(context) || context[start] != context[end] {
			return result
		}

		result = context[start:end+1]
		length++
	}
}