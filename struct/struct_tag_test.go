package test

import (
	"fmt"
	"reflect"
	"testing"
)

type structTag struct {
	Name string `default:"Unknown Name"`
	Sex string `Man:"1" Woman:"2"`
}

func TestStructTag(t *testing.T) {
	s := structTag{}

	fmt.Println(reflect.TypeOf(s).Field(0).Tag.Get("default"))
	// 获取不到的话，不会报错
	fmt.Println(reflect.TypeOf(s).Field(0).Tag.Get("default0"))

	fmt.Println(reflect.TypeOf(s).Field(1).Tag.Get("Man"))
	fmt.Println(reflect.TypeOf(s).Field(1).Tag.Get("Woman"))
}