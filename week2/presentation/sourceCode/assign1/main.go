package main

import (
	"assign1/calculator"
	numop "assign1/calculator/operator"
	strop "assign1/concat/operator"
	"fmt"
)

func main() {
	calculator.Calc()

	s := strop.Plus("concat", "string")
	n := numop.Add(1, 2)

	fmt.Printf("result of assign1/concat/operator.Plus() : %s\n", s)
	fmt.Printf("result of assign1/calculator/operator.Add() : %d\n", n)
}
