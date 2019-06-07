package test

import (
	"fmt"
	"strconv"
	"testing"
)

const (
	JobStudent = "student"
	JobTeacher = "teacher"
)

type People struct {
	name string
	age int
	job string
}

func (p *People) Construct(name string, age int, job string) *People {
	p.name = name
	p.age = age
	p.job = job

	return p
}

func (p *People) String() string {
	age := strconv.Itoa(p.age)
	return "My Name is : " + p.name + ". My Age is " + age + ". My Job is " + p.job
}

func (p *People) Hello() {
	fmt.Println("Hello " + p.name)
}

// student class
type StructStudent struct {
	People
}
func (p *StructStudent) Construct(name string, age int) *StructStudent {
	p.People.Construct(name, age, JobStudent)
	return p
}

// teacher class
type StructTeacher struct {
	People
}
func (p *StructTeacher) Construct(name string, age int) *StructTeacher {
	p.People.Construct(name, age, JobTeacher)

	return p
}

// test struct extends
func TestStructExtends(t *testing.T) {
	tom := StructStudent{}
	tom.Construct("tom", 18)
	fmt.Println(tom.String())
	tom.Hello()

	lisa := StructTeacher{}
	lisa.Construct("lisa", 32)
	fmt.Println(lisa.String())
	lisa.Hello()
}
