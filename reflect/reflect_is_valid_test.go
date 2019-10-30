package rand

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type IsValidTest struct {
}
// 判断是否为nil
func TestIsValid(t *testing.T) {
	a := 1
	b := 0
	var c int
	d := errors.New("hello")
	var e map[string]string
	var f IsValidTest

	fmt.Println(reflect.ValueOf(a).IsValid())
	fmt.Println(reflect.ValueOf(b).IsValid())
	fmt.Println(reflect.ValueOf(c).IsValid())
	fmt.Println(reflect.ValueOf(d).IsValid())
	fmt.Println(reflect.ValueOf(e).IsValid())
	fmt.Println(reflect.ValueOf(e["hello"]).IsValid())
	fmt.Println(reflect.ValueOf(nil).IsValid())
	fmt.Println(reflect.ValueOf(f).IsValid())
}