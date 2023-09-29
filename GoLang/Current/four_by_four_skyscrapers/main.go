package main

import (
	"fmt"
)

var allOption = [][]int{{1, 2, 3, 4}, {2, 1, 3, 4}, {3, 2, 1, 4}, {2, 3, 1, 4}, {3, 1, 2, 4}, {1, 3, 2, 4}, {4, 1, 2, 3}, {1, 4, 2, 3}, {2, 1, 4, 3}, {1, 2, 4, 3}, {2, 4, 1, 3}, {4, 2, 1, 3}, {3, 4, 1, 2}, {4, 3, 1, 2}, {1, 4, 3, 2}, {4, 1, 3, 2}, {1, 3, 4, 2}, {3, 1, 4, 2}, {2, 3, 4, 1}, {3, 2, 4, 1}, {4, 3, 2, 1}, {3, 4, 2, 1}, {4, 2, 3, 1}, {2, 4, 3, 1}}
var oneOptions = [][]int{}
var twoOptions = [][]int{}
var threeOptions = [][]int{}
var fourOptions = [][]int{{1, 2, 3, 4}}

var colMap = make(map[int][][]int, 4)
var rowMap = make(map[int][][]int, 4)

var clues = []int{
	0, 0, 1, 2,
	0, 2, 0, 0,
	0, 3, 0, 0,
	0, 1, 0, 0}

// var clues = []int{
// 	3, 0, 0, 0,
// 	0, 0, 0, 0,
// 	0, 0, 0, 0,
// 	0, 0, 0, 0}

func main() {
	fmt.Println(SolvePuzzle(clues))
}

func SolvePuzzle(clues []int) [][]int {
	createOptions()
	createSets()

	for i, v := range clues {
		if v == 0 {
			continue
		}

		// if i < 4 {
		// 	handleClue(true, false, i, 0, v)
		// }

		if i < 4 {
			handleClue(true, false, i, 0, v)
		} else if i < 8 {
			handleClue(false, true, 3, i-4, v)
		} else if i < 12 {
			handleClue(true, true, 11-i, 3, v)
		} else {
			handleClue(false, false, 0, 15-i, v)
		}
	}

	fmt.Println("col 1: ", colMap[0])
	fmt.Println("col 2: ", colMap[1])
	fmt.Println("col 3: ", colMap[2])
	fmt.Println("col 4: ", colMap[3])
	fmt.Println("-------------")
	fmt.Println("row 1: ", rowMap[0])
	fmt.Println("row 2: ", rowMap[1])
	fmt.Println("row 3: ", rowMap[2])
	fmt.Println("row 4: ", rowMap[3])

	// rowCountMap := make(map[string]int)

	// for _, v := range rowMap {
	// 	for _, k := range v {
	// 		s := fmt.Sprintf("%d%d%d%d", k[0], k[1], k[2], k[3])
	// 		_, ok := rowCountMap[s]
	// 		if ok {
	// 			rowCountMap[s]++
	// 		} else {
	// 			rowCountMap[s] = 1
	// 		}
	// 	}
	// }

	// fmt.Println(rowCountMap)

	return nil
}

func createOptions() {
	for i := 0; i < 4; i++ {
		colMap[i] = allOption
		rowMap[i] = allOption
	}
}

func createSets() {
	for _, v := range allOption {
		if v[0] == 4 {
			oneOptions = append(oneOptions, v)
		}
		if (v[0] == 1 && v[1] == 4) || (v[0] == 2 && v[1] == 4) || (v[0] == 3 && v[1] == 4) || (v[0] == 3 && v[1] == 2) {
			twoOptions = append(twoOptions, v)
		}
	}
	twoOptions = append(twoOptions, []int{2, 1, 4, 3}, []int{3, 1, 4, 2}, []int{3, 1, 2, 4})
	threeOptions = append(threeOptions, []int{1, 3, 4, 2}, []int{1, 2, 4, 3}, []int{2, 3, 4, 1}, []int{2, 1, 3, 4}, []int{2, 3, 1, 4})
}

func handleClue(col, rev bool, ci, ri, clue int) {
	var mapSetOne, mapSetTwo, newSetOne, newSetTwo, newSetThree [][]int
	mapSetNew := make(map[int][][]int)
	var index int
	var i0, i1, i2, i3 = 0, 1, 2, 3
	if rev {
		i0, i1, i2, i3 = 3, 2, 1, 0
	}

	fmt.Sprintf("%d%d%d%d", i0, i1, i2, i3)

	if col {
		mapSetOne = colMap[ci]
		mapSetTwo = rowMap[ri]

		index = ci
	} else {
		mapSetOne = rowMap[ri]
		mapSetTwo = colMap[ci]

		index = ri
	}

	switch clue {
	case 1:
		for _, v := range mapSetOne {
			if v[i0] == 4 {
				newSetOne = append(newSetOne, v)
			}
		}

		for _, v := range mapSetTwo {
			if v[index] != 4 {
				continue
			}
			newSetTwo = append(newSetTwo, v)
		}

		var newInd int

		if col {
			if newSetOne != nil {
				colMap[ci] = newSetOne
			}
			if newSetTwo != nil {
				rowMap[ri] = newSetTwo
			}
			newInd = ci
			mapSetNew = colMap
		} else {
			if newSetOne != nil {
				rowMap[ri] = newSetOne
			}
			if newSetTwo != nil {
				colMap[ci] = newSetTwo
			}
			newInd = ri
			mapSetNew = rowMap
		}

		for i, v := range mapSetNew {
			if i == newInd {
				continue
			}
			var newVals [][]int
			for _, k := range v {
				if k[i0] == 4 {
					continue
				}
				newVals = append(newVals, k)
			}
			if col {
				colMap[i] = newVals
			} else {
				rowMap[i] = newVals
			}
		}
	case 2:
		for _, v := range mapSetOne {
			for _, k := range twoOptions {
				if v[0] == k[i0] && v[1] == k[i1] && v[2] == k[i2] && v[3] == k[i3] {
					newSetOne = append(newSetOne, v)
				}
			}
		}

		for _, v := range mapSetTwo {
			for _, k := range twoOptions {
				if v[0] == k[0] && v[1] == k[1] && v[2] == k[2] && v[3] == k[3] {
					newSetTwo = append(newSetTwo, v)
				}
			}
		}
	case 3:
		fmt.Println(ri)
		for _, v := range mapSetOne {
			for _, k := range threeOptions {
				if v[0] == k[i0] && v[1] == k[i1] && v[2] == k[i2] && v[3] == k[i3] {
					newSetOne = append(newSetOne, v)
					break
				}
			}
		}

		for _, v := range mapSetTwo {
			if v[index] == 4 || v[index] == 3 {
				continue
			}
			newSetTwo = append(newSetTwo, v)
		}
		for _, v := range rowMap[ri-1] {
			if v[index] == 4 {
				continue
			}
			newSetThree = append(newSetThree, v)
		}
		// case 4:
	}

	if col {
		if newSetOne != nil {
			colMap[ci] = newSetOne
		}
		if newSetTwo != nil {
			rowMap[ri] = newSetTwo
		}
		if newSetThree != nil {
			rowMap[ri-1] = newSetThree
		}
	} else {
		if newSetOne != nil {
			rowMap[ri] = newSetOne
		}
		if newSetTwo != nil {
			colMap[ci] = newSetTwo
		}
		// if newSetThree != nil {
		// 	fmt.Println("do something")
		// }
	}

}

// switch clue {
// case 1:
// 	for _, v := range oldSet {
// 		for _, k := range oneOptions {
// 			if v[i0] == k[0] && v[i1] == k[1] && v[i2] == k[2] && v[i3] == k[3] {
// 				newSet = append(newSet, v)
// 			}
// 		}
// 		if index == 2 {
// 			var newOppSet [][]int
// 			for _, k := range rowMap[0] {
// 				if k[2] != 4 {
// 					continue
// 				}
// 				newOppSet = append(newOppSet, k)
// 			}
// 			rowMap[0] = newOppSet
// 		}
// 		if index == 13 {
// 			var newOppSet [][]int
// 			for _, k := range colMap[0] {
// 				if k[2] != 4 {
// 					continue
// 				}
// 				newOppSet = append(newOppSet, k)
// 			}
// 			colMap[0] = newOppSet
// 		}
// 	}
// case 2:
// 	for _, v := range oldSet {
// 		for _, k := range twoOptions {
// 			if v[i0] == k[0] && v[i1] == k[1] && v[i2] == k[2] && v[i3] == k[3] {
// 				newSet = append(newSet, v)
// 			}
// 		}
// 		if index == 3 {
// 			var newOppSet [][]int
// 			for _, k := range rowMap[0] {
// 				if k[3] == 4 {
// 					continue
// 				}
// 				newOppSet = append(newOppSet, k)
// 			}
// 			rowMap[0] = newOppSet
// 		}
// 		if index == 5 {
// 			var newOppSet [][]int
// 			for _, k := range colMap[3] {
// 				if k[1] == 4 {
// 					continue
// 				}
// 				newOppSet = append(newOppSet, k)
// 			}
// 			colMap[3] = newOppSet
// 		}
// 	}
// case 3:
// 	for _, v := range oldSet {
// 		for _, k := range threeOptions {
// 			if v[i0] == k[0] && v[i1] == k[1] && v[i2] == k[2] && v[i3] == k[3] {
// 				newSet = append(newSet, v)
// 			}
// 		}
// 	}
// 	if index == 9 {
// 		var newOppSetOne, newOppSetTwo [][]int
// 		for _, k := range rowMap[3] {
// 			if k[2] == 4 || k[2] == 3 {
// 				continue
// 			}
// 			newOppSetOne = append(newOppSetOne, k)
// 		}
// 		rowMap[1] = newOppSetOne
// 		for _, k := range rowMap[2] {
// 			if k[2] == 4 {
// 				continue
// 			}
// 			newOppSetTwo = append(newOppSetTwo, k)
// 		}
// 		rowMap[2] = newOppSetTwo
// 	}
// case 4:
// 	newSet = [][]int{{1, 2, 3, 4}}
// }

// if s == "col" {
// 	colMap[num] = newSet
// } else {
// 	rowMap[num] = newSet
// }

// for i := 0; i < 4; i++ {
// 	if len(rowMap[i]) == 2 {
// one, two, three, four := true, true, true, true
// if rowMap[i][0][0] == rowMap[i][1][0] {
// 	one = false
// }
// if rowMap[i][0][1] == rowMap[i][1][1] {
// 	two = false
// }
// if rowMap[i][0][2] == rowMap[i][1][2] {
// 	two = false
// }
// if rowMap[i][0][3] == rowMap[i][1][3] {
// 	three = false
// }
// fmt.Println(one, two, three, four)
// }
// if len(colMap[i]) == 2 {
// 	fmt.Println(colMap[i])
// 	var compareSet [][]int
// 	for i := 0; i < len(colMap[2]); i++ {
// 		for j := 0; j < 4; j++ {
// 			for _, v := range rowMap[j] {
// 				if v[2] == colMap[i][2][j] {
// 					compareSet = append(compareSet, v)
// 				}
// 			}
// 		}
// 	}
// 	for _, v := range colMap[2] {
// 		answer := true
// 		for _, k := range compareSet {
// 			if v[0] == k[0] && v[1] == k[1] && v[2] == k[2] && v[3] == k[3] {
// 				answer = false
// 				break
// 			}
// 		}
// 		if answer {
// 			fmt.Println(v)
// 		}
// 	}
// one, two, three, four := true, true, true, true
// if colMap[i][0][0] == colMap[i][1][0] {
// 	one = false
// }
// if colMap[i][0][1] == colMap[i][1][1] {
// 	two = false
// }
// if colMap[i][0][2] == colMap[i][1][2] {
// 	two = false
// }
// if colMap[i][0][3] == colMap[i][1][3] {
// 	three = false
// }
// fmt.Println(one, two, three, four)
// }
// }
