package array

import (
	"fmt"
	"testing"
)

// https://sanyuesha.com/2017/07/23/go-defer/

// multi defer
func TestMulti(t *testing.T) {
	defer func() {fmt.Println(1)}()
	defer func() {fmt.Println(2)}()
	defer func() {fmt.Println(3)}()
}