package unittest

// 执行命令 go test -v github.com/yinggaozhen/GoLearn/mytest/testing/unittest
import (
	"testing"
)

// go test -v github.com/yinggaozhen/GoLearn/mytest/testing/unittest -test.run ^TestUserIsJack$
func TestUserIsJack(t *testing.T) {
	u := User{
		Name: "Jack",
	}

	if u.GetName() != "Jack" {
		t.Error("名字应该叫 Jack")
	}
}

// go test -v github.com/yinggaozhen/GoLearn/mytest/testing/unittest -test.run ^TestUserIsBoy$
func TestUserIsBoy(t *testing.T) {
	u := User{
		Name: "Lisa",
		Sex: 2,
	}

	if u.IsBoy() {
		t.Error("Lisa 是女的呀")
	}

	t.Fail()
}