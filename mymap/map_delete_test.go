package array

import (
	"fmt"
	"testing"
)

type H map[string]string

func (h H) Get(name string) string {
	if v, ok := h[name]; ok {
		return v
	}

	return "nil"
}

func TestMapDelete(t *testing.T) {
	h := make(H)
	h["hello"] = "world"

	fmt.Println(h.Get("hello"))
	delete(h, "hello")
	fmt.Println(h.Get("hello"))
	delete(h, "hello")
}