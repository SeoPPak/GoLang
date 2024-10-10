package operator

import "fmt"

func Div(a int, b int) int {
	fmt.Printf("assign1/calculator/operator.Div() excuted\n")
	defer fmt.Printf("assign1/calculator/operator.Div() finished\n\n\n")
	return a / b
}
