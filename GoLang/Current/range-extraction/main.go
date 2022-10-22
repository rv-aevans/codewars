package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Solution([]int{-38, -37, -29, -28, -27, -26, -25, -20, -19, -15, -13, -12, -6}))
}

func Solution(list []int) (s string) {
	for i, d := range list {
		if i == 0 {
			s += fmt.Sprintf("%d", d)
		} else if i == len(list)-1 {
			if list[i-1] != list[i-2]+1 {
				s = strings.TrimSuffix(s, "-")
				s += fmt.Sprintf(",%d", d)
			} else if list[i] == list[i-1]+1 {
				s += fmt.Sprintf("%d", d)
			} else {
				s += fmt.Sprintf("%d,%d", list[i-1], d)
			}
		} else if d != list[i-1]+1 && string(s[len(s)-1]) == "-" {
			if list[i-1] != list[i-2]+1 {
				s = strings.TrimSuffix(s, "-")
				s += fmt.Sprintf(",%d,%d", list[i-1], d)
			} else {
				s += fmt.Sprintf("%d,%d", list[i-1], d)
			}
		} else if d == list[i-1]+1 && string(s[len(s)-1]) != "-" {
			if d+1 == list[i+1] {
				s += "-"
			} else if i == len(list)-2 {
				s += ","
			} else {
				s += fmt.Sprintf(",%d", d)
			}
		} else if d == list[i-1]+1 && string(s[len(s)-1]) == "-" && i != len(list)-1 {
			continue
		} else {
			s += fmt.Sprintf(",%d", d)
		}
	}
	return
}
