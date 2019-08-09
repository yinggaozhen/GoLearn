package array

import (
	"fmt"
	"testing"
	"time"
)

func TestStampSub(t *testing.T) {
	start := time.Unix(1564588800, 0)
	end := time.Unix(1564675200, 0)

	// 第一种减法，是用于直接用time的sub方法去减
	fmt.Println(time.Time.Sub(end, start).Seconds())

	// 第二种减法，是用对象减对象
	fmt.Println(end.Sub(start).Seconds())
}