package unittest

// 执行命令 go test -v github.com/yinggaozhen/GoLearn/mytest/testing/coverage
import (
	"testing"
)

func TestUser(t *testing.T) {
	u := User{
		Name: "Jack",
	}

	if u.GetName() != "Jack" {
		t.Error("名字应该叫 Jack")
	}
}