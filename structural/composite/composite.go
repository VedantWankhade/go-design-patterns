package composite

import "fmt"

type searchable interface {
	search(string)
}

type folder struct {
	name  string
	items []searchable
}

func (f *folder) search(s string) {
	fmt.Printf("seraching folder %s recursively for keyword %s\n", f.name, s)
	for _, i := range f.items {
		i.search(s)
	}
}

func (f *folder) add(s searchable) {
	f.items = append(f.items, s)
}

type file struct {
	name string
}

func (f *file) search(s string) {
	fmt.Printf("searching file %s for keyword %s\n", f.name, s)
}
