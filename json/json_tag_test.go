package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type testJsonTagStudent struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Language []string `json:"language"`
}

func TestJsonTagStruct(t *testing.T) {
	student := `{"name":"guozhen","age":19,"language":["hello","world"]}`

	var decode testJsonTagStudent
	error2 := json.Unmarshal([]byte(student), &decode)
	if error2 == nil {
		fmt.Println(decode.Name)
		fmt.Println(decode.Age)
		fmt.Println(decode.Language)
	} else {
		fmt.Println(error2)
	}

	encode, _ := json.Marshal(decode)
	fmt.Println(string(encode))
}