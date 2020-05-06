package array

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestDeadlineContext(t *testing.T) {
	deadCtx, deadFn := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond * 3500))
	incr3(deadCtx)

	defer deadFn()
}

func incr3(ctx context.Context){
	var counter int
	for {
		counter += rand.Intn(5)
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			fmt.Println("finish")
			return
		default:
			fmt.Println("level up to " + strconv.Itoa(counter))
		}
	}
}