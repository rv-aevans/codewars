package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(Race(720, 850, 70))
}

func Race(v1, v2, g int) [3]int {
	if v1 > v2 {
		return [3]int{-1, -1, -1}
	}
	v1Fps, v2Fps := float64(v1)/60/60, float64(v2)/60/60
	var seconds int

	for i := 1; i > 0; i++ {
		if (v2Fps * float64(i)) > v1Fps*float64(i)+float64(g) {
			seconds = i - 1
			break
		}
	}

	log.Print(v1, v2, g)

	return [3]int{seconds / 60 / 60, seconds / 60 % 60, seconds % 60}
}
