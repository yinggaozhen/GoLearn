package array

import (
	"bytes"
	"fmt"
	"testing"
)

func TestStringAdd(t *testing.T) {
	a := "hello "
	b := "world"

	c := a + b

	fmt.Println(c)
}

func TestStringAdd2(t *testing.T) {
	c := bytes.Buffer{}

	c.WriteString("hello ")
	c.WriteString("world")

	fmt.Println(c.String())
	c.Reset()
	fmt.Println(c.String())
}

func TestStringAdd3(t *testing.T) {
	c := &bytes.Buffer{}

	fmt.Fprintln(c, "hello ")
	fmt.Fprintf(c, "%s", "world")

	fmt.Println(c.String())

}