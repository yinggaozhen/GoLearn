package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() (number int) {
	defer func() {
		number++
	}()

	return 1
}