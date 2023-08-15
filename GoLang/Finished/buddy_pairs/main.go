package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(Buddy(8983, 13355))
}

func Buddy(start, limit int) []int {
	possibles := make(map[int]int)
	for i := start; i <= limit; i++ {
		possibles[getSum(i)] = i
	}

	for j := start + 1; j <= limit; j++ {
		if n, ok := possibles[getSum(j)-1]; ok && n != j {
			res := []int{j, n}
			sort.Ints(res)
			return res
		}
	}

	return nil
}

func getSum(n int) (sum int) {
	if n == 1 {
		return 0
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if i == (n / i) {
				sum += i
			} else {
				sum += (i + n/i)
			}
		}
	}
	return sum + 1
}
