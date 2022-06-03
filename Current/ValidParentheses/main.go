package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ValidParentheses("(())"))
}

func ValidParentheses(parens string) bool {
	fmt.Println(parens)
	splitParen := strings.Split(parens, "")
	var open, closed int
	for _, char := range splitParen {
		if char == "(" {
			open++
		} else if char == (")") {
			closed++
		}
	}
	if len(splitParen) == 0 {
		return true
	} else if len(splitParen) == 1 {
		return false
	} else if open != closed {
		return false
	} else if splitParen[len(splitParen)-2] == ")" && splitParen[len(splitParen)-1] == "(" {
		return false
	}
	return true
}
