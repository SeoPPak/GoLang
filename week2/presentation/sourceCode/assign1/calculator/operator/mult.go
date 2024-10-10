package operator

import "fmt"

func Mult(a int, b int) int {
	fmt.Printf("assign1/calculator/operator.Mult() excuted\n")
	defer fmt.Printf("assign1/calculator/operator.Mult() finished\n\n\n")
	return a * b
}
