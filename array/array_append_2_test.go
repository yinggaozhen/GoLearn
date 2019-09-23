package array

import (
	"fmt"
	"testing"
)

// append 2
func TestAppend2(t *testing.T) {
	var ta = []int{1, 2, 3, 4}
	fmt.Println(append(ta[1:1], ta[1:]...))
}