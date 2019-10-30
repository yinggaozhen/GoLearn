package rand

import (
	"bytes"
	"fmt"
	"sort"
	"testing"
)

type person struct {
	Name string
	Age int
}

type SortPerson []person

func (m SortPerson) Less(i, j int) bool {
	return m[i].Age < m[j].Age
}

func (m SortPerson) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m SortPerson) Len() int {
	return len(m)
}

func (m SortPerson) String() string {
	buf := bytes.Buffer{}
	for _, v := range m {
		fmt.Fprintf(&buf, "%d %s\n", v.Age, v.Name)
	}

	return buf.String()
}

var _ sort.Interface = SortPerson{}

// 按照Key进行排序
func TestSortFuncAsKey(t *testing.T) {
	var sortPersonSlice = SortPerson{
		{Name:"Age18Person", Age:18},
		{Name:"Age21erson", Age:21},
		{Name:"Age16Person", Age:16},
		{Name:"Age24Person", Age:24},
		{Name:"Age188Person", Age:188},
		{Name:"Age0Person", Age:0},
	}

	sort.Sort(sortPersonSlice)

	fmt.Println(sortPersonSlice)
}