package singleton

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	c := GetCounter()
	fmt.Println(c == nil)
	fmt.Println(c)
	c.AddOne()
	fmt.Println(c)
	c1 := GetCounter()
	fmt.Println(c1 == nil)
	fmt.Println(c1)
	c1.AddOne()
	fmt.Println(c1)
}
