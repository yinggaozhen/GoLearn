package array

import (
	"fmt"
	"testing"
)

type MyLevel uint

func (l *MyLevel) Replace(newLevel MyLevel) {
	*l = newLevel
}

func TestMyStringToString(t *testing.T) {
	l := MyLevel(1)
	fmt.Println(l)

	l.Replace(233)
	fmt.Println(l)
}