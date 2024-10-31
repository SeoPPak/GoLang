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

	plus := make(chan int)
	mult := make(chan int)
	finish := make(chan int)

	//num1Chan := make(chan int, 3)
	//num2Chan := make(chan int, 3)

	outputChan := make(chan [2]int, 10)
	//plusOutput := make(chan int)
	//multOutput := make(chan int)

	go func(num1, num2 int, plus, mult, finish chan int, outputChan chan [2]int) {
		for {
			select {
			case <-plus:
				var add [2]int
				add[0] = 0
				add[1] = num1 + num2
				outputChan <- add
			case <-mult:
				var mul [2]int
				mul[0] = 1
				mul[1] = num1 * num2
				outputChan <- mul
			case <-finish:
				close(outputChan)
				return
			}
		}
	}(num1, num2, plus, mult, finish, outputChan)

	//for i := 0; i < 3; i++ {
	//	num1Chan <- num1
	//	num2Chan <- num2
	//}

	plus <- 1

	mult <- 2

	for i := 0; i < 2; i++ {
		res := <-outputChan
		if res[0] == 0 {
			fmt.Println("덧셈 결과는 : ", res[1])
		} else {
			fmt.Println("곱셈 결과는 : ", res[1])
		}
	}

	finish <- 3
	close(finish)
	close(outputChan)
	//close(num1Chan)
	//close(num2Chan)
}
