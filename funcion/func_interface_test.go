package array

import (
	"fmt"
	"reflect"
	"testing"
)

// Map Push
func TestFuncInterface(t *testing.T) {
	any(1)
	any("world")
	any(false)
}

func any(args interface{}) {
	argsType := reflect.TypeOf(args).String()

	switch argsType {
	case "string":
		fmt.Println("hello" + args.(string))
	case "int":
		fmt.Println(100 + args.(int))
	default:
		fmt.Println("type:" + argsType)
	}
}
