package array

import (
	"fmt"
	"reflect"
	"testing"
)

type CaseStudent struct {
	Name string
}
func TestMapKeyCase(t *testing.T) {
	m := map[string]string{
		"name": "jack",
	}

	// var s CaseStudent
	// nameValue := reflect.Indirect(reflect.ValueOf(s)).FieldByName("Name")

	nameValue := reflect.ValueOf("Name")
	fmt.Println(reflect.ValueOf(m).MapIndex(nameValue).IsValid())
}
