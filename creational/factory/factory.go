package factory

import "fmt"

type paymentMethod interface {
	pay(float64)
}

type paymentMethodType int

const (
	paymentTypePaypal paymentMethodType = iota
	paymentTypeBilldesk
)

type paypal struct{}

func (p *paypal) pay(amount float64) {
	fmt.Println("payed thru paypal")
}

type billdesk struct{}

func (b *billdesk) pay(amount float64) {
	fmt.Println("payed thru billdesk")
}

func getPaymentMethodFactory(pt paymentMethodType) paymentMethod {
	switch pt {
	case paymentTypeBilldesk:
		return &billdesk{}
	case paymentTypePaypal:
		return &paypal{}
	}
	return nil
}
