package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(f([]string{":D", ";D", ":)", ";~)", ":(", "~)"}))
}

func f(s []string) (res int) {
	return len(regexp.MustCompile(`[;:][~-]?[D)]`).FindAll([]byte(strings.Join(s, ",")), -1))
}
