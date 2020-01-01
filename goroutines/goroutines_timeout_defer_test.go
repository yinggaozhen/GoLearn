package array

import (
	"fmt"
	"testing"
	"time"
)

func getDelayHandler() (after <-chan time.Time, handler func()){
	after = time.After(3*time.Second)
	handler = func() {
		time.Sleep(5*time.Second)
		fmt.Println("hello world")
	}
	return
}

func TestTimeoutDefer(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func () {
			defer func() {
				if e := recover(); e != nil {
					fmt.Println(e)
				}
			}()

			timeoutDefer()
		}()

	}

	time.Sleep(time.Second*100)
}

func timeoutDefer() {
	after, handler := getDelayHandler()

	errorChan := make(chan error)
	go func() {
		defer func() {
			errorChan <- nil
			fmt.Println("???")
		}()

		handler()
	}()

	select {
	case <-errorChan:
		fmt.Println("error chan")
	case <-after:
		panic("timeout le")
	}
}