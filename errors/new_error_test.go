package array

import (
	"errors"
	"fmt"
	"testing"
)

// 抛出一个新的异常
func TestFilePath(t *testing.T) {
	fmt.Println("start new error")
	error := errors.New("hello new error")
	fmt.Println(error)
}


// 抛出一个新的异常
func TestMultiReturnError(t *testing.T) {
	text, error := testMultiReturnError()
	fmt.Println("=== TestMultiReturnError start")
	fmt.Println(text)
	fmt.Println(error)
	fmt.Println("=== TestMultiReturnError end")
}

func testMultiReturnError() (string, error) {
	e := errors.New("multi return")
	content := "multi text"

	return content, e
}