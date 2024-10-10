package operator

import "fmt"

func Sub(a int, b int) int {
	fmt.Printf("assign1/calculator/operator.Sub() excuted\n")
	defer fmt.Printf("assign1/calculator/operator.Sub() fineshed\n\n\n")
	return a - b
}
