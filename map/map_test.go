package array

import (
	"fmt"
	"testing"
)

// Map Push
func TestMapNew(t *testing.T) {
	// 1. new map
	jack := make(map[string]string)
	jack["name"] = "jack"
	jack["age"] = "18"

	tom := map[string]string{"name": "tom", "age": "18", "typo": "typo"}

	// 2. get property
	name := jack["name"]
	fmt.Println(name)

	unknown, ok := jack["age_unknown"]
	fmt.Println("age_unknown : ", unknown)
	fmt.Println("age_unknown_ok", ok)

	// 3. foreach
	for key, value := range jack {
		fmt.Println("key", key)
		fmt.Println("value", value)
	}

	// 4. delete
	delete(tom, "typo")
	for key, value := range tom {
		fmt.Println("key", key)
		fmt.Println("value", value)
	}
}
