package array

import (
	"fmt"
	"testing"
)

// ================ 1. 测试普通回调
type StringCallback func(content string) bool
func TestStringCallback(t *testing.T) {
	callback := func(content string) bool {
		return len(content) > 0
	}

	invokeStringCallback(callback, "hello")
	invokeStringCallback(callback, "")
}
func invokeStringCallback(callback StringCallback, content string) {
	fmt.Println("string callback: ", callback(content))
}

// ================ 2.测试结构体成员方法
type Student struct {
	name string
	age int
	scope *Student
}
type StudentNameCallback func() string

func TestMethodCallback(t *testing.T) {
	tom := Student{}
	tom.construct("tom", 18)

	invokeMethodCallback(tom.getName)
}

func invokeMethodCallback(callback StudentNameCallback) {
	result := callback()
	fmt.Println("method callback: ", result)
}

func (student *Student) getName() string {
	return student.name
}

func (student *Student) construct(name string, age int) {
	student.name = name
	student.age = age
}