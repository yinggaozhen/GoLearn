package array

import (
	"fmt"
	"testing"
)

// https://sanyuesha.com/2017/07/23/go-defer/

func TestFilePath(t *testing.T) {
	i := 1

	defer func(i int) {
		fmt.Println(i)
	}(i)

	i++
	fmt.Println(i)

	return
}