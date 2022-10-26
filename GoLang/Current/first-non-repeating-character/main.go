package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FirstNonRepeating("moonmEn"))
}

func FirstNonRepeating(str string) (res string) {
	for _, s := range str {
		if strings.Count(strings.ToUpper(str), strings.ToUpper(string(s))) == 1 {
			return string(s)
		}
	}
	return
}
