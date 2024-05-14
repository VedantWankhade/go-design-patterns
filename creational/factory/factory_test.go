package factory

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	m1 := getPaymentMethodFactory(paymentTypePaypal)
	m2 := getPaymentMethodFactory(paymentTypeBilldesk)
	m1.pay(2)
	m2.pay(2)
	fmt.Printf("%T\n", m1)
	fmt.Printf("%T\n", m2)
}
