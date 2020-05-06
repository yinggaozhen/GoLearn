package array

import (
	"fmt"
	"strconv"
	"testing"
)

// 测试channel select
func TestChannelSelect(t *testing.T) {
	channel := make(chan int, 100)

	select {
	case <-channel:
		fmt.Println("channel output")
	default:
	}

	fmt.Println("select finish")
}

// 测试channel select
func TestChannelSelectDeadlock(t *testing.T) {
	channel := make(chan int, 100)

	select {
	// deadlock
	case <-channel:
		fmt.Println("channel output")
	}

	fmt.Println("select finish")
}

// 测试channel select
func TestChannelSelectNormal(t *testing.T) {
	channel := make(chan int, 100)
	channel <- 1

	select {
	// deadlock
	case r := <-channel:
		fmt.Println("channel output : " + strconv.Itoa(r))
	default:
		fmt.Println("default")
	}

	fmt.Println("select finish")
}

// 测试channel range
func TestChannelRange(t *testing.T) {
	channel := make(chan int, 100)
	channel <- 1
	channel <- 2
	channel <- 3

	// close(channel) // 如果不加上这一句，后边range channel会出现死锁

	fmt.Println(len(channel)) // 3

	for cr := range channel {
		l := len(channel)
		fmt.Println("channel output : " + strconv.Itoa(cr))
		fmt.Println("channel len : " + strconv.Itoa(l))

		// 如果不明确channel是否关闭，则需要加一个len
		// if l == 0 { break }
	}

	fmt.Println("range finish")
}