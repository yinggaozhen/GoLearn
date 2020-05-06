package array

import (
	"context"
	"fmt"
	"testing"
)

func TestValueContext(t *testing.T) {
	ctx1 := context.WithValue(context.Background(), "hello", "world")
	ctx2 := context.WithValue(ctx1, "nihao", 2)

	fmt.Println(ctx2.Value("hello"), ctx2.Value("nihao"), ctx2.Value("???"))

	ctx2 = context.WithValue(ctx2, "hello", "shijie")
	fmt.Println(ctx1.Value("hello"), ctx2.Value("hello"))

	ctx1 = context.WithValue(ctx1, "hello", "world")
	fmt.Println(ctx1.Value("hello"), ctx2.Value("hello"))
}