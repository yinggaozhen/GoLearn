package array

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrlEncode(t *testing.T) {
	// http%3A%2F%2Fwww.baidu.com%3Fa%3D2%26c%3D2
	queryEncode := url.QueryEscape("http://www.baidu.com?a=2&c=2")
	fmt.Println(queryEncode)
}

func TestUrlDecode(t *testing.T) {
	encode := url.QueryEscape("http://www.baidu.com?a=2&c=2")
	fmt.Println("encode url : " + encode)

	decode, _ := url.QueryUnescape(encode)
	fmt.Println("decode url : " + decode)
}