package test

import (
	"fmt"
	"testing"
)

type human interface {
	speak()
}

type student map[string]string

var _ human = student(nil)

func (student) speak() {
	fmt.Println("student speak")
}

func TestMapExtends(t *testing.T) {
	xiaoming := &student{}
	xiaoming.speak()
}