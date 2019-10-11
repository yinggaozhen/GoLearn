package array

import (
	"fmt"
	"testing"
)

const (
	ALevel uint32 = iota
	BLevel
	CLevel
	DLevel
	ELevel
)

const (
	A2Level uint32 = iota << 2
	B2Level
	C2Level
	D2Level
	E2Level
)

const (
	A3Level uint32 = 16
	B3Level
	C3Level
	D3Level
	E3Level
)

const (
	A4Level uint32 = iota
	B4Level
	_
	_
	E4Level
)

const (
	A5Level = iota
	B5Level = 10086
	C5Level = iota
	D5Level = iota
	E5Level
)

func TestIota(t *testing.T) {
	fmt.Println(ALevel, BLevel, CLevel, DLevel, ELevel)
}

func TestIota2(t *testing.T) {
	// 0 4 8 12 16
	fmt.Println(A2Level, B2Level, C2Level, D2Level, E2Level)
}

func TestIota3(t *testing.T) {
	// 16 16 16 16 16
	fmt.Println(A3Level, B3Level, C3Level, D3Level, E3Level)
}

func TestIota4(t *testing.T) {
	// 0 1 4
	// 2 和 3 被跳过
	fmt.Println(A4Level, B4Level, E4Level)
}

func TestIota5(t *testing.T) {
	// 0 10086 2 3 4
	// D5还是2，不会因为中间被插入所以不累加
	fmt.Println(A5Level, B5Level, C5Level, D5Level, E5Level)
}