package main

import (
	"fmt"

	lab "github.com/SeoPPak/GoLang/week2/Lab/lab05/calc"
)

func main() {
	var num1 int
	var num2 int
	var op string

	fmt.Println("Calculator")
	fmt.Println("첫 번째 숫자 입력: ")
	fmt.Scan(&num1)
	fmt.Println("두 번째 숫자 입력: ")
	fmt.Scan(&num2)
	fmt.Println("연산자 입력(+, -, *, /) : ")
	fmt.Scan(&op)

	lab.Calc(num1, num2, op)
}
