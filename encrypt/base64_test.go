package array

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	// aGVsbG8=
	encodeString := base64.StdEncoding.EncodeToString([]byte("hello"))
	fmt.Println(encodeString)

	decodeString, _ := base64.StdEncoding.DecodeString(encodeString)
	fmt.Println(string(decodeString))
}