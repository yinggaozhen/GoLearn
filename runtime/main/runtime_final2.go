package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Cache = *wrapper

var pool sync.Pool
var once sync.Once

type wrapper struct {
	*cache
}

type cache struct {
	content string
	stop    chan struct{}
	onStopped func()
}

func newCache() *cache {
	pool.New = func() interface{} {
		return &cache{
			content: "some thing",
			stop: make(chan struct{}),
		}
	}

	c := pool.Get().(*cache)
	once.Do(func() {
		go c.run()
	})

	return c
}

func NewCache(i int) Cache {
	w := &wrapper{
		cache : newCache(),
	}
	runtime.SetFinalizer(w, stop)
	return w
}

func stop(w *wrapper) {
	fmt.Println("stop")
	w.cache.Stop()
}

func (c *cache) run() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("hello")
		case <-c.stop:
			if c.onStopped != nil {
				c.onStopped()
			}
			return
		}
	}
}

func (c *cache) Stop() {
	c.stop <- struct{}{}
}

func main() {
	var i int
	for {
		NewCache(i)
		time.Sleep(time.Second)
	}
}