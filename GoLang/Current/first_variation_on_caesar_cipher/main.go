package main

import (
	"fmt"
	"math"
	"unicode"
	"unicode/utf8"
)

func main() {
	fmt.Println(MovingShift("I should have known that you would have a perfect answer for me!!!", 1))
}

func MovingShift(s string, shift int) []string {
	n, count, curShift, rev := math.Ceil(float64(len(s))/5), 0, 1, false
	var allStrings []string
	var ns string
	for i := 1; i <= len(s); i++ {
		r, _ := utf8.DecodeRune([]byte{s[i-1]})

		if string(r) == " " {
			if rev {
				rev = false
			} else {
				rev = true
			}
			ns += string(r)
			continue
		}

		count++

		if rev {
			curShift = 0 - count
		}

		if unicode.IsUpper(r) && int(s[i-1])+curShift > 91 {
			ns += string((int(s[i-1]) + curShift - int(math.Ceil(float64(count)/26))*26))
		} else if unicode.IsLower(r) && int(s[i-1])+curShift > 122 {
			ns += string((int(s[i-1]) + curShift - int(math.Ceil(float64(count)/26))*26))
		} else {
			ns += string(int(s[i-1]) + curShift)
		}

		if i%int(n) == 0 {
			allStrings = append(allStrings, ns)
			ns = ""
		}
	}
	return append(allStrings, ns)
}

func DemovingShift(arr []string, shift int) string {
	return ""
}
