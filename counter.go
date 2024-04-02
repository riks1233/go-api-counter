package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	count int64
	mut   *sync.Mutex
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{
		count: 0,
		mut:   &sync.Mutex{},
	}
}

func (c *SafeCounter) Increment() {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.count++
}

func (c *SafeCounter) Decrement() {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.count--
}

func (c *SafeCounter) StringCount() string {
	c.mut.Lock()
	defer c.mut.Unlock()
	return fmt.Sprintf("%v", c.count)
}
