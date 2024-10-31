package main

func Plus(num1, num2 int, output chan int) {
	output <- num1 + num2
}

func Mult(num1, num2 int, output chan int) {
	output <- num1 * num2
}
