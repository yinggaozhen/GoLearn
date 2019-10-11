package array

import (
	"fmt"
	"testing"
)

const (
	ALevel uint32 = iota
	BLevel
	CLevel
)

func TestIota(t *testing.T) {
	fmt.Println(ALevel, BLevel, CLevel)
}