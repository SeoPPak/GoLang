package concat

import (
	"assign1/concat/operator"
	"fmt"
)

func Concat(a string, b string) string {
	fmt.Printf("assign1/concat.Concat() excuted\n")
	defer fmt.Printf("assign1/concat.Concat() finished\n\n\n")
	return operator.Plus(a, b)
}
