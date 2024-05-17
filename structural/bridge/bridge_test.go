package bridge

import "testing"

func TestBridge(t *testing.T) {
	hp := &com{}
	hp.setPrinter(&oldPrinter{})
	epson := &com{}
	epson.setPrinter(&newPrinter{})

	hp.print()
	epson.print()
}
