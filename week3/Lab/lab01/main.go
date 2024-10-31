package main

import (
	"fmt"
)

func main() {
	var num1, num2 int

	fmt.Printf("첫 번째 정수를 입력하세요 : ")
	fmt.Scan(&num1)
	fmt.Printf("두 번째 정수를 입력하세요 : ")
	fmt.Scan(&num2)

	outputPlus := make(chan int)
	outputMult := make(chan int)
	go Plus(num1, num2, outputPlus)
	go Mult(num1, num2, outputMult)

	fmt.Printf("덧셈 결과는 : %d\n", <-outputPlus)
	fmt.Printf("곱셈 결과는 : %d\n", <-outputMult)
}
