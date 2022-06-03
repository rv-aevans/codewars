package main

import (
	"fmt"
	"unicode"
)

type MyString string

func main() {
	var inputString MyString = "WHAT DOES THE FOX SAY"

	fmt.Println(inputString.IsUpperCase())
}

func (s MyString) IsUpperCase() bool {
	for _, char := range s {
		if unicode.IsLower(char) {
			return false
		}
	}
	return true
}
