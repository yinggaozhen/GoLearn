package os

import (
	"fmt"
	"os"
	"testing"
)

func TestOsEnv(t *testing.T) {
	fmt.Println(os.Environ())
	fmt.Println(os.Getenv("PATH"))
}