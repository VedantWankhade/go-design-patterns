package factory

import "fmt"

type chair interface {
	do()
}

type modernChair struct{}

func (c *modernChair) do() {
	fmt.Println("modernChair")
}

type classicChair struct{}

func (c *classicChair) do() {
	fmt.Println("classicChair")
}

type sofa interface {
	do()
}

type modernSofa struct{}

func (c *modernSofa) do() {
	fmt.Println("modernSofa")
}

type classicSofa struct{}

func (c *classicSofa) do() {
	fmt.Println("classicSofa")
}

type furnitureFactory interface {
	getSofa() sofa
	getChair() chair
}

type modernFactory struct{}

func (f *modernFactory) getSofa() sofa {
	return &modernSofa{}
}

func (f *modernFactory) getChair() chair {
	return &modernChair{}
}

type classicFactory struct{}

func (f *classicFactory) getSofa() sofa {
	return &classicSofa{}
}

func (f *classicFactory) getChair() chair {
	return &classicChair{}
}

type factoryType int

const (
	typeModern factoryType = iota
	typeClassic
)

func getFactory(t factoryType) furnitureFactory {
	switch t {
	case typeModern:
		return &modernFactory{}
	case typeClassic:
		return &classicFactory{}
	}
	return nil
}
