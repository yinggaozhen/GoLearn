package array

import (
	"fmt"
	"testing"
)

type HeaderSource map[string]string

func TestTypeInvoke(t *testing.T) {
	hs := HeaderSource{
		"hello": "world",
	}

	// 第一眼这个写法看上去很不理解，为什么type能作为函数进行调用。
	// 其实这个只是类型转换
	hs2 := HeaderSource(hs)
	fmt.Println(hs2)

	hs["hello"] = "world2"
	fmt.Println(hs)
	fmt.Println(hs2)
}