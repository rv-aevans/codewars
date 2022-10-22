package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(PosAverage("444996, 699990, 666690, 096904, 600644, 640646, 606469, 409694, 666094, 606490"))
}

func PosAverage(s string) float64 {
	var count int
	splitStr := strings.Split(s, ", ")
	total := (len(splitStr) * (len(splitStr) - 1)) / 2
	for i, word := range splitStr {
		if i > len(splitStr[0])-1 {
			break
		}
		for j := 0; j < 6; j++ {
			if word[j] == splitStr[j+1][j] {
				count++
			}
		}
	}
	return float64(count) / float64(total)
}
