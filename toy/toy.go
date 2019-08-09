package main

import (
	"fmt"
	"reflect"
	"strings"
)

type T struct {
	A string
	B struct {
		C int   `yaml:"c"`
		D []int `yaml:",flow"`
	}
}

func (t T) echo(path string) {
	nodes := strings.Split(path, ".")

	p := reflect.ValueOf(t)
	for _, node := range nodes {
		field := p.FieldByName(node)
		if !field.IsValid() {
			fmt.Println("panic")
			return
		}

		p = field
	}

	fmt.Println(p)
}

func main() {
	t := T{
		A: "tom",
		B: struct {
			C int   `yaml:"c"`
			D []int `yaml:",flow"`
		}{
			C: 10086,
			D: []int{0, 1},
		},
	}

	t.echo("A")
	t.echo("B.C")
	t.echo("B.D")
	t.echo("B.E")
}