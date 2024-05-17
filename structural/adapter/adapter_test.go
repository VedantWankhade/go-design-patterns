package adapter

import "testing"

func TestAdapter(t *testing.T) {
	m := mac{}
	p := pc{}
	useLightningPort(&m)
	// useLightningPort(&p) // &p doesnt have lightning method
	pAdapter := &pcLightningAdapter{&p}
	useLightningPort(pAdapter)
}
