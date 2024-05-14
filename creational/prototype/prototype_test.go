package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	fileProto := file{"hello"}
	folderProto := folder{[]*file{{"yes"}, {"no"}}}

	f1 := fileProto.clone()
	f2 := f1.clone()

	fo1 := folderProto.clone()
	fo2 := fo1.clone()

	fmt.Println(f1 == f2)
	fmt.Println(f1 == &fileProto)
	fmt.Println(fo1 == fo2)
	fmt.Println(fo1 == &folderProto)

	fmt.Printf("%+v\n%+v", *f1, *fo1)
}
