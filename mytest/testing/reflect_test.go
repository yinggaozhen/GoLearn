package mytest

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Name string
}

func (u User) GetName() string {
	return u.Name
}

// 1. 测试获取 Method
func TestReflectMethod(t *testing.T) {
	u := User{
		Name: "guozhen",
	}

	rt := reflect.TypeOf(u)
	for i := 0; i < rt.NumMethod(); i++ { //这里同样通过t.NumMethod来获取它拥有的方法的数量，来决定循环的次数
		m := rt.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)
	}
}

// 2. 测试获取成员变量
func TestReflectField(t *testing.T) {
	u := User{
		Name: "guozhen",
	}

	rt := reflect.TypeOf(u)
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		fmt.Printf("%6s:%v\n", f.Name, f.Type)
	}
}