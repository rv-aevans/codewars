package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	m := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	fmt.Println(ValidateSolution(m))
}

var sudokuInts = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func ValidateSolution(m [][]int) bool {
	k := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	for i, row := range m {
		sort.Ints(row)
		if !reflect.DeepEqual(row, sudokuInts) || row[i] == 0 {
			return false
		}
		for j := 0; j < 9; j++ {
			k[j] = append(k[j], row[i])
		}
		fmt.Println(row)
	}

	for _, column := range k {
		keys := make(map[int]bool)
		list := []int{}

		for _, entry := range column {
			if _, value := keys[entry]; !value {
				keys[entry] = true
				list = append(list, entry)
			}
		}

		sort.Ints(list)
		if !reflect.DeepEqual(list, sudokuInts) {
			return false
		}
	}
	return true
}
