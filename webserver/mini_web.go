package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
}

func main() {
	server := new(http.Server)
	server.Addr = ":8080"

	c := MyHandler{}
	server.Handler = c
	error := server.ListenAndServe()

	fmt.Println(error)
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 处理每次 http 请求
	// 在这里可根据 URL 进行路由
	w.Write([]byte("hello aaaa "))
	w.Write([]byte(r.URL.String()))
}