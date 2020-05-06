package array

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestCancelContext(t *testing.T) {
	cancelCtx, cancelFn := context.WithCancel(context.Background())
	incrC := incr(cancelCtx)

	for c := range incrC {
		if c >= 50 {
			cancelFn()
			break
		}
	}
}

func incr(ctx context.Context) chan int {
	counterC := make(chan int)

	go func() {
		var counter int

		for {
			counter += rand.Intn(5)
			time.Sleep(time.Second)

			select {
			case <-ctx.Done():
				fmt.Println("incr done")
				return
			default:
				counterC <- counter
				fmt.Println("level up to " + strconv.Itoa(counter))
			}
		}
	}()

	return counterC
}