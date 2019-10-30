package test

import (
	"fmt"
	"testing"
)

type MapExtends2 struct {
}

func (m *MapExtends2) Echo() {
	fmt.Println("hello")
}

func TestMapExtends2(t *testing.T) {
	var m *MapExtends2

	m.Echo()
}