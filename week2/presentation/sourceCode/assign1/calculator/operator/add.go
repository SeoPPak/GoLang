package operator

import "fmt"

func Add(a int, b int) int {
	fmt.Printf("assign1/calculator/operator.Add() excuted\n")
	defer fmt.Printf("assign1/calculator/operator.Add() fineshed\n\n\n")
	return a + b
}
