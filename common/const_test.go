package array

import (
	"fmt"
	"testing"
)

type ConstTestType string

const (
	Boy  = 1
	Girl = 2
	UnknownA
	UnknownB
	UnknownC = 3
	UnknownD
	UnknownE
	UnknownF ConstTestType = "F"
	UnknownG
)
func TestConst(t *testing.T) {
	fmt.Println(Boy)  // 1
	fmt.Println(Girl) // 2
	fmt.Println(UnknownA) // 2
	fmt.Println(UnknownB) // 3
	fmt.Println(UnknownC) // 3
	fmt.Println(UnknownD) // 3
	fmt.Println(UnknownE) // 3
	fmt.Println(UnknownF) // F
	fmt.Println(UnknownG) // F
}