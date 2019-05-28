package main

import (
	"fmt"
	"math"
	"strconv"
)

const POS_INT int = math.MaxInt32
const NEG_INT int = -math.MaxInt32

func main() {
	fmt.Println(Reverse(123))
	fmt.Println(Reverse(-123))
	fmt.Println(Reverse(120))
	fmt.Println(Reverse(1534236469))
}

func Reverse(x int) int {
	var isNegative bool = false
	if x < 0 {
		isNegative = true
	}

	var numString = strconv.Itoa(int(math.Abs(float64(x))))
	revereString := numReverse(numString)

	if isNegative {
		revereString = "-" + revereString
	}

	i, err := strconv.Atoi(revereString)
	if err != nil {
		return 0
	}

	if i > POS_INT || i < NEG_INT {
		return 0
	}

	return i
}

func numReverse(reverse string) string {
	var result string

	for _, v := range reverse {
		result = string(v) + result
	}

	return result
}
