package main

import "fmt"

func main() {
	fmt.Println(Transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// fmt.Println(Transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// fmt.Println(Transpose([][]int{{1, 2, 3}}))
}

func Transpose(matrix [][]int) [][]int {
	res := make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			res[i] = append(res[i], matrix[j][i])
		}
	}
	return res
}
