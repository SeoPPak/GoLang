package calculator

import (
	"assign1/calculator/operator"
	"fmt"
)

func Calc() int {
	fmt.Printf("assign1/calculator.Calc() excuted\n")
	defer fmt.Printf("assign1/calculator.Calc() finished\n\n\n")
	var a, b int
	var op string

	fmt.Printf("Enter two numbers: ")
	fmt.Scan(&a, &b)

	fmt.Printf("Enter the operation: ")
	fmt.Scan(&op)

	var res int

	switch op {
	case "+":
		res = operator.Add(a, b)
	case "-":
		res = operator.Sub(a, b)
	case "*":
		res = operator.Mult(a, b)
	case "/":
		res = operator.Div(a, b)
	default:
		fmt.Println("Invalid operation")
	}

	fmt.Printf("result of Clac() is\n%d %s %d = %d\n", a, op, b, res)
	return res
}
