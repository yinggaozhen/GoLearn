// +build tag_3

package tags

import (
	"fmt"
	"testing"
)

// go test ./mytest/tags --tags tag_3
func TestTags3_before(t *testing.T) {
	fmt.Println("TestTags3_before")
}

func TestTags3(t *testing.T) {
	panic("tags 3")
}

func TestTags3_2(t *testing.T) {
	fmt.Println("TestTags3_2")
}