package singletonsafe

import (
	"fmt"
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	threads := sync.WaitGroup{}
	threads.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			c := GetCounter()
			c.AddOne()
			fmt.Println("in thread", i, c)
			threads.Done()
		}(i)
	}
	threads.Wait()
}
