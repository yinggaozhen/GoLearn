package test

import (
	"fmt"
	"testing"
)

type Binding interface {
	Name() string
}

type JsonBinding struct {
}

func (b JsonBinding) Name() string {
	return "json"
}

func TestBinding(t *testing.T) {
	b := JsonBinding{}

	fmt.Println(getName(b))
}

func getName(b Binding) string {
	return b.Name()
}