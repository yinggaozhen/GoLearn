package array

import (
	"fmt"
	"testing"
)

// Map Interface
func TestMathDiv(t *testing.T) {
	a1 := 16
	b1 := 10
	c1 := a1 / b1
	fmt.Println(c1) // 1

	a2 := 9
	b2 := 10
	c2 := a2 / b2
	fmt.Println(c2) // 0
}