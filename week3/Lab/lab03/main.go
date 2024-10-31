package main

//#cgo LDFLAGS: -L. -lcfibonacci
//int fibonacci(int n);
import "C"
import "fmt"

func main() {
	fmt.Println(C.fibonacci(C.int(10)))
}
