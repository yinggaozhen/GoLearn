package rand

import (
	"bytes"
	"fmt"
	"sort"
	"testing"
)

type ColPerson struct {
	Name string
	Age int
	Score int
}

type SortColPerson []ColPerson
type byAge SortColPerson
type byScore SortColPerson

func (m byAge) Less(i, j int) bool { return m[i].Age < m[j].Age }
func (m byAge) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m byAge) Len() int {
	return len(m)
}

func (m byScore) Less(i, j int) bool { return m[i].Score < m[j].Score }
func (m byScore) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m byScore) Len() int {
	return len(m)
}

func (m SortColPerson) String() string {
	buf := bytes.Buffer{}
	for _, v := range m {
		_, _ = fmt.Fprintf(&buf, "%s %d %d\n", v.Name, v.Age, v.Score)
	}

	return buf.String()
}

var _ sort.Interface = SortPerson{}

func TestSortFuncByAge(t *testing.T) {
	var sortPersonSlice = SortColPerson{
		{Name:"Age18Person", Age:18, Score:100},
		{Name:"Age21erson", Age:21, Score:92},
		{Name:"Age16Person", Age:16, Score:98},
		{Name:"Age24Person", Age:24, Score:66},
		{Name:"Age188Person", Age:188, Score:87},
		{Name:"Age0Person", Age:0, Score:75},
	}

	sort.Sort(byAge(sortPersonSlice))
	fmt.Println(sortPersonSlice)
}

func TestSortFuncByScore(t *testing.T) {
	var sortPersonSlice = SortColPerson{
		{Name:"Age18Person", Age:18, Score:100},
		{Name:"Age21erson", Age:21, Score:92},
		{Name:"Age16Person", Age:16, Score:98},
		{Name:"Age24Person", Age:24, Score:66},
		{Name:"Age188Person", Age:188, Score:87},
		{Name:"Age0Person", Age:0, Score:75},
	}

	sort.Sort(byScore(sortPersonSlice))
	fmt.Println(sortPersonSlice)
}