package assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssert(t *testing.T) {
	a := assert.New(t)

	expect := 1
	actual := 2
	a.Equal(expect, actual, "not equal")
}