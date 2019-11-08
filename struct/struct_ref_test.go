package test

import (
	"fmt"
	"testing"
)

type students struct {
	peoples []string
}

func (s *students) add(name string) {
	s.peoples = append(s.peoples, name)
}

func (s students) add2(name string) {
	s.peoples = append(s.peoples, name)
}

func TestAppendStudent(t *testing.T) {
	s := students{}
	s.add("jack")
	s.add("tom")
	fmt.Println(s)

	s.add2("jack2")
	s.add2("tom2")
	fmt.Println(s)
}