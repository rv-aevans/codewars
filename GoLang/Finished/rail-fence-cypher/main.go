package main

import (
	"fmt"
	"strings"
)

const (
	testString = "abcdefghi"
	testRails  = 4
)

func main() {
	Encode(testString, testRails)
}

func Encode(text string, rails int) string {

	textSlice := strings.Split(text, "")

	output := ""

	if rails == 1 {
		fmt.Println(text)
		return text
	}

	for i := 1; i <= rails; i++ {
		index := 0
		for j := 0; j < len(textSlice); j++ {
			if j == 0 {
				output = output + textSlice[i-1]
			} else {
				// index = index + (currentMultiplier * 2)
				output = output + textSlice[index]
			}
			if index < len(textSlice) {
				output = output + textSlice[index]
			}
		}
	}

	fmt.Println(output)
	return output
}
