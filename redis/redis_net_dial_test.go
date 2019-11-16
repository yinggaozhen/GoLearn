package redis

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"
)

type EmptyContext struct {
	context.Context
}

func TestDialRedis(t *testing.T) {
	netDialer := &net.Dialer{
		Timeout:   3 * time.Second,
		KeepAlive: 5 * time.Minute,
	}

	conn, e := netDialer.Dial("tcp", "127.0.0.1:6379")
	if e != nil {
		fmt.Println(e)
		return
	}

	cmd := []byte("get nihao\n")
	n, e := conn.Write(cmd)
	fmt.Println(n, e)

	var result []byte
	n2, e := conn.Read(result)
	fmt.Println(n2, e)

	fmt.Println(result)
}