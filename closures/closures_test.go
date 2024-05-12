package closures

import (
	"fmt"
	"testing"
)

func TestClosure(t *testing.T) {
	c1 := addN(1)
	c2 := addN(2)
	c3 := addN(1)
	c4 := addN(2)
	fmt.Println(c1(2), c2(2))
	fmt.Println(c1(2), c2(2))
	fmt.Println(c3(2), c4(2))
	fmt.Println(c3(2), c4(2))
}
