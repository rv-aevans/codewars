package main

import (
	"fmt"
)

func main() {
	fmt.Println(productFib(4895))
	// fmt.Println(fib(144))
}

// func productFib(n uint64) [3]uint64 {
// 	m := int(math.Sqrt(float64(n))) * 2
// 	var i int
// 	for i := 1; i < m; i++ {
// 		if i != 1 && uint64(fib(i)*fib(i-1)) == n {
// 			return [3]uint64{uint64(fib(i - 1)), uint64(fib(i)), uint64(1)}
// 		}
// 	}
// 	return [3]uint64{uint64(fib(i - 2)), uint64(fib(i - 1)), uint64(0)}
// }

// func fib(n int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	return fib(n-1) + fib(n-2)
// }

func productFib(n uint64) [3]uint64 {
	f1, f2 := uint64(0), uint64(1)
	for f1*f2 < n {
		f1, f2 = f2, f2+f1
	}
	if n == f1*f2 {
		return [3]uint64{f1, f2, 1}
	}
	return [3]uint64{f1, f2, 0}
}
