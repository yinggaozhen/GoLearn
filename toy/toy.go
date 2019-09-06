package main

import "fmt"

func main() {
	func1()
	fmt.Println("func1 fin")

	func2()
	fmt.Println("func2 fin")
}

func func1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}

		fmt.Println("OKOKOKOOKOK")
	}()

	panic(1)
}

func func2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}

		fmt.Println("OKOKOKOOKOK")
	}()

	panic(2)
}