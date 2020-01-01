package main

import (
	"errors"
	"fmt"
)

type ErrorType struct {
	cause error
	class string
}

func CreateType(t string) ErrorType {
	return ErrorType{
		class: t,
	}
}

func (t *ErrorType) New(args... interface{}) *ErrorType {
	e := args[0]

	switch e.(type) {
	case error :
		t.cause = e.(error)
	case string :
		t.cause = errors.New(e.(string))
	case nil:
		t.cause = nil
	default:
		panic("invalid error")
	}

	return t
}

func (t *ErrorType) IsTypeOf(nt ErrorType) bool {
	return t.class == nt.class
}

func (t *ErrorType) String() string {
	if t.cause != nil {
		return t.cause.Error()
	}
	return "empty cause"
}

func main() {
	var StateError = CreateType("StateError")
	var InvalidArgumentsError = CreateType("InvalidArgumentsError")

	// 1
	e1 := StateError.New("test e1 error")
	fmt.Println(e1.IsTypeOf(StateError))
	fmt.Println(e1.IsTypeOf(InvalidArgumentsError))
	fmt.Println(e1)

	// 2
	e2 := StateError.New(errors.New("test 2"))
	fmt.Println(e2)

	e3 := StateError.New(nil)
	fmt.Println(e3)
}