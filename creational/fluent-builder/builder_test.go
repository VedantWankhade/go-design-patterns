package fluentbuilder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	builder1 := Builder{}
	house1 := builder1.WithWall("wood").WithRooms(4).WithGarden("big").Build()
	fmt.Println(house1)
	builder2 := Builder{}
	house2 := builder2.WithTerrace().WithWall("paper").Build()
	fmt.Println(house2)
}
