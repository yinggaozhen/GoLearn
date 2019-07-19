package test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

var pool sync.Pool
var variable int

// test struct extends
func TestSyncPool(t *testing.T) {
	pool.New = func() interface{} {
		fmt.Println("调用了呀 " + strconv.Itoa(variable))
		return variable
	}

	variable := pool.Get().(int)
	pool.Put(variable)

	variable2 := pool.Get().(int)
	pool.Put(variable2)

	variable3 := pool.Get().(int)
	pool.Put(variable3)

	variable4 := pool.Get().(int)
	pool.Put(variable4)

	variable5 := pool.Get().(int)
	pool.Put(variable5)

	fmt.Println(variable)
}
