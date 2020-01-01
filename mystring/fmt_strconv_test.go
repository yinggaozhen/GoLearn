package array

import (
	"fmt"
	"strconv"
	"testing"
)

func TestConv(t *testing.T) {
	fmt.Println(strconv.Atoi("111"))
	fmt.Println(strconv.Atoi("-1"))
	fmt.Println(strconv.Atoi("1.222"))
	fmt.Println(strconv.Atoi("1.222xxx"))
	fmt.Println(strconv.Atoi("xxx"))
	fmt.Println(strconv.Atoi("1123xxxx"))
}

func TestConvFloat(t *testing.T) {
	fmt.Println(strconv.ParseFloat("111", 64))
	fmt.Println(strconv.ParseFloat("111.11", 64))
	fmt.Println(strconv.ParseFloat("111.11111", 64))
	fmt.Println(strconv.ParseFloat("0111.11111", 64))
	fmt.Println(strconv.ParseFloat("-0111111111.11111", 64))
}

func TestConvFloat2(t *testing.T) {
	text1 := 111.11611
	text2 := fmt.Sprintf("%.2f", text1)
	fmt.Println(text2)
}