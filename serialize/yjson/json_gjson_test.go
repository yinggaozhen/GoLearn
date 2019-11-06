package yjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseJson(t *testing.T) {
	var jsonString = `{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":["Sara","Alex","Jack"],"fav.movie":"Deer Hunter","friends":[{"first":"Dale","last":"Murphy","age":44,"nets":["ig","fb","tw"]},{"first":"Roger","last":"Craig","age":68,"nets":["fb","tw"]},{"first":"Jane","last":"Murphy","age":47,"nets":["ig","tw"]}]}`
	var parseResult interface{}

	json.Unmarshal([]byte(jsonString), &parseResult)

	fmt.Println(parseResult.(map[string]interface{})["name"])
}

func TestYJson(t *testing.T) {
	var jsonString = `{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":["Sara","Alex","Jack"],"fav.movie":"Deer Hunter","friends":[{"first":"Dale","last":"Murphy","age":44,"nets":["ig","fb","tw"]},{"first":"Roger","last":"Craig","age":68,"nets":["fb","tw"]},{"first":"Jane","last":"Murphy","age":47,"nets":["ig","tw"]}]}`

	r, _ := Parse(jsonString)

	fmt.Println(r.Get("name"))
}

func TestArray(t *testing.T) {
	var jsonString = `[1, 2, 3]`

	r, _ := Parse(jsonString)

	fmt.Println(r.Get(0))
	fmt.Println(r.Get(1))
	fmt.Println(r.Get("2"))
	fmt.Println(r.Get("4"))
}

func TestYJsonObjectPath(t *testing.T) {
	var jsonString = `{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":["Sara","Alex","Jack"],"fav.movie":"Deer Hunter","friends":[{"first":"Dale","last":"Murphy","age":44,"nets":["ig","fb","tw"]},{"first":"Roger","last":"Craig","age":68,"nets":["fb","tw"]},{"first":"Jane","last":"Murphy","age":47,"nets":["ig","tw"]}]}`

	r, _ := Parse(jsonString)

	fmt.Println(r.Get("name.first"))
	fmt.Println(r.Get("children"))
	fmt.Println(r.Get("children.0"))
}