package bridge

import "fmt"

type printer interface {
	printFile()
}

type (
	oldPrinter struct{}
	newPrinter struct{}
)

func (p *oldPrinter) printFile() {
	fmt.Println("OLD printer printing")
}

func (p *newPrinter) printFile() {
	fmt.Println("NEW printer printing")
}

type com struct {
	p printer
}

func (c *com) setPrinter(p printer) {
	c.p = p
}

func (c *com) print() {
	c.p.printFile()
}
