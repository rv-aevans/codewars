package main

import "fmt"

func main() {
	fmt.Println(TwoSum([]int{3, 8, 12, 19, 45, 62, 83, 197}, 107))
}

func TwoSum(numbers []int, target int) [2]int {
	for i, numOne := range numbers {
		for j, numTwo := range numbers {
			if i == j {
				continue
			} else {
				if numOne+numTwo == target {
					return [2]int{i, j}
				}
			}
		}
	}
	return [2]int{0, 0}
}
