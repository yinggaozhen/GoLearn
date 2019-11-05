package now

import (
	"fmt"
	"testing"
)

var n = Now{}

func TestNowBeginningOfMinute(t *testing.T) {
	fmt.Println(n.beginningOfMinute())
}

func TestNowBeginningOfHour(t *testing.T) {
	fmt.Println(n.BeginningOfHour())
}

func TestNowBeginningOfDay(t *testing.T) {
	fmt.Println(n.BeginningOfDay())
}

func TestNowBeginningOfWeekday(t *testing.T) {
	fmt.Println(n.BeginningOfWeek())
}

func TestNowBeginningOfMonth(t *testing.T) {
	fmt.Println(n.BeginningOfMonth())
}
