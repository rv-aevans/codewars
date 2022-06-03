package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Smallest(261235))
}

func Smallest(n int64) []int64 {
	var numSlice []string
	smallestNum := n
	stringNum := strconv.Itoa(int(n))
	for _, num := range stringNum {
		numSlice = append(numSlice, string(rune(num)))
	}
	for i := 0; i <= len(numSlice); i++ {
		var newSlice []string
		if i == len(numSlice) {
			newSlice = numSlice[:len(numSlice)-1]
		} else if i > 0 {
			newSlice = append(newSlice, numSlice[:i-1]...)
			newSlice = append(newSlice, numSlice[i:]...)
		} else {
			newSlice = append(newSlice, numSlice[i+1:]...)
		}
		for j, digit := range numSlice {
			if len(newSlice) == j {
				newSlice = append(newSlice, digit)
			}
			fmt.Println(newSlice)
		}
	}
	return []int64{smallestNum}
}
