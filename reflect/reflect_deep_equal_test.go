package rand

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeepEqual1(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{2, 1, 3}


	ai := getInterface(a)
	bi := getInterface(b)

	// false，数组顺序不同也会被认为是不同
	fmt.Println(reflect.DeepEqual(ai, bi))
}

func getInterface(i interface{}) interface{} {
	return i
}