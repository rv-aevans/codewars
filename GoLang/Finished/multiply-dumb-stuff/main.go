package main

import "fmt"

func main() {
	fmt.Println(Comp([]int{121, 144, 19, 161, 19, 144, 19, 11}, []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}))
}

func Comp(arrayOne []int, arrayTwo []int) bool {
	for _, num := range arrayOne {
		for i, numTwo := range arrayTwo {
			if num*num == numTwo {
				break
			}
			if i+1 >= len(arrayTwo) {
				return false
			}
		}
	}
	return true
}
