package main

import (
	"fmt"
)

var allOption = [][]int{{1, 2, 3, 4}, {2, 1, 3, 4}, {3, 2, 1, 4}, {2, 3, 1, 4}, {3, 1, 2, 4}, {1, 3, 2, 4}, {4, 1, 2, 3}, {1, 4, 2, 3}, {2, 1, 4, 3}, {1, 2, 4, 3}, {2, 4, 1, 3}, {4, 2, 1, 3}, {3, 4, 1, 2}, {4, 3, 1, 2}, {1, 4, 3, 2}, {4, 1, 3, 2}, {1, 3, 4, 2}, {3, 1, 4, 2}, {2, 3, 4, 1}, {3, 2, 4, 1}, {4, 3, 2, 1}, {3, 4, 2, 1}, {4, 2, 3, 1}, {2, 4, 3, 1}}
var solMap = make(map[int]bool, 16)
var rowMap = make(map[int][][]int, 4)
var colMap = make(map[int][][]int, 4)

var clues = []int{
	0, 0, 1, 2,
	0, 2, 0, 0,
	0, 3, 0, 0,
	0, 1, 0, 0}

// var clues = []int{
// 	4, 0, 0, 0,
// 	0, 0, 0, 0,
// 	0, 0, 0, 0,
// 	0, 0, 0, 0}

func main() {
	fmt.Println(SolvePuzzle(clues))
}

func SolvePuzzle(clues []int) [][]int {
	createOptions()
	for i, v := range clues {
		if v == 0 {
			continue
		}
		if i < 4 {
			handleClue("col", false, i, v)
		} else if i < 8 {
			handleClue("row", true, i, v)
		} else if i < 12 {
			handleClue("col", true, i, v)
		} else {
			handleClue("row", false, i, v)
		}
	}

	fmt.Println(colMap[2])

	return nil
}

func createOptions() {
	for i := 0; i < 4; i++ {
		rowMap[i] = allOption
		colMap[i] = allOption
	}
	for i := 0; i < 16; i++ {
		solMap[i] = false
	}
}

func handleClue(s string, rev bool, index, clue int) {
	var newSet [][]int
	var oldSet [][]int
	var num int
	var i0, i1, i2, i3 = 0, 1, 2, 3

	switch {
	case s == "col" && index < 8:
		num = index
		oldSet = colMap[num]
	case s == "row" && index < 8:
		num = index - 4
		oldSet = rowMap[num]
	case s == "col" && index >= 8:
		num = 11 - index
		oldSet = colMap[num]
	case s == "row" && index >= 8:
		num = 15 - index
		oldSet = rowMap[num]
	}

	if rev {
		i0, i1, i2, i3 = 3, 2, 1, 0
	}

	switch clue {
	case 1:
		for _, v := range oldSet {
			if v[i0] == 4 {
				newSet = append(newSet, v)
			}
		}
	case 2:

		for _, v := range oldSet {
			if v[i0] == 4 || (v[i0] == 1 && v[i1] != 4) || (v[i0] == 2 && v[i1] == 3) {
				continue
			}
			newSet = append(newSet, v)
		}
	case 3:
		for _, v := range oldSet {
			if v[i0] == 4 || v[i0] == 3 || v[i1] == 4 || (v[i0] == 1 && v[i1] == 2 && v[i2] == 3 && v[i3] == 4) {
				continue
			}
			newSet = append(newSet, v)
		}
	case 4:
		newSet = [][]int{{1, 2, 3, 4}}
	}

	if s == "col" {
		colMap[num] = newSet
	} else {
		rowMap[num] = newSet
	}
}
