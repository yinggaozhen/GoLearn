package main

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
	"time"
)

type ConnPool struct {
	poolSize int
	idleConnsLen int

	optPoolSize int
	optMinIdleConns int

	connsMu      sync.Mutex
}

var p ConnPool

func main() {
	p = ConnPool{
		optPoolSize: 10,
		optMinIdleConns: 10,
	}

	checkMinIdleConns()

	time.Sleep(2 * time.Second)
}

func checkMinIdleConns() {
	for p.poolSize < p.optPoolSize && p.idleConnsLen < p.optMinIdleConns {
		p.poolSize++
		p.idleConnsLen++
		go func(size int) {
			err := addIdleConn(size)
			if err != nil {
				fmt.Println("dial failed")
				p.connsMu.Lock()
				p.poolSize--
				p.idleConnsLen--
				p.connsMu.Unlock()
				fmt.Printf("dial process %d\n", p.poolSize)
			}
		}(p.poolSize)
	}
}

func addIdleConn(size int) error {
	// ... dial

	// Simulation dial failed
		return errors.New("dial failed")

}