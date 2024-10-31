package main

import (
	"assign2/primechecker"

	"assign1/calculator"
)

func main() {
	res := calculator.Calc()
	primechecker.Check(int64(res))
}
