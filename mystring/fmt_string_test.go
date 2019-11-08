package array

import (
	"bytes"
	"fmt"
	"testing"
)

func TestFmt(t *testing.T) {
	w := bytes.Buffer{}
	fmt.Fprint(&w, 1)
	fmt.Println(w.String())

	w.Reset()
	fmt.Fprint(&w, 1.222)
	fmt.Println(w.String())

	w.Reset()
	fmt.Fprint(&w, 16564894879)
	fmt.Println(w.String())

	w.Reset()
	fmt.Fprint(&w, "hello world")
	fmt.Println(w.String())
}