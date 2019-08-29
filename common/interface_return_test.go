package array

import (
	"strconv"
	"testing"
)

type InterfaceReturn interface {
	GetName() string
}

type InterfaceReturnImpl struct {
	rn InterfaceReturn
}

func (impl InterfaceReturnImpl) GetName() string {
	return "nihao"
}

type InterfaceReturnString string

func (i InterfaceReturnString) GetName() string {
	return string(i)
}

func TestInterfaceReturn(t *testing.T) {
}

func getInterfaceReturn(number int) InterfaceReturn {
	if number%2 == 0 {
		return InterfaceReturnString(strconv.Itoa(number))
	} else {
		return &InterfaceReturnImpl{}
	}
}