package facade

import (
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	homeTheater := NewHomeTheaterFacade()
	homeTheater.WatchMovie()
	fmt.Println("Enjoy the movie!")
	homeTheater.EndMovie()
}
