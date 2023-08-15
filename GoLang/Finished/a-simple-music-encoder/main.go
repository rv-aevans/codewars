package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(solution([]int{1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
	fmt.Println(solution([]int{1, 1, 1, 1, 2, 3}))
}

func solution(raw []int) string {
	var solution string
	var prev, cur, next int
	count := 1
	for i := 0; i < len(raw); i++ {
		cur = raw[i]
		if i > 0 {
			prev = raw[i-1]
		}
		if i < len(raw)-1 {
			next = raw[i+1]
		}
		if cur == prev && cur == next {
			count++
			continue
		}
		if cur == prev && cur != next {
			count++
			solution += fmt.Sprintf("%d*%d,", cur, count)
			count = 1
			continue
		}
		if cur != prev && cur != next {
			solution += fmt.Sprintf("%d,", cur)
		}
		if i == len(raw)-1 {
			solution += fmt.Sprintf("%d,", cur)
		}
	}
	return strings.TrimSuffix(solution, ",")
}
