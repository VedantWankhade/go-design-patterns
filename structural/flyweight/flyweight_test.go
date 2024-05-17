package flyweight

import (
	"math/rand"
	"testing"
)

func TestFlyweight(t *testing.T) {
	forest := newforest()

	treeTypes := []struct {
		name, color, texture string
	}{
		{"Oak", "Green", "Rough"},
		{"Pine", "Dark Green", "Smooth"},
	}

	for i := 0; i < 10; i++ {
		t := treeTypes[rand.Intn(len(treeTypes))]
		forest.plantTree(rand.Intn(100), rand.Intn(100), t.name, t.color, t.texture)
	}

	forest.Draw()
}
