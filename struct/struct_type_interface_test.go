package test

import (
	"fmt"
	"testing"
)

type IRoute interface {
	GET()
	POST()
}

type Router struct {
}

var _ IRoute = Router{}

func (r Router) GET() {
	fmt.Println("GET implement me")
}

func (r Router) POST() {
	fmt.Println("POST implement me")
}

func TestTypeInterface(t *testing.T) {
	router := &Router{}
	router.GET()
}