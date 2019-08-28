package array

import (
	"fmt"
	"testing"
)

var NilFunc func(string)

func TestEmptyFunc(t *testing.T) {
	if NilFunc == nil {
		NilFunc = func(s string) {
			fmt.Println(s)
		}

		NilFunc("empty func")
	} else {
		NilFunc("hello")
	}
}