package array

import (
	"fmt"
	"testing"
)

// Map Push
func TestMapGet(t *testing.T) {
	// 1. new map
	jack := make(map[string]string)
	jack["name"] = "jack"
	jack["age"] = "18"

	name, err := jack["name"]
	fmt.Println(name, err)

	age, err := jack["age"]
	fmt.Println(age, err)

	// 获取不存在的key,不会报错，err返回是false
	hello, err := jack["hello"]
	fmt.Println(hello, err)
}
