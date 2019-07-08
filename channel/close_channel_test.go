package array

import (
	"testing"
)

// 关闭channel
func TestCloseChannel(t *testing.T) {
	channel := make(chan int, 100)

	channel <- 11 // [11]
	channel <- 12 // [11, 12]

	// close 之后取数据是正常的
	close(channel)

	v1, state1 := <- channel
	println(v1, state1) // 11 true

	v2, state2 := <- channel
	println(v2, state2) // 12 true

	v3, state3 := <- channel
	println(v3, state3) // 13 false。channel 被关闭之后，继续读取不会报错

	// close 之后，加数据的话，就会抛出异常
	channel <- 14
	println(len(channel))
}