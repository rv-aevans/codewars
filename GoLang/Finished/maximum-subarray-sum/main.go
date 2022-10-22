package main

import "fmt"

func main() {
	fmt.Println(MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func MaximumSubarraySum(numbers []int) (max int) {
	for i := range numbers {
		new := 0
		for _, n := range numbers[i:] {
			new += n
			if new > max {
				max = new
			}
		}
	}
	if max <= 0 {
		return 0
	}
	return
}
