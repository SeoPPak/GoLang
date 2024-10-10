package primechecker

import (
	"fmt"

	"github.com/otiai10/primes"
)

func Check(a int64) {
	fmt.Printf("assign2/primechecker.Check() excuted\n")
	defer fmt.Printf("assign2/primechecker.Check() finished\n\n\n")
	f := primes.Factorize(int64(a))
	fmt.Println("primes:", len(f.Powers()) == 1)
}
