package array

import (
	"fmt"
	"testing"
)

type VStudent struct {
	name string
	age int
}

func TestFmtV(t *testing.T) {
	s := VStudent{
		name: "jack",
		age: 18,
	}

	fmt.Printf("%T\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%#v\n", s)
	fmt.Println(s)
}