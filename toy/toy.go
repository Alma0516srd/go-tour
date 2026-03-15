package toy

type Toy struct {
	Name   string
	Weight float64
	onHand bool
	sold   bool
}

func New(name string, weight float64) *Toy {
	return &Toy{
		Name:   name,
		Weight: weight,
	}
}

func (toy *Toy) OnHand(onhand bool) bool {
	toy.onHand = onhand
	return toy.onHand
}

func (toy *Toy) Sold(newSold bool) bool {
	toy.sold = newSold
	return toy.sold
}
