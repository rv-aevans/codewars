package main

import "fmt"

func main() {
	fmt.Println(SpinningRings(3, 2))
}

func SpinningRings(innerMax, outerMax int) int {
	if innerMax == 0 && outerMax == 0 {
		return 0
	}
	for i, j, k := 1, innerMax, 1; ; i++ {
		if j == k {
			return i
		}
		if j == 0 {
			j = innerMax
		} else {
			j--
		}
		if k == outerMax {
			k = 0
		} else {
			k++
		}
	}
}
