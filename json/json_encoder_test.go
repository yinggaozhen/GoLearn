package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

type testJsonEncoderStudent struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Language []string `json:"language"`
}

func TestJsonEncoder(t *testing.T) {
	var network bytes.Buffer
	var encoder = json.NewEncoder(&network)

	guozhen := testJsonEncoderStudent{
		Name : `guozhen`,
		Age : 18,
		Language: []string{"hello", "world"},
	}
	err := encoder.Encode(guozhen)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(network.String()) // {"name":"guozhen","age":18,"language":["hello","world"]}
	}

	// =======================

	xiaoming := testJsonEncoderStudent{
		Name : `xiaoming`,
		Age : 25,
		Language: []string{"hello", "world"},
	}
	err2 := encoder.Encode(xiaoming)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		// TODO 在guozhen基础上，多一个xiaoming
		// {"name":"guozhen","age":18,"language":["hello","world"]}
		// {"name":"xiaoming","age":25,"language":["hello","world"]}
		fmt.Println(network.String())
	}

	// =======================

	laozhang := testJsonEncoderStudent{
		Name : `laozhang`,
		Age : 68,
		Language: []string{"hello", "world"},
	}
	network.Reset()
	err3 := encoder.Encode(laozhang)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		// TODO 由于network被重置，所以只输出一个
		// {"name":"laozhang","age":68,"language":["hello","world"]}
		fmt.Println(network.String())
	}
}