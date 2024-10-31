package main

import (
	"assign2/primechecker"

	//로컬에서 작성한 또 다른 패키지를 임포트
	//"assign1/calculator"
)

func main() {
	res := calculator.Calc()
	primechecker.Check(int64(res))
}
