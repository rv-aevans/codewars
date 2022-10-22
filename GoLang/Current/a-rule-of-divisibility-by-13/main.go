package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Thirt(1234567))
}

func Thirt(n int) int {
	len := len(fmt.Sprintf("%d", n))
	var multipliers []int
	for i := 0; i < len; i++ {
		multipliers = append(multipliers, int(math.Pow(10, float64(i)))%13)
	}
	return recur(multipliers, n)
}

func recur(multis []int, n int) (new int) {
	strSlice := strings.Split(fmt.Sprintf("%d", n), "")
	for i, d := range strSlice {
		i++
		num, _ := strconv.Atoi(d)
		new += num * multis[len(strSlice)-i]
	}
	if new == n {
		return new
	}
	return recur(multis, new)
}
