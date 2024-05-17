package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	clscPzz := &classicPizza{}
	fmt.Println(clscPzz.getPrice())

	clscPzzWTmt := &tommatoTopping{clscPzz}
	fmt.Println(clscPzzWTmt.getPrice())

	clscPzzWTmtAndChz := &chesseTopping{clscPzzWTmt}
	fmt.Println(clscPzzWTmtAndChz.getPrice())
}
