package calculator

import (
	"fmt"
)

func Sub(a int, b int) int {
	return a - b
}

func Mult(a int, b int) int {
	return a * b
}

func Div(a int, b int) int {
	return a / b
}

func Calc() {
	var a, b int
	var op string
	fmt.Println("숫자 2개를 차례로 입력하세요: ")
	fmt.Scan(&a, &b)
	fmt.Println("연산자를 입력하세요: ")
	fmt.Scan(&op)

	var res int

	if op == "+" {
		res = Add(a, b)
	} else if op == "-" {
		res = Sub(a, b)
	} else if op == "*" {
		res = Mult(a, b)
	} else if op == "/" {
		res = Div(a, b)
	} else {
		fmt.Println("잘못된 연산자입니다.")
	}

	fmt.Println(a, op, b, " = ", res)
}
