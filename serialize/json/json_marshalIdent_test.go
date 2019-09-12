package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type testMarshalIdentStudent struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Language []string `json:"language"`
}

func TestMarshalIdentJson(t *testing.T) {
	student := testMarshalIdentStudent{
		Name: "guozhen",
		Age: 26,
		Language: []string{"hello"},
	}

	encode, error1 := json.MarshalIndent(student, "", "  ")
	if error1 == nil {
		fmt.Println(string(encode))
	} else {
		fmt.Println(error1)
	}
}