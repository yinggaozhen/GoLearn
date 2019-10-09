package array

import (
	"fmt"
	"testing"
)

func TestClosureFunc(t *testing.T) {
	f := closureFunc2(1, 2)
	r := f()
	fmt.Println(r)
}

func closureFunc2(a int, b int) func() int {
	return func() int {
		return a+b
	}
}
