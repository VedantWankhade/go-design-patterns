package flyweight

import "fmt"

// flyweight holding instrinsic values
type treeType struct {
	name, color, texture string
}

func newTreeType(name, color, texture string) *treeType {
	return &treeType{name, color, texture}
}

// flyweight factory
type treeFactory struct {
	treeTypes map[string]*treeType
}

func newTreeFactory() *treeFactory {
	return &treeFactory{treeTypes: make(map[string]*treeType)}
}

func (f *treeFactory) getTreeType(name, color, texture string) *treeType {
	key := name + "_" + color + "_" + texture
	if _, exists := f.treeTypes[key]; !exists {
		f.treeTypes[key] = newTreeType(name, color, texture)
	}
	return f.treeTypes[key]
}

// extrinsic values
type tree struct {
	x, y        int
	treeTypeRef *treeType
}

func newTree(x, y int, treeType *treeType) *tree {
	return &tree{x: x, y: y, treeTypeRef: treeType}
}

func (t *tree) Draw() {
	fmt.Printf("tree of type %s at (%d, %d)\n", t.treeTypeRef.name, t.x, t.y)
}

type forest struct {
	trees   []*tree
	factory *treeFactory
}

func newforest() *forest {
	return &forest{factory: newTreeFactory()}
}

func (f *forest) plantTree(x, y int, name, color, texture string) {
	treeType := f.factory.getTreeType(name, color, texture)
	tree := newTree(x, y, treeType)
	f.trees = append(f.trees, tree)
}

func (f *forest) Draw() {
	for _, tree := range f.trees {
		tree.Draw()
	}
}
