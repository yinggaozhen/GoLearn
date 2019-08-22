package test

import (
	"fmt"
	"sync"
	"testing"
)

type PeoplePool struct {
	sync.Pool
}

type BookPool struct {
	sync.Pool
}

// test struct extends
func TestSyncPoolMulti(t *testing.T) {
	p := PeoplePool{}
	p.Pool.New = func() interface{} {
		fmt.Println("people instance")
		return "People"
	}

	fmt.Println(p.Pool.Get().(string)) // 会重新New

	p.Pool.Put("hello people")
	fmt.Println(p.Pool.Get().(string))

	fmt.Println(p.Pool.Get().(string)) // 会重新New

	//  book

	b := BookPool{}
	b.Pool.New = func() interface{} {
		fmt.Println("book instance")
		return "Book"
	}

	fmt.Println(b.Pool.Get().(string)) // 会重新New

	b.Pool.Put("hello book")
	fmt.Println(b.Pool.Get().(string))

	fmt.Println(b.Pool.Get().(string)) // 会重新New
}
