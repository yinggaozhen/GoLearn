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
	variable = 8888

	pool.New = func() interface{} {
		fmt.Println("我被调用了 " + strconv.Itoa(variable))
		return variable
	}

	variable := pool.Get().(int) // 我被调用了,8888
	fmt.Println(variable) // 8888

	pool.Put(10086)
	fmt.Println(pool.Get().(int)) // 10086
	fmt.Println(pool.Get().(int)) // 我被调用了，10086
	fmt.Println(pool.Get().(int)) // 我被调用了，10086

	pool.Put(10087)
	fmt.Println(pool.Get().(int)) // 10087

	pool.Put("variable5")
	fmt.Println(pool.Get().(string))
}
