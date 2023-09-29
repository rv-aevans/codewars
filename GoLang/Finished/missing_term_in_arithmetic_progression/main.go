package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(solution([]int{-1, -3, -4}))
}

func solution(n []int) int {
	m := math.Min(math.Abs(float64(n[1]-n[0])), math.Abs(float64(n[2]-n[1])))
	for i, v := range n {
		if v < 0 && n[i+1] != v-int(m) {
			return v - int(m)
		} else if v > 0 && n[i+1] != v+int(m) {
			return v + int(m)
		}
	}
	return 0
}
