package array

import (
	"fmt"
	"testing"
)

// https://sanyuesha.com/2017/07/23/go-defer/

func TestDeferTrap(t *testing.T) {
	fmt.Println(trap())
}

func trap()(i int) {
	defer func() {
		i++
	}()

	return 1
}