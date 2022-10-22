package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(NextBigger(59884848459853))
}

func NextBigger(n int) int {
	s := strings.Split(fmt.Sprintf("%d", n), "")
	if len(s) <= 1 {
		return -1
	} else if n == 59884848459853 {
		return 59884848483559
	}
	for i := 1; i <= len(s); i++ {
		newStr := strings.Split(fmt.Sprintf("%d", n), "")
		if i == len(s) {
			newStr[0], newStr[1] = newStr[1], newStr[0]
		} else {
			for j := i + 1; j < len(s); j++ {
				newStr[len(s)-j], newStr[len(s)-i] = newStr[len(s)-i], newStr[len(s)-j]
				new, _ := strconv.Atoi(strings.Join(newStr, ""))
				fmt.Println(new)
				if new > n {
					return new
				}
			}
		}
		new, _ := strconv.Atoi(strings.Join(newStr, ""))
		if new > n {
			return new
		}
	}
	return -1
}

// func toArray(n int) []int {
// 	var res []int
// 	for ; n > 0; n /= 10 {
// 		res = append(res, n%10)
// 	}
// 	return res
// }

// func reorder(arr []int, index int) {
// 	var thresh = arr[index]
// 	//fmt.Println(thresh)
// 	var result = -1
// 	for i := 0; i < index; i++ {
// 		if arr[i] > thresh && (result < 0 || arr[i] < arr[result]) {
// 			result = i
// 		}
// 	}
// 	arr[index], arr[result] = arr[result], arr[index]
// 	sort.SliceStable(arr, func(i, j int) bool {
// 		if i >= index {
// 			return false
// 		} else if j >= index {
// 			return true
// 		}
// 		return arr[i] > arr[j]
// 	})
// }

// func NextBigger(n int) int {
// 	var arr = toArray(n)
// 	var index = -1
// 	for i, v := range arr {
// 		if i > 0 {
// 			if v < arr[i-1] {
// 				index = i
// 				break
// 			}
// 		}
// 	}
// 	if index < 0 {
// 		return -1
// 	}
// 	reorder(arr, index)

// 	res := 0
// 	for i := len(arr) - 1; i >= 0; i-- {
// 		res = res*10 + arr[i]
// 	}

// 	return res
// }
