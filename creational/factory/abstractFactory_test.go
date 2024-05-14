package factory

import "testing"

func TestAbstractFactory(testing *testing.T) {
	mf := getFactory(typeClassic)
	cf := getFactory(typeModern)
	mc := mf.getChair()
	ms := mf.getSofa()
	cc := cf.getChair()
	cs := cf.getSofa()

	mc.do()
	ms.do()
	cc.do()
	cs.do()
}
