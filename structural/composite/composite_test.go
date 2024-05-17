package composite

import "testing"

func TestComposite(t *testing.T) {
	fldr := folder{name: "home"}
	userdir := &folder{name: "username"}
	userdir.add(&file{name: "zshrc"})
	userdir.add(&file{name: "bashrc"})
	fldr.add(userdir)
	fldr.add(&folder{name: "local"})

	fldr.search("something")
}
