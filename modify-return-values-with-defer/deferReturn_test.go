package modifyreturnvalueswithdefer

import (
	"fmt"
	"testing"
)

func TestDeferReturn(t *testing.T) {
	fmt.Println(addTwo(3))
	fmt.Println(addTwoDefer(3))
}
