package test

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func runtimeFuncNameGuozhen() {
}

func nameOfFunction(f interface{}) {
	fmt.Println(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
}

// test struct extends
func TestRuntimeFuncName(t *testing.T) {
	nameOfFunction(runtimeFuncNameGuozhen)
}
