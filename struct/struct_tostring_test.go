package test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

const (
	Man = 1
	Woman = 2
)

type Student struct {
	name string
	age int
	extra interface{}
}

// test struct extends
func TestInterface(t *testing.T) {
	jack := Student{
		name : "jack",
	}
	fmt.Println(jack)
	fmt.Println(jack.isGirl())

	tom := Student{
		name : "tom",
		age: 18,
	}
	fmt.Println(tom)
	fmt.Println(tom.isGirl())

	lisa := Student{
		name : "lisa",
		age: 18,
		extra: Woman,
	}
	fmt.Println(lisa)
	fmt.Println(lisa.isGirl())
}

func (student Student) String() string {
	extra := student.extra

	var extraType string
	if extra != nil {
		extraType = reflect.TypeOf(extra).String()
	}

	switch extraType {
		case "int":
			sex := "boy"

			if extra.(int) == Woman {
				sex = "girl"
			}

			student.extra = sex
		case "string":
			student.extra = "sex : " + student.extra.(string)
		default:
			student.extra = "default"
	}

	name := student.name
	age := strconv.Itoa(student.age)

	return "name : " + name + " age : " + age + " extra : " + student.extra.(string)
}

func (student Student) isGirl() bool {
	extra := student.extra

	var extraType string
	if extra != nil {
		extraType = reflect.TypeOf(extra).String()
	}

	switch extraType {
	case "int":
		sex := "boy"

		if extra.(int) == Woman {
			sex = "girl"
		}

		student.extra = sex
	case "string":
		student.extra = "sex : " + student.extra.(string)
	default:
		student.extra = "default"
	}

	return student.extra == "girl"
}
