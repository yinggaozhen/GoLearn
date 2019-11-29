package array

import (
	"fmt"
	"testing"
)

func EmptyArray(a []string) bool {
	return len(a) == 0
}

func TestEmptyArray(t *testing.T) {
	fmt.Println([]string(nil))
	fmt.Println(EmptyArray(nil))
}
