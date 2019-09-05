package rand

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandInt(t *testing.T) {
	fmt.Println(rand.Int())
}

func TestRandSeedInt(t *testing.T) {
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(1000))
}