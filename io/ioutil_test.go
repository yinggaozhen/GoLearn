package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestRead(t *testing.T) {
	_ = ioutil.WriteFile("/tmp/hello", []byte("nihao"), 0644)
	content, _ := ioutil.ReadFile("/tmp/hello")

	fmt.Println(string(content))
}