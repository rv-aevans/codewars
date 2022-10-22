package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(ToCamelCase("the-stealth-Warrior"))
}

func ToCamelCase(s string) string {
	r := regexp.MustCompile("-|_").Split(s, -1)
	for i, w := range r[1:] {
		r[i+1] = strings.Title(w)
	}
	return strings.Join(r, "")
}
