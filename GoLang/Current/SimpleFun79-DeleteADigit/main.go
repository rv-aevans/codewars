package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(DeleteDigit(1001))
}

func DeleteDigit(n int) int {
	var highest int
	for i := 0; i < len(strings.Split(fmt.Sprintf("%d", n), "")); i++ {
		strSlice := strings.Split(fmt.Sprintf("%d", n), "")
		strSlice = append(strSlice[:i], strSlice[i+1:]...)
		newInt, _ := strconv.Atoi(strings.Join(strSlice[:], ""))
		if newInt > highest {
			highest = newInt
		}
	}
	return highest
}
