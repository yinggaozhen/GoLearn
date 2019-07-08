package test

import (
	"sync"
	"testing"
)

var once sync.Once

// test struct extends
func TestOnce(t *testing.T) {
	count := 0

	once.Do(func() {
		count++
		println(count)
	})

	once.Do(func() {
		count++
		println(count)
	})
	once.Do(func() {
		count++
		println(count)
	})
	once.Do(func() {
		count++
		println(count)
	})
	once.Do(func() {
		count++
		println(count)
	})
	once.Do(func() {
		count++
		println(count)
	})

}
