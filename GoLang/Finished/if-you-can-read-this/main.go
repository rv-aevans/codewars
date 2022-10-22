package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(ToNato("go for it!"))
}

func ToNato(s string) string {
	natoMap := make(map[string]string)
	alpha := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", ",", ".", "!", "?"}
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu", ",", ".", "!", "?"}
	for i, char := range alpha {
		natoMap[char] = nato[i]
	}
	splitString := strings.Split(s, "")
	var spacedString string
	for _, char := range splitString {
		lowerChar := strings.ToLower(char)
		spacedString = spacedString + natoMap[lowerChar] + " "
	}
	pattern := regexp.MustCompile(`\s+`)
	res := pattern.ReplaceAllString(spacedString, " ")
	return strings.TrimSuffix(res, " ")
}
