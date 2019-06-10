package array

import (
	"fmt"
	"testing"
)

// Map Interface
func TestMapInterface(t *testing.T) {
	// string map

	// not exist
	fmt.Println(isset(map[interface{}]interface{} {"hello" : "world"}, "hello")) // true
	fmt.Println(isset(map[interface{}]interface{} {"hello" : "world"}, "no-exist")) // false
	fmt.Println(isset(map[interface{}]interface{} {11 : "world"}, 11)) // true
	fmt.Println(isset(map[interface{}]interface{} {"hello" : "world"}, "hello")) // false
}

func isset(input map[interface{}]interface{}, key interface{}) bool {
	// setType := reflect.TypeOf(input).Kind().String()
	return input[key] != nil
}