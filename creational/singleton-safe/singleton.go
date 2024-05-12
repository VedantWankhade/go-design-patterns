package singletonsafe

import (
	"fmt"
	"sync"
)

type counter int

var (
	mut sync.Mutex
	c   *counter
)

func GetCounter() *counter {
	mut.Lock()
	defer mut.Unlock()
	if c == nil {
		c = new(counter)
		return c
	}
	return c
}

func (c *counter) AddOne() {
	mut.Lock()
	defer mut.Unlock()
	*c++
}

func (c *counter) String() string {
	mut.Lock()
	defer mut.Unlock()
	return fmt.Sprintf("count: %d", *c)
}
