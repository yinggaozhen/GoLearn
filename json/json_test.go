package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type testJsonStudent struct {
	Name string
	Age int
	Language []string
}

func TestJson(t *testing.T) {
	encode, error1 := json.Marshal([]int{1, 2, 3})
	if error1 == nil {
		fmt.Println(string(encode))
	} else {
		fmt.Println(error1)
	}

	var decode []int
	error2 := json.Unmarshal([]byte(encode), &decode)
	if error2 == nil {
		fmt.Println(decode)
		fmt.Println(append(decode, 4))
	} else {
		fmt.Println(error2)
	}
}

func TestJsonStruct(t *testing.T) {
	student := testJsonStudent{
		Name: "guozhen",
		Age: 18,
		Language: []string{"hello", "world"},
	}

	encode, error1 := json.Marshal(student)
	if error1 == nil {
		fmt.Println(string(encode))
	} else {
		fmt.Println(error1)
	}

	var decode testJsonStudent
	error2 := json.Unmarshal([]byte(encode), &decode)
	if error2 == nil {
		fmt.Println(decode.Name)
		fmt.Println(decode.Age)
		fmt.Println(decode.Language)
	} else {
		fmt.Println(error2)
	}
}