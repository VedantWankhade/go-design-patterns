package decorator

type iPizza interface {
	getPrice() int
}

type classicPizza struct{}

func (p *classicPizza) getPrice() int {
	return 10
}

type tommatoTopping struct {
	pizza iPizza
}

func (p *tommatoTopping) getPrice() int {
	return p.pizza.getPrice() + 5
}

type chesseTopping struct {
	pizza iPizza
}

func (p *chesseTopping) getPrice() int {
	return p.pizza.getPrice() + 15
}
