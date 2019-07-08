package array

import (
	"testing"
)

// Map Push
func TestFuncReturnInterface(t *testing.T) {
	hello := hello()
	println(hello)
}

func hello() (world string) {
	world = "world"
	return
}
