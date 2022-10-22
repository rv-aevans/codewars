package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(PrimeBefAfter(97))
}

func PrimeBefAfter(n int) [2]int {
	b, a := 0, 0
	for i := 1; i > 0; i++ {
		if big.NewInt(int64(n - i)).ProbablyPrime(0) {
			b = n - i
			break
		}
	}
	for i := 1; i > 0; i++ {
		if big.NewInt(int64(n + i)).ProbablyPrime(0) {
			a = n + i
			break
		}
	}
	return [2]int{b, a}
}
