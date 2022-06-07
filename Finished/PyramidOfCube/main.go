package main

import "fmt"

func main() {
	fmt.Println(FindHeight(35))
}

func FindHeight(cubes int) int {
	if cubes == 0 {
		return 0
	} else {
		return calcLayers(2, 1, 1, cubes)
	}
}

func calcLayers(layer, used, prev, max int) int {
	used += prev + layer
	if used == max {
		return layer
	} else if used > max {
		return layer - 1
	} else {
		return calcLayers(layer+1, used, prev+layer, max)
	}
}
