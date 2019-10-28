package array

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringTitle(t *testing.T) {
	fmt.Println(strings.Title("hello"))
	fmt.Println(strings.Title("hello world"))
	fmt.Println(strings.Title("HELLO WORLD NIHAO"))
	fmt.Println(strings.Title("HELLO WORLD NIHAO 123"))
}