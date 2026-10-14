package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {

	count := SafeCounter{v: make(map[string]int)}

	for range 1000 {
		go count.Inc("somekey")

	}

	time.Sleep(time.Second)
	fmt.Println(count.Value("somekey"))

}
