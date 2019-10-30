package array

import (
	"fmt"
	"testing"
)

func TestMapKey(t *testing.T) {
	var hello = make(map[string]string)
	n := "name"

	hello[n] = "jack"

	fmt.Println(hello)
}
