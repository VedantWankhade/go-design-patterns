package fluentbuilder

type House struct {
	Wall    string
	Garden  string
	Rooms   int
	Windows int
	Terrace bool
}

type Builder struct {
	h House
}

func (b *Builder) Build() *House {
	return &b.h
}

func (b *Builder) WithWall(s string) *Builder {
	b.h.Wall = s
	return b
}

func (b *Builder) WithRooms(n int) *Builder {
	b.h.Rooms = n
	return b
}

func (b *Builder) WithWindows(n int) *Builder {
	b.h.Windows = n
	return b
}

func (b *Builder) WithGarden(s string) *Builder {
	b.h.Garden = s
	return b
}

func (b *Builder) WithTerrace() *Builder {
	b.h.Terrace = true
	return b
}
