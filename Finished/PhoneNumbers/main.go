package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

func CreatePhoneNumber(numbers [10]uint) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("(%d)", numbers[:3]), "[", ""), "]", ""), " ", "") + " " + strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%d-%d", numbers[3:6], numbers[6:10]), "[", ""), "]", ""), " ", "")
}
