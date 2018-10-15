package money

type Doller struct {
	amount int
}

func (d *Doller) Times(amount int) (product Doller) {
	product = Doller{d.amount * amount}
	return 
}
