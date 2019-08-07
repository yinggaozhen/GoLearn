package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type testJsonDecoderStudent struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Language []string `json:"language"`
}

func TestJsonDecoder(t *testing.T) {
	student := &testJsonDecoderStudent{}

	err := json.NewDecoder(strings.NewReader(`{"name":"guozhen","age":19,"language":["hello","world"]}`)).Decode(student)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(student.Name)
		fmt.Println(student.Age)
		fmt.Println(student.Language)
	}
}