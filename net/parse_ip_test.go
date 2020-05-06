package nettest

import (
	"fmt"
	"net"
	"testing"
)

// 解析IP
func TestParseIp(t *testing.T) {
	fmt.Println(net.ParseIP("172.24.24.14"))
	fmt.Println(net.ParseIP("127.0.0.1"))
}