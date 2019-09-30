package controller

import "fmt"

type HelloController struct {
}

// 1. test comment
// 2. test comment
// 3. @router /hello
func (c HelloController) Echo() {
	fmt.Println("echo")
}