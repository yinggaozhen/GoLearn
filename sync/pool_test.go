package test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

var pool sync.Pool
var variable int

// pool 是当没有的值的时候会New
// 被Put一次之后，下次再取就会从pool取出来，然后清空
func TestSyncPool(t *testing.T) {
	variable = 8888

	pool.New = func() interface{} {
		fmt.Println("我被调用了 " + strconv.Itoa(variable))
		return variable
	}

	fmt.Println("=========== 第一阶段 ===========")

	fmt.Println(pool.Get().(int)) // 8888
	fmt.Println(pool.Get().(int)) // 8888
	fmt.Println(pool.Get().(int)) // 8888

	fmt.Println("=========== 第二阶段 ===========")

	pool.Put(10086)
	fmt.Println(pool.Get().(int)) // 10086
	fmt.Println(pool.Get().(int)) // 我被调用了，8888
	fmt.Println(pool.Get().(int)) // 我被调用了，8888

	fmt.Println("=========== 第三阶段 ===========")

	pool.Put(10087)
	fmt.Println(pool.Get().(int)) // 10087
	fmt.Println(pool.Get().(int)) // 10087

	fmt.Println("=========== 第四阶段 ===========")

	pool.Put("variable5")
	fmt.Println(pool.Get().(string))
	fmt.Println(pool.Get().(string))
	fmt.Println(pool.Get().(string))
}
