package array

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 抛出一个新的异常
func TestChannelGoRoutines(t *testing.T) {
	myChan := make(chan int, 10)

	fmt.Println("xxxxxxxxx111111")

	go getChan(myChan)
	for i := 0; i < 100; i++ {
		go putChan(myChan, i)
	}

	fmt.Println("xxxxxxxxx222222")
	time.Sleep(time.Second * 1)
}

func putChan(c chan int, i int) {
	fmt.Println("putChan" + strconv.Itoa(i))
	c <- i
}

func getChan(myChan chan int) {
	for {
		select {
		case number := <- myChan:
			fmt.Println(number)
		}
	}
}