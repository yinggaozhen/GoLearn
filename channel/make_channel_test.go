package array

import (
	"testing"
)

// 抛出一个新的异常
func TestMakeChannel(t *testing.T) {
	channel := make(chan int, 100)

	channel <- 11 // [11]
	channel <- 12 // [11, 12]
	channel <- 13 // [11, 12, 13]

	// 11 <- [12, 13]
	member := <- channel
	println(member)
	println(len(channel), <- channel)
	println(len(channel), <- channel)
	println(len(channel))
}