package test

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

type TestIOWriteStruct struct {
	http.ResponseWriter
	code int
}

func (io *TestIOWriteStruct) Write(message []byte) (int, error) {
	fmt.Println(string(message))
	return 10086, nil
}

func (io *TestIOWriteStruct) WriteHeader(statusCode int) {
	io.code = statusCode
}

// test struct extends
func TestIOWrite(t *testing.T) {
	w := &TestIOWriteStruct{}

	n, err := io.WriteString(w, "Hello World!")
	fmt.Println(n, err)
}
