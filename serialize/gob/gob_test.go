package test

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

type GobEcho struct {
	Name string
}

func (echo GobEcho) echo() {
	fmt.Println(echo.Name)
}

// go 原生的序列化函数，可以序列化方法
func TestGob(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	echo := GobEcho{
		Name: "jack",
	}

	error := gob.NewEncoder(buf).Encode(echo)
	if error != nil {
		fmt.Println(error)
	}

	result := new(GobEcho)
	gob.NewDecoder(buf).Decode(result)

	result.echo()
}