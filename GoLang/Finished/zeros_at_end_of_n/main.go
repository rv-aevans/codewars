package main

import (
	"math/big"
)

func main() {
	CountZeros(500)
}

func CountZeros(n int) int {
	total := big.NewInt(1)
	for i := n; i > 0; i -= 2 {
		total.Mul(total, big.NewInt(int64(i)))
	}
	s := total.String()
	count := 0
	for i := range s {
		if isZero(string(s[len(s)-1-i])) {
			count++
		} else {
			break
		}
	}
	return count
}

func isZero(s string) bool {
	return s == "0"
}
