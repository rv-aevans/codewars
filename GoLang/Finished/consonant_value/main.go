package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(solve("abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes"))
}

func solve(s string) (res int) {
	for _, v := range regexp.MustCompile("[^aeiou]+").FindAllStringSubmatch(s, -1) {
		n := 0
		for _, r := range strings.Join(v[:], ",") {
			n += int(r - 96)
		}
		if n > res {
			res = n
		}
	}
	return
}
