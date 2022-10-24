package main

import "fmt"

func main() {
	// fmt.Println(snail([][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}))
	// fmt.Println(snail([][]int{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 10, 11, 12}, {13, 14, 15, 16, 17, 18}, {19, 20, 21, 22, 23, 24}, {25, 26, 27, 28, 29, 30}, {31, 32, 33, 34, 35, 36}}))
	// fmt.Println(snail([][]int{{1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14}, {15, 16, 17, 18, 19, 20, 21}, {22, 23, 24, 25, 26, 27, 28}, {29, 30, 31, 32, 33, 34, 35}, {36, 37, 38, 39, 40, 41, 42}, {43, 44, 45, 46, 47, 48, 49}}))
	fmt.Println(snail([][]int{{}}))
}

func snail(snailMap [][]int) []int {
	res := []int{}
	for i := 0; i < len(snailMap)/2; i++ {
		for j := i; j < len(snailMap)-i; j++ {
			res = append(res, snailMap[i][j])
		}
		for j := i + 1; j < len(snailMap)-i; j++ {
			res = append(res, snailMap[j][len(snailMap)-i-1])
		}
		for j := len(snailMap) - 1 - i; j > i; j-- {
			res = append(res, snailMap[len(snailMap)-i-1][j-1])
		}
		for j := len(snailMap) - 2 - i; j > i; j-- {
			res = append(res, snailMap[j][i])
		}
	}
	if len(snailMap)%2 == 0 || len(snailMap[0]) == 0 {
		return res
	}
	return append(res, snailMap[len(snailMap)/2][len(snailMap)/2])
}

/*

1  2  3  4  5  6
7  8  9  10 11 12
13 14 15 16 17 18
19 20 21 22 23 24
25 26 27 28 29 30
31 32 33 34 35 36

1  2  3  4  5  6  7
8  9  10 11 12 13 14
15 16 17 18 19 20 21
22 23 24 25 26 27 28
29 30 31 32 33 34 35
36 37 38 39 40 41 42
43 44 45 46 47 48 49

*/
