package decoder

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	type Person struct {
		Name   string
		Age    int
		Emails []string
		Extra  map[string]string
	}

	input := map[string]interface{}{
		"name":   "Mitchell",
		"age":    91,
		"emails": []string{"one", "two", "three"},
		"extra": map[string]string{
			"twitter": "mitchellh",
		},
	}

	var result Person
	err := Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", result)
}

func TestDecodeTag(t *testing.T) {
	type Person struct {
		Name   string `guozhen:"name2"`
		Age    int `guozhen:"age2"`
		Emails []string `guozhen:"emails2"`
		Extra  map[string]string `guozhen:"extra2"`
	}

	input := map[string]interface{}{
		"name2":   "Mitchell",
		"age2":    91,
		"emails2": []string{"one", "two", "three"},
		"extra2": map[string]string{
			"twitter": "mitchellh",
		},
	}

	var result Person
	err := Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", result)
}