package fluentbuilder

type SingleBuilder struct {
	h *House
}

func GetSingleBuilder() *SingleBuilder {
	return &SingleBuilder{&House{}}
}

func (b *SingleBuilder) Build() *House {
	return b.h
}

func (b *SingleBuilder) WithWall(s string) *SingleBuilder {
	b.h.Wall = s
	return b
}

func (b *SingleBuilder) WithRooms(n int) *SingleBuilder {
	b.h.Rooms = n
	return b
}

func (b *SingleBuilder) WithWindows(n int) *SingleBuilder {
	b.h.Windows = n
	return b
}

func (b *SingleBuilder) WithGarden(s string) *SingleBuilder {
	b.h.Garden = s
	return b
}

func (b *SingleBuilder) WithTerrace() *SingleBuilder {
	b.h.Terrace = true
	return b
}
