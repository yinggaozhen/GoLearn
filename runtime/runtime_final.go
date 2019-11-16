package test

import (
	"fmt"
	"runtime"
	"time"
)

type wrapper struct {
	*Cache
}

type Cache struct {
	value map[string]string
	timer *time.Ticker
	stop chan bool
}

func New() *wrapper {
	c := &Cache{
		value : make(map[string]string),
		stop  : make(chan bool),
	}

	w := &wrapper{
		c,
	}

	go w.Cache.run()

	return w
}

func (c *Cache) Set(key string, value string) {
	c.value[key] = value
}

func (c *Cache) run() {
	c.timer = time.NewTicker(time.Second)
	runtime.SetFinalizer(c, func(c2 *Cache) {
		fmt.Println("SetFinalizer")
		c2.stop <- true
	})

	for {
		select {
		case <-c.timer.C:
			fmt.Println("hello")
		case <-c.stop:
			fmt.Println("stop chan")
		}
	}
}

func main() {
	for {
		w := New()
		w.Cache.Set("hello", "world")
		w = nil
		time.Sleep(time.Second)
	}

}