package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ValidParentheses("()()"))
}

func ValidParentheses(parens string) bool {
	openCount := 0
	for _, char := range strings.Split(parens, "") {
		if char == "(" {
			openCount++
		} else if char == ")" && openCount == 0 {
			return false
		} else if char == ")" {
			openCount--
		}
	}

	return openCount == 0
}
