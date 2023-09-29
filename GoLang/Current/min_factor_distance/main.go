package main

import "fmt"

func main() {
	fmt.Println(MinDistance(345817))
}

func MinDistance(n int) int {
	var factors []int
	res := n

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}

	if len(factors) == 2 {
		return n - 1
	}

	for _, n := range factors {
		for j := 1; j < len(factors)-1; j++ {
			if factors[j]-n > 0 && factors[j]-n < res {
				res = factors[j] - n
			}
		}
	}

	return res
}
