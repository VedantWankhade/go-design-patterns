package fluentbuilder

import (
	"fmt"
	"testing"
)

func TestSingleBuilder(t *testing.T) {
	builder := GetSingleBuilder()
	house1 := builder.WithWall("wood").WithRooms(4).WithGarden("big").Build()
	fmt.Println(house1)
	house2 := builder.WithTerrace().WithWall("paper").Build()
	fmt.Println(house2)
}
