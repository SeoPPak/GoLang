package main

type Snack struct {
	Name  string
	Price int
}

type Beverage struct {
	Name  string
	Price int
}

func (s *Snack) SaleAndGetPrice() int {
	discount := 0.9

	return int(float64(s.Price) * discount)
}

func (b *Beverage) SaleAndGetPrice() int {
	discount := 0.8

	return int(float64(b.Price) * discount)
}
