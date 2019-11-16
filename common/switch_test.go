package array

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	a := 1
	// switch 只要选中一个，其他分支便不再生效
	switch a {
	case 1:
		fmt.Println(a)
		a++
	case 2:
		fmt.Println(a)
		a++
	case 3:
		fmt.Println(a)
		a++
	}

	fmt.Println(a)
}