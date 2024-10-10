package main

import (
	"fmt"
)

func main() {
	chips := Snack{"Pringles", 4000}
	crackers := Snack{"Ace", 2500}

	soda := Beverage{"Sprite", 1800}
	coffee := Beverage{"TOP", 2700}

	var total int = 0
	total += chips.SaleAndGetPrice()
	total += crackers.SaleAndGetPrice()
	total += soda.SaleAndGetPrice()
	total += coffee.SaleAndGetPrice()

	fmt.Println(total)
}
