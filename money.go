package money

type Doller struct {
	amount int
}

func (d *Doller) Times(amount int) {
	d.amount *= 2
}
