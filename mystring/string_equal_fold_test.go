package array

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringEqualFold(t *testing.T) {
	fmt.Println(strings.EqualFold("name", "Name"))
	fmt.Println(strings.EqualFold("NAME", "Name"))
	fmt.Println(strings.EqualFold("NAME", "Name2"))
}