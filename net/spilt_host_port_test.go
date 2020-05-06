package nettest

import (
	"fmt"
	"net"
	"testing"
)

func TestSpiltHostPort(t *testing.T) {
	fmt.Println(net.SplitHostPort("172.24.24.14"))
	fmt.Println(net.SplitHostPort("172.24.24.14:8080"))
	fmt.Println(net.SplitHostPort("172.24.24.14:"))
}