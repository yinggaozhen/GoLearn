package array

import (
	"fmt"
	"testing"
)

// Map Push
func TestMapNew(t *testing.T) {
	// 1. new map
	var student = make(map[string]string)
	student["name"] = "jack"
	student["age"] = "18"

	// 2. get property
	name := student["name"]
	fmt.Println(name)

	unknown, ok := student["age_unknown"]
	fmt.Println("age_unknown : ", unknown)
	fmt.Println("age_unknown_ok", ok)

	// 3. foreach
	for key, value := range student {
		fmt.Println("key", key)
		fmt.Println("value", value)
	}
}
