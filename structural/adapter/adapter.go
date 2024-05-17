package adapter

import "fmt"

type lightningPort interface {
	attachLightningPort()
}

func useLightningPort(com lightningPort) {
	com.attachLightningPort()
}

type mac struct{}

func (m *mac) attachLightningPort() {
	fmt.Println("mac attaching lightning port")
}

type pc struct{}

func (p *pc) attachUSBPort() {
	fmt.Println("pc attaching usb port")
}

type pcLightningAdapter struct {
	p *pc
}

func (p *pcLightningAdapter) attachLightningPort() {
	fmt.Println("using pc lightning adapter")
	p.p.attachUSBPort()
}
