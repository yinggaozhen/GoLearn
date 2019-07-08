package test

import (
	"testing"
)

type arrayExtends []int

func (a arrayExtends) getSecondsMember() int {
	if len(a) < 2 {
		return -1
	}

	return a[1]
}

// test struct extends
func TestStructExtends2(t *testing.T) {
	array := arrayExtends{1, 2, 3}
	println(array.getSecondsMember())
}
