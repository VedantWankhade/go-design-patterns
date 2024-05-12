package singleton

import "fmt"

type counter int

var c *counter

func GetCounter() *counter {
	if c == nil {
		c = new(counter)
		return c
	}
	return c
}

func (c *counter) AddOne() {
	*c++
}

func (c *counter) String() string {
	return fmt.Sprintf("count: %d", *c)
}
