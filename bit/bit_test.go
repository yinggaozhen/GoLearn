package bit

import (
	"fmt"
	"math"
	"testing"
)

func TestBit(t *testing.T) {
	fmt.Println(1 << 2)
	fmt.Println(1 << 3)
	fmt.Println(1 << 4)
	fmt.Println(^(1 << 4))
	fmt.Println(^(1 << 5))
	fmt.Println(^(math.MaxUint64 << 5) & (math.MaxUint64 << 2))
}