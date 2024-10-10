package operator

import "fmt"

func Plus(a string, b string) string {
	fmt.Printf("assign1/concat.Concat() excuted\n")
	defer fmt.Printf("assign1/concat.Concat() finished\n\n\n")
	res := a + b

	return res
}
