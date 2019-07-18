package array

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestInterfaceType(t *testing.T) {
	fmt.Println(add(1, 2))
	fmt.Println(add("3", 7))
}

func add(a interface{}, b interface{}) int {
	return transToInt(a) + transToInt(b)
}

func transToInt(i interface{}) int {
	varType := reflect.TypeOf(i).String()

	switch varType {
	case "int":
		return i.(int)
	case "string":
		inta, _ := strconv.Atoi(i.(string))
		return inta
	default:
		return -1
	}
}