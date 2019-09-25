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
	case channel <- 1:
		fmt.Println("channel 1")
	case channel <- 2:
		fmt.Println("channel 2")
	case o := <-channel:
		fmt.Println("default" + strconv.Itoa(o))
	}


	//channel <- 1111
	//<- channel
}