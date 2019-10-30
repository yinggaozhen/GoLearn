// +build args_build

package args

import (
	"flag"
	"fmt"
	"testing"
)

// cd ./mytest/args
// go test --tags args_build --args -who jack -age 18

func TestMain(m *testing.M) {
	var who string
	var age int
	var air string

	flag.StringVar(&who, "who", "defaultName", "Usage : whoami")
	flag.IntVar(&age, "age", -1, "Usage : how old are you")
	flag.StringVar(&air, "air", "defaultAir", "Usage : non-exist args")

	flag.Parse()

	fmt.Println(who) // jack
	fmt.Println(age) // 18
	fmt.Println(air) // defaultAir
}