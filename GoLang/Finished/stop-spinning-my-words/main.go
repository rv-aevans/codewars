package main

import (
	"strings"
)

var testString = "My fabulous test string"

func main() {
	solution(testString)
}

func solution(text string) string {
	splitText := strings.Split(text, " ")

	solution := ""

	for _, word := range splitText {
		if len(word) >= 5 {
			wordToConcat := ""
			for _, l := range word {
				wordToConcat = string(l) + wordToConcat
			}
			solution = solution + wordToConcat + " "
		} else {
			solution = solution + word + " "
		}
	}

	return strings.TrimSpace(solution)
}
