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

	// 因为channel数据已经取完了，在未close情况下，如果继续取回报错
	// println(<- channel)
}