package prototype

type clonable interface {
	clone()
}

type file struct {
	content string
}

func (f *file) clone() *file {
	return &file{f.content}
}

type folder struct {
	files []*file
}

func (f *folder) clone() *folder {
	nf := new(folder)
	for _, f := range f.files {
		nf.files = append(nf.files, f.clone())
	}
	return nf
}
