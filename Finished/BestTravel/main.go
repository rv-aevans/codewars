package main

import "fmt"

var (
	// townDistances = []int{50, 55, 57, 58, 60, 30, 70, 90}
	townDistances = []int{50, 55, 58, 61}
	maximumMiles  = 174
	numberOfTowns = 3
)

func main() {
	solution(numberOfTowns, maximumMiles, townDistances)
}

func solution(k, t int, ls []int) int {
	arr := getAllPossibles(k, ls)
	result := -1
	for _, v := range arr {
		if result < v && v <= t {
			result = v
		}
		fmt.Println("3rd: ", result)
	}
	return result
}

func getAllPossibles(k int, ls []int) []int {
	if k == 1 {
		return ls
	}

	result := []int{}
	for i := 0; i < len(ls)-k+1; i++ {
		arr := getAllPossibles(k-1, ls[i+1:])
		for _, v := range arr {
			result = append(result, ls[i]+v)
		}
		fmt.Println("Last Append: ", result)
	}
	fmt.Println("2nd: ", result)
	return result
}

// func solution(townDistances []int, maxMiles int, stops int) {
// 	highToLowDistances := sortIntsHighToLow(townDistances)

// 	// pick 3 random numbers???

// }

// func sortIntsHighToLow(ints []int) []int {
// 	sort.Ints(ints)

// 	var highToLowInts []int

// 	for i := len(ints) - 1; i >= 0; i-- {
// 		highToLowInts = append(highToLowInts, ints[i])
// 	}

// 	return highToLowInts
// }
