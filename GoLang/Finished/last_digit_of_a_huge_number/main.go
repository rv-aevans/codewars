package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(LastDigit([]int{99, 2}))
}

// func LastDigit(ns []int) int {
// 	if len(ns) == 0 {
// 		return 1
// 	}

// 	pow := int(math.Pow(float64(ns[0]%10), float64(ns[1])))

// 	return pow % 10
// }

func LastDigit(ns []int) int {
	if len(ns) == 0 {
		return 1
	} else if len(ns) == 1 {
		return ns[len(ns)-1] % 10
	}

	var cur int

	for i := len(ns) - 1; i > 0; i-- {
		cur = twoNumbers(ns[i-1], ns[i])
	}

	fmt.Println(cur)

	return cur % len(fmt.Sprintf("%d", ns[1]))
}

func twoNumbers(r1, r2 int) int {
	r2len := len(fmt.Sprintf("%d", r2))

	fmt.Println(r2len)

	return int(math.Pow(float64(r1%r2len), float64(r2)))
}
