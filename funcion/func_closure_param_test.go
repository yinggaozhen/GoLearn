package array

import (
	"testing"
)

type ClosureParamFunc func()

type ClosureParam struct {
	name string
}

func (c *ClosureParam) ClosureParamFunc() {
	println(c.name)
}

func TestFuncClosureParam(t *testing.T) {
	object := ClosureParam{
		name: "jack",
	}
	helloClosure(object.ClosureParamFunc)
}

func helloClosure(callable ClosureParamFunc) {
	callable()
}
