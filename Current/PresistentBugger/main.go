package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := 39
	fmt.Println(Persistence(num, 1))
}

func Persistence(num, count int) int {
	var numSlice []int
	var total int
	if num < 10 {
		return 1
	} else {
		count++
		for _, num := range strings.Split(fmt.Sprintf("%d", num), "") {
			digString, _ := strconv.Atoi(num)
			numSlice = append(numSlice, digString)
		}
		for i, _ := range numSlice {
			a[]
		}
		if total >= 10 {
			Persistence(total, count)
		} else {
			return count
		}
	}
	return 0
}
