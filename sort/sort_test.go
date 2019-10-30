package rand

import (
	"fmt"
	"sort"
	"testing"
)

// 排序
func TestSort(t *testing.T) {
	sortedString := []string{"b", "c", "a"}
	sort.Strings(sortedString)
	fmt.Println(sortedString)
}