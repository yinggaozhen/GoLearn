package array

import (
	"fmt"
	"testing"
	"time"
)

// 循环触发
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	count := 0

	for {
		select {
		case t := <- ticker.C:
			fmt.Println(t.Format(time.RFC3339))
			count++

			if count >= 5 {
				fmt.Println("ticker stop")
				ticker.Stop()
				return
			}
		}
	}
}