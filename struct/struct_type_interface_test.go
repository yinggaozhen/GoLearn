package test

import (
	"testing"
)

type IRoute interface {
	GET() IRoute
	POST() IRoute
}

type Router struct {
}

func (r Router) GET() IRoute {
	panic("implement me")
}

func (r Router) POST() IRoute {
	panic("implement me")
}

var _ IRoute = &Router{}

func TestTypeInterface(t *testing.T) {

}