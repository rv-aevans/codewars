package main

import "fmt"

func main() {
	fmt.Println(Pyramid(5))
}

func Pyramid(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		sub := []int{}
		for j := 0; j <= i; j++ {
			sub = append(sub, 1)
		}
		res[i] = sub
	}
	return res
}
