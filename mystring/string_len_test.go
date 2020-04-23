package array

import (
	"fmt"
	"testing"
)

// 字符串长度
func TestLen(t *testing.T) {
	fmt.Println(len("123"))
	fmt.Println(len("中文啊"))
	fmt.Println(len([]rune("中文啊")))

	fmt.Println(len("中文啊123x西安市大「『"))
	fmt.Println(len([]rune("中文啊123x西安市大「『")))
}

type HHH struct {
	shuzu []int
}

func (h *HHH) add(i int) {
	h.shuzu = append(h.shuzu, i)
}

func TestHHHH(t *testing.T) {
	hh := &HHH{}
	hh.add(1)
	fmt.Println(hh.shuzu)

	hh.add(2)
	fmt.Println(hh.shuzu)
}