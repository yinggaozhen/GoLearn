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