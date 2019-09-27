package test

import (
	"fmt"
	"testing"
)

type SturctTest struct {
	NameChange string
	PointerNameChange string
}

// 不是引用，值传递
func (s SturctTest) setName(name string) {
	s.NameChange = name
}

// 其实就是引用了
func (s *SturctTest) setNamePointer(name string) {
	s.PointerNameChange = name
}

// test struct extends
func TestStructExtends3(t *testing.T) {
	o := SturctTest{
		NameChange: "Tom",
		PointerNameChange: "Tom",
	}

	o.setName("TomChange")
	fmt.Println(o.NameChange)

	o.setNamePointer("TomChange")
	fmt.Println(o.PointerNameChange)
}
