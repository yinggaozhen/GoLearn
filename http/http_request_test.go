package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// test http get
func TestGet(t *testing.T) {
	r, _ := http.Get("http://www.baidu.com")
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(b))
}

// test http get
func TestPost(t *testing.T) {
	r, _ := http.Post("http://www.baidu.com", "application/x-www-form-urlencoded", strings.NewReader("hello"))
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(b))
}
