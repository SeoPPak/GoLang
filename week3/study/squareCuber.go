package main

import (
	"fmt"
)

func squarerCuber(sqInChan, sqOutChan, cuInChan, cuOutChan, exitChan chan int) {
	var squareX int
	var cubeX int
	for {
		select {
		case squareX = <-sqInChan:
			sqOutChan <- squareX * squareX
		case cubeX = <-cuInChan:
			cuOutChan <- cubeX * cubeX * cubeX
		case <-exitChan:
			return
		}
	}
}

func main() {
	/*
		sqInChan := make(chan int, 10)
		cuInChan := make(chan int, 10)
		sqOutChan := make(chan int, 10)
		cuOutChan := make(chan int, 10)
		exitChan := make(chan int)
	*/

	inputChan := make(chan int, 10)
	finishChan := make(chan int)
	outputChan := make(chan int, 10)

	//go squarerCuber(sqInChan, sqOutChan, cuInChan, cuOutChan, exitChan)
	go func(inputChan, finishChan chan int) {
		for {
			select {
			case x := <-inputChan:
				outputChan <- x * x
			case _ = <-finishChan:
				return
			}
		}
	}(inputChan, finishChan)
	for i := 0; i < 10; i++ {
		inputChan <- i
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-outputChan)
	}
	finishChan <- 1
}
