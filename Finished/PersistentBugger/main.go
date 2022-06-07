package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Persistence(999))
}

func Persistence(num int) int {
	if num < 10 {
		return 0
	} else {
		return intToStringToInt(num, 1)
	}
}

func intToStringToInt(num, count int) int {
	var intSlice []int
	for _, num := range strings.Split(fmt.Sprintf("%d", num), "") {
		numInt, _ := strconv.Atoi(num)
		intSlice = append(intSlice, numInt)
	}

	total := 1
	for _, val := range intSlice {
		total *= val
	}

	if total < 10 {
		return count
	} else {
		count++
		return intToStringToInt(total, count)
	}
}
