package test

import (
	"fmt"
	"testing"
)

type MyObject struct {
	name string
}

// test struct extends
func TestObjectInterface(t *testing.T) {
	o := MyObject{
		name: "jack",
	}

	fakeChangeName(o)
	fmt.Println(o.name)

	realChangeName(&o)
	fmt.Println(o.name)
}

func fakeChangeName(o MyObject) {
	o.name = "tom"
}

func realChangeName(o *MyObject) {
	o.name = "tom"
}
