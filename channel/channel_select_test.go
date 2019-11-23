package array

import (
	"fmt"
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
	case <-channel:
		fmt.Println("channel output")
	default:
		fmt.Println("default")
	}

	fmt.Println("select finish")
}