package calc

import (
	"fmt"
)

func Calc(num1 int, num2 int, op string) {
	switch op {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Print(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("0으로 나눌 수 없습니다.")
		} else {
			fmt.Print(float32(num1) / float32(num2))
		}
	default:
		fmt.Println("잘못된 연산자")
	}
}
